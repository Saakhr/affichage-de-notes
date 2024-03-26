package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	gowebly "github.com/gowebly/helpers"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer(apiCFG ApiConf) error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8080"))
	if err != nil {
		return err
	}

	// Create a new Echo server.
	echo := echo.New()

	// Add Echo middlewares.
	echo.Use(middleware.Logger())
	echo.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// Handle static files.
	echo.Static("/static", "./static")

	// Handle index page view.
	echo.GET("/", apiCFG.middlewareAuth(apiCFG.indexViewHandler))

	// Handle API endpoints.
	echo.POST("/api/login", apiCFG.LoginApiHandler)
	echo.GET("/api/add_etu", apiCFG.middlewareAdminPer(apiCFG.add_etu_page))
	echo.POST("/api/add_etu", apiCFG.middlewareAdminPer(apiCFG.add_etu))
	echo.GET("/api/mng_etu", apiCFG.middlewareAdminPer(apiCFG.mng_etu_page))
	echo.PUT("/api/mng_etu", apiCFG.middlewareAdminPer(apiCFG.edit_etu))
	echo.DELETE("/api/mng_etu", apiCFG.middlewareAdminPer(apiCFG.delete_etu))
	echo.GET("/api/add_grp", apiCFG.middlewareAdminPer(apiCFG.add_grp_page))
	echo.POST("/api/add_grp", apiCFG.middlewareAdminPer(apiCFG.add_grp))
	echo.GET("/api/mng_grp", apiCFG.middlewareAdminPer(apiCFG.mng_grp_page))
	echo.PUT("/api/mng_grp", apiCFG.middlewareAdminPer(apiCFG.edit_grp))
	echo.DELETE("/api/mng_grp", apiCFG.middlewareAdminPer(apiCFG.delete_etu))
	echo.GET("/api/add_spec", apiCFG.middlewareAdminPer(add_spec_page))
	echo.POST("/api/add_spec", apiCFG.middlewareAdminPer(apiCFG.add_spec))
	echo.GET("/api/mng_spec", apiCFG.middlewareAdminPer(apiCFG.mng_spec_page))
	echo.PUT("/api/mng_spec", apiCFG.middlewareAdminPer(apiCFG.edit_spec))
	echo.DELETE("/api/mng_spec", apiCFG.middlewareAdminPer(apiCFG.delete_spec))
	echo.GET("/api/add_pro", apiCFG.middlewareAdminPer(add_prof_page))
	echo.POST("/api/add_pro", apiCFG.middlewareAdminPer(apiCFG.add_prof))
	echo.GET("/api/mng_pro", apiCFG.middlewareAdminPer(apiCFG.mng_prof_page))
	echo.PUT("/api/mng_pro", apiCFG.middlewareAdminPer(apiCFG.edit_prof))
	echo.DELETE("/api/mng_pro", apiCFG.middlewareAdminPer(apiCFG.delete_prof))
	echo.GET("/api/add_mod", apiCFG.middlewareAdminPer(apiCFG.add_mod_page))
	echo.POST("/api/add_mod", apiCFG.middlewareAdminPer(apiCFG.add_mod))
	echo.GET("/api/mng_mod", apiCFG.middlewareAdminPer(apiCFG.mng_mod_page))
	echo.PUT("/api/mng_mod", apiCFG.middlewareAdminPer(apiCFG.edit_mod))
	echo.DELETE("/api/mng_mod", apiCFG.middlewareAdminPer(apiCFG.delete_mod))

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      echo, // handle all Echo routes
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
