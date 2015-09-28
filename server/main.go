package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/caarlos0/getantibody"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
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
			return c.String(http.StatusBadRequest, err.Error())
		}
		url, err := getantibody.DownloadURL(v, os, arch)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.Redirect(http.StatusSeeOther, url)
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

	e.Run(":3000")
}
