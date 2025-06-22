package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Obtener token del sistema
	// Se debe ejecutar export BEARER_TOKEN="" en la terminal antes de correr el código
	token := os.Getenv("BEARER_TOKEN")
	if token == "" {
		fmt.Println("Ocurrió un error al recibir el token")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Se debe crear cliente personalizado para realizar solicitud HTTP
		client := &http.Client{}
		//Solicitud a API
		req, err := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list", nil)
		if err != nil {
			http.Error(w, "Error creando solicitud", http.StatusInternalServerError)
			return
		}
		// Añadir token
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error obteniendo datos", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		// Leer respuesta
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error leyendo cuerpo de respuesta", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bodyBytes)

	})
	http.ListenAndServe(":9000", nil)
}
