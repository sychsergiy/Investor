package file

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	CreateWorkDir()
	code := m.Run()
	CleanupWorkDir()
	os.Exit(code)
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
	WriteFile(t, filename, string(content))
	// test write some content
	n, err = NewPlainFile(getFullPath(filename)).Write(content)
	if err != nil {
		t.Errorf("Unpected file write err: %s", err)
	} else {
		if n != len(content) {
			t.Errorf("Not expected number of bytes written to file")
		} else {
			written := ReadFile(filename)
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
	WriteFile(t, filename, "some_text")
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
	CreateDir(t, dirName)
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
	WriteFile(t, filename, "")
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
	fullPath := getFullPath(filename)
	f := NewPlainFile(fullPath)
	err := f.Create()
	if err != nil {
		t.Errorf("Unpexected error during file creation: %s", err)
	}

	t.Cleanup(func() {
		err := os.Remove(getFullPath(filename))
		if err != nil {
			log.Fatalf("Failed to remove file: %s due to err: %s", fullPath, err)
		}
	})

	filename2 := "test_create_2.txt"
	WriteFile(t, filename2, "initial_text")
	err = NewPlainFile(getFullPath(filename2)).Create()
	if err != nil {
		t.Errorf("Not expected err: %s", err)
	} else {
		content := ReadFile(filename2)
		if len(content) != 0 {
			t.Error("Empty expected to recreated")
		}
	}
}

func TestJsonFile_Create(t *testing.T) {
	filename := "test_json_file_create_1.json"
	jf := NewJsonFile(PlainFile{getFullPath(filename)})
	err := jf.Create()
	if err != nil {
		t.Errorf("Unepxected err: %s", err)
	} else {
		content := ReadFile(filename)
		if string(content) != "{}" {
			t.Errorf("Json file with empty map expected")
		}
	}
}

func TestJsonFile_Write(t *testing.T) {
	// setup
	filename := "test_json_file_write_1.json"
	WriteFile(t, filename, "")

	// test write
	jf := NewJsonFile(PlainFile{getFullPath(filename)})
	err := jf.WriteJson("test")
	if err != nil {
		t.Errorf("Unepxected err: %s", err)
	} else {
		content := ReadFile(filename)
		if string(content) != "\"test\"" {
			t.Errorf("Json file with empty map expected")
		}
	}

	err = jf.WriteJson(map[string]int{"test": 1})
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		content := ReadFile(filename)
		if string(content) != "{\"test\":1}" {
			t.Errorf("Unexpected content written to json file")
		}
	}
}

func TestCreateIfNotExists(t *testing.T) {
	filename := "test_create_if_not_exists.json"
	jf := NewJsonFile(NewPlainFile(getFullPath(filename)))
	created, err := CreateIfNotExists(jf)
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if created != true {
			t.Errorf("Should return true when new file created")
		}
	}

	// setup existent file
	filename2 := "test_create_if_not_exists_2.json"
	WriteFile(t, filename2, "")
	// test with existent file
	jf2 := NewJsonFile(NewPlainFile(getFullPath(filename2)))
	created, err = CreateIfNotExists(jf2)
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if created != false {
			t.Errorf("Should return false when file already exists")
		}
	}
}
