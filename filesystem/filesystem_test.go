package filesystem

import (
	"github.com/orian/files"

	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFilesystemStoreCreate(t *testing.T) {
	n, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("cannot create a temp dir: %s", err)
		t.FailNow()
	}
	var f files.FileStore = &FilesystemStore{n}
	w, err := f.Create("napis")
	if err != nil {
		t.Errorf("cannot create file: %s", err)
	}
	if w == nil {
		t.Errorf("writer cannot be null")
	}
	defer w.Close()
	fmt.Fprint(w, "some data")
}

func TestFilesystemStoreDelete(t *testing.T) {
	name := "napis"
	n, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("cannot create a temp dir: %s", err)
		t.FailNow()
	}
	var f files.FileStore = &FilesystemStore{n}
	w, err := f.Create(name)
	if err != nil {
		t.Errorf("cannot create file: %s", err)
	}
	if w == nil {
		t.Errorf("writer cannot be null")
	}
	w.Close()

	if err := f.Delete(name); err != nil {
		t.Errorf("cannot delete file: %s", err)
	}

	if _, err := f.Get(name); !os.IsNotExist(err) {
		t.Errorf("expeted a IsNotExist")
	}
}

func TestFilesystemStoreContent(t *testing.T) {
	name := "napis"
	n, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("cannot create a temp dir: %s", err)
		t.FailNow()
	}
	var f files.FileStore = &FilesystemStore{n}
	w, err := f.Create(name)
	if err != nil {
		t.Errorf("cannot create file: %s", err)
	}
	if w == nil {
		t.Errorf("writer cannot be null")
	}
	data := "some data"
	fmt.Fprint(w, data)
	w.Close()
	r, err := f.Get(name)
	if os.IsNotExist(err) {
		t.Errorf("got error: %s", err)
	}
	body, _ := ioutil.ReadAll(r)
	if data != string(body) {
		t.Errorf("want: %s, got: %s", data, body)
	}
}
