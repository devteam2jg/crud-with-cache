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

func NewInitializer(infra *Infra, router router.Router) *Initializer {
	return &Initializer{infra: infra, Router: router}
}

func (i *Initializer) InitFeedService() {
	mysqlRepo := feedinfra.NewFeedMySQLRepository(i.infra.RDB)
	cacheRepo := feedinfra.NewFeedCache(mysqlRepo, i.infra.Redis)
	useCase := feeddomain.NewFeedUseCase(cacheRepo)
	feedcontroller.NewFeedController(i.Router, useCase)
}
