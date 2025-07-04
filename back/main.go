package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Nos conectamos a la BD
	conexionBD()

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

	http.HandleFunc("/api/datos", enviarDatos)
	http.HandleFunc("/cargarDatos", cargarDatos)
	http.HandleFunc("/calcularDatos", enviarCalificaciones)
	http.HandleFunc("/prueba", handler2)
	fmt.Println("Servidor escuchando en http://localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Probando url")
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
