package server

import (
	"net/http"
	"teste/internals/core/domain"
	"teste/internals/infra/controller"

	"github.com/go-chi/chi"
)

type Server struct {
	router       *chi.Mux
	server       *http.Server
	bancoUseCase domain.BancoUseCase
	port         string
}

func NewServer(bancoUseCase domain.BancoUseCase, port string) *Server {
	return &Server{
		router:       chi.NewRouter(),
		bancoUseCase: bancoUseCase,
		port:         port,
	}
}

func (server *Server) ConfigureRoutes() {
	bancoController := controller.NewBancoController(server.bancoUseCase)

	server.router.Post("/banco", bancoController.Create)
	server.router.Get("/banco/{id}", bancoController.GetByID)

}

func (server *Server) Start() error {
	server.server = &http.Server{
		Addr:    ":" + server.port,
		Handler: server.router,
	}

	return server.server.ListenAndServe()
}
