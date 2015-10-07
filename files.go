// Package files profides a files API and filesystem and Google App Engine
// implementation.

package files

import (
	"fmt"
	"io"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type FileStore interface {
	Create(string) (io.WriteCloser, error)
	Get(string) (io.ReadCloser, error)
	Delete(string) error
}
