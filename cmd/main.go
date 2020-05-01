package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/hthl85/example-gen-thumbnails-func/thumbnails"
	"log"
	"os"
)

const (
	portKey = "PORT"
	portVal = "8080"
)

func main() {
	funcframework.RegisterEventFunction("/", thumbnails.GenThumbnails)
	// Use PORT environment variable, or default to 8080.
	port := portVal
	if envPort := os.Getenv(portKey); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
