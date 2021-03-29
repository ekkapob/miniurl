package service

import (
	"context"
	"miniurl/api/models"
	"miniurl/pkg/utils"
	"os"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

var (
	DB                     *pg.DB
	RD                     *redis.Client
	popularURLTimelapse    = 15
	cachedExpiresInSeconds = 3600
)

type URLService interface {
	InsertURL(url models.URL) error
	GetURLFromShortURL(shortURL string) (models.URL, error)
	GetCounter() (int, error)
	UpdateHit(url models.URL) error
	GetURLs(options map[string]string) (urls []models.URL, total int, err error)
	DeleteURL(url models.URL) (int, string, error)

	CacheURL(url models.URL)
	GetCachedURL(shortURL string) (string, error)
	DeleteCache(shortURL string)
}

type service struct {
	mu sync.Mutex
}

func NewURLService(db *pg.DB, rd *redis.Client) URLService {
	DB = db
	RD = rd
	popularURLTimelapse = utils.GetInt(
		os.Getenv("POPULAR_URL_TIMELAPSE_MINS"),
		popularURLTimelapse,
	)
	cachedExpiresInSeconds = utils.GetInt(
		os.Getenv("REDIS_URL_EXPIRE_SECONDS"),
		cachedExpiresInSeconds,
	)
	return &service{}
}

func (s *service) InsertURL(url models.URL) error {
	_, err := DB.Model(&url).Insert()
	return err
}

func (s *service) GetURLFromShortURL(shortURL string) (models.URL, error) {
	var url models.URL
	err := DB.Model(&url).
		Where(
			`short_url = ? AND
			now() <= created_at + expires_in_seconds * interval '1 second'`,
			shortURL,
		).
		Select()

	return url, err
}

func (s *service) UpdateHit(url models.URL) error {
	s.mu.Lock()
	url.Hits++
	_, err := DB.Model(&url).WherePK().Update()
	s.mu.Unlock()
	return err
}

func (s *service) GetURLs(options map[string]string) (urls []models.URL, total int, err error) {
	query := DB.Model(&urls)
	i, err := utils.GetIntFromMap(options, "limit")
	if err == nil {
		query.Limit(i)
	}
	i, err = utils.GetIntFromMap(options, "page")
	if err == nil {
		query.Offset(i)
	}

	if v, ok := options["orderBy"]; ok {
		if v == "expired_date" {
			query.OrderExpr(
				"created_at + expires_in_seconds * interval '1 second'" +
					" " +
					options["orderDirection"],
			)
		} else {
			query.Order(v + " " + options["orderDirection"])
		}
	}

	count, err := query.SelectAndCount()
	return urls, count, err
}

func (s *service) DeleteURL(url models.URL) (int, string, error) {
	var shortURLs []string
	r, err := DB.Model(&url).WherePK().Returning("short_url").Delete(&shortURLs)
	rowAffected := r.RowsAffected()

	var shortURL string
	if rowAffected > 0 {
		shortURL = shortURLs[0]
	}
	return rowAffected, shortURL, err
}

func (s *service) GetCounter() (int, error) {
	var counter int
	s.mu.Lock()
	_, err := DB.QueryOne(
		pg.Scan(&counter),
		`SELECT nextval('url_counter')`,
	)
	s.mu.Unlock()
	return counter, err
}

func (s *service) CacheURL(url models.URL) {
	lastModifiedLapse := time.Since(url.LastModifiedAt.Local())
	urlTimelapseDuration := time.Duration(popularURLTimelapse) * time.Minute
	expiresInSecondsDuration := time.Duration(cachedExpiresInSeconds) * time.Second

	if lastModifiedLapse < urlTimelapseDuration {
		RD.Set(
			context.Background(),
			url.ShortURL,
			url.FullURL,
			expiresInSecondsDuration,
		)
	}
}

func (s *service) GetCachedURL(shortURL string) (string, error) {
	return RD.Get(context.Background(), shortURL).Result()
}

func (s *service) DeleteCache(shortURL string) {
	RD.Del(context.Background(), shortURL)
}
