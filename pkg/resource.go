package pkg

import (
	"fmt"
	"github.com/juju/errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Resource should be implemented by any resource model so it can be loaded and parsed by the grule parser.
type Resource interface {
	Load() ([]byte, error)
	String() string
}

// NewReaderResource will create a new Resource using a common reader.
func NewReaderResource(reader io.Reader) Resource {
	return &ReaderResource{Reader: reader}
}

// ReaderResource is a struct that will hold the reader instance.
type ReaderResource struct {
	Reader io.Reader
}

// Load will load the resource into byte array.
func (res *ReaderResource) Load() ([]byte, error) {
	return ioutil.ReadAll(res.Reader)
}

// String will state the resource source.
func (res *ReaderResource) String() string {
	return "Reader resource. Source unknown."
}

// NewFileResource will create a new Resource using a file located in path.
func NewFileResource(path string) Resource {
	return &FileResource{
		Path: path,
	}
}

// FileResource is a struct that will hold the file path and readed data bytes.
type FileResource struct {
	Path  string
	Bytes []byte
}

// Load will load the resource into byte array.
// The load byte array will be cached by the FileResource. So Calling
// Load multiple time will only load the file once on the first call.
// If you wish to reload the file, simply create new instance using NewFileResource function.
func (res *FileResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	}
	data, err := ioutil.ReadFile(res.Path)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.Bytes = data
	return res.Bytes, nil
}

// String will state the resource file path.
func (res *FileResource) String() string {
	return fmt.Sprintf("File resource at %s", res.Path)
}

// NewBytesResource will create a new Resource using a byte array.
func NewBytesResource(bytes []byte) Resource {
	return &BytesResource{
		Bytes: bytes,
	}
}

// BytesResource is a struct that will hold the byte array data
type BytesResource struct {
	Bytes []byte
}

// Load will load the resource into byte array.
func (res *BytesResource) Load() ([]byte, error) {
	return res.Bytes, nil
}

// String will state the resource byte array.
func (res *BytesResource) String() string {
	return fmt.Sprintf("Byte array resources %d bytes", len(res.Bytes))
}

// NewURLResource will create a new Resource using a resource as located in the url
func NewURLResource(url string) Resource {
	return &URLResource{
		URL: url,
	}
}

// URLResource is a struct that will hold the byte array data and URL source
type URLResource struct {
	URL   string
	Bytes []byte
}

// String will state the resource url.
func (res *URLResource) String() string {
	return fmt.Sprintf("URL resource at %s", res.URL)
}

// Load will load the resource into byte array. This resource will cache the obtained result byte arrays.
// So calling this function multiple times only call the URL once at the first time.
// If you want to refresh the load, you simply create a new instance of URLResource using
// NewURLResource
func (res *URLResource) Load() ([]byte, error) {
	if res.Bytes != nil {
		return res.Bytes, nil
	}
	resp, err := http.Get(res.URL)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.Bytes = data
	return res.Bytes, nil
}
