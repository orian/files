package appengine

import (
	"github.com/orian/files"
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"

	"io"
)

var _ files.FileStore = NewApi("a", nil)

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
	return nil, nil
}

func (a *AppengineStore) Delete(name string) error {
	return nil
}
