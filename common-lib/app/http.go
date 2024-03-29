package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/gogoclouds/gogo-services/common-lib/app/conf"
	"github.com/gogoclouds/gogo-services/common-lib/app/server/http/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/web/gin/validator"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"net/http"
	"time"

	"github.com/gogoclouds/gogo-services/common-lib/app/logger"

	"github.com/gin-gonic/gin"
)

func RunHttpServer(app *App, register func(a *App, e *gin.Engine)) {
	app.Wg.Add(1)
	defer app.Wg.Done()

	e := gin.New()
	e.Use(gin.Logger()) // TODO -> zap.Logger
	e.Use(middleware.Recovery())
	e.Use(middleware.LoggerResponseFail())

	binding.Validator = new(validator.DefaultValidator)
	healthApi(e)     // provide health API
	register(app, e) // register router

	srv := &http.Server{Addr: app.Opts.Conf.Server.Http.Addr, Handler: e}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Panicf("listen: %s\n", err)
		}
	}()
	logger.Infof("http server running %s", app.Opts.Conf.Server.Http.Addr)
	<-app.Exit
	logger.Info("Shutting down http server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Http server forced to shutdown: %w", err)
	}
	logger.Info("http server exiting")
}

// healthApi http check-up API
func healthApi(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		msg := fmt.Sprintf("%s [env=%s] %s, is active", conf.Conf.Name, conf.Conf.Env, conf.Conf.Version)
		r.Reply(c, msg)
	})
}
