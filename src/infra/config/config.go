package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	Envs *Config // Envs is global vars Config.
	once sync.Once
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type LogConf struct {
	Name string
}

type SqlDbConf struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
}

type RedisConf struct {
	Host string
	Port string
}

type MongoDbConf struct {
	Dsn    string
	DbName string
}

type OauthConf struct {
	Google OauthGoogleConf
}

type OauthGoogleConf struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// Config ...
type Config struct {
	App   AppConf
	Http  HttpConf
	Log   LogConf
	SqlDb SqlDbConf
	Redis RedisConf
	Oauth OauthConf
}

// NewConfig ...
func Make() Config {
	logrus.Info("loading config...")
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	sqldb := SqlDbConf{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	redis := RedisConf{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	logs := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	oauth := OauthConf{
		Google: OauthGoogleConf{
			ClientID:     os.Getenv("OAUTH_GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_GOOGLE_CLIENT_SECRET"),
			CallbackURL:  os.Getenv("OAUTH_GOOGLE_CALLBACK_URL"),
		},
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	dBMaxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err == nil {
		sqldb.MaxOpenConn = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err == nil {
		sqldb.MaxIdleConn = dBMaxIdleConn
	}

	dBMaxIdleTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxIdleTimeConnSeconds = int64(dBMaxIdleTimeConnSeconds)
	}

	dBMaxLifeTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxLifeTimeConnSeconds = int64(dBMaxLifeTimeConnSeconds)
	}

	config := Config{
		App:   app,
		Http:  http,
		Log:   logs,
		SqlDb: sqldb,
		Redis: redis,
		Oauth: oauth,
	}

	once.Do(func() {
		logrus.Info("config loaded")
		Envs = &config
	})

	return config
}
