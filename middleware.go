package asciiart_middleware

import (
	"math"
	"net/http"
	"strconv"
)

func NewAsciiArtMiddleware(prefix string, message []string) func(next http.Handler) http.Handler {
	multiplier := int64(math.Pow10(int(math.Ceil(math.Log10(float64(len(message)))))))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for i, s := range message {
				w.Header().Add(prefix+strconv.FormatInt(multiplier+int64(i), 10), "# "+s)
			}

			next.ServeHTTP(w, r)
		})
	}
}
