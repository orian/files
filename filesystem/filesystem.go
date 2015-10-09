package filesystem

import (
	"github.com/orian/files"
	"golang.org/x/net/context"

	"io"
	"os"
	"path"
)

var _ files.FileStore = &FilesystemStore{}
var _ files.Generator = &FilesystemStore{}

type writer struct {
	f     *os.File
	tmpP  string
	destP string
}

func (w *writer) Write(p []byte) (n int, err error) {
	return w.f.Write(p)
}

func (w *writer) Close() error {
	return w.f.Close()
}

type FilesystemStore struct {
	Dir string
}

func (f *FilesystemStore) Generate(ctx context.Context) files.FileStore {
	return f
}

func (f *FilesystemStore) Create(name string) (io.WriteCloser, error) {
	return os.Create(path.Join(f.Dir, name))
}

func (f *FilesystemStore) Get(name string) (io.ReadCloser, error) {
	r, err := os.Open(path.Join(f.Dir, name))
	if os.IsNotExist(err) {
		return nil, files.ErrNotFound
	}
	return r, err
}

func (f *FilesystemStore) Delete(name string) error {
	err := os.Remove(path.Join(f.Dir, name))
	if os.IsNotExist(err) {
		return files.ErrNotFound
	}
	return err
}
