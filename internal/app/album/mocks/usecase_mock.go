// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_album is a generated GoMock package.
package mock_album

import (
	models "2021_1_Noskool_team/internal/app/album/models"
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

// AddAlbumToFavorites mocks base method.
func (m *MockUsecase) AddAlbumToFavorites(userID, albumID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlbumToFavorites", userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAlbumToFavorites indicates an expected call of AddAlbumToFavorites.
func (mr *MockUsecaseMockRecorder) AddAlbumToFavorites(userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlbumToFavorites", reflect.TypeOf((*MockUsecase)(nil).AddAlbumToFavorites), userID, albumID)
}

// AddAlbumToMediateka mocks base method.
func (m *MockUsecase) AddAlbumToMediateka(userID, albumID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlbumToMediateka", userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAlbumToMediateka indicates an expected call of AddAlbumToMediateka.
func (mr *MockUsecaseMockRecorder) AddAlbumToMediateka(userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlbumToMediateka", reflect.TypeOf((*MockUsecase)(nil).AddAlbumToMediateka), userID, albumID)
}

// DeleteAlbumFromMediateka mocks base method.
func (m *MockUsecase) DeleteAlbumFromMediateka(userID, albumID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbumFromMediateka", userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlbumFromMediateka indicates an expected call of DeleteAlbumFromMediateka.
func (mr *MockUsecaseMockRecorder) DeleteAlbumFromMediateka(userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbumFromMediateka", reflect.TypeOf((*MockUsecase)(nil).DeleteAlbumFromMediateka), userID, albumID)
}

// DelteAlbumFromFavorites mocks base method.
func (m *MockUsecase) DelteAlbumFromFavorites(userID, albumID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelteAlbumFromFavorites", userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelteAlbumFromFavorites indicates an expected call of DelteAlbumFromFavorites.
func (mr *MockUsecaseMockRecorder) DelteAlbumFromFavorites(userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelteAlbumFromFavorites", reflect.TypeOf((*MockUsecase)(nil).DelteAlbumFromFavorites), userID, albumID)
}

// GetAlbumByID mocks base method.
func (m *MockUsecase) GetAlbumByID(albumID int) (*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumByID", albumID)
	ret0, _ := ret[0].(*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumByID indicates an expected call of GetAlbumByID.
func (mr *MockUsecaseMockRecorder) GetAlbumByID(albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumByID", reflect.TypeOf((*MockUsecase)(nil).GetAlbumByID), albumID)
}

// GetAlbums mocks base method.
func (m *MockUsecase) GetAlbums() ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbums")
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbums indicates an expected call of GetAlbums.
func (mr *MockUsecaseMockRecorder) GetAlbums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbums", reflect.TypeOf((*MockUsecase)(nil).GetAlbums))
}

// GetAlbumsByMusicianID mocks base method.
func (m *MockUsecase) GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumsByMusicianID", musicianID)
	ret0, _ := ret[0].(*[]models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumsByMusicianID indicates an expected call of GetAlbumsByMusicianID.
func (mr *MockUsecaseMockRecorder) GetAlbumsByMusicianID(musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumsByMusicianID", reflect.TypeOf((*MockUsecase)(nil).GetAlbumsByMusicianID), musicianID)
}

// GetAlbumsByTrackID mocks base method.
func (m *MockUsecase) GetAlbumsByTrackID(trackID int) (*[]models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumsByTrackID", trackID)
	ret0, _ := ret[0].(*[]models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumsByTrackID indicates an expected call of GetAlbumsByTrackID.
func (mr *MockUsecaseMockRecorder) GetAlbumsByTrackID(trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumsByTrackID", reflect.TypeOf((*MockUsecase)(nil).GetAlbumsByTrackID), trackID)
}

// GetFavoriteAlbums mocks base method.
func (m *MockUsecase) GetFavoriteAlbums(userID int, pagination *models0.Pagination) ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteAlbums", userID, pagination)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteAlbums indicates an expected call of GetFavoriteAlbums.
func (mr *MockUsecaseMockRecorder) GetFavoriteAlbums(userID, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteAlbums", reflect.TypeOf((*MockUsecase)(nil).GetFavoriteAlbums), userID, pagination)
}

// SearchAlbums mocks base method.
func (m *MockUsecase) SearchAlbums(searchQuery string) ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAlbums", searchQuery)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlbums indicates an expected call of SearchAlbums.
func (mr *MockUsecaseMockRecorder) SearchAlbums(searchQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlbums", reflect.TypeOf((*MockUsecase)(nil).SearchAlbums), searchQuery)
}
