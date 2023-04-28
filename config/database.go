package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dikyayodihamzah/simrs.git/exception"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DB *mongo.Client = ConnectDB()

	host          = os.Getenv("DB_HOST")
	port          = os.Getenv("DB_PORT")
	username      = os.Getenv("DB_USERNAME")
	password      = os.Getenv("DB_PASSWORD")
	setTimeout, _ = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
)

func NewDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(setTimeout)*time.Second)
}

func ConnectDB() *mongo.Client {
	ctx, cancel := NewDBContext()
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, host, port)
	option := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(option)
	exception.PanicIfError(err)

	err = client.Connect(ctx)
	exception.PanicIfError(err)

	//ping database
	err = client.Ping(ctx, readpref.Primary())
	exception.PanicIfError(err)

	return client
}
