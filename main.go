package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"./controllers"
	"./drivers"
	"./middlewares"
	"./routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	gotenv "github.com/subosito/gotenv"
)

// Gobal DB handler
var db *sql.DB

func init() {
	// Load environment variables on startup
	err := gotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	// Open DB handler and defer close
	db = drivers.ConnectDB()
	defer db.Close()

	// Initialize controller, middleware and routes
	controller := controllers.Controller{}
	middleware := middlewares.Middleware{}
	routes := routes.Routes{}

	// Gorilla mux router
	r := mux.NewRouter()

	// Get routes and pass the router, controller, middleware and db
	routes.Get(r, controller, middleware, db)

	// Start server on port specified in the SERVER_PORT environment variable
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))
}
