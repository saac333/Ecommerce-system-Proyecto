package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-system/cart"
	"ecommerce-system/payment"
	"ecommerce-system/product"
	"ecommerce-system/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoints de Producto
	r.HandleFunc("/product", product.AddProduct).Methods("POST")
	r.HandleFunc("/product/{id}", product.GetProduct).Methods("GET")

	// Endpoints de Usuario
	r.HandleFunc("/user", user.RegisterUser).Methods("POST")
	r.HandleFunc("/user/{id}", user.GetUser).Methods("GET")

	// Endpoints de Carrito
	r.HandleFunc("/cart/{userId}", cart.AddToCart).Methods("POST")
	r.HandleFunc("/cart/{userId}", cart.ViewCart).Methods("GET")

	// Endpoints de Pago
	r.HandleFunc("/payment/{userId}", payment.ProcessPayment).Methods("POST")

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
