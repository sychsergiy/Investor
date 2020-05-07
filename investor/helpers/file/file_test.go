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
	n, err := NewPlainFile(GetFilePath("not_exists_file")).Write([]byte(""))
	switch _, err := NewPlainFile(GetFilePath("not_exists_file")).Write([]byte("")); err.(type) {
	default:
		t.Errorf("PathError expected got %s", err)
	case *os.PathError:
	}

	// setup
	filename := "write_test_1.txt"
	content := []byte("some_text")
	WriteTextToFile(t, filename, string(content))
	// test write some content
	n, err = NewPlainFile(GetFilePath(filename)).Write(content)
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
	fullPath := GetFilePath("read_test_1.txt")
	switch _, err := NewPlainFile(fullPath).Read(); err.(type) {
	case *os.PathError:
	default:
		t.Errorf("Path err expected")
	}

	// file exists with content
	// setup
	filename := "read_test_2.txt"
	WriteTextToFile(t, filename, "some_text")
	content, err := NewPlainFile(GetFilePath(filename)).Read()
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
	exists, err := NewPlainFile(GetFilePath("not_existent")).Exists()
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
	exists, err = NewPlainFile(GetFilePath(dirName)).Exists()
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
	WriteTextToFile(t, filename, "")
	// test true
	exists, err = NewPlainFile(GetFilePath(filename)).Exists()
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
	fullPath := GetFilePath(filename)
	f := NewPlainFile(fullPath)
	err := f.Create()
	if err != nil {
		t.Errorf("Unpexected error during file creation: %s", err)
	}

	t.Cleanup(func() {
		err := os.Remove(GetFilePath(filename))
		if err != nil {
			log.Fatalf("Failed to remove file: %s due to err: %s", fullPath, err)
		}
	})

	filename2 := "test_create_2.txt"
	WriteTextToFile(t, filename2, "initial_text")
	err = NewPlainFile(GetFilePath(filename2)).Create()
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
	jf := NewJsonFile(PlainFile{GetFilePath(filename)})
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
	WriteTextToFile(t, filename, "")

	// test write
	jf := NewJsonFile(PlainFile{GetFilePath(filename)})
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
	jf := NewJsonFile(NewPlainFile(GetFilePath(filename)))
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
	WriteTextToFile(t, filename2, "")
	// test with existent file
	jf2 := NewJsonFile(NewPlainFile(GetFilePath(filename2)))
	created, err = CreateIfNotExists(jf2)
	if err != nil {
		t.Errorf("Unexpected err: %s", err)
	} else {
		if created != false {
			t.Errorf("Should return false when file already exists")
		}
	}
}
