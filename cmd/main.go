package main

import (
	"github.com/DmitryKuzmenec/ImgPrettify/config"
	"github.com/DmitryKuzmenec/ImgPrettify/handlers"
	"github.com/DmitryKuzmenec/ImgPrettify/repositories"
	"github.com/DmitryKuzmenec/ImgPrettify/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Error(err)
	}

	repo := repositories.NewImgRepo(config)
	service := services.NewImgSvc(repo)
	handler := handlers.NewImgHandler(service)

	e := echo.New()

	e.Static("/", "frontend/build")

	users := e.Group("/v1")
	users.POST("/image/pretty", handler.Pretty)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "frontend/build",
		Index: "index.html",
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
