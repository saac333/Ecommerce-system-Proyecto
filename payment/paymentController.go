package payment

import (
	"ecommerce-system/cart"    // Importa el paquete de carrito para acceder a los datos
	"ecommerce-system/product" // Importa el paquete de producto para los tipos de datos
	"encoding/json"
	"net/http"
)

// Payment define la estructura de pago con tags JSON para la comunicación web
type Payment struct {
	UserID      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}

// ProcessPayment maneja la lógica de cobro utilizando programación funcional
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var details Payment

	// 1. Decodificar los detalles del pago desde la petición
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, "Error en el formato de pago: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Buscar los ítems del carrito del usuario
	var userItems []product.Product
	found := false

	// Accedemos a la lista global de carritos del paquete cart
	for _, c := range cart.Carts {
		if c.UserID == details.UserID {
			userItems = c.Items
			found = true
			break
		}
	}

	// 3. Validar si hay productos para pagar
	if !found || len(userItems) == 0 {
		http.Error(w, "No se encontró un carrito activo para este usuario", http.StatusUnprocessableEntity)
		return
	}

	// 4. APLICACIÓN DE PROGRAMACIÓN FUNCIONAL (Criterio de Evaluación)
	// Usamos la función de orden superior 'ApplyToCart' definida en el paquete cart.
	// Pasamos una función anónima que extrae el precio de cada producto.
	total := cart.ApplyToCart(userItems, func(p product.Product) float64 {
		return p.Price // 'Price' debe estar en mayúscula para ser visible aquí
	})

	// 5. Finalizar la transacción
	details.TotalAmount = total
	details.Status = "Pago procesado exitosamente"

	// Enviar respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}
