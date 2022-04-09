package main

import (
	"github.com/anil8753/onesheds/handlers/auth"
	"github.com/anil8753/onesheds/handlers/faq"
	"github.com/anil8753/onesheds/handlers/knowledgebase"
	"github.com/anil8753/onesheds/handlers/order"
	"github.com/anil8753/onesheds/handlers/profile"
	"github.com/anil8753/onesheds/handlers/review"
	"github.com/anil8753/onesheds/handlers/warehouse"
	"github.com/anil8753/onesheds/interfaces"
	"github.com/anil8753/onesheds/ledger"
	"github.com/anil8753/onesheds/middlewares"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine         *gin.Engine
	HAuth          *auth.Auth
	HProfile       *profile.Profile
	HWarehouse     *warehouse.Warehouse
	HOrder         *order.Handler
	HReview        *review.Handler
	HFAQ           *faq.Handler
	HKnowledgeBase *knowledgebase.Handler
}

func InitRoutes(engine *gin.Engine, db interfaces.Database, ledger *ledger.Ledger) {

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:         engine,
		HAuth:          &auth.Auth{Database: db, Ledger: ledger},
		HProfile:       &profile.Profile{Database: db, Ledger: ledger},
		HWarehouse:     &warehouse.Warehouse{Database: db, Ledger: ledger},
		HOrder:         &order.Handler{Database: db, Ledger: ledger},
		HReview:        &review.Handler{Database: db, Ledger: ledger},
		HFAQ:           &faq.Handler{Database: db, Ledger: ledger},
		HKnowledgeBase: &knowledgebase.Handler{Database: db, Ledger: ledger},
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

	// Profile
	protected.GET("/profile", r.HProfile.GetProfileHandler())
	protected.PUT("/profile", r.HProfile.UpdateProfileHandler())

	// warehouse
	protected.POST("/warehouse", r.HWarehouse.CreateWarehouseHandler())
	protected.GET("/warehouse", r.HWarehouse.GetWarehousesHandler())
	protected.PUT("/warehouse", r.HWarehouse.UpdateWarehouseHandler())

	protected.GET("/warehouse/queryall", r.HWarehouse.QuaryAllHandler())
	protected.GET("/warehouse/query", r.HWarehouse.QueryHandler())
	protected.GET("/warehouse/querypagination", r.HWarehouse.QueryPaginationHandler())

	// order
	protected.POST("/order", r.HOrder.NewOrder())
	protected.GET("/order", r.HOrder.GetAllOrders())

	// reviews
	protected.GET("/reviews/:warehouse_id", r.HReview.GetAllWarehouseReviews())
	protected.POST("/review", r.HReview.AddUserRating())
	protected.POST("/review_reply", r.HReview.AddReply())
	protected.GET("/review/:review_id", r.HReview.GetReview())

	// faq
	protected.GET("/faq/:warehouse_id", r.HFAQ.GetAllFAQ())
	protected.POST("/faq/add", r.HFAQ.AddFAQ())
	protected.PUT("/faq/question", r.HFAQ.UpdateFAQQuestion())
	protected.PUT("/faq/answer", r.HFAQ.UpdateFAQAnswer())
	protected.PUT("/faq/delete", r.HFAQ.DeleteFAQ())

	// knowledgebase
	protected.GET("/knowledgebase/:warehouse_id", r.HKnowledgeBase.GetAll())
	protected.POST("/knowledgebase/question", r.HKnowledgeBase.AddQuestion())
	protected.POST("/knowledgebase/answer", r.HKnowledgeBase.AddAnswer())

}
