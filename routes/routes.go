package routes

import (
	"net/http"

	"github.com/AlwiLion/controllers"
)

func SetupRoutes() {

	http.HandleFunc("/users", controllers.Handle)
}
