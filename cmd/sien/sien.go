package main

import (
	"sien_apis/internal/interactor/pkg/connect"
	"sien_apis/internal/interactor/pkg/util/log"
	"sien_apis/internal/router"
	"sien_apis/internal/router/article"
	"sien_apis/internal/router/login"
	"sien_apis/internal/router/user"

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
