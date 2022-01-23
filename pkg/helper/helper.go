package helper

import (
	"fmt"
	"os"
	"strconv"
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
		if isError(err) {
			fmt.Println("File Created Failed", path)
			return
		}
		defer file.Close()
		fmt.Println("File Created Successfully", path)
	}
}

func DeleteHealthFile() {
	var err = os.Remove(path)
	if isError(err) {
		fmt.Println("File deleting failed")
		return
	}
	fmt.Println("File Deleted")

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
