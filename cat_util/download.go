package cat_util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func GetList(filename string) ([]string, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, fmt.Errorf("Failed to read in list of cats => %s", err.Error())
	}

	return strings.Split(string(buf), "\n"), nil
}

func GetCat(url string) error {
	resp, err := http.Get("http://" + url)
	if err != nil {
		return fmt.Errorf("Failed to download cat => %s", err.Error())
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read HTTP response => %s", err.Error())
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	if ioutil.WriteFile(filename, buf, 0744) != nil {
		return fmt.Errorf("Failed to write to file, %s=> %s", filename, err.Error())
	}

	return nil
}
