package controllers

import (
	"spe/database"
	"spe/repositories"
	"spe/services"
)

type Controlador struct {
	SetorCtrl                 SetorControlador
	AutenticacaoCtrl          AutenticacaoControlador
	TecnicoAdministrativoCtrl TecnicoAdministrativoControlador
	BolsistaCtrl              BolsistaControlador
}

func NovoControlador() *Controlador {
	// Repositórios.
	setorRepo := repositories.NovoSetorRepositorio(database.DB)
	cargoRepo := repositories.NovoCargoRepositorio(database.DB)
	mebroRepo := repositories.NovoMembroRepositorio(database.DB)
	vinculoRepo := repositories.NovoVinculoRepositorio(database.DB)
	tecnicoAdministrativoRepo := repositories.NovoTecnicoAdministrativoRepositorio(database.DB)
	bolsistaRepo := repositories.NovoBolsistaRepositorio(database.DB)

	// Serviços.
	setorServ := services.NovoSetorServico(setorRepo)
	cargoServ := services.NovoCargoServico(cargoRepo)
	autenticacaoServ := services.NovoAutenticacaoServico(mebroRepo)
	membroServ := services.NovoMembroServico(mebroRepo)
	vinculoServ := services.NovoVinculoServico(vinculoRepo)
	tecnicoAdministrativoServ := services.NovoTecnicoAdministrativoServico(tecnicoAdministrativoRepo)
	bolsistaServ := services.NovoBolsistaServico(bolsistaRepo)

	return &Controlador{
		SetorCtrl:        NovoSetorControlador(setorServ),
		AutenticacaoCtrl: NovoAutenticacaoControlador(autenticacaoServ),
		TecnicoAdministrativoCtrl: NovoTecnicoAdministrativoControlador(
			tecnicoAdministrativoServ,
			membroServ,
			cargoServ,
			setorServ,
			vinculoServ,
		),
		BolsistaCtrl: NovoBolsistaControlador(
			bolsistaServ,
			membroServ,
			cargoServ,
			setorServ,
			vinculoServ,
		),
	}
}
