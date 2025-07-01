package main

import (
	"ai_writer/internal/interactor/pkg/connect"
	"ai_writer/internal/interactor/pkg/util/log"
	"ai_writer/internal/router"
	"ai_writer/internal/router/article"
	"ai_writer/internal/router/login"
	"ai_writer/internal/router/user"

	"github.com/apex/gateway"
)

func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}

	engine := router.Default()
	article.GetRouter(engine, db)
	user.GetRouter(engine, db)
	login.GetRouter(engine, db)

	log.Fatal(gateway.ListenAndServe(":8080", engine))
}
