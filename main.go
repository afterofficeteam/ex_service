package main

import (
	"context"
	"database/sql"

	usecases "ex_service/src/app/usecases"
	userUC "ex_service/src/app/usecases/user"
	"ex_service/src/infra/config"
	oauthGoogleInteg "ex_service/src/infra/integration/oauthgoogle"
	ms_log "ex_service/src/infra/log"
	postgres "ex_service/src/infra/persistence/postgres"
	userRepo "ex_service/src/infra/persistence/postgres/user"
	"ex_service/src/interface/rest"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	conf := config.Make()

	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)

	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	userRepository := userRepo.NewUserRepository(postgresdb.Conn)
	OauthGoogleIntegration := oauthGoogleInteg.NewOauthGoogleService()

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			UserUC: userUC.NewUserUseCase(userRepository, OauthGoogleIntegration),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
