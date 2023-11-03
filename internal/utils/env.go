package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDatabaseEnv() (map[string]interface{}, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	sslMode := os.Getenv("SSLMODE")

	var infoDB = map[string]interface{}{
		"host": dbHost,
		"port": dbPort,
		"user": dbUser,
		"password": dbPassword,
		"name": dbName,
		"sslmode": sslMode,
	}

	return infoDB, nil
}