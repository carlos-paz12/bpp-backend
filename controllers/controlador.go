package controllers

import (
	"spe/database"
	"spe/repositories"
	"spe/services"
)

type Controlador struct {
	AutenticacaoCtrl          AutenticacaoControlador
	BolsistaCtrl              BolsistaControlador
	TecnicoAdministrativoCtrl TecnicoAdministrativoControlador
}

func NovoControlador() *Controlador {
	// Repositórios.
	cargoRepo := repositories.NovoCargoRepositorio(database.DB)
	setorRepo := repositories.NovoSetorRepositorio(database.DB)
	mebroRepo := repositories.NovoMembroRepositorio(database.DB)
	vinculoRepo := repositories.NovoVinculoRepositorio(database.DB)
	tecnicoAdministrativoRepo := repositories.NovoTecnicoAdministrativoRepositorio(database.DB)
	bolsistaRepo := repositories.NovoBolsistaRepositorio(database.DB)

	// Serviços.
	cargoServ := services.NovoCargoServico(cargoRepo)
	setorServ := services.NovoSetorServico(setorRepo)
	autenticacaoServ := services.NovoAutenticacaoServico(mebroRepo)
	membroServ := services.NovoMembroServico(mebroRepo)
	vinculoServ := services.NovoVinculoServico(vinculoRepo)
	tecnicoAdministrativoServ := services.NovoTecnicoAdministrativoServico(tecnicoAdministrativoRepo)
	bolsistaServ := services.NovoBolsistaServico(bolsistaRepo)

	return &Controlador{
		AutenticacaoCtrl: NovoAutenticacaoControlador(autenticacaoServ),
		BolsistaCtrl: NovoBolsistaControlador(
			bolsistaServ,
			membroServ,
			cargoServ,
			setorServ,
			vinculoServ,
		),
		TecnicoAdministrativoCtrl: NovoTecnicoAdministrativoControlador(
			tecnicoAdministrativoServ,
			membroServ,
			cargoServ,
			setorServ,
			vinculoServ,
		),
	}
}
