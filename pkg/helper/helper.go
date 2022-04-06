package helper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/appneuroncompany/light-logger/clogger"
)

var path = "healthy.txt"

func ResolvePath(host string, port string) string {
	h := os.Getenv(host)
	p, _ := strconv.Atoi(os.Getenv(port))
	return fmt.Sprintf("%s:%d", h, p)
}

func CreateHealthFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			clogger.Error(&map[string]interface{}{
				"message: ": fmt.Sprintf("Fail create file on path: %s", path),
				"error: ":   err})
			return
		}
		defer file.Close()
		clogger.Info(&map[string]interface{}{
			"info message: ": fmt.Sprintf("File created successfully on path: %s", path)})
	}
}

func DeleteHealthFile() {
	var err = os.Remove(path)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"message: ": fmt.Sprintf("Fail delete file on path: %s", path),
			"error: ":   err,
		})
		return
	}
	clogger.Info(&map[string]interface{}{
		"info message: ": fmt.Sprintf("File deleted successfully on path: %s", path)})
}
