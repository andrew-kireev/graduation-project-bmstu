package main

import (
	"2021_1_Noskool_team/configs"
	albumRepository "2021_1_Noskool_team/internal/app/album/repository"
	albumUsecase "2021_1_Noskool_team/internal/app/album/usecase"
	musicHttp "2021_1_Noskool_team/internal/app/music/delivery/http"
	musiciansRepository "2021_1_Noskool_team/internal/app/musicians/repository"
	musicianUsecase "2021_1_Noskool_team/internal/app/musicians/usecase"
	playlistRepository "2021_1_Noskool_team/internal/app/playlists/repository"
	playlistUsecase "2021_1_Noskool_team/internal/app/playlists/usecase"
	searchUsecase "2021_1_Noskool_team/internal/app/search/usecase"
	trackRepository "2021_1_Noskool_team/internal/app/tracks/repository"
	trackUsecase "2021_1_Noskool_team/internal/app/tracks/usecase"
	"2021_1_Noskool_team/internal/clients/s3_music"
	"2021_1_Noskool_team/internal/pkg/server"
	"2021_1_Noskool_team/internal/pkg/utility"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

const (
	configPath = "configs/config.toml"
)

func main() {
	//time.Sleep(40 * time.Second)

	local := true
	config := configs.NewConfig(local)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	musicDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	musicianskRep := musiciansRepository.NewMusicRepository(musicDBCon)
	musUsecase := musicianUsecase.NewMusicsUsecase(musicianskRep)

	tracDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	trackRep := trackRepository.NewTracksRepository(tracDBCon)

	trackUse := trackUsecase.NewTracksUsecase(trackRep)

	albumDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	albumRep := albumRepository.NewAlbumsRepository(albumDBCon)

	albumsUse := albumUsecase.NewAlbumcUsecase(albumRep)
	playlistDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	playlistRep := playlistRepository.NewPlaylistRepository(playlistDBCon)
	playlistUse := playlistUsecase.NewPlaylistUsecase(playlistRep)

	client, err := s3_music.New()
	if err != nil {
		logrus.Error(err)
	}

	searhUse := searchUsecase.NewSearchUsecase(trackRep, albumRep, musicianskRep, playlistRep)
	handler := musicHttp.NewFinalHandler(config,
		trackUse,
		musUsecase,
		albumsUse,
		playlistUse,
		searhUse,
		client,
	)

	fmt.Println("Нормально запустились")
	err = server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
	fmt.Println("Закончили работу")
}
