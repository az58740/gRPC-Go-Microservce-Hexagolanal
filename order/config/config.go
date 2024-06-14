package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}
func GetDataSourceURL() string {
	fmt.Println(getEnvironmentValue("DATA_SOURCE_URL"))
	return getEnvironmentValue("DATA_SOURCE_URL")
}
func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

func getEnvironmentValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file err:%v", err)
	}

	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable id missing", key)
	}
	return os.Getenv(key)
}
