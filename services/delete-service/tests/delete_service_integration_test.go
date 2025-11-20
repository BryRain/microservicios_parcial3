package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoURI() string {
	// Tomar la URI del entorno SIEMPRE que exista
	uri := os.Getenv("MONGO_URI")
	if uri != "" {
		return uri
	}

	// DEFAULT LOCAL SIN AUTENTICACIÓN (para desarrollo local)
	return "mongodb://localhost:27017"
}

func TestMongoConnection(t *testing.T) {
	uri := getMongoURI()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		t.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	// Intentar ping para verificar autenticación y conexión
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("No se pudo hacer ping a MongoDB: %v", err)
	}
}


