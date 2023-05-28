package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Mongo struct {
	*mongo.Client
}

func getTlsMongoUri(baseMongoUri string) string {
	isMongoTlsEnabled := os.Getenv("MONGO_TLS_ENABLED")
	tlsRootCaPath := os.Getenv("ROOT_CA_PATH")
	tlsCertificatePath := os.Getenv("CERTIFICATE_PATH")
	tlsPrivateKeyPath := os.Getenv("PRIVATE_KEY_PATH")

	if isMongoTlsEnabled != "true" {
		log.Println("tls is not enabled, returning base mongouri isMongoTlsEnabled:", isMongoTlsEnabled)
		return baseMongoUri
	}

	tlsArg := fmt.Sprintf("%v=%v", "tls", "true")
	tlsCertificateFileArg := fmt.Sprintf("%v=%v", "tlsCertificateFile", tlsCertificatePath)
	tlsPrivateKeyFileArg := fmt.Sprintf("%v=%v", "tlsPrivateKeyFile", tlsPrivateKeyPath)
	tlsRootCaFileArg := fmt.Sprintf("%v=%v", "tlsCAFile", tlsRootCaPath)

	mongoUriArgs := []string{baseMongoUri, tlsArg, tlsPrivateKeyFileArg, tlsCertificateFileArg, tlsRootCaFileArg}
	mongo_uri := strings.Join(mongoUriArgs, "&")

	return mongo_uri
}

func NewMongoClient() Mongo {
	uri := getTlsMongoUri(os.Getenv("MONGO_URI"))

	_, err := connstring.ParseAndValidate(uri)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	connectionOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(connectionOptions)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Connected to MongoDB")

	return Mongo{client}
}