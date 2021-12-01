package main

import (
	"context"
	"fmt"
	"github.com/aljorhythm/golang-playground/sockets"
	"github.com/aljorhythm/golang-playground/storage"
	"io/ioutil"
	"log"
	"net/http"
)

func generateUploadFile(store storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("File Upload Endpoint Hit")

		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Server Received File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := ioutil.ReadAll(file)
		store.Store(context.Background(), "file1", fileBytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}

func main() {
	pingServer := sockets.NewPingServer()
	httpServer := http.NewServeMux()
	config := readConfig()
	log.Printf("config %#v", config)

	storage := getStorage(config)

	httpServer.HandleFunc("/upload", generateUploadFile(storage))
	httpServer.HandleFunc("/open", generateDownloadFile(storage))

	httpServer.Handle("/", pingServer.Handler)

	log.Println("Running server")
	if err := http.ListenAndServe(":8080", httpServer); err != nil {
		log.Fatalf("%#v", err)
	}
}

func getStorage(config Config) storage.Storage {
	if config.DigitalOceanStore != nil {
		log.Printf("loading digital ocean store")
		digitalOceanConfig := config.DigitalOceanStore
		bucketProps := storage.BucketProperties{
			Name:     digitalOceanConfig.Space,
			Location: "",
		}

		endpoint := digitalOceanConfig.Client.Endpoint
		accessKey := digitalOceanConfig.Client.AccessKey
		secKey := digitalOceanConfig.Client.SecretKey

		store, err := storage.NewSpaceStore(context.Background(), endpoint, accessKey, secKey, bucketProps)

		if err != nil {
			log.Fatalf("cannot initialize #%v", err)
			return nil
		} else {
			return store
		}
	} else {
		return storage.NewInmemoryStore()
	}
}

func generateDownloadFile(s storage.Storage) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := s.Retrieve(context.Background(), "file1")
		if err != nil {
			log.Println("error upload file %#v", err)
		}
		writer.Write(data)
	}
}
