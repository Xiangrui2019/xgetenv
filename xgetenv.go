package xgetenv

import (
	"log"
	"os"
)

func XGetenv(key string, default_value ...string) string {
	if r := os.Getenv(key); r != "" {
		return r
	} else {
		if len(default_value) == 0 {
			return ""
		}

		return default_value[0]
	}

	return ""
}
