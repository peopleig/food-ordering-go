package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peopleig/food-ordering-go/pkg/controllers"
	"github.com/peopleig/food-ordering-go/pkg/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	staticFileDirectory := http.Dir("web/static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	// userController := controllers.NewUserController()

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/signup", controllers.SignupHandler).Methods("GET", "POST")
	router.Handle("/menu", middleware.JWTMiddleware(http.HandlerFunc(controllers.MenuHandler))).Methods("GET", "POST")
	router.Handle("/chef", middleware.JWTMiddleware(middleware.AllowChefAccess(http.HandlerFunc(controllers.ChefHandler)))).Methods("GET", "PATCH", "POST")
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	// router.HandleFunc("/users/add", userController.CreateUser).Methods("POST")
	// router.HandleFunc("/users/update/{id}", userController.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/delete/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:8001")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /users")
	fmt.Println("  GET  /users/{id}")
	fmt.Println("  POST /users/add")
	fmt.Println("  PUT  /users/update/{id}")
	fmt.Println("  DELETE /users/delete/{id}")
}
