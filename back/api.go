package main

import (
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	datos, err := obtenerRegistros()
	if err != nil {
		http.Error(w, "Error al obtener los registros de la BD: ", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)

}
