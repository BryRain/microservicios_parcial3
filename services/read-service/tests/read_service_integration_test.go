package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnection(t *testing.T) {
	uri := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASS") +
		"@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("Ping falló: %v", err)
	}
}

