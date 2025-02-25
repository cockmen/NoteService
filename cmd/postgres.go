package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type DB struct {
	Name     string `yaml:"POSTGRES_DB"`
	User     string `yaml:"POSTGRES_USER"`
	Password string `yaml:"POSTGRES_PASSWORD"`
	Port     string `yaml:"PORT"`
	Host     string `yaml:"HOST"`
}

func PostgresConnection() (*sql.DB, error) {
	config := getDBconfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDBconfig() DB {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {

		log.Fatalf("Ошибка чтения конфига:%v", err)
	}
	var config DB

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Ошибка разбора конфига:%v", err)
	}
	return config
}
