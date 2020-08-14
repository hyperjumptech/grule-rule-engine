package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4"
)

// ResouceBundle is a helper struct to help load multiple resource at once.
type ResouceBundle interface {
	Load() ([]Resource, error)
	MustLoad() []Resource
}

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

// NewFileResourceBundle creates new instance of FileResourceBundle struct
// basePath denotes the directory location where the file is located.
// pathPattern are list of paths that filters the files. Its important that
// the pattern will include the base path as it filter.
// For example, if the base path is "/some/base/path"
// The pattern to accept all GRL file is "/some/base/path/**/*.grl".
// This will accept all *.grl files under /some/base/path and its directories.
func NewFileResourceBundle(basePath string, pathPattern ...string) *FileResourceBundle {
	return &FileResourceBundle{
		BasePath:    basePath,
		PathPattern: pathPattern,
	}
}

// FileResourceBundle is a helper struct to load multiple files all at once by specifying
// the root location of the file and the file pattern to look for.
// It will look into sub-directories for the file with pattern matching.
type FileResourceBundle struct {
	// The base path where all the
	BasePath string
	// List Glob like file pattern.
	// *.grl           <- matches abc.grl but not /anyfolder/abc.grl
	// **/*.grl        <- matches abc/def.grl or abc/def/ghi.grl or abc/def/.grl
	// /abc/**/*.grl   <- matches /abc/def.grl or /abc/def/ghi.drl
	PathPattern []string
}

// Load all file resources that locateed under BasePath that conform to the PathPattern.
func (bundle *FileResourceBundle) Load() ([]Resource, error) {
	return bundle.loadPath(bundle.BasePath)
}

// MustLoad function is the same as Load with difference that it will panic if any error is raised
func (bundle *FileResourceBundle) MustLoad() []Resource {
	resources, err := bundle.Load()
	if err != nil {
		panic(err)
	}
	return resources
}

func (bundle *FileResourceBundle) loadPath(path string) ([]Resource, error) {
	logrus.Tracef("Enter directory %s", path)

	finfos, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}
	ret := make([]Resource, 0)
	for _, finfo := range finfos {
		fulPath := fmt.Sprintf("%s/%s", path, finfo.Name())
		fulPath, _ = filepath.Abs(fulPath)
		if finfo.IsDir() {
			gres, err := bundle.loadPath(fulPath)
			if err != nil {
				return nil, err
			}
			ret = append(ret, gres...)
		} else {
			for _, pattern := range bundle.PathPattern {
				matched, err := doublestar.PathMatch(pattern, fulPath)
				if err != nil {
					return nil, err
				}
				if matched {
					logrus.Debugf("Loading file %s", fulPath)
					bytes, err := ioutil.ReadFile(fulPath)
					if err != nil {
						return nil, err
					}
					gress := &FileResource{
						Path:  fulPath,
						Bytes: bytes,
					}
					ret = append(ret, gress)
					break
				}
			}
		}
	}
	return ret, nil
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
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res.Bytes = data
	return res.Bytes, nil
}

// NewGITResourceBundle will create a new instance of GITResourceBundle
// url is the GIT http/https url.
// pathPattern are list of file pattern (glob) to filter files located in the repository
func NewGITResourceBundle(url string, pathPattern ...string) *GITResourceBundle {
	return &GITResourceBundle{
		URL:         url,
		PathPattern: pathPattern,
	}
}

// GITResourceBundle is a helper struct to load multiple files from GIT all at once by specifying
// the necessary information needed to communicate to the GIT server.
// It will look into sub-directories, in the git, for the file with pattern matching.
type GITResourceBundle struct {
	// GIT Repository HTTPS URL
	URL string
	// The Ref name to checkout, if you dont know, let it empty
	RefName string
	// The remote name. IF you left it empty, it will use origin
	Remote string
	// Specify the user name if your repository requires user/password authentication
	User string
	// Password for authentication
	Password string
	// File path pattern to load in your git. The path / is the root on the repository.
	PathPattern []string
}

func (bundle *GITResourceBundle) loadPath(url, path string, fs billy.Filesystem) ([]Resource, error) {
	logrus.Tracef("Enter directory %s", path)
	finfos, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	ret := make([]Resource, 0)
	for _, finfo := range finfos {
		fulPath := fmt.Sprintf("%s/%s", path, finfo.Name())
		if path == "/" && finfo.IsDir() {
			fulPath = fmt.Sprintf("/%s", finfo.Name())
		}
		if finfo.IsDir() {
			gres, err := bundle.loadPath(url, fulPath, fs)
			if err != nil {
				return nil, err
			}
			ret = append(ret, gres...)
		} else {
			for _, pattern := range bundle.PathPattern {
				matched, err := doublestar.Match(pattern, fulPath)
				if err != nil {
					return nil, err
				}
				if matched {
					logrus.Debugf("Loading git file %s", fulPath)
					f, err := fs.Open(fulPath)
					if err != nil {
						return nil, err
					}
					bytes, err := ioutil.ReadAll(f)
					if err != nil {
						return nil, err
					}
					gress := &GITResource{
						URL:   url,
						Path:  fulPath,
						Bytes: bytes,
					}
					ret = append(ret, gress)
					break
				}
			}
		}
	}
	return ret, nil
}

// MustLoad is the same as Load, the difference is it will panic if an error is raised during fetching resources.
func (bundle *GITResourceBundle) MustLoad() []Resource {
	resources, err := bundle.Load()
	if err != nil {
		panic(err)
	}
	return resources
}

// GITResource resource implementation that loaded from GIT
type GITResource struct {
	URL   string
	Path  string
	Bytes []byte
}

// String will state the resource url.
func (res *GITResource) String() string {
	return fmt.Sprintf("From GIT URL [%s] %s", res.URL, res.Path)
}

// Load will load the resource into byte array. This implementation will no re-load resources from git when this method
// is called, it simply return the loaded data.
func (res *GITResource) Load() ([]byte, error) {
	return res.Bytes, nil
}
