package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MGYOSBEL/mongo-ops-exporter/config"
	"github.com/MGYOSBEL/mongo-ops-exporter/mongo"
)

const (
	DEFAULT_SLOWMS = 100
	DEFAULT_LIMIT  = 10
)

type API struct {
	Config  *config.Config
	Clients []*mongo.DBClient
	Cache   map[string]interface{}
}

func (a *API) ListDatabases(w http.ResponseWriter, r *http.Request) {
	var dbs []string
	for _, db := range a.Clients {
		dbs = append(dbs, db.Name)
	}
	json.NewEncoder(w).Encode(dbs)
}

func (a *API) GetSlowOps(w http.ResponseWriter, r *http.Request) {
	slowMS, _ := strconv.Atoi(r.URL.Query().Get("slowms"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if slowMS == 0 {
		slowMS = DEFAULT_SLOWMS
	}
	if limit == 0 {
		limit = DEFAULT_LIMIT
	}

	allOps := make(map[string][]mongo.Operation)
	for _, db := range a.Clients {
		ops, err := db.GetSlowOps(slowMS, limit)
		if err == nil {
			allOps[db.Name] = append(allOps[db.Name], ops...)
		}
	}
	json.NewEncoder(w).Encode(allOps)
}

func (a *API) GetSlowOpsPerDb(w http.ResponseWriter, r *http.Request) {
	slowMS, _ := strconv.Atoi(r.URL.Query().Get("slowms"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if slowMS == 0 {
		slowMS = DEFAULT_SLOWMS
	}
	if limit == 0 {
		limit = DEFAULT_LIMIT
	}

	database := r.PathValue("database")
	allOps := make(map[string][]mongo.Operation)
	for _, db := range a.Clients {
		if db.Name != database {
			continue
		}
		ops, err := db.GetSlowOps(slowMS, limit)
		if err == nil {
			allOps[db.Name] = append(allOps[db.Name], ops...)
		}
	}
	json.NewEncoder(w).Encode(allOps)
}
