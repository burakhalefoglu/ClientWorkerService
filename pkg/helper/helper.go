package helper

import (
	"fmt"
	"os"
	"strconv"
)

func ResolvePath(host string, port string) string {
	h := os.Getenv(host)
	p, _ := strconv.Atoi(os.Getenv(port))
	return fmt.Sprintf("%s:%d", h, p)
}
