package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FileScanner struct {
	Root string
}

func (s *FileScanner) PathExists(target string) bool {
	files, err := ioutil.ReadDir(s.Root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == target {
			return true
		}
	}
	return false
}

func (s *FileScanner) PathExistsByRoot(root, path string) bool {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == path {
			return true
		}
	}
	return false
}

func (s *FileScanner) FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (s *FileScanner) FindRepo(path string) (string, error) {
	target := fmt.Sprintf("%s/%s", s.Root, path)
	if s.PathExists(path) {
		files, err := ioutil.ReadDir(target)
		if err != nil {
			return "", err
		}
		for _, file := range files {
			root := fmt.Sprintf("%s/%s", target, file.Name())
			if file.IsDir() && s.PathExistsByRoot(root, ".git") {
				return root, nil
			}
		}
	}
	return "", nil
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
