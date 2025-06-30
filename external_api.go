package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func extraerDatos() ([]Data, error) {

	var datos []Data
	var result APIResponse
	urlInicial := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

	// Se debe crear cliente personalizado para realizar solicitud HTTP
	client := &http.Client{}

	urlCompleta := ""
	next := ""
	parametros := url.Values{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar .env: ", err)
	}
	// Bearer token desde .env
	token := os.Getenv("token")

	for {
		// Verificar si está el campo next para asignar parametro
		if next != "" {
			parametros.Set("next_page", next)
			urlCompleta = urlInicial + "?" + parametros.Encode()
		} else {
			urlCompleta = urlInicial
		}

		// Armar una solicitud GET a la API
		req, err := http.NewRequest("GET", urlCompleta, nil)
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

		// Decodificar respuesta JSON
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}

		datos = append(datos, result.Items...)
		if result.NextPage == "" {
			break
		}

		next = result.NextPage
		fmt.Println("Siguiente página:", result.NextPage)
	}

	// Devolver información obtenida
	return datos, nil
}
