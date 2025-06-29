package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func extraerDatos() ([]Data, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar .env: ", err)
	}
	// Bearer token desde .env
	token := os.Getenv("token")

	// Se debe crear cliente personalizado para realizar solicitud HTTP
	client := &http.Client{}

	// Armar una solicitud GET a la API
	req, err := http.NewRequest("GET",
		"https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list", nil)

	if err != nil {
		return nil, err
	}

	// Añadir token a solicitud GET
	req.Header.Add("Authorization", "Bearer "+token)

	// Realizar solicitud y guardar respuesta
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result APIResponse
	// Decodificar respuesta JSON
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	// Devolver información obtenida
	return result.Items, nil
}
