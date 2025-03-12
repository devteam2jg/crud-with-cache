package app

import (
	feedcontroller "crud-with-cache/pkg/feedsvc/controller"
	feeddomain "crud-with-cache/pkg/feedsvc/domain"
	feedinfra "crud-with-cache/pkg/feedsvc/infra"
	"crud-with-cache/router"
)

type Initializer struct {
	infra  *Infra
	Router router.Router
}

func NewInitializer(infra *Infra) *Initializer {
	return &Initializer{infra: infra}
}

func (i *Initializer) InitFeedService() {
	mysqlRepo := feedinfra.NewFeedMySQLRepository(i.infra.RDB)
	useCase := feeddomain.NewFeedUseCase(mysqlRepo)
	feedcontroller.NewFeedController(i.Router, useCase)
}
