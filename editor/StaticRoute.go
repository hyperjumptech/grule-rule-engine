package editor

import (
	"embed"
	"fmt"
	mux "github.com/hyperjumptech/hyper-mux"
	"github.com/sirupsen/logrus"
	"grule-rule-engine/editor/mime"
	"net/http"
	"os"
	"strings"
)

//go:embed statics
var fs embed.FS

var (
	errFileNotFound = fmt.Errorf("file not found")
)

type FileData struct {
	Bytes       []byte
	ContentType string
}

func IsDir(path string) bool {
	for _, s := range GetPathTree("static") {
		if s == "[DIR]"+path {
			return true
		}
	}

	return false
}

func GetPathTree(path string) []string {
	logrus.Debugf("Into %s", path)
	var entries []os.DirEntry
	var err error
	if strings.HasPrefix(path, "./") {
		entries, err = fs.ReadDir(path[2:])
	} else {
		entries, err = fs.ReadDir(path)
	}
	ret := make([]string, 0)
	if err != nil {
		return ret
	}
	logrus.Debugf("Path %s %d etries", path, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			ret = append(ret, "[DIR]"+path+"/"+e.Name())
			ret = append(ret, GetPathTree(path+"/"+e.Name())...)
		} else {
			ret = append(ret, path+"/"+e.Name())
		}
	}

	return ret
}

func GetFile(path string) (*FileData, error) {
	bytes, err := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	mimeType, err := mime.MimeForFileName(path)
	if err != nil {

		return &FileData{
			Bytes:       bytes,
			ContentType: http.DetectContentType(bytes),
		}, nil
	}

	return &FileData{
		Bytes:       bytes,
		ContentType: mimeType,
	}, nil
}

func InitializeStaticRoute(router *mux.HyperMux) {
	for _, p := range GetPathTree("statics") {
		if !strings.HasPrefix(p, "[DIR]") {
			path := p[len("statics"):]
			router.AddRoute(path, http.MethodGet, StaticHandler(p))
		}
	}

	router.AddRoute("/", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Location", "/index.html")
		writer.WriteHeader(http.StatusMovedPermanently)
	})
}

func StaticHandler(path string) func(writer http.ResponseWriter, request *http.Request) {
	fData, err := GetFile(path)
	if err != nil {

		return func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", fData.ContentType)
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(fData.Bytes)
	}
}
