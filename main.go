package main

import (
	"os"

	"github.com/NicolasDeveloper/store-catalog-api/api"
)

func main() {
	api.StartUp(os.Getenv("PORT"))
}
