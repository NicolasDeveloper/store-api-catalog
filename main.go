// Package classification Catalog API.
//
// Catalog API to manager products
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Nicolas Silva<nicolas.senac15@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta
package main

import (
	"os"

	"github.com/NicolasDeveloper/store-catalog-api/api"
)

func main() {
	api.StartUp(os.Getenv("PORT"))
}
