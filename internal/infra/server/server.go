package server

import (
	"net/http"
	"teste/internal/core/domain"

	"github.com/go-chi/chi"
)

type Server struct {
	router            *chi.Mux
	server            *http.Server
	bancoController   domain.BancoController
	usuarioController domain.UsuarioController
	port              string
}

func NewServer(bancoController domain.BancoController, usuarioController domain.UsuarioController, port string) *Server {
	return &Server{
		router:            chi.NewRouter(),
		bancoController:   bancoController,
		usuarioController: usuarioController,
		port:              port,
	}
}

func (server *Server) ConfigureRoutes() {

	server.router.Post("/banco", server.bancoController.Create)
	server.router.Get("/banco", server.bancoController.ListAll)
	server.router.Get("/banco/{id}", server.bancoController.GetByID)
	server.router.Delete("/banco/{id}", server.bancoController.DeleteByID)
	server.router.Put("/banco/{id}", server.bancoController.Update)

	server.router.Post("/usuario", server.usuarioController.Create)
	server.router.Get("/usuario", server.usuarioController.ListAll)
}

func (server *Server) Start() error {
	server.server = &http.Server{
		Addr:    ":" + server.port,
		Handler: server.router,
	}

	return server.server.ListenAndServe()
}
