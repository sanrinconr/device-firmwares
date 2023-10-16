package dependencies

import (
	"fmt"
	"os"

	"github.com/sanrinconr/device-firmwares/dc-nearshore/database"
)

func resolveDatabase() (database.Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	return database.New(dsn)
}
