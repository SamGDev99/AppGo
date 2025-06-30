package main

import (
	"fmt"     // para imprimir en consola
	"log"     // para generar logs
	"os"      // para sacar las variables del enviroment
	"strconv" // para convertir los strings de los números
	"strings" // para eliminar simbolo de $ y ,

	"github.com/joho/godotenv" // para asignar el .env
	"gorm.io/driver/postgres"  // driver de bd
	"gorm.io/gorm"             // orm
)

var DB *gorm.DB

func conexionBD() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar .env: ", err)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, name, ssl)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectarse a la BD", err)
	}

	err = DB.AutoMigrate(&Data{})
	if err != nil {
		log.Fatal("Error al crear la tabla: ", err)
	}

}

func guardarDatos(datos []Data) error {
	result := DB.CreateInBatches(&datos, 100)
	if result.Error != nil {
		return result.Error
	}
	/*
		for _, v := range datos {
			result := DB.Create(&v)
			if result.Error != nil {
				return result.Error
			}
			fmt.Printf("Insertando: %+v\n", v)
		}
	*/
	return nil
}

func consultarDatos() (int64, error) {
	var registros int64
	result := DB.Model(&Data{}).Count(&registros)
	if result.Error != nil {
		return 0, result.Error
	}
	return registros, nil
}

func obtenerRegistros() ([]DataResponse, error) {
	var datos []Data

	result := DB.Model(&Data{}).Find(&datos)
	if result.Error != nil {
		return nil, result.Error
	}
	var respuesta []DataResponse
	for _, v := range datos {

		target_from, err := parseCleanFloat(v.Target_from)
		if err != nil {
			log.Printf("Error al convertir target_from: %v", err)
			return nil, err
		}
		target_to, err := parseCleanFloat(v.Target_to)
		if err != nil {
			log.Printf("Error al convertir target_from: %v", err)
			return nil, err
		}

		respuesta = append(respuesta, DataResponse{
			ID:          v.ID,
			Ticker:      v.Ticker,
			Company:     v.Company,
			Brokerage:   v.Brokerage,
			Action:      v.Action,
			Rating_from: v.Rating_from,
			Rating_to:   v.Rating_to,
			Target_from: target_from,
			Target_to:   target_to,
		})
	}

	return respuesta, nil

}

func parseCleanFloat(value string) (float64, error) {
	cleaned := strings.ReplaceAll(value, "$", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	return strconv.ParseFloat(cleaned, 64)
}
