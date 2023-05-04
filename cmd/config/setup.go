package config

import (
	"flag"
	"service-routes/cmd/handler"
	"service-routes/internal/domain/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "./data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	Setup(configPath)
}

func Run(configPath string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	conf := GetConfig()
	setupDB(conf)
	ioc := genIoc(conf)
	e = handler.NewCommentEntry(e, ioc["comment"].(usecase.CommentsUseCase), AuthVerify)
	e = handler.NewResourceEntry(e, ioc["resource"].(usecase.ResourcesUseCase), AuthVerify)
	e = handler.NewRoutesEntry(e, ioc["routes"].(usecase.RoutesUseCase), AuthVerify)
	e = handler.NewStepsEntry(e, ioc["steps"].(usecase.StepsUseCase), AuthVerify)
	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
