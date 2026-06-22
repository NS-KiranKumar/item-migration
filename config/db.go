package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Item_Migration"
	}

	// Context Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Mongo Client Options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Ping Database
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	fmt.Println("MongoDB Connected Successfully")

	// Database Instance
	DB = client.Database(dbName)
	fmt.Println("Mongo URI:", mongoURI)
	fmt.Println("DB Name:", dbName)
}
