package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func enviarDatos(w http.ResponseWriter, r *http.Request) {

	datos, err := obtenerRegistros()
	if err != nil {
		http.Error(w, "Error al obtener los registros de la BD: ", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)

}

func cargarDatos(w http.ResponseWriter, r *http.Request) {
	// Consultamos si ya cargamos la información de la API en la BD
	registros, err := consultarDatos()
	if err != nil {
		log.Fatal("Error al consultar los registros: ", err)
	}
	// Si no hemos cargado nada, procedemos a consultar la API y cargar la información
	if registros == 0 {
		datos, err := extraerDatos()
		if err != nil {
			log.Fatal("Ocurrió un error al extraer los datos de la API", err)
		}

		err = guardarDatos(datos)
		if err != nil {
			log.Fatal("Error al guardar la información: ", err)
		}

		fmt.Fprint(w, "Se extraen datos de la API y se guardan de forma exitosa")
	} else {
		fmt.Fprint(w, "La BD ya tiene información. No se consulta la API")
	}
}
