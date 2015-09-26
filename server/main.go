package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/caarlos0/env"
	"github.com/caarlos0/getantibody"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type Config struct {
	Port string `env:"PORT" envDefault:"3000"`
}

func main() {
	var config Config
	env.Parse(&config)

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/distributions", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, getantibody.Distributions())
	})

	e.Get("/latest/:os/:arch", func(c *echo.Context) error {
		os := c.Param("os")
		arch := c.Param("arch")
		v, err := getantibody.LatestRelease()
		if err != nil {
			return err
		}
		return c.Redirect(
			http.StatusSeeOther,
			getantibody.DownloadURL(v, os, arch),
		)
	})

	// frontend
	assetHandler := http.FileServer(rice.MustFindBox("static").HTTPBox())
	e.Get("/", func(c *echo.Context) error {
		assetHandler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})
	e.Get("/static/*", func(c *echo.Context) error {
		http.StripPrefix("/static/", assetHandler).
			ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	e.Run(":" + config.Port)
}
