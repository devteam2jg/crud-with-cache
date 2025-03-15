package subscriber

import "crud-with-cache/router"

func NewServer(infra *Infra) *Server {
	return &Server{
		Infra: infra,
	}
}

type Server struct {
	Infra *Infra
}

func (s Server) RegisterRouter(router router.Router) {
	Initializer := NewInitializer(s.Infra, router)
	Initializer.InitCommentBuffer()
}

func MiddlewareInitializer(i *Infra) {

}
