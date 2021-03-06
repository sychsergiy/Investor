package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

const WorkDir = "__FILE_SYSTEM_RELATED_TESTS_DIR__"

func CleanupWorkDir() {
	err := os.RemoveAll(WorkDir)
	if err != nil {
		log.Fatalf("Failed to cleanup work dir due to err: %s", err)
	}
}

func GetFilePath(p string) string {
	return path.Join(WorkDir, p)
}

func ReadFile(filename string) []byte {
	fullPath := GetFilePath(filename)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Failed to read file with path: %s due to err: %s", fullPath, err)
	}
	return content
}

func CreateDir(t *testing.T, dirName string) {
	fullPath := GetFilePath(dirName)
	err := os.Mkdir(fullPath, os.ModeDir)
	if err != nil {
		checkErr(err, fmt.Sprintf("Failed to create dir with path %s", fullPath))
	}

	t.Cleanup(func() {
		err := os.Remove(fullPath)
		checkErr(err, fmt.Sprintf("Failed to remove dir with path %s", fullPath))
	})
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s. Root error: %s", message, err)
	}
}

func WriteBytesToFile(t *testing.T, filename string, bytes []byte) {
	fullPath := GetFilePath(filename)
	err := ioutil.WriteFile(fullPath, bytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write file with path: %s due to err: %s", fullPath, err)
	}

	t.Cleanup(func() {
		err := os.Remove(fullPath)
		checkErr(err, fmt.Sprintf("Failed to remove file with path: %s", fullPath))
	})
}

func WriteTextToFile(t *testing.T, filename, text string) {
	WriteBytesToFile(t, filename, []byte(text))
}

func CreateWorkDir() {
	err := os.MkdirAll(WorkDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create a dir with path: %s due to err: %s", WorkDir, err)
	}

}
