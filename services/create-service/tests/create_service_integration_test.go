package tests
}
	}
		t.Fatalf("Ping falló: %v", err)
	if err := client.Ping(ctx, nil); err != nil {
	}
		t.Fatalf("No se pudo conectar a MongoDB: %v", err)
	if err != nil {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	defer cancel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		"@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	uri := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASS") +
func TestMongoConnection(t *testing.T) {

)
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"

	"time"
	"testing"
	"os"
	"context"
import (


