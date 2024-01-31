package tools

import (
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func SetCors(logger *logrus.Logger) *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedOrigins: []string{
			"http://localhost:5173",
			"https://uptimerobot.com",
			"http://www.uptimerobot.com",
			os.Getenv("ORIGIN_ALLOWED"),
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
			"Accept",
			"Origin",
			"X-Requested-With",
			"X-CSRF-Token",
			"Set-Cookie",
		},
		Debug:  false,
		Logger: logger,
	})
	return c
}
