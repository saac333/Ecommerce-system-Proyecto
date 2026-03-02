package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Product usa campos exportados con tags JSON para mantener la encapsulación funcional
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var Products []Product

type ProductManager interface {
	AddProduct(p Product) error
	GetProduct(id int) (Product, error)
}

type ProductService struct{}

func (ps *ProductService) AddProduct(p Product) error {
	if p.Price <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}
	Products = append(Products, p)
	return nil
}

func (ps *ProductService) GetProduct(id int) (Product, error) {
	for _, p := range Products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("producto no encontrado")
}

var service = &ProductService{}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.AddProduct(newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, err := service.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}
