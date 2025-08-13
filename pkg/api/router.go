package api

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/peopleig/food-ordering-go/pkg/controllers"
	"github.com/peopleig/food-ordering-go/pkg/middleware"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func Run() *mux.Router {
	router := SetupRouter()
	PrintRoutes()
	return router
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	utils.DefinePath(router)
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
	adminRouter.HandleFunc("/dish", controllers.AdminDishHandler).Methods("GET", "POST")

	billRouter := protected.PathPrefix("/bill").Subrouter()
	billRouter.Use(middleware.AllowAdminandIdAccess)
	billRouter.HandleFunc("/", controllers.BillPayerHandler).Methods("POST")
	billRouter.HandleFunc("/{status}/{order_id}", controllers.SingleBillHandler).Methods("GET")

	protected.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:8001")
}
