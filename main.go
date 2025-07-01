package main

import (
	"ai_writer/internal/interactor/pkg/connect"
	"fmt"
	"net/http"

	_ "ai_writer/api"
	"ai_writer/internal/interactor/pkg/util/log"
	"ai_writer/internal/router"
	"ai_writer/internal/router/article"
	"ai_writer/internal/router/login"
	"ai_writer/internal/router/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main is run all api form localhost port 8080

//	@title			AI WRITER APIs
//	@version		0.1
//	@description	AI WRITER APIs
//	@termsOfService

//	@contact.name
//	@contact.url
//	@contact.email

//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.en.html

// @host		writer.t.api.qplan.ai
// @BasePath	/sien/v1.0
// @schemes	https
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

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8081/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8081", engine))
}
