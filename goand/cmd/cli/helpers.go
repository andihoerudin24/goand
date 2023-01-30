package main

import (
	"github.com/joho/godotenv"
	"os"
)

func setUp() {
	err := godotenv.Load()
	if err != nil {
		exitGraceFuly(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGraceFuly(err)
	}
	cel.RootPath = path
	cel.DB.DataType = os.Getenv("DATABASE_TYPE")
}
