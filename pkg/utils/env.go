package utils

import "os"

func GetEnv(key, d string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}

	return v
}
