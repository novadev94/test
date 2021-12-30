package server

import (
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/chikob3/test/api"
	"github.com/chikob3/test/inject"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	conform   = modifiers.New()
	decoder   = form.NewDecoder()
	validate  = validator.New()
	pathRegex = regexp.MustCompile(`^\/?([A-z\d]+)/([A-z\d]+)/([A-z\d]+)/([A-z\d]+)\/?$`)
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	deployTime := time.Now()
	engine := gin.Default()
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployTime": deployTime,
			"timestamp":  time.Now(),
		})
	})
	server := &Server{
		engine: engine,
	}
	server.registerRoutes()
	return server
}

func (a *Server) registerRoutes() {
	injector := inject.NewInjector("mainnet")
	// TODO provide services then register routes
	binanceService := injector.ProvideService()
	binanceApi := api.NewApi(binanceService)

	lendingApi := injector.ProvideLendingApi()

	v1 := a.engine.Group("v1")
	v1.GET("tokenPrice", binanceApi.TokenPrice)

	v1.GET("lending", lendingApi.Overview)
}

func (a *Server) Serve(port string) error {
	return a.engine.Run(fmt.Sprintf(":%s", port))
}
