package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	ServerPort         string   `mapstructure:"PORT"`
	MongoURI           string   `mapstructure:"MONGO_URI"`
	DBName             string   `mapstructure:"DB_NAME"`
	JWTSecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int      `mapstructure:"JWT_EXPIRATION_HOURS"`
	EnableCache        bool     `mapstructure:"ENABLE_CACHE"`
	RedisAddr          string   `mapstructure:"REDIS_ADDR"`
	RedisPassword      string   `mapstructure:"REDIS_PASSWORD"`
	LogLevel           string   `mapstructure:"LOG_LEVEL"`
	LogFormat          string   `mapstructure:"LOG_FORMAT"`
	CookieDomains      []string `mapstructure:"COOKIE_DOMAINS"`
	SecureCookie       bool     `mapstructure:"SECURE_COOKIE"`
	AllowedOrigins     []string `mapstructure:"ALLOWED_ORIGINS"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Ignore if .env doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	// ------------------------------------------------
	// TEMP DEBUG
	// ------------------------------------------------

	config.MongoURI = "mongodb+srv://isedeemmanuel26_db_user:gxS3ceOIiPNaEsz8@cluster0.6kasww3.mongodb.net/much_todo_db?appName=Cluster0"
	config.DBName = "much_todo_db"

	// ------------------------------------------------
	// Read environment variables manually
	// ------------------------------------------------

	if val := os.Getenv("JWT_SECRET_KEY"); val != "" {
		config.JWTSecretKey = val
	}

	if val := os.Getenv("REDIS_ADDR"); val != "" {
		config.RedisAddr = val
	}

	if val := os.Getenv("PORT"); val != "" {
		config.ServerPort = val
	}

	// ******** THIS FIXES YOUR CRASH ********

	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {

		if val == "*" {
			config.AllowedOrigins = []string{"*"}
		} else {
			config.AllowedOrigins = strings.Split(val, ",")
		}

	} else {
		config.AllowedOrigins = []string{"*"}
	}

	return config, nil
}
