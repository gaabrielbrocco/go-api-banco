package server

import (
	"net/http"
	"teste/internal/core/domain"
	"teste/internal/infra/controller"

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
	server.router.Get("/banco", bancoController.ListAll)
	server.router.Get("/banco/{id}", bancoController.GetByID)
	server.router.Delete("/banco/{id}", bancoController.DeleteByID)
	server.router.Put("/banco/{id}", bancoController.Update)

}

func (server *Server) Start() error {
	server.server = &http.Server{
		Addr:    ":" + server.port,
		Handler: server.router,
	}

	return server.server.ListenAndServe()
}
