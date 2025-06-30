package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Nos conectamos a la BD
	conexionBD()

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

		fmt.Println("Se extraen datos de la API y se guardan de forma exitosa")
	} else {
		fmt.Println("La BD ya tiene información. No se consulta la API")
	}

	/*
		// Probando inserción de un solo dato
			info := Data{
				Ticker:      "AKBA",
				Company:     "Akebia Therapeutics",
				Brokerage:   "HC Wainwright",
				Action:      "initiated by",
				Rating_from: "Buy",
				Rating_to:   "Buy",
				Target_from: "$8.00",
				Target_to:   "$8.00",
			}
			result := DB.Create(&info)
			if result.Error != nil {
				log.Fatal("Ocurrió un error al crear el registro: ", result.Error)
			} else {
				fmt.Println("Se guardó registro con el ID: ", info.ID)
			}
	*/

	http.HandleFunc("/api/datos", handler)
	http.HandleFunc("/prueba", handler2)
	fmt.Println("Servidor escuchando en http://localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Probando url")
}
