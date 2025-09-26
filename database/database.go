package database

import (
	"spe/migrations"
	"spe/seeders"

	"log"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	// Connection to database.
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("banco de dados conectado.")

	DB = db

	// Migrações.
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.CriaTabelaCargos(),
		migrations.CriaTabelaSetores(),
		migrations.CriaTabelaMembros(),
		migrations.CriaTabelaVinculos(),
		migrations.CriaTabelaTecnicosAdministrativos(),
		migrations.CriaTabelaBolsistas(),
	})
	if err := m.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("migrações aplicadas.")

	// Seeders.
	if err := seeders.PopulaSetores(db); err != nil {
		log.Fatal(err)
	}

	if err := seeders.PopulaCargos(db); err != nil {
		log.Fatal(err)
	}

	if err := seeders.CriaMembroRedesInicial(db); err != nil {
		log.Fatal(err)
	}

	log.Println("seeds aplicadas.")
}
