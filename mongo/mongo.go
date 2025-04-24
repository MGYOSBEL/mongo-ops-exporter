package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBClient struct {
	Name   string
	Client *mongo.Client
}

type CurrentOperations struct {
	Inprog []Operation `bson:"inprog"`
	Ok     int         `bson:"ok"`
}

type Operation struct {
	Host             string         `bson:"host,omitempty"`
	Client           string         `bson:"client,omitempty"`
	AppName          string         `bson:"appName,omitempty"`
	Op               string         `bson:"op,omitempty"`
	Command          map[string]any `bson:"command,omitempty"`
	MicrosecsRunning int32          `bson:"microsecs_running,omitempty"`
}

func NewClient(uri, name string) (*DBClient, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &DBClient{Name: name, Client: client}, nil
}

func (db *DBClient) GetSlowOps(slowMS int, limit int) ([]Operation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentOp := db.Client.Database("admin").RunCommand(ctx, bson.D{
		{"currentOp", 1},
		{"$all", true},
	})

	var ops CurrentOperations
	if err := currentOp.Decode(&ops); err != nil {
		return nil, err
	}

	if len(ops.Inprog) == 0 {
		return nil, nil
	}

	var slowOps []Operation
	for _, op := range ops.Inprog {
		if op.MicrosecsRunning/1000 >= int32(slowMS) {
			slowOps = append(slowOps, op)
		}
		if len(slowOps) >= limit {
			break
		}
	}
	return slowOps, nil
}
