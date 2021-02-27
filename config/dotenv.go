package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("Error loading .env file")
	}
}
func StringEnvVariable(key string) string {
	loadEnv()
	return os.Getenv(key)
}

func Uint32EnvVariable(key string) uint32 {
	loadEnv()
	ui32, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		log.Panicf("Error parsing to uint32")
		return 0
	}
	return uint32(ui32)
}

func Uint8EnvVariable(key string) uint8 {
	loadEnv()
	ui8, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		log.Panicf("Error parsing to uint8")
		return 0
	}
	return uint8(ui8)
}
