package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/order"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/profile"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/warehouse"
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
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
	HOrder     *order.Handler
}

func InitRoutes(engine *gin.Engine, db interfaces.Database, ledger *ledger.Ledger) {

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:     engine,
		HAuth:      &auth.Auth{Database: db, Ledger: ledger},
		HProfile:   &profile.Profile{Database: db, Ledger: ledger},
		HWarehouse: &warehouse.Warehouse{Database: db, Ledger: ledger},
		HOrder:  &order.Handler{Database: db, Ledger: ledger},
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

	protected.GET("/warehouse", r.HWarehouse.QuaryAllHandler())
	protected.GET("/warehouse/query", r.HWarehouse.QueryHandler())
	protected.GET("/warehouse/querypagination", r.HWarehouse.QueryPaginationHandler())
}
