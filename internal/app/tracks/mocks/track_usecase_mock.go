// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_tracks is a generated GoMock package.
package mock_tracks

import (
	models "2021_1_Noskool_team/internal/app/tracks/models"
	models0 "2021_1_Noskool_team/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// AddDeleteTrackToMediateka mocks base method.
func (m *MockUsecase) AddDeleteTrackToMediateka(userID, trackID int, operationType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeleteTrackToMediateka", userID, trackID, operationType)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDeleteTrackToMediateka indicates an expected call of AddDeleteTrackToMediateka.
func (mr *MockUsecaseMockRecorder) AddDeleteTrackToMediateka(userID, trackID, operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeleteTrackToMediateka", reflect.TypeOf((*MockUsecase)(nil).AddDeleteTrackToMediateka), userID, trackID, operationType)
}

// AddToHistory mocks base method.
func (m *MockUsecase) AddToHistory(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToHistory", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToHistory indicates an expected call of AddToHistory.
func (mr *MockUsecaseMockRecorder) AddToHistory(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToHistory", reflect.TypeOf((*MockUsecase)(nil).AddToHistory), userID, trackID)
}

// AddTrackToFavorites mocks base method.
func (m *MockUsecase) AddTrackToFavorites(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrackToFavorites", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrackToFavorites indicates an expected call of AddTrackToFavorites.
func (mr *MockUsecaseMockRecorder) AddTrackToFavorites(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrackToFavorites", reflect.TypeOf((*MockUsecase)(nil).AddTrackToFavorites), userID, trackID)
}

// CheckTrackInFavorite mocks base method.
func (m *MockUsecase) CheckTrackInFavorite(userID, trackID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTrackInFavorite", userID, trackID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckTrackInFavorite indicates an expected call of CheckTrackInFavorite.
func (mr *MockUsecaseMockRecorder) CheckTrackInFavorite(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTrackInFavorite", reflect.TypeOf((*MockUsecase)(nil).CheckTrackInFavorite), userID, trackID)
}

// CheckTrackInMediateka mocks base method.
func (m *MockUsecase) CheckTrackInMediateka(userID, trackID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTrackInMediateka", userID, trackID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckTrackInMediateka indicates an expected call of CheckTrackInMediateka.
func (mr *MockUsecaseMockRecorder) CheckTrackInMediateka(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTrackInMediateka", reflect.TypeOf((*MockUsecase)(nil).CheckTrackInMediateka), userID, trackID)
}

// DeleteTrackFromFavorites mocks base method.
func (m *MockUsecase) DeleteTrackFromFavorites(trackID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackFromFavorites", trackID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackFromFavorites indicates an expected call of DeleteTrackFromFavorites.
func (mr *MockUsecaseMockRecorder) DeleteTrackFromFavorites(trackID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackFromFavorites", reflect.TypeOf((*MockUsecase)(nil).DeleteTrackFromFavorites), trackID, userID)
}

// GetBillbordTopCharts mocks base method.
func (m *MockUsecase) GetBillbordTopCharts() ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillbordTopCharts")
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillbordTopCharts indicates an expected call of GetBillbordTopCharts.
func (mr *MockUsecaseMockRecorder) GetBillbordTopCharts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillbordTopCharts", reflect.TypeOf((*MockUsecase)(nil).GetBillbordTopCharts))
}

// GetFavoriteTracks mocks base method.
func (m *MockUsecase) GetFavoriteTracks(userID int, pagination *models0.Pagination) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteTracks", userID, pagination)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteTracks indicates an expected call of GetFavoriteTracks.
func (mr *MockUsecaseMockRecorder) GetFavoriteTracks(userID, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteTracks", reflect.TypeOf((*MockUsecase)(nil).GetFavoriteTracks), userID, pagination)
}

// GetHistory mocks base method.
func (m *MockUsecase) GetHistory(userID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", userID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockUsecaseMockRecorder) GetHistory(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockUsecase)(nil).GetHistory), userID)
}

// GetTop20Tracks mocks base method.
func (m *MockUsecase) GetTop20Tracks() ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTop20Tracks")
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTop20Tracks indicates an expected call of GetTop20Tracks.
func (mr *MockUsecaseMockRecorder) GetTop20Tracks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTop20Tracks", reflect.TypeOf((*MockUsecase)(nil).GetTop20Tracks))
}

// GetTopTrack mocks base method.
func (m *MockUsecase) GetTopTrack() ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopTrack")
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopTrack indicates an expected call of GetTopTrack.
func (mr *MockUsecaseMockRecorder) GetTopTrack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopTrack", reflect.TypeOf((*MockUsecase)(nil).GetTopTrack))
}

// GetTrackByID mocks base method.
func (m *MockUsecase) GetTrackByID(trackID int) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrackByID", trackID)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrackByID indicates an expected call of GetTrackByID.
func (mr *MockUsecaseMockRecorder) GetTrackByID(trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrackByID", reflect.TypeOf((*MockUsecase)(nil).GetTrackByID), trackID)
}

// GetTrackByMusicianID mocks base method.
func (m *MockUsecase) GetTrackByMusicianID(musicianID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrackByMusicianID", musicianID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrackByMusicianID indicates an expected call of GetTrackByMusicianID.
func (mr *MockUsecaseMockRecorder) GetTrackByMusicianID(musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrackByMusicianID", reflect.TypeOf((*MockUsecase)(nil).GetTrackByMusicianID), musicianID)
}

// GetTracksByAlbumID mocks base method.
func (m *MockUsecase) GetTracksByAlbumID(albumID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByAlbumID", albumID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByAlbumID indicates an expected call of GetTracksByAlbumID.
func (mr *MockUsecaseMockRecorder) GetTracksByAlbumID(albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByAlbumID", reflect.TypeOf((*MockUsecase)(nil).GetTracksByAlbumID), albumID)
}

// GetTracksByGenreID mocks base method.
func (m *MockUsecase) GetTracksByGenreID(genreID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByGenreID", genreID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByGenreID indicates an expected call of GetTracksByGenreID.
func (mr *MockUsecaseMockRecorder) GetTracksByGenreID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByGenreID", reflect.TypeOf((*MockUsecase)(nil).GetTracksByGenreID), genreID)
}

// GetTracksByTittle mocks base method.
func (m *MockUsecase) GetTracksByTittle(trackTittle string) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByTittle", trackTittle)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByTittle indicates an expected call of GetTracksByTittle.
func (mr *MockUsecaseMockRecorder) GetTracksByTittle(trackTittle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByTittle", reflect.TypeOf((*MockUsecase)(nil).GetTracksByTittle), trackTittle)
}

// GetTracksByUserID mocks base method.
func (m *MockUsecase) GetTracksByUserID(userID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByUserID", userID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByUserID indicates an expected call of GetTracksByUserID.
func (mr *MockUsecaseMockRecorder) GetTracksByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByUserID", reflect.TypeOf((*MockUsecase)(nil).GetTracksByUserID), userID)
}

// SearchTracks mocks base method.
func (m *MockUsecase) SearchTracks(searchQuery string) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTracks", searchQuery)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTracks indicates an expected call of SearchTracks.
func (mr *MockUsecaseMockRecorder) SearchTracks(searchQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTracks", reflect.TypeOf((*MockUsecase)(nil).SearchTracks), searchQuery)
}

// UploadAudio mocks base method.
func (m *MockUsecase) UploadAudio(trackID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAudio", trackID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAudio indicates an expected call of UploadAudio.
func (mr *MockUsecaseMockRecorder) UploadAudio(trackID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAudio", reflect.TypeOf((*MockUsecase)(nil).UploadAudio), trackID, audioPath)
}

// UploadPicture mocks base method.
func (m *MockUsecase) UploadPicture(trackID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPicture", trackID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPicture indicates an expected call of UploadPicture.
func (mr *MockUsecaseMockRecorder) UploadPicture(trackID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPicture", reflect.TypeOf((*MockUsecase)(nil).UploadPicture), trackID, audioPath)
}