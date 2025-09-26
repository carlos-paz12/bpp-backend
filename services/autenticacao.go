package services

import (
	"spe/models/dto"
	"spe/repositories"

	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var assinatura = []byte(os.Getenv("API_SECRET"))

type AutenticacaoServico interface {
	Autenticar(req *dto.ReqAutenticacaoDTO) (*dto.ResAutenticacaoDTO, error)
}

type AutenticacaoServicoImpl struct {
	membroRepo repositories.MembroRepositorio
}

func NovoAutenticacaoServico(membroRepo repositories.MembroRepositorio) AutenticacaoServico {
	return &AutenticacaoServicoImpl{
		membroRepo: membroRepo,
	}
}

func (serv *AutenticacaoServicoImpl) Autenticar(req *dto.ReqAutenticacaoDTO) (*dto.ResAutenticacaoDTO, error) {
	membro, err := serv.membroRepo.BuscarPeloNomeUsuario(req.NomeUsuario)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(membro.SenhaHash), []byte(req.Senha)) != nil {
		return nil, errors.New("nome de usuário ou senha inválidos")
	}

	cargoDoMembro, err := serv.descobrirCargo(membro.ID)

	claims := &dto.ReivindicacoesJWT{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Cargo:    cargoDoMembro,
		MembroID: membro.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenAssinado, err := token.SignedString(assinatura)

	if err != nil {
		return nil, err
	}

	return &dto.ResAutenticacaoDTO{
		Token: tokenAssinado,
		Cargo: cargoDoMembro,
		Membro: dto.MembroDTO{
			MembroID:     membro.ID,
			NomeCompleto: membro.NomeCompleto,
			NomeUsuario:  membro.NomeUsuario,
			Email:        membro.Email,
			Celular:      membro.Celular,
			Matricula:    membro.Matricula,
		},
	}, nil
}

func (serv *AutenticacaoServicoImpl) descobrirCargo(mid uint) (string, error) {
	cargos, err := serv.membroRepo.BuscarCargosPeloMembroID(mid)
	if err != nil {
		return "", err
	}

	if len(cargos) == 0 {
		return "", errors.New("membro não está associado a nenhum cargo")
	}

	// Retorna só a primeira role pq no início do sistema cada usuário só terá uma role.
	return cargos[0].Nome, nil
}
