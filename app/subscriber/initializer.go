package subscriber

import (
	"context"
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

func (i *Initializer) InitCommentBuffer() {
	mysqlRepo := commentinfra.NewMySQLRepository(i.infra.RDB)
	redisBuffer := commentinfra.NewSubscriberBuffer(mysqlRepo, i.infra.Buffer)
	err := redisBuffer.WaitForMessage(context.Background())
	if err != nil {
		panic(err)
	}
}
