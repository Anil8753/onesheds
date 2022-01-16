package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/db"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/profile"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/warehouse"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/middlewares"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine     *gin.Engine
	HAuth      *auth.Auth
	HProfile   *profile.Profile
	HWarehouse *warehouse.Warehouse
}

func InitRoutes(engine *gin.Engine) {

	db := db.New()
	ledger := ledger.Ledger{}
	ledger.Init()

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:     engine,
		HAuth:      &auth.Auth{Database: db, Ledger: &ledger},
		HProfile:   &profile.Profile{Database: db, Ledger: &ledger},
		HWarehouse: &warehouse.Warehouse{Database: db, Ledger: &ledger},
	}

	api := r.Engine.Group("/api")

	// public
	public := api.Group("/v1")
	public.POST("/signup", r.HAuth.SignupHandler())
	public.POST("/signin", r.HAuth.SigninHandler())
	public.POST("/refreshtoken", r.HAuth.RefreshTokenHandler())

	// protected
	protected := api.Group("/v1")
	protected.Use(middlewares.JwtAuth())

	protected.GET("/profile", r.HProfile.GetProfileHandler())
	protected.GET("/warehouse", r.HWarehouse.GetWarehouseHandler())
}