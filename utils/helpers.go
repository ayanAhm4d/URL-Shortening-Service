package utils

import (
	"os"
	"strings"

	"github.com/asaskevich/govalidator"
)

func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url
}

func IsValidURL(url string) bool {
	return govalidator.IsURL(url)
}

func GetBaseURL() string {
	return os.Getenv("BASE_URL")
}
