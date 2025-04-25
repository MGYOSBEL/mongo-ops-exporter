package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/MGYOSBEL/mongo-ops-exporter/api"
	"github.com/MGYOSBEL/mongo-ops-exporter/config"
	"github.com/MGYOSBEL/mongo-ops-exporter/mongo"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config.file", "config.yaml", "The file configuration path")
	flag.StringVar(&configFile, "f", "config.yaml", "The file configuration path")
}

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var clients []*mongo.DBClient
	for _, db := range cfg.Databases {
		client, err := mongo.NewClient(db.URI, db.Name)
		if err != nil {
			log.Printf("Skipping DB %s: %v", db.Name, err)
			continue
		}
		clients = append(clients, client)
	}

	api := &api.API{
		Config:  cfg,
		Clients: clients,
		Cache:   make(map[string]interface{}),
	}

	http.HandleFunc("/databases", api.ListDatabases)
	http.HandleFunc("/slowops/{database}", api.GetSlowOpsPerDb)
	http.HandleFunc("/slowops", api.GetSlowOps)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
