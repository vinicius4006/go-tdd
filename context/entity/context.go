package entity

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyResponseWriter struct {
	Written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.Written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.Written = true
	return 0, errors.New("não implementado")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.Written = true
}

type SpyStore struct {
	Response string
	T        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.Response {
			select {
			case <-ctx.Done():
				s.T.Log("spy store foi cancelado")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}
