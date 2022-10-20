package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	albumHttp "2021_1_Noskool_team/internal/app/album/delivery/http"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	musicHttp "2021_1_Noskool_team/internal/app/musicians/delivery/http"
	"2021_1_Noskool_team/internal/app/playlists"
	playlistHttp "2021_1_Noskool_team/internal/app/playlists/delivery/http"
	"2021_1_Noskool_team/internal/app/search"
	searchHttp "2021_1_Noskool_team/internal/app/search/delivery/http"
	"2021_1_Noskool_team/internal/app/tracks"
	trackHttp "2021_1_Noskool_team/internal/app/tracks/delivery/http"
	"2021_1_Noskool_team/internal/clients/s3_music"
	"2021_1_Noskool_team/internal/pkg/monitoring"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
)

type MusicHandler struct {
	router          *mux.Router
	tracksHandler   *trackHttp.TracksHandler
	musicianHandler *musicHttp.MusiciansHandler
	albumsHandler   *albumHttp.AlbumsHandler
	playlistHandler *playlistHttp.PlaylistsHandler
	searchHandler   *searchHttp.SearchHandler
	s3Client        s3_music.S3MusicBucket
}

func (h MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewFinalHandler(config *configs.Config,
	tracksUsecase tracks.Usecase,
	musicUsecase musicians.Usecase,
	albumsUsecase album.Usecase,
	playlistUsecase playlists.Usecase,
	searchUsecase search.Usecase,
	s3Client s3_music.S3MusicBucket,
) *MusicHandler {
	handler := &MusicHandler{
		router:   mux.NewRouter(),
		s3Client: s3Client,
	}

	metricks := monitoring.RegisterMetrics(handler.router)

	logrus.Info(config.MediaFolder)

	handler.router.PathPrefix("/api/v1/music/data/").
		Handler(
			http.StripPrefix(
				"/api/v1/music/data/", http.FileServer(http.Dir(config.MediaFolder))))

	sanitizer := bluemonday.UGCPolicy()

	handler.router.Use(middleware.LoggingMiddleware(metricks))

	//handler.router.HandleFunc("/api/v1/music/data/{trackUrl}",
	//	handler.GetFiles).Methods(http.MethodGet, http.MethodOptions)

	musicRouter := handler.router.PathPrefix("/api/v1/music/musician/").Subrouter()
	tracksRouter := handler.router.PathPrefix("/api/v1/music/track/").Subrouter()
	albumsRouter := handler.router.PathPrefix("/api/v1/music/album/").Subrouter()
	searchRouter := handler.router.PathPrefix("/api/v1/music/search/").Subrouter()
	playlistsRouter := handler.router.PathPrefix("/api/v1/music/playlist/").Subrouter()
	handler.musicianHandler = musicHttp.NewMusicHandler(musicRouter, config, musicUsecase)
	handler.tracksHandler = trackHttp.NewTracksHandler(tracksRouter, config, tracksUsecase)
	handler.albumsHandler = albumHttp.NewAlbumsHandler(albumsRouter, config, albumsUsecase,
		tracksUsecase, musicUsecase)
	handler.searchHandler = searchHttp.NewSearchHandler(searchRouter, config, searchUsecase, sanitizer)
	handler.playlistHandler = playlistHttp.NewPlaylistsHandler(playlistsRouter, config, playlistUsecase,
		albumsUsecase)

	handler.router.HandleFunc("/api/v1/music/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("main main page"))
	})

	CORSMiddleware := middleware.NewCORSMiddleware(config)
	handler.router.Use(CORSMiddleware.CORS)
	handler.router.Use(middleware.PanicMiddleware(metricks))
	return handler
}

// TODO Протестить с фронтом
func (h *MusicHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	trackUrl := mux.Vars(r)["trackUrl"]

	trackFile, err := h.s3Client.GetObject(r.Context(), s3_music.MusicBucket, trackUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(trackFile)
}
