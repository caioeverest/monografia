package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"time"
)

type dbConfig struct {
	Host 	string 	`json:"host"`
	User	string 	`json:"user"`
	Pass 	string 	`json:"pass"`
	DbName 	string	`json:"db_name"`
}

func getConfig() *dbConfig {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	return &dbConfig{
		Host: fmt.Sprintf("%s:%s", host, port),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}

func Conectar() *mgo.Database {
	connection, err := prepareConection()
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func prepareConection() (*mgo.Database, error) {
	confEnv := getConfig()
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{confEnv.Host},
		//ReplicaSetName: "replicasetkey123",
		Timeout:  time.Minute,
		Database: confEnv.DbName,
		Username: confEnv.User,
		Password: confEnv.Pass,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return session.DB(confEnv.DbName), nil
}

func ConnectionTest() error {
	if _, err := prepareConection(); err != nil {
		return err
	}
	return nil
}