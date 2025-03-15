package subscriber

import (
	commentcontroller "crud-with-cache/pkg/comment/controller"
	commentdomain "crud-with-cache/pkg/comment/domain"
	commentinfra "crud-with-cache/pkg/comment/infra"
	"crud-with-cache/router"
)

type Initializer struct {
	infra  *Infra
	Router router.Router
}

func NewInitializer(infra *Infra, router router.Router) *Initializer {
	return &Initializer{infra: infra, Router: router}
}

func (i *Initializer) InitCommentService() {
	mysqlRepo := commentinfra.NewMySQLRepository(i.infra.RDB)
	redisBuffer := commentinfra.NewBuffer(mysqlRepo, i.infra.Buffer)
	useCase := commentdomain.NewCommentUseCase(redisBuffer)
	commentcontroller.NewCommentController(i.Router, useCase)
}
