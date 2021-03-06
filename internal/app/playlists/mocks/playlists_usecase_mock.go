// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_playlists is a generated GoMock package.
package mock_playlists

import (
	models "2021_1_Noskool_team/internal/app/playlists/models"
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

// AddPlaylistToMediateka mocks base method.
func (m *MockUsecase) AddPlaylistToMediateka(userID, playlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaylistToMediateka", userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlaylistToMediateka indicates an expected call of AddPlaylistToMediateka.
func (mr *MockUsecaseMockRecorder) AddPlaylistToMediateka(userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaylistToMediateka", reflect.TypeOf((*MockUsecase)(nil).AddPlaylistToMediateka), userID, playlistID)
}

// AddTrackToPlaylist mocks base method.
func (m *MockUsecase) AddTrackToPlaylist(playlistID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrackToPlaylist", playlistID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrackToPlaylist indicates an expected call of AddTrackToPlaylist.
func (mr *MockUsecaseMockRecorder) AddTrackToPlaylist(playlistID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrackToPlaylist", reflect.TypeOf((*MockUsecase)(nil).AddTrackToPlaylist), playlistID, trackID)
}

// CreatePlaylist mocks base method.
func (m *MockUsecase) CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", playlist)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockUsecaseMockRecorder) CreatePlaylist(playlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockUsecase)(nil).CreatePlaylist), playlist)
}

// DeletePlaylistFromUser mocks base method.
func (m *MockUsecase) DeletePlaylistFromUser(userID, playlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylistFromUser", userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylistFromUser indicates an expected call of DeletePlaylistFromUser.
func (mr *MockUsecaseMockRecorder) DeletePlaylistFromUser(userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylistFromUser", reflect.TypeOf((*MockUsecase)(nil).DeletePlaylistFromUser), userID, playlistID)
}

// DeleteTrackFromPlaylist mocks base method.
func (m *MockUsecase) DeleteTrackFromPlaylist(playlistID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackFromPlaylist", playlistID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackFromPlaylist indicates an expected call of DeleteTrackFromPlaylist.
func (mr *MockUsecaseMockRecorder) DeleteTrackFromPlaylist(playlistID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackFromPlaylist", reflect.TypeOf((*MockUsecase)(nil).DeleteTrackFromPlaylist), playlistID, trackID)
}

// GetMediateka mocks base method.
func (m *MockUsecase) GetMediateka(userID int) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediateka", userID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMediateka indicates an expected call of GetMediateka.
func (mr *MockUsecaseMockRecorder) GetMediateka(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediateka", reflect.TypeOf((*MockUsecase)(nil).GetMediateka), userID)
}

// GetPlaylistByID mocks base method.
func (m *MockUsecase) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylistByID", playlistID)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylistByID indicates an expected call of GetPlaylistByID.
func (mr *MockUsecaseMockRecorder) GetPlaylistByID(playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylistByID", reflect.TypeOf((*MockUsecase)(nil).GetPlaylistByID), playlistID)
}

// GetPlaylistByUID mocks base method.
func (m *MockUsecase) GetPlaylistByUID(UID string) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylistByUID", UID)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylistByUID indicates an expected call of GetPlaylistByUID.
func (mr *MockUsecaseMockRecorder) GetPlaylistByUID(UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylistByUID", reflect.TypeOf((*MockUsecase)(nil).GetPlaylistByUID), UID)
}

// GetPlaylists mocks base method.
func (m *MockUsecase) GetPlaylists() ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylists")
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylists indicates an expected call of GetPlaylists.
func (mr *MockUsecaseMockRecorder) GetPlaylists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylists", reflect.TypeOf((*MockUsecase)(nil).GetPlaylists))
}

// GetPlaylistsByGenreID mocks base method.
func (m *MockUsecase) GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylistsByGenreID", genreID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylistsByGenreID indicates an expected call of GetPlaylistsByGenreID.
func (mr *MockUsecaseMockRecorder) GetPlaylistsByGenreID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylistsByGenreID", reflect.TypeOf((*MockUsecase)(nil).GetPlaylistsByGenreID), genreID)
}

// UpdatePlaylistDescription mocks base method.
func (m *MockUsecase) UpdatePlaylistDescription(playlist *models.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylistDescription", playlist)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlaylistDescription indicates an expected call of UpdatePlaylistDescription.
func (mr *MockUsecaseMockRecorder) UpdatePlaylistDescription(playlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylistDescription", reflect.TypeOf((*MockUsecase)(nil).UpdatePlaylistDescription), playlist)
}

// UpdatePlaylistTittle mocks base method.
func (m *MockUsecase) UpdatePlaylistTittle(playlist *models.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylistTittle", playlist)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlaylistTittle indicates an expected call of UpdatePlaylistTittle.
func (mr *MockUsecaseMockRecorder) UpdatePlaylistTittle(playlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylistTittle", reflect.TypeOf((*MockUsecase)(nil).UpdatePlaylistTittle), playlist)
}

// UploadPicture mocks base method.
func (m *MockUsecase) UploadPicture(playlistID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPicture", playlistID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPicture indicates an expected call of UploadPicture.
func (mr *MockUsecaseMockRecorder) UploadPicture(playlistID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPicture", reflect.TypeOf((*MockUsecase)(nil).UploadPicture), playlistID, audioPath)
}
