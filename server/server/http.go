package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HohenzoIIern/APZ234/Lab3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

type StructAddingDisk struct {
	Server_id int `json:"server_id"`
	Disk_id   int `json:"disk_id"`
}

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandlerServer(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListServer(store, rw)
		} else if r.Method == "POST" {
			handleServerCreate(r, rw, store)
		} else if r.Method == "PATCH" {
			handleAddDiskToServer(r, store, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func HttpHandlerDisk(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListDisk(store, rw)
		} else if r.Method == "POST" {
			handleDiskCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleServerCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Server
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateServer(&c)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListServer(store *Store, rw http.ResponseWriter) {
	res, err := store.ListServers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleListDisk(store *Store, rw http.ResponseWriter) {
	res, err := store.ListDisk()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleDiskCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Disk
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateDisk(&c)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleAddDiskToServer(r *http.Request, store *Store, rw http.ResponseWriter) {
	var c StructAddingDisk
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	server, err := store.AddDiskToServer(c.Server_id, c.Disk_id)
	fmt.Println(server)
	if err == nil {
		tools.WriteJsonOk(rw, &server)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
