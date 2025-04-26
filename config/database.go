package config

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type SQLConfig struct {
	Enabled  bool
	Host     string
	Port     string
	Database string
	Username string
	Password string
}


type MongoConfig struct {
	Enabled  bool
	URI      string
	Database string
}

var (
	
	MySQLConfig    SQLConfig
	PostgresConfig SQLConfig
	MongoConf    MongoConfig

	
	MySQLDB    *gorm.DB
	PostgresDB *gorm.DB
	MongoDB    *mongo.Client
)

func LoadDatabaseConfig() {
	
	drivers := strings.Split(GetEnv("DB_DRIVERS", "mysql"), "+")

	
	MySQLConfig = SQLConfig{
		Enabled:  contains(drivers, "mysql") || contains(drivers, "all"),
		Host:     GetEnv("DB_MYSQL_HOST", "127.0.0.1"),
		Port:     GetEnv("DB_MYSQL_PORT", "3306"),
		Database: GetEnv("DB_MYSQL_DATABASE", "laravel_like"),
		Username: GetEnv("DB_MYSQL_USERNAME", "root"),
		Password: GetEnv("DB_MYSQL_PASSWORD", ""),
	}

	
	PostgresConfig = SQLConfig{
		Enabled:  contains(drivers, "pgsql") || contains(drivers, "all"),
		Host:     GetEnv("DB_PGSQL_HOST", "127.0.0.1"),
		Port:     GetEnv("DB_PGSQL_PORT", "5432"),
		Database: GetEnv("DB_PGSQL_DATABASE", "laravel_like"),
		Username: GetEnv("DB_PGSQL_USERNAME", "postgres"),
		Password: GetEnv("DB_PGSQL_PASSWORD", ""),
	}

	
	MongoConf = MongoConfig{
		Enabled:  contains(drivers, "mongo") || contains(drivers, "all"),
		URI:      GetEnv("DB_MONGO_URI", "mongodb://localhost:27017"),
		Database: GetEnv("DB_MONGO_DATABASE", "laravel_like"),
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ConnectDatabases() {
	var err error

	
	if MySQLConfig.Enabled {
		dsn := MySQLConfig.Username + ":" + MySQLConfig.Password + "@tcp(" + MySQLConfig.Host + ":" + MySQLConfig.Port + ")/" + MySQLConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
		MySQLDB, err = gorm.Open(mysql.Open(dsn), nil)
		if err != nil {
			log.Fatal("Failed to connect to MySQL:", err)
		}
		log.Println("Connected to MySQL database")
	}

	
	if PostgresConfig.Enabled {
		dsn := "host=" + PostgresConfig.Host + " user=" + PostgresConfig.Username + " password=" + PostgresConfig.Password + " dbname=" + PostgresConfig.Database + " port=" + PostgresConfig.Port + " sslmode=disable TimeZone=Asia/Jakarta"
		PostgresDB, err = gorm.Open(postgres.Open(dsn), nil)
		if err != nil {
			log.Fatal("Failed to connect to PostgreSQL:", err)
		}
		log.Println("Connected to PostgreSQL database")
	}

	
	if MongoConf.Enabled { 
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		MongoDB, err = mongo.Connect(ctx, options.Client().ApplyURI(MongoConf.URI)) 
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		err = MongoDB.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}
		log.Println("Connected to MongoDB")
	}

	
	if !MySQLConfig.Enabled && !PostgresConfig.Enabled && !MongoConf.Enabled { 
		log.Fatal("Tidak ada database yang diaktifkan. Silakan setidaknya aktifkan satu driver di DB_DRIVERS")
	}
}


func GetSQLDB() *gorm.DB {
	if MySQLConfig.Enabled {
		return MySQLDB
	}
	return PostgresDB 
}

func GetMongoDB() *mongo.Client {
	return MongoDB
}

func GetMongoDatabase() *mongo.Database {
	return MongoDB.Database(MongoConf.Database)
}

