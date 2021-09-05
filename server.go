package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alexandebruce/test-server-golang/db/logging"
	"github.com/Alexandebruce/test-server-golang/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Application struct {
	//pgDb       *sql.DB
	server          *http.Server
	mongoLog        *mongo.Client
}

func (a *Application) Run() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()
}

func createApp(postgresConnString string, mongoConnString string) *Application {
	/*pgDb, err := dbConnection(postgresConnString, "postgres")
	if err != nil {
		panic(err)
	}*/

	mongoLog, err := mongoConnection(mongoConnString)
	if err != nil {
		panic(err)
	}

	logRepository := logging.NewLogRepository(mongoLog, "PetLog", "AllMessages")

	router := mux.NewRouter()
	router.Handle("/ping", handlers.PingHandler(logRepository)).Methods(http.MethodPost)
	//router.Handle("/carriage", handlers.ServiceHandler(carriageRepo)).Methods(http.MethodPost)
	//router.Handle("/search", authorize.NewAuthHandler(handlers.SearchHandler(searchService))).
		//Methods(http.MethodPost)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	log.Printf("HTTP server started on %d\n", 8080)
	return &Application{
		//pgDb:       pgDb,
		server:          &server,
		mongoLog:        mongoLog,
	}
}

func dbConnection(connection, driverName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, connection)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Can't connect to DB: %s", err))
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("Can't ping DB: %s", err))
	}
	return db, nil
}

func mongoConnection(connection string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {
		return &mongo.Client{}, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
