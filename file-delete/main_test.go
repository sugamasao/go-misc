package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteParentDirectory(t *testing.T) {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	path := filepath.Join(dir, "foo", "bar")
	os.MkdirAll(path, 0777)
	tmpfn := filepath.Join(path, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(tmpfn); os.IsNotExist(err) {
		t.Errorf("tmpfile is Not Created")
	}
	actual := DeleteParentDirectory(path)
	if actual != nil {
		t.Errorf("got: %v\nwant: %v", actual, nil)
	}
	if _, err := os.Stat(tmpfn); os.IsExist(err) {
		t.Errorf("Not Delete tmpfile. got: %v", err)
	}
	if _, err := os.Stat(filepath.Dir(path)); os.IsExist(err) {
		t.Errorf("Not Delete tmpfile. got: %v", err)
	}
	if _, err := os.Stat(dir); err != nil {
		t.Errorf("Not Delete tmpfile. got: %v", err)
	}
}
