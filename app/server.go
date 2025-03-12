package app

func NewServer() {

}

type Server struct {
	Infra *Infra
}

func (s Server) RegisterRouter() {
	Initializer := NewInitializer(s.Infra)
	Initializer.InitFeedService()
}

func MiddlewareInitializer(i *Infra) {

}
