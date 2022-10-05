package env

import (
	"fmt"
	"testing"
)

func TestEnvMap(t *testing.T) {
	envMap := GetEnvMap()
	for k, v := range envMap {
		fmt.Println(k, " -> ", v)
	}
}
