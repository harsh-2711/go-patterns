package connection

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

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

var mongoClient *mongo.Client

func ConnectMongo() {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	if mongoClient != nil {
		return
	}

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

	// ctx, cancel := context.WithTimeout(context.Background(), 25 * time.Second)
	// defer cancel()
	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
		panic(err)
	}

	mongoClient = client

	log.Println("Connected to MongoDB")
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		ConnectMongo()
	}

	return mongoClient
}

// DEPRECATED

var mongoSession *mgo.Session

func ConnectMongoWithMGO() {
	rootCerts := x509.NewCertPool()

	ca, err := os.ReadFile("/Users/harshpatel/Desktop/certs/ca.crt")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	ok := rootCerts.AppendCertsFromPEM([]byte(ca))
	if !ok {
		panic("failed to parse root certificate")
	}

	clientCerts := []tls.Certificate{}
	cert, err := tls.LoadX509KeyPair("/Users/harshpatel/Desktop/certs/developer.crt", "/Users/harshpatel/Desktop/certs/developer.key")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	clientCerts = append(clientCerts, cert)

	tlsConfig := &tls.Config{
		RootCAs:            rootCerts,
		Certificates:       clientCerts,
		InsecureSkipVerify: true,
	}

	dialInfo, err := mgo.ParseURL(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	dialInfo.Direct = true

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Println(err)
		}
		return conn, err
	}

	mongoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Connected to MongoDB")
}

func GetMongoSession() *mgo.Session {
	return mongoSession.Clone()
}
