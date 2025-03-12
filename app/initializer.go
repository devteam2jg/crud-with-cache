package app

import (
	feedcontroller "crud-with-cache/pkg/feedsvc/controller"
	feeddomain "crud-with-cache/pkg/feedsvc/domain"
	feedinfra "crud-with-cache/pkg/feedsvc/infra"
)

type Initializer struct {
	infra *Infra
}

func NewInitializer(infra *Infra) *Initializer {
	return &Initializer{infra: infra}
}

func (i *Initializer) InitFeedService() {
	mysqlrepo := feedinfra.NewFeedMySQLRepository(i.infra.RDB)
	usecase := feeddomain.NewFeedUseCase(mysqlrepo)
	feedcontroller.NewFeedController(usecase)
}
