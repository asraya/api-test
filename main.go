package main

import (
	"api-test/database/migration"
	"api-test/database/seeder"
	"api-test/engine"
	"os"
)

func main() {
	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "rollback" {
		migration.RunRollback()
	} else if dbEvent == "migration" {
		migration.RunMigration()
	} else if dbEvent == "seeder" {
		migration.RunMigration()
		seeder.RunSeeder()
	}

	r := engine.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
