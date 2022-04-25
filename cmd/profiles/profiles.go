package main

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/profiles/delivery/grpc"
	profiles "2021_1_Noskool_team/internal/app/profiles/delivery/http"
	"2021_1_Noskool_team/internal/app/profiles/repository/postgresDB"
	"2021_1_Noskool_team/internal/app/profiles/usecase"
	"2021_1_Noskool_team/internal/pkg/utility"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	//time.Sleep(40 * time.Second)
	flag.Parse()

	local := true
	config := configs.NewConfig(local)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	profDBCon, err := utility.CreatePostgresConnection(config.ProfileDB)
	if err != nil {
		logrus.Error(err)
	}
	profRep := postgresDB.NewProfileRepository(profDBCon)
	profUsecase := usecase.NewProfilesUsecase(profRep)
	sanitizer := bluemonday.UGCPolicy()

	go grpc.StartSessionsGRPCServer("127.0.0.1:7777", config, profUsecase)

	s := profiles.New(config, profUsecase, sanitizer)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
