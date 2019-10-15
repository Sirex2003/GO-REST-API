package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.WithFields(logrus.Fields{
			"host":        r.Host,
			"remote_addr": r.RemoteAddr,
			"method":      r.Method,
			"work_time":   time.Since(start),
			"path":        r.URL.Path,
		}).Info()
	})
}
