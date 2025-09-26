package controllers

import (
	"spe/database"
	"spe/repositories"
	"spe/services"
)

type Controlador struct {
	AutenticacaoCtrl AutenticacaoControlador
	BolsistaCtrl     BolsistaControlador
}

func NovoControlador() *Controlador {
	// Repositórios.
	cargoRepo := repositories.NovoCargoRepositorio(database.DB)
	setorRepo := repositories.NovoSetorRepositorio(database.DB)
	mebroRepo := repositories.NovoMembroRepositorio(database.DB)
	bolsistaRepo := repositories.NovoBolsistaRepositorio(database.DB)
	vinculoRepo := repositories.NovoVinculoRepositorio(database.DB)

	// Serviços.
	cargoServ := services.NovoCargoServico(cargoRepo)
	setorServ := services.NovoSetorServico(setorRepo)
	autenticacaoServ := services.NovoAutenticacaoServico(mebroRepo)
	membroServ := services.NovoMembroServico(mebroRepo)
	bolsistaServ := services.NovoBolsistaServico(bolsistaRepo)
	vinculoServ := services.NovoVinculoServico(vinculoRepo)

	return &Controlador{
		AutenticacaoCtrl: NovoAutenticacaoControlador(autenticacaoServ),
		BolsistaCtrl: NovoBolsistaControlador(
			bolsistaServ,
			membroServ,
			cargoServ,
			setorServ,
			vinculoServ,
		),
	}
}
