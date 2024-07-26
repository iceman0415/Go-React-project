package main

import (
	"net/http"
	"os"

	"github.com/bogdanguranda/go-react-example/api"
	"github.com/bogdanguranda/go-react-example/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	portAPI = "8080"
)

func main() {
	dbMySQL := createDBConnection()
	defer dbMySQL.Close()

	startAPIServer(dbMySQL)
}

func createDBConnection() db.DB {
	mySQLPass := os.Getenv("MYSQL_PASS")
	if mySQLPass == "" {
		logrus.Fatal("Env var MYSQL_PASS was not set!")
	}

	dbMySQL, err := db.NewMySqlDB(mySQLPass)
	if err != nil {
		logrus.Fatal(errors.Wrapf(err, "failed to start MySQL"))
	}

	return dbMySQL
}

func startAPIServer(dbMySQL db.DB) {
	appAPI := api.NewDefaultAPI(dbMySQL)

	router := mux.NewRouter()
	router.HandleFunc("/app/people", appAPI.ListPersons).Methods(http.MethodGet)
	router.HandleFunc("/app/people", appAPI.CreatePerson).Methods(http.MethodPost)
	router.HandleFunc("/app/people", appAPI.DeletePerson).Methods(http.MethodDelete)

	router.HandleFunc("/app/people/{email}", appAPI.GetPerson).Methods(http.MethodGet)
	router.HandleFunc("/app/people/{email}", appAPI.UpdatePerson).Methods(http.MethodPut)

	logrus.Info("REST API server listening on port " + portAPI)

	if err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)); err != nil {
		logrus.Fatal("Failed to listen and serve on port " + portAPI)
	}
}
