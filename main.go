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
		handler(w, r, token)
	})

	http.HandleFunc("/prueba", handler2)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, _ *http.Request, token string) {
	resp, err := solicitudAPI(token)
	if err != nil {
		http.Error(w, "Error en solicitud a la API: "+err.Error(), http.StatusInternalServerError)
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

}

func solicitudAPI(token string) (*http.Response, error) {
	// Se debe crear cliente personalizado para realizar solicitud HTTP
	client := &http.Client{}
	//Solicitud a API
	req, err := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list", nil)
	if err != nil {
		return nil, err
	}
	// Añadir token
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Probando url")
}
