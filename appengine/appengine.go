package appengine

import (
	"github.com/orian/files"
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"

	"io"
)

var _ files.FileStore = NewApi("", nil)

type AppengineStoreConfig struct {
	Bucket string
}

func (cfg *AppengineStoreConfig) Api(c context.Context) *AppengineStore {
	return &AppengineStore{cfg, c}
}

func NewApi(bucket string, c context.Context) *AppengineStore {
	return &AppengineStore{&AppengineStoreConfig{bucket}, c}
}

// It's context dependend.
type AppengineStore struct {
	Cfg *AppengineStoreConfig
	Ctx context.Context
}

func (a *AppengineStore) Create(name string) (io.WriteCloser, error) {
	return storage.NewWriter(a.Ctx, a.Cfg.Bucket, name), nil
}

func (a *AppengineStore) Get(name string) (io.ReadCloser, error) {
	r, err := storage.NewReader(a.Ctx, a.Cfg.Bucket, name)
	if err == storage.ErrObjectNotExist || err == storage.ErrBucketNotExist {
		// TODO log an error
		return nil, files.ErrNotFound
	}
	return r, err
}

func (a *AppengineStore) Delete(name string) error {
	err := storage.DeleteObject(a.Ctx, a.Cfg.Bucket, name)
	if err == storage.ErrObjectNotExist || err == storage.ErrBucketNotExist {
		return files.ErrNotFound
	}
	return err
}
