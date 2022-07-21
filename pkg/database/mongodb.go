package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConfig Database Config
type DBConfig struct {
	Host string `envconfig:"HOST"`
	Name string `envconfig:"NAME"`
}

type MongoContainer struct {
	testcontainers.Container
	URI string
}

var DBConfigTest = DBConfig{
	Host: "mongodb://localhost:27017/testdatabase",
	Name: "testdatabase",
}

var mongoPort = "27017"

func InitMongoDB(cfg DBConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Host))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return client, nil
}

func TruncateDataForTest(conn *mongo.Database) error {
	var collections = []string{"users"}
	for _, value := range collections {
		err := truncateDataForTest(context.Background(), conn, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func truncateDataForTest(ctx context.Context, conn *mongo.Database, collectionName string) error {
	err := conn.Collection(collectionName).Drop(ctx)
	if err != nil {
		return err
	}
	return nil
}

func StartMongoContainer(ctx context.Context) (*MongoContainer, error) {

	mongoPorts, _ := nat.NewPort("", mongoPort)

	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.4.3",
		ExposedPorts: []string{mongoPort},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(5 * time.Minute),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, mongoPorts)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("mongodb://%s:%s/testdatabase", ip, mappedPort.Port())
	return &MongoContainer{
		Container: container,
		URI:       uri,
	}, nil
}
