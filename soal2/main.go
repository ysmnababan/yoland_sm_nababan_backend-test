package main

import (
	"soal2/config"
	"soal2/murid"
	"soal2/murid/api/repository"
	"soal2/murid/api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.Config();
	e := echo.New();

	repo:= &repository.DB{DB: db}
	service := &service.Service{MuridRepository: repo}
	controller := murid.Controller{MuridService: service}

	e.GET("/murids", controller.GetAllMurid);
	e.Logger.Fatal(e.Start(":5050"));
}