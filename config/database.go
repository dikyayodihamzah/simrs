package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dikyayodihamzah/simrs/exception"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
)

func NewDB() *mongo.Client {
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	fmt.Println(dsn)
	option := options.Client().ApplyURI(dsn)

	client, err := mongo.NewClient(option)
	exception.PanicIfError(err)

	err = client.Connect(context.Background())
	exception.PanicIfError(err)

	//ping database
	err = client.Ping(context.Background(), readpref.Primary())
	exception.PanicIfError(err)

	log.Println("Connected to database")
	return client
}
