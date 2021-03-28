package db

import (
	"context"
	"miniurl/api/models"
	"os"
	"strconv"
	"time"
)

var popularURLTimelapse = 15
var cachedExpiresInSeconds = 3600

func init() {
	popularURLTimelapse = setIntValue(
		os.Getenv("POPULAR_URL_TIMELAPSE_MINS"),
		popularURLTimelapse,
	)
	cachedExpiresInSeconds = setIntValue(
		os.Getenv("REDIS_URL_EXPIRE_SECONDS"),
		cachedExpiresInSeconds,
	)
}

func setIntValue(value string, defaultValue int) int {
	i, err := strconv.Atoi(value)
	if err == nil && i >= 0 {
		return i
	}
	return defaultValue
}

func (c *Context) CacheURL(url models.URL) {
	lastModifiedLapse := time.Since(url.LastModifiedAt.Local())
	urlTimelapseDuration := time.Duration(popularURLTimelapse) * time.Minute
	expiresInSecondsDuration := time.Duration(cachedExpiresInSeconds) * time.Second

	if lastModifiedLapse < urlTimelapseDuration {
		c.RD.Set(
			context.Background(),
			url.ShortURL,
			url.FullURL,
			expiresInSecondsDuration,
		)
	}
}

func (c *Context) DeleteCache(shortURL string) {
	c.RD.Del(context.Background(), shortURL)
}

func (c *Context) GetCachedURL(shortURL string) (string, error) {
	return c.RD.Get(context.Background(), shortURL).Result()
}
