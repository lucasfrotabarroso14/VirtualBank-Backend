package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	//string conexao com o mysql
	ConnectionStringDB = " "

	//porta  onde a api vai estar rodando
	Port = 0

	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente
func LoadEnvs() {

	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}
	// ou podia ter botado embaixo : (localhost:3399)
	ConnectionStringDB = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3399)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
