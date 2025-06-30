package main

import (
	"fmt" // para imprimir en consola
	"log" // para generar logs
	"os"  // para sacar las variables del enviroment

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
		respuesta = append(respuesta, DataResponse{
			Ticker:      v.Ticker,
			Company:     v.Company,
			Brokerage:   v.Brokerage,
			Action:      v.Action,
			Rating_from: v.Rating_from,
			Rating_to:   v.Rating_to,
			Target_from: v.Target_from,
			Target_to:   v.Target_to,
		})
	}

	return respuesta, nil

}
