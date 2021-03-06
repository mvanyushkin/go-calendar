package main

import (
	"flag"
	server "github.com/mvanyushkin/go-calendar/api"
	"github.com/mvanyushkin/go-calendar/internal"
	"github.com/mvanyushkin/go-calendar/internal/config"
	"github.com/mvanyushkin/go-calendar/internal/store"
	"github.com/mvanyushkin/go-calendar/logger"
	"github.com/mvanyushkin/go-calendar/pkg/calendar"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	configFilePath := flag.String("config", "", "settings file")
	flag.Parse()
	if configFilePath == nil {
		defaultConfigFileName := "local_config.json"
		configFilePath = &defaultConfigFileName
	}

	cfg, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("The config file is broken: %v", err.Error())
	}

	logger.SetupLogger(cfg.LogFile, cfg.LogLevel)
	log.Info("application started.")
	err = serve(cfg.HttpListen, cfg.ConnectionString)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Info("application is shutdown")
}

func serve(binding string, connectionString string) error {
	lis, err := net.Listen("tcp", binding)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	calendar.RegisterCalendarServer(grpcServer, server.CalendarHandler{
		Calendar: internal.NewCalendar(store.NewDatabaseEventStore(connectionString)),
	})
	return grpcServer.Serve(lis)
}
