package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/faq"
	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/order"
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
	Engine           *gin.Engine
	HandlerAuth      *auth.Auth
	HandlerProfile   *profile.Profile
	HandlerWarehouse *warehouse.Warehouse
	HandlerOrder     *order.Handler
	HandlerReview    *review.Handler
	HandlerFAQ       *faq.Handler
}

func InitRoutes(engine *gin.Engine, db interfaces.Database, ledger *ledger.Ledger) {

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:           engine,
		HandlerAuth:      &auth.Auth{Database: db, Ledger: ledger},
		HandlerProfile:   &profile.Profile{Database: db, Ledger: ledger},
		HandlerWarehouse: &warehouse.Warehouse{Database: db, Ledger: ledger},
		HandlerOrder:     &order.Handler{Database: db, Ledger: ledger},
		HandlerReview:    &review.Handler{Database: db, Ledger: ledger},
		HandlerFAQ:       &faq.Handler{Database: db, Ledger: ledger},
	}

	api := r.Engine.Group("/api")

	// public
	public := api.Group("/v1")
	public.POST("/signup", r.HandlerAuth.SignupHandler())
	public.POST("/signin", r.HandlerAuth.SigninHandler())
	public.POST("/refreshtoken", r.HandlerAuth.RefreshTokenHandler())

	// protected
	protected := api.Group("/v1")
	protected.Use(middlewares.JwtAuth())

	protected.GET("/profile", r.HandlerProfile.GetProfileHandler())

	protected.GET("/warehouse", r.HandlerWarehouse.QuaryAllHandler())
	protected.GET("/warehouse/query", r.HandlerWarehouse.QueryHandler())
	protected.GET("/warehouse/querypagination", r.HandlerWarehouse.QueryPaginationHandler())

	protected.POST("/order", r.HandlerOrder.NewOrder())
	protected.GET("/order", r.HandlerOrder.GetAllOrders())

	protected.GET("/review/warehouse/:warehouse_id", r.HandlerReview.GetAllWarehouseReviews())
	protected.GET("/review/:review_id", r.HandlerReview.GetReview())
	protected.POST("/review", r.HandlerReview.AddUserRating())
	protected.POST("/review_reply", r.HandlerReview.AddReply())

	protected.GET("/faq/warehouse/:warehouse_id", r.HandlerFAQ.GetAllFAQ())
	protected.POST("/faq", r.HandlerFAQ.AddFAQ())
	protected.PUT("/faq/question", r.HandlerFAQ.UpdateFAQQuestion())
	protected.PUT("/faq/answer", r.HandlerFAQ.UpdateFAQAnswer())
}
