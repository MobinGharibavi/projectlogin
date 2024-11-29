// main.go
package main

import (
	"net/http"
	"project/controllers"
)

func main() {
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)

	http.ListenAndServe(":8080", nil)
}
