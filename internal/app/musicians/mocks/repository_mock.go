// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_musicians is a generated GoMock package.
package mock_musicians

import (
	models "2021_1_Noskool_team/internal/app/musicians/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddMusicianToFavorites mocks base method.
func (m *MockRepository) AddMusicianToFavorites(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMusicianToFavorites", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMusicianToFavorites indicates an expected call of AddMusicianToFavorites.
func (mr *MockRepositoryMockRecorder) AddMusicianToFavorites(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMusicianToFavorites", reflect.TypeOf((*MockRepository)(nil).AddMusicianToFavorites), userID, musicianID)
}

// AddMusicianToMediateka mocks base method.
func (m *MockRepository) AddMusicianToMediateka(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMusicianToMediateka", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMusicianToMediateka indicates an expected call of AddMusicianToMediateka.
func (mr *MockRepositoryMockRecorder) AddMusicianToMediateka(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMusicianToMediateka", reflect.TypeOf((*MockRepository)(nil).AddMusicianToMediateka), userID, musicianID)
}

// CheckMusicianInFavorite mocks base method.
func (m *MockRepository) CheckMusicianInFavorite(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMusicianInFavorite", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckMusicianInFavorite indicates an expected call of CheckMusicianInFavorite.
func (mr *MockRepositoryMockRecorder) CheckMusicianInFavorite(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMusicianInFavorite", reflect.TypeOf((*MockRepository)(nil).CheckMusicianInFavorite), userID, musicianID)
}

// CheckMusicianInMediateka mocks base method.
func (m *MockRepository) CheckMusicianInMediateka(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMusicianInMediateka", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckMusicianInMediateka indicates an expected call of CheckMusicianInMediateka.
func (mr *MockRepositoryMockRecorder) CheckMusicianInMediateka(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMusicianInMediateka", reflect.TypeOf((*MockRepository)(nil).CheckMusicianInMediateka), userID, musicianID)
}

// DeleteMusicianFromFavorites mocks base method.
func (m *MockRepository) DeleteMusicianFromFavorites(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMusicianFromFavorites", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMusicianFromFavorites indicates an expected call of DeleteMusicianFromFavorites.
func (mr *MockRepositoryMockRecorder) DeleteMusicianFromFavorites(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMusicianFromFavorites", reflect.TypeOf((*MockRepository)(nil).DeleteMusicianFromFavorites), userID, musicianID)
}

// DeleteMusicianFromMediateka mocks base method.
func (m *MockRepository) DeleteMusicianFromMediateka(userID, musicianID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMusicianFromMediateka", userID, musicianID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMusicianFromMediateka indicates an expected call of DeleteMusicianFromMediateka.
func (mr *MockRepositoryMockRecorder) DeleteMusicianFromMediateka(userID, musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMusicianFromMediateka", reflect.TypeOf((*MockRepository)(nil).DeleteMusicianFromMediateka), userID, musicianID)
}

// GetGenreForMusician mocks base method.
func (m *MockRepository) GetGenreForMusician(nameMusician string) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenreForMusician", nameMusician)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenreForMusician indicates an expected call of GetGenreForMusician.
func (mr *MockRepositoryMockRecorder) GetGenreForMusician(nameMusician interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenreForMusician", reflect.TypeOf((*MockRepository)(nil).GetGenreForMusician), nameMusician)
}

// GetMusicianByAlbumID mocks base method.
func (m *MockRepository) GetMusicianByAlbumID(albumID int) (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicianByAlbumID", albumID)
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicianByAlbumID indicates an expected call of GetMusicianByAlbumID.
func (mr *MockRepositoryMockRecorder) GetMusicianByAlbumID(albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicianByAlbumID", reflect.TypeOf((*MockRepository)(nil).GetMusicianByAlbumID), albumID)
}

// GetMusicianByID mocks base method.
func (m *MockRepository) GetMusicianByID(musicianID int) (*models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicianByID", musicianID)
	ret0, _ := ret[0].(*models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicianByID indicates an expected call of GetMusicianByID.
func (mr *MockRepositoryMockRecorder) GetMusicianByID(musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicianByID", reflect.TypeOf((*MockRepository)(nil).GetMusicianByID), musicianID)
}

// GetMusicianByPlaylistID mocks base method.
func (m *MockRepository) GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicianByPlaylistID", playlistID)
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicianByPlaylistID indicates an expected call of GetMusicianByPlaylistID.
func (mr *MockRepositoryMockRecorder) GetMusicianByPlaylistID(playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicianByPlaylistID", reflect.TypeOf((*MockRepository)(nil).GetMusicianByPlaylistID), playlistID)
}

// GetMusicianByTrackID mocks base method.
func (m *MockRepository) GetMusicianByTrackID(trackID int) (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicianByTrackID", trackID)
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicianByTrackID indicates an expected call of GetMusicianByTrackID.
func (mr *MockRepositoryMockRecorder) GetMusicianByTrackID(trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicianByTrackID", reflect.TypeOf((*MockRepository)(nil).GetMusicianByTrackID), trackID)
}

// GetMusicians mocks base method.
func (m *MockRepository) GetMusicians() (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicians")
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicians indicates an expected call of GetMusicians.
func (mr *MockRepositoryMockRecorder) GetMusicians() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicians", reflect.TypeOf((*MockRepository)(nil).GetMusicians))
}

// GetMusiciansByGenre mocks base method.
func (m *MockRepository) GetMusiciansByGenre(genre string) (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusiciansByGenre", genre)
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusiciansByGenre indicates an expected call of GetMusiciansByGenre.
func (mr *MockRepositoryMockRecorder) GetMusiciansByGenre(genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusiciansByGenre", reflect.TypeOf((*MockRepository)(nil).GetMusiciansByGenre), genre)
}

// GetMusiciansFavorites mocks base method.
func (m *MockRepository) GetMusiciansFavorites(userID int) ([]*models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusiciansFavorites", userID)
	ret0, _ := ret[0].([]*models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusiciansFavorites indicates an expected call of GetMusiciansFavorites.
func (mr *MockRepositoryMockRecorder) GetMusiciansFavorites(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusiciansFavorites", reflect.TypeOf((*MockRepository)(nil).GetMusiciansFavorites), userID)
}

// GetMusiciansMediateka mocks base method.
func (m *MockRepository) GetMusiciansMediateka(userID int) ([]*models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusiciansMediateka", userID)
	ret0, _ := ret[0].([]*models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusiciansMediateka indicates an expected call of GetMusiciansMediateka.
func (mr *MockRepositoryMockRecorder) GetMusiciansMediateka(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusiciansMediateka", reflect.TypeOf((*MockRepository)(nil).GetMusiciansMediateka), userID)
}

// GetMusiciansTop4 mocks base method.
func (m *MockRepository) GetMusiciansTop4() (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusiciansTop4")
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusiciansTop4 indicates an expected call of GetMusiciansTop4.
func (mr *MockRepositoryMockRecorder) GetMusiciansTop4() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusiciansTop4", reflect.TypeOf((*MockRepository)(nil).GetMusiciansTop4))
}

// SearchMusicians mocks base method.
func (m *MockRepository) SearchMusicians(searchQuery string) ([]*models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMusicians", searchQuery)
	ret0, _ := ret[0].([]*models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMusicians indicates an expected call of SearchMusicians.
func (mr *MockRepositoryMockRecorder) SearchMusicians(searchQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMusicians", reflect.TypeOf((*MockRepository)(nil).SearchMusicians), searchQuery)
}
