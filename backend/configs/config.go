package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"net"
	"os"
)

// LoadEnv loads the environment variables from .env file
func EnvMongoURI() string {
	value , ok := os.LookupEnv("MONGO_URI")
	if (ok == true) {
		return value
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Environment variable {MONGO_URI} not found!")
	}
	return os.Getenv("MONGO_URI")
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func EnvPort() string {
	value , ok := os.LookupEnv("PORT")
	if (ok == true) {
		return value
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Environment variable {PORT} not found!")
	}
	return os.Getenv("PORT")
}

func EnvHostAddress() string {
	value , ok := os.LookupEnv("HOST_ADDRESS")
	if (ok == true) {
		return value
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Environment variable {HOST_ADDRESS} not found!")
	}
	return os.Getenv("HOST_ADDRESS")
}



func ConnectDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(EnvMongoURI())

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

var DB = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("example").Collection(collectionName)
	return collection
}
