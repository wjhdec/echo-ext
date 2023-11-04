package echoext_test

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoext "github.com/wjhdec/echo-ext/v2"
)

func TestFileBinder(t *testing.T) {
	type fileRequest struct {
		File *multipart.FileHeader `form:"file"`
	}

	dir, err := os.MkdirTemp(".", "out-")
	defer os.RemoveAll(dir)
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.Binder = echoext.NewFileBinder(e.Binder)
	e.POST("/demo", func(c echo.Context) error {
		freq := new(fileRequest)
		if err := c.Bind(freq); err != nil {
			return err
		}
		f, err := freq.File.Open()
		if err != nil {
			return err
		}
		defer f.Close()
		outputFile := path.Join(dir, freq.File.Filename)

		of, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer of.Close()
		if _, err = io.Copy(of, f); err != nil {
			return err
		}
		return nil
	})
	reqFile := "README.md"
	body, contentType, err := formBody(reqFile)
	panicErr(err)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/demo", body)
	req.Header.Set(echo.HeaderContentType, contentType)
	e.ServeHTTP(rec, req)
	t.Logf("body: %s", rec.Body.String())
	if rec.Code != 200 {
		t.Errorf("code should be 200, current: %d", rec.Code)
	}
	reqHash, err := fileMd5(reqFile)
	panicErr(err)
	recHash, err := fileMd5(path.Join(dir, filepath.Base(reqFile)))
	panicErr(err)
	if reqHash != recHash {
		t.Errorf("hash not equal, req: %s, rec: %s", reqHash, recHash)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func formBody(path string) (body *bytes.Buffer, contentType string, err error) {
	body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return
	}
	sample, err := os.Open(path)
	if err != nil {
		return
	}
	if _, err = io.Copy(part, sample); err != nil {
		return
	}
	contentType = writer.FormDataContentType()
	return
}

func fileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
