package cart

import (
	"ecommerce-system/product"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cart struct {
	UserID int               `json:"user_id"`
	Items  []product.Product `json:"items"`
}

var Carts []Cart

// AGREGA ESTA FUNCIÓN AL FINAL: Es la que falta y causa el error en payment
func ApplyToCart(items []product.Product, f func(product.Product) float64) float64 {
	var res float64
	for _, item := range items {
		res += f(item)
	}
	return res
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	var newProd product.Product
	if err := json.NewDecoder(r.Body).Decode(&newProd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, c := range Carts {
		if c.UserID == uID {
			Carts[i].Items = append(Carts[i].Items, newProd)
			json.NewEncoder(w).Encode(Carts[i])
			return
		}
	}

	newCart := Cart{UserID: uID, Items: []product.Product{newProd}}
	Carts = append(Carts, newCart)
	json.NewEncoder(w).Encode(newCart)
}

func ViewCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uID, _ := strconv.Atoi(vars["userId"])
	for _, c := range Carts {
		if c.UserID == uID {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "Carrito no encontrado", http.StatusNotFound)
}
