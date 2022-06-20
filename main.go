package main

import (
	"api-test/config"
	"os"
)

func main() {

	r := config.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))

}
