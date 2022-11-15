package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ElladanTasartir/golang-mongodb/pkg/http"
	"github.com/ElladanTasartir/golang-mongodb/pkg/storage"
)

func main() {
	client := storage.Connect(getEnv("MONGODB_HOST"), getEnv("MONGODB_DATABASE"))

	port, err := strconv.Atoi(getEnv("PORT"))
	if err != nil {
		log.Fatalf("Port is not a number %d", port)
	}

	defer client.CloseConnection()

	http.Run(port, client)
}

func getEnv(key string) string {
	return os.Getenv(key)
}
