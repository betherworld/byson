package main

import (
	"strings"

	"fin4-core/server/assetservice"
	"fin4-core/server/datatype"
	"fin4-core/server/dbservice"
	"fin4-core/server/emailer"
	"fin4-core/server/env"
	"fin4-core/server/ethereum"
	"fin4-core/server/filestorage"
	"fin4-core/server/logger"
	"fin4-core/server/routes"
	"fin4-core/server/timelineservice"
	"fin4-core/server/tokenservice"
	"fin4-core/server/userservice"
)

func main() {
	environment := strings.ToLower(env.MustGetenv("ENVIRONMENT"))

	//TODO this should be changed ASAP
	baseDir := "/home/parallels/Projects/goprojects/src/fin4-core"
	env.Load(baseDir, false)

	cfg := datatype.Config{
		Environment:        environment,
		DBMigrationPath:    env.MustGetenv("DB_MIGRATIONS_PATH"),
		ServerPort:         getPort(),
		DataSourceName:     env.MustGetenv("DATA_SOURCE_NAME"),
		TestDataSourceName: env.Getenv("TEST_DATA_SOURCE_NAME"),
		AwsSesKey:          env.Getenv("AWS_SES_KEY"),
		AwsSesSecret:       env.Getenv("AWS_SES_SECRET"),
		AwsEmailFrom:       env.Getenv("EMAIL_FROM"),
		SlackWebhookURL:    env.Getenv("SLACK_WEBHOOK_URL"),
	}

	logger.Setup(cfg)

	emailer.Setup(cfg)

	// connect to database
	db := dbservice.MustConnect(cfg.DataSourceName)
	userService := userservice.NewService(db)
	tokenService := tokenservice.NewService(db)
	timelineService := timelineservice.NewService(db)
	assetService := assetservice.NewService(db)
	// connect with Amazon AWS (used in our live instance)
	fs := filestorage.GetStorage(
		env.Getenv("AWS_SES_KEY"),
		env.Getenv("AWS_SES_SECRET"),
	)
	// connect with an instance of ethereum
	ethereum := ethereum.MustNewEthereum()
	sc := datatype.ServiceContainer{
		Config:          cfg,
		AssetService:    assetService,
		TimelineService: timelineService,
		UserService:     userService,
		TokenService:    tokenService,
		FileStorage:     fs,
		Ethereum:        ethereum,
	}
	// start listening to api calls
	routes.SetupRouting(sc).Run(":" + cfg.ServerPort)
}

func getPort() string {
	port := env.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
