package utils

import (
	"fmt"
	"os"
)

func SetEnv(k, v string) {
	fmt.Printf("  + %s%s%s='%s'\n", Yellow, k, End, v)
	os.Setenv(k, v)
}
