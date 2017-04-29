package serverstats

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
)

type ServerStats struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"requestCount"`
	Statuses     map[string]int `json:"statuses"`

	mutex sync.RWMutex
}

func NewServerStats() *ServerStats {
	return &ServerStats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

func (s *ServerStats) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if err = next(ctx); err != nil {
			ctx.Error(err)
		}

		s.mutex.Lock()
		defer s.mutex.Unlock()

		s.RequestCount++

		status := strconv.Itoa(ctx.Response().Status)
		s.Statuses[status]++

		return nil
	}
}

func (s *ServerStats) Handler(ctx echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return ctx.JSON(http.StatusOK, s)
}
