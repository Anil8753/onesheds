package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/middlewares"
	"github.com/anil8753/onesheds/apps/warehousemen/service/profile"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine   *gin.Engine
	HAuth    *auth.Auth
	HProfile *profile.Profile
}

func InitRoutes(engine *gin.Engine, dep *HandlerDependency) {

	// server static files (web app)
	engine.Use(static.Serve("/", static.LocalFile("./www", false)))

	r := &Routes{
		Engine:   engine,
		HAuth:    &auth.Auth{Dep: dep},
		HProfile: &profile.Profile{Dep: dep},
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
}
