package routes

import (
	"basic_api/controllers"
	"basic_api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRouter() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowCredentials: true,
	})
	hand := c.Handler(r)
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/signin", controllers.Signin).Methods("POST")
	r.HandleFunc("/logout", middleware.IsAuthorized(controllers.Logout)).Methods("POST")
	r.HandleFunc("/user", middleware.IsAuthorized(controllers.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.IsAuthorized(controllers.Deleteuser)).Methods("DELETE")

	//todos
	r.HandleFunc("/todo", middleware.IsAuthorized(controllers.GetTodos)).Methods("GET")
	r.HandleFunc("/todo/{todoid}", middleware.IsAuthorized(controllers.GetTodo)).Methods("GET")
	r.HandleFunc("/todo", middleware.IsAuthorized(controllers.CreateTodo)).Methods("POST")
	r.HandleFunc("/todo/{todoid}", middleware.IsAuthorized(controllers.UpdateTodo)).Methods("PUT")
	r.HandleFunc("/todo/{todoid}", middleware.IsAuthorized(controllers.DeleteTodo)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", hand))
}
