package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
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

func createDir(t *testing.T, dirName string) {
	fullPath := getFullPath(dirName)
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
	// test false without file
	exists, err := NewPlainFile(getFullPath("not_existent")).Exists()
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if exists != false {
			t.Errorf("File exsists false result expected")
		}
	}

	// setup
	dirName := "test_exists"
	createDir(t, dirName)
	exists, err = NewPlainFile(getFullPath(dirName)).Exists()
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if exists != false {
			t.Errorf("File exists with directory path should return false")
		}
	}
	// test false when dir

	// setup
	filename := "test_exists_2.txt"
	writeFile(t, filename, "")
	// test true
	exists, err = NewPlainFile(getFullPath(filename)).Exists()
	if err != nil {
		t.Errorf("Unpected err: %s", err)
	} else {
		if exists != true {
			t.Errorf("File exists true result expected")
		}
	}

}

func TestPlainFile_Create(t *testing.T) {
	filename := "test_create_1.txt"
	f, err := NewPlainFile(getFullPath(filename)).Create()
	if err != nil {
		t.Errorf("Unpexected error during file creation: %s", err)
	} else {
		if !strings.HasSuffix(f.Name(), filename) {
			t.Errorf("Unexpected filename: %s", filename)
		}
	}

	filename = "test_create_2.txt"
	writeFile(t, filename, "initial_text")
	_, err = NewPlainFile(getFullPath(filename)).Create()
	if err != nil {
		t.Errorf("Not expected err: %s", err)
	} else {
		content := readFile(filename)
		if len(content) != 0 {
			t.Error("Empty expected to recreated")
		}
	}
}
