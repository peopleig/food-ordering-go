package api

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("/error", controllers.ErrorController).Methods("GET")
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/menu", controllers.MenuHandler).Methods("GET", "POST")
	protected.HandleFunc("/api/menu", controllers.ApiMenuHandler).Methods("GET")

	chefRouter := protected.PathPrefix("/chef").Subrouter()
	chefRouter.Use(middleware.AllowChefAccess)
	chefRouter.HandleFunc("", controllers.ChefHandler).Methods("GET", "PATCH", "POST")

	adminRouter := protected.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AllowAdminAccess)
	adminRouter.HandleFunc("", controllers.AdminHandler).Methods("GET")
	adminRouter.HandleFunc("/{user_id}", controllers.AdminApproveHandler).Methods("PATCH", "DELETE")
	adminRouter.HandleFunc("/dish", controllers.AdminDishHandler).Methods("GET", "POST")

	protected.HandleFunc("/bill", controllers.GetMyBills).Methods("GET")
	billRouter := protected.PathPrefix("/bill").Subrouter()
	billRouter.Use(middleware.AllowAdminandIdAccess)
	billRouter.HandleFunc("/{order_id}", controllers.BillPayerHandler).Methods("POST")
	billRouter.HandleFunc("/{status}/{order_id}", controllers.SingleBillHandler).Methods("GET")

	protected.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(controllers.Custom404Handler)
	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:8001")
}
