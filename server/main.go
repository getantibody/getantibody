package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/getantibody"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/caarlos0/env"
)

type Config struct {
	Port string `env:"PORT" envDefault:"3000"`
}

const download = "https://github.com/caarlos0/antibody/releases/download/%s/antibody_%s_%s.tar.gz"

func main() {
	var config Config
	env.Parse(&config)

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/latest/:os/:arch", func(c *echo.Context) error {
		os := c.Param("os")
		arch := c.Param("arch")
		v, err := getantibody.LatestRelease()
		if err != nil {
			return err
		}
		return c.Redirect(
			http.StatusSeeOther,
			fmt.Sprintf(download, v, os, arch),
		)
	})

	e.Run(":" + config.Port)
}
