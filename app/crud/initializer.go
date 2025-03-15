package crud

import (
	commentcontroller "crud-with-cache/pkg/comment/controller"
	commentdomain "crud-with-cache/pkg/comment/domain"
	commentinfra "crud-with-cache/pkg/comment/infra"
	feedcontroller "crud-with-cache/pkg/feed/controller"
	feeddomain "crud-with-cache/pkg/feed/domain"
	feedinfra "crud-with-cache/pkg/feed/infra"
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
	mysqlRepo := feedinfra.NewMySQLRepository(i.infra.RDB)
	cacheRepo := feedinfra.NewCache(mysqlRepo, i.infra.Cache)
	useCase := feeddomain.NewFeedUseCase(cacheRepo)
	feedcontroller.NewFeedController(i.Router, useCase)
}

func (i *Initializer) InitCommentService() {
	mysqlRepo := commentinfra.NewMySQLRepository(i.infra.RDB)
	redisBuffer := commentinfra.NewBuffer(mysqlRepo, i.infra.Buffer)
	redisCache := commentinfra.NewCache(redisBuffer, i.infra.Cache)
	useCase := commentdomain.NewCommentUseCase(redisCache)
	commentcontroller.NewCommentController(i.Router, useCase)
}
