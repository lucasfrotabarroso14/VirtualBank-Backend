package main

import (
	"fmt"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/config"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()

	r := router.Gerar() //meu arquivo de rotas
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permitir qualquer origem
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	fmt.Println("Escutando na porta", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler))

}

// migrate -path migrations -database "sua_string_de_conexao" up (para subir as migrations)
