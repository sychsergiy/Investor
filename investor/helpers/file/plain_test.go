package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

const WorkDirEnvVar = "TESTS_WORK_DIR"

func TestMain(m *testing.M) {
	createWorkDir()
	code := m.Run()
	cleanupWorkDir()
	os.Exit(code)
}

func getWorkDirPath() string {
	value, ok := os.LookupEnv(WorkDirEnvVar)
	if !ok {
		log.Fatalf("Please set %s env var", WorkDirEnvVar)
	}
	return value
}

func cleanupWorkDir() {
	err := os.RemoveAll(getWorkDirPath())
	if err != nil {
		log.Fatalf("Failed to cleanup work dir due to err: %s", err)
	}
}

func getFullPath(p string) string {
	return path.Join(getWorkDirPath(), p)
}

func readFile(filename string) []byte {
	fullPath := getFullPath(filename)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Failed to read file with path: %s due to err: %s", fullPath, err)
	}
	return content
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s. Root error: %s", message, err)
	}
}

func writeFile(t *testing.T, filename, text string) {
	fullPath := getFullPath(filename)
	err := ioutil.WriteFile(fullPath, []byte(text), 0644)
	if err != nil {
		log.Fatalf("Failed to write file with path: %s due to err: %s", fullPath, err)
	}

	t.Cleanup(func() {
		err := os.Remove(fullPath)
		checkErr(err, fmt.Sprintf("Failed to remove file with path: %s", fullPath))
	})
}

func createWorkDir() {
	p := getWorkDirPath()
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create a dir with path: %s due to err: %s", p, err)
	}

}

func TestPlainFile_Write(t *testing.T) {
	// test file not exists
	n, err := NewPlainFile(getFullPath("not_exists_file")).Write([]byte(""))
	switch _, err := NewPlainFile(getFullPath("not_exists_file")).Write([]byte("")); err.(type) {
	default:
		t.Errorf("PathError expected got %s", err)
	case *os.PathError:
	}

	// setup
	filename := "write_test_1.txt"
	content := []byte("some_text")
	writeFile(t, filename, string(content))
	// test write some content
	n, err = NewPlainFile(getFullPath(filename)).Write(content)
	if err != nil {
		t.Errorf("Unpected file write err: %s", err)
	} else {
		if n != len(content) {
			t.Errorf("Not expected number of bytes written to file")
		} else {
			written := readFile(filename)
			if string(written) != string(content) {
				t.Errorf("Unexpected content writtne to file")
			}
		}
	}
}

func TestPlainFile_Read(t *testing.T) {
	// file not exists
	fullPath := getFullPath("read_test_1.txt")
	switch _, err := NewPlainFile(fullPath).Read(); err.(type) {
	case *os.PathError:
	default:
		t.Errorf("Path err expected")
	}

	// file exists with content
	// setup
	filename := "read_test_2.txt"
	writeFile(t, filename, "some_text")
	content, err := NewPlainFile(getFullPath(filename)).Read()
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if string(content) != "some_text" {
			t.Errorf("Unexpected file content")
		}
	}
}

func TestPlainFile_Exists(t *testing.T) {

}

func TestPlainFile_Delete(t *testing.T) {

}
