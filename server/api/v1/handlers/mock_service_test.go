package handlers

import (
	"miniurl/api/models"
	"net/http"
)

type MockService struct {
	Counter        int
	InsertURLError error
	GetURLData     models.URL
	GetURLsData    []models.URL
}

func (m *MockService) CreateURL(w http.ResponseWriter, r *http.Request) {
}
func (m *MockService) GetCounter() (int, error) {
	return m.Counter, nil
}
func (m *MockService) InsertURL(url models.URL) error {
	return m.InsertURLError
}
func (m *MockService) GetURLFromShortURL(shortURL string) (models.URL, error) {
	return m.GetURLData, nil
}
func (s *MockService) UpdateHit(url models.URL) error {
	return nil
}
func (s *MockService) GetURLs(options map[string]string) (urls []models.URL, total int, err error) {
	return s.GetURLsData, 0, nil
}
func (s *MockService) DeleteURL(url models.URL) (int, string, error) {
	return 0, "", nil
}
func (s *MockService) CacheURL(url models.URL) {
}
func (s *MockService) GetCachedURL(shortURL string) (string, error) {
	return "", nil
}
func (s *MockService) DeleteCache(shortURL string) {
}

func (s *MockService) InsertBlacklistURL(url string) error {
	return nil
}

func (s *MockService) GetBlacklistURLs() ([]models.BlacklistURL, error) {
	return []models.BlacklistURL{}, nil
}

func (s *MockService) DeleteBlacklistURL(id int) (int, error) {
	return 0, nil
}
func (s *MockService) FindBlacklistURL(url *models.BlacklistURL) error {
	return nil
}
