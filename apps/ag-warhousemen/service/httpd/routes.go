package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/faq"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/profile"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/review"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/warehouse"
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/middlewares"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine        *gin.Engine
	HAuth         *auth.Auth
	HProfile      *profile.Profile
	HWarehouse    *warehouse.Asset
	HandlerReview *review.Handler
	HandlerFAQ    *faq.Handler
}

func InitRoutes(engine *gin.Engine, db interfaces.Database, ledger *ledger.Ledger) {

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:        engine,
		HAuth:         &auth.Auth{Database: db, Ledger: ledger},
		HProfile:      &profile.Profile{Database: db, Ledger: ledger},
		HWarehouse:    &warehouse.Asset{Database: db, Ledger: ledger},
		HandlerReview: &review.Handler{Database: db, Ledger: ledger},
		HandlerFAQ:    &faq.Handler{Database: db, Ledger: ledger},
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

	// user
	//protected.GET("/identity", r.HProfile.GetIdentityHandler())
	protected.GET("/profile", r.HProfile.GetProfileHandler())
	protected.PUT("/profile", r.HProfile.UpdateProfileHandler())

	// warehouse
	protected.POST("/warehouse", r.HWarehouse.CreateWarehouseHandler())
	protected.GET("/warehouse", r.HWarehouse.GetWarehousesHandler())
	protected.PUT("/warehouse", r.HWarehouse.UpdateWarehouseHandler())

	// reviews
	protected.GET("/review/warehouse/:warehouse_id", r.HandlerReview.GetAllWarehouseReviews())
	protected.GET("/review/:review_id", r.HandlerReview.GetReview())
	protected.POST("/review_reply", r.HandlerReview.AddReply())

	// faq
	protected.GET("/faq/warehouse/:warehouse_id", r.HandlerFAQ.GetAllFAQ())
	protected.POST("/faq", r.HandlerFAQ.AddFAQ())
	protected.PUT("/faq/question", r.HandlerFAQ.UpdateFAQQuestion())
	protected.PUT("/faq/answer", r.HandlerFAQ.UpdateFAQAnswer())
}
