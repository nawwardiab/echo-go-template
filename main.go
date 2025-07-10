package main

import (
	"fmt"
	"log"
	"os"

	"echo-server/internal/db"
	"echo-server/internal/handler"
	"echo-server/internal/repository"
	"echo-server/internal/service"
	"echo-server/internal/view"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	// load config from .env
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize db connection
	// retrieve db connection string values from .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")

	// put them together
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPwd, dbHost, dbPort, dbName)

	// connect to db
	dbConn, dbConnErr := db.NewDB(connStr)
	if dbConnErr != nil {
		log.Fatalf("failed to connect to DB: %v", dbConnErr)
	}
	defer dbConn.Close()

	
	// wire repositories & services
	// user repo and service
	userRepo := repository.NewUserRepo(dbConn)
	userSvc := service.NewUserService(userRepo)

	// products repo and service
	productRepo := repository.NewProductRepo(dbConn)
	productSvc := service.NewProductService(productRepo)
	
	
	// wire handlers
	uh := handler.NewAuthHandler(userSvc)
	ph := handler.NewProductHandler(productSvc)
	ch := handler.NewCartHandler(productSvc)
	
	// bootstrap echo
	e:= echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	
	// create session
	sessKey := os.Getenv("SESSION_KEY")
	store := sessions.NewCookieStore([]byte(sessKey))
	e.Use(session.Middleware(store))


	// assign view/templates.go to echo renderer
	rend, rendErr := view.New("templates")
	if rendErr != nil {
		log.Fatalf("view: %v", rendErr)
	}
	e.Renderer = rend

	// Serve Static files
	staticDir := os.Getenv("STATIC_DIR")
	e.Static("/staticFiles", staticDir)

	// wire endpoints' handlers
	// Home
	e.GET("/",              handler.ViewHome)

	// User Auth
	e.GET("/login",         uh.LoginForm)
	e.POST("/login",        uh.LoginSubmit)
	e.GET("/register",      uh.RegisterForm)
	e.POST("/register",     uh.RegisterSubmit)
	e.GET("/logout",        uh.Logout)

	// Products
	e.GET("/products",           ph.ListProducts)
	e.GET("/products/:id",       ph.ListProductDetails)

	// Cart
	e.POST("/cart/add",          ch.AddToCart)
	e.POST("/cart/remove",       ch.RemoveFromCart)
	e.GET("/cart",               ch.ViewCart)

	// start server
	serverPort := os.Getenv("SERVER_PORT")
	serverHost := os.Getenv("SERVER_HOST")
	addStr := serverHost + ":" + serverPort
	e.Logger.Fatal(e.Start(addStr))
}