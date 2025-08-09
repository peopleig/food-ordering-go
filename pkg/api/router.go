package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peopleig/food-ordering-go/pkg/controllers"
	"github.com/peopleig/food-ordering-go/pkg/middleware"
)

func Run() *mux.Router {
	router := SetupRouter()
	PrintRoutes()
	return router
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	staticFileDirectory := http.Dir("web/static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/signup", controllers.SignupHandler).Methods("GET", "POST")

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/menu", controllers.MenuHandler).Methods("GET", "POST")

	chefRouter := protected.PathPrefix("/chef").Subrouter()
	chefRouter.Use(middleware.AllowChefAccess)
	chefRouter.HandleFunc("", controllers.ChefHandler).Methods("GET", "PATCH", "POST")

	adminRouter := protected.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AllowAdminAccess)
	adminRouter.HandleFunc("", controllers.AdminHandler).Methods("GET")
	adminRouter.HandleFunc("/{user_id}", controllers.AdminApproveHandler).Methods("PATCH")

	billRouter := protected.PathPrefix("/bill").Subrouter()
	billRouter.Use(middleware.AllowAdminandIdAccess)
	billRouter.HandleFunc("", controllers.BillHandler).Methods("GET")
	billRouter.HandleFunc("", controllers.BillPayerHandler).Methods("POST")

	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:8001")
}
