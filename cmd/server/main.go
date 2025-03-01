package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/vdscruz/ytools-api/config"
	"github.com/vdscruz/ytools-api/handlers"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic("Erro ao carregar configuração: " + err.Error())
	}

	startServer()
}

func startServer() {
	r := router.New()
	r.GET("/process/{id}", handlers.HandlerYtDescription)
	r.GET("/stream/{id}", handlers.HandlerStream)

	port := ":8081"
	log.Printf("🚀 ytools-api rodando na porta %s", port)

	if err := fasthttp.ListenAndServe(":8081", r.Handler); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
