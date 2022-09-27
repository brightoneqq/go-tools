package env

import (
	"os"
	"strings"
)

func getEnvMap() map[string]string {
	getEnv := func(data []string,
		getKV func(item string) (key, val string)) map[string]string {

		items := make(map[string]string)
		for _, item := range data {
			key, val := getKV(item)
			items[key] = val
		}
		return items
	}

	envMap := getEnv(os.Environ(),
		func(item string) (key, val string) {
			splits := strings.Split(item, "=")
			key = splits[0]
			val = splits[1]
			return
		})
	return envMap
}
