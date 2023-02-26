package repository

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

const (
	petTypeTable = "pet_type"
	petCardTable = "pet_card"
	breedTable   = "breed"
)

type Config struct {
	DebugMode bool
	DB        *DB
	Listen    *Listen  `yaml:"listen"`
	Gateway   *Gateway `yaml:"gateway"`
	VersionDB int      `yaml:"db_version"`
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Listen struct {
	Port string `yaml:"port"`
	IP   string `yaml:"ip"`
}

type Gateway struct {
	Port  string `yaml:"port"`
	IP    string `yaml:"ip"`
	Label string `yaml:"label"`
}

type Service struct {
	Name      string `json:"name"`
	Port      string `json:"port"`
	IP        string `json:"ip"`
	Label     string `json:"label"`
	Endpoints []struct {
		URL       string   `json:"url"`
		Protected bool     `json:"protected"`
		Methods   []string `json:"methods"`
	} `json:"endpoints"`
}

func NewPostgresDB(cfg DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetConfig(debugMode bool) *Config {

	logger := log.Default()
	logger.Print("Read application configuration")
	instance := &Config{DB: &DB{}, DebugMode: debugMode}
	if err := cleanenv.ReadConfig("./conf/config.yml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		logger.Print(help)
		logger.Fatal(err)
	}
	if debugMode {
		dbConfigName := "DBConfig"
		if err := cleanenv.ReadConfig(fmt.Sprintf("./conf/db/%s.yml", dbConfigName), instance.DB); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Print(help)
			logger.Fatal(err)
		}
	} else {
		instance.DB = &DB{
			Host:     getEnv("POSTGRES_HOST", ""),
			Port:     getEnv("POSTGRES_PORT", ""),
			DBName:   getEnv("POSTGRES_DB_NAME", ""),
			Username: getEnv("POSTGRES_USER", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
		}
	}

	return instance
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
