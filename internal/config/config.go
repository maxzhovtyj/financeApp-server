package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	httpHostEnv      = "HTTP_HOST"
	mongoURIEnv      = "MONGO_URI"
	mongoUserEnv     = "MONGO_USER"
	mongoPasswordEnv = "MONGO_PASS"
	passwordSaltEnv  = "PASSWORD_SALT"
	jwtSigningKeyEnv = "JWT_SIGNING_KEY"
)

type (
	Config struct {
		Mongo MongoConfig
		HTTP  HTTPConfig
		Auth  AuthConfig
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}

	MongoConfig struct {
		URI      string
		User     string
		Password string
		Database string `mapstructure:"databaseName"`
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}
)

func Init() (*Config, error) {
	var cfg Config

	log.Info("loading configuration file...")
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	log.Info("initialize config from env file...")
	setFromEnv(&cfg)

	log.Info("unmarshalling config file...")
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv(httpHostEnv)

	cfg.Auth.PasswordSalt = os.Getenv(passwordSaltEnv)
	cfg.Auth.JWT.SigningKey = os.Getenv(jwtSigningKeyEnv)

	cfg.Mongo.URI = os.Getenv(mongoURIEnv)
	cfg.Mongo.User = os.Getenv(mongoUserEnv)
	cfg.Mongo.Password = os.Getenv(mongoPasswordEnv)
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	return nil
}

func parseConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return nil
}
