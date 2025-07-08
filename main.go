package main

import (
	"fmt"
	"log"

	"echo-server/internal/config"
	"echo-server/internal/db"
	"echo-server/internal/handler"
	"echo-server/internal/repository"
	"echo-server/internal/service"
	"echo-server/internal/session"
	"echo-server/internal/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Makefile
// I need targets, dependencies for commands:
// migrate up: T ./migrations up
// migrate down: T ./migrations down
// run: T main.go
// compile: T
// clean:
// reload:

func main(){
	// 1. load config from config.yaml
	cfg, configErr := config.Load("config.yaml")
	if configErr != nil {
		log.Fatalf("failed to load cofig: %v", configErr)
	}

	// 2. initialize db connection
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.USER, cfg.DB.PWD, cfg.DB.HOST, cfg.DB.PORT, cfg.DB.DBNAME)
	dbConn, dbConnErr := db.NewDB(connStr)
	if dbConnErr != nil {
		log.Fatalf("failed to connect to DB: %v", dbConnErr)
	}
	defer dbConn.Close()

	// 3. initialize session store
	sess := session.NewSession(cfg.Session.Key)

	// 4. wire repositories & services
	// 	4.1 user repo and service
	userRepo := repository.NewUserRepo(dbConn)
	userSvc := service.NewUserService(userRepo)
	// 	4.2 products repo and service
	productRepo := repository.NewProductRepo(dbConn)
	productSvc := service.NewProductService(productRepo)

	
	// 5. wire handlers
	// 	5.1 initialize handlers and pass arguments
	uh := handler.NewAuthHandler(userSvc, sess)
	ph := handler.NewProductHandler(productSvc, sess)
	hh := handler.NewHomeHandler(sess)
	ch := handler.NewCartHandler(sess, productSvc)
	
	// 6. bootstrap echo
	e:= echo.New()
	// echo.WrapMiddleware(middleware.ServeStatic(cfg.StaticDir))
	// e.Pre()
	// e.Pre(middleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))


	// assign view/templates.go to echo renderer
	rend, rendErr := view.New("templates")
	if rendErr != nil {
		log.Fatalf("view: %v", rendErr)
	}
	e.Renderer = rend

	// 7. Serve Static files
	e.Static("/staticFiles", cfg.StaticDir)

	// 8. wire handlers
	e.GET("/",              hh.ViewHome)                
	e.GET("/login",         uh.LoginForm)
	e.POST("/login",        uh.LoginSubmit)
	e.GET("/register",      uh.RegisterForm)
	e.POST("/register",     uh.RegisterSubmit)
	e.GET("/logout",        uh.Logout)

	e.GET("/products",           ph.ListProducts)
	e.GET("/products/:id",       ph.ListProductDetails)
	e.POST("/cart/add",          ch.AddToCart)
	e.POST("/cart/remove",       ch.RemoveFromCart)
	e.GET("/cart",               ch.ViewCart)

	addStr := cfg.Server.HOST + ":" + cfg.Server.PORT
	// 9. start server
	e.Logger.Fatal(e.Start(addStr))
}