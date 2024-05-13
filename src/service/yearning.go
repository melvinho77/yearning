package service

import (
	"Yearning-go/src/model"
	_ "Yearning-go/src/model"
	"Yearning-go/src/router"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cookieY/yee"
	"github.com/cookieY/yee/middleware"
)

var f embed.FS

var html string

func loadDBInit() {
	model.DB().First(&model.GloPer)
	_ = json.Unmarshal(model.GloPer.Message, &model.GloMessage)
	_ = json.Unmarshal(model.GloPer.Ldap, &model.GloLdap)
	_ = json.Unmarshal(model.GloPer.Other, &model.GloOther)
	_ = json.Unmarshal(model.GloPer.AuditRole, &model.GloRole)
}

func StartYearning(port string) {
	go cronTabMaskQuery()
	go cronTabTotalTickets()
	loadDBInit()
	e := yee.New()
	e.Pack("/front", f, "dist")
	e.Use(middleware.Cors())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recovery())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9,
	}))
	e.SetLogLevel(model.TransferLogLevel())
	e.GET("/", func(c yee.Context) error {
		return c.HTML(http.StatusOK, html)
	})
	router.AddRouter(e)

	e.Run(fmt.Sprintf("%s", port))
}
