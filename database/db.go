package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
    stringDeConexao := "host=" + os.Getenv("HOST") +
		" user=" + os.Getenv("root") +
		" password=" + os.Getenv("root") +
		" dbname=" + os.Getenv("root") +
		" port=" + os.Getenv("5432") +
		" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
