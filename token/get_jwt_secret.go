package token

import (
	"os"

	"callumkloos.dev/reelingit/logger"
)

func GetJWTSecret(logger logger.Logger) string {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		jwtSecret = "default-secret-for-dev"
		logger.Info("JWT_SECRET NOT SET, using default development secret")
	} else {
		logger.Info("JWT_SECRET found")
	}

	return jwtSecret
}
