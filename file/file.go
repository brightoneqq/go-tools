package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var ENTER byte = '\n'

func CheckExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func IsDir(file string) (bool, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return false, fmt.Errorf("%T : NO STAT FOR FILE %s", err, file)
	}
	_, err2 := filepath.Abs(file)
	if err2 != nil {
		return false, fmt.Errorf("%T : NO Abs FOR FILE %s", err2, file)
	}
	if fileInfo.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

func ReadLineAct(file *os.File, act func(line string) error) error {
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString(ENTER)
		line = strings.TrimSpace(line)

		actErr := act(line)
		if actErr != nil {
			return err
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return err
			}
		}
	}
	return nil
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
