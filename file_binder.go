package echoext

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type fileBinder struct {
	base echo.Binder
}

func (b *fileBinder) Bind(i any, c echo.Context) error {
	if err := b.base.Bind(i, c); err != nil {
		return err
	}
	cType := c.Request().Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(cType, echo.MIMEApplicationForm) || strings.HasPrefix(cType, echo.MIMEMultipartForm) {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		return b.bindFile(i, form.File)
	}
	return nil
}

func (b *fileBinder) bindFile(i any, filesMap map[string][]*multipart.FileHeader) error {
	iValue := reflect.Indirect(reflect.ValueOf(i))
	if iValue.Kind() != reflect.Struct {
		return fmt.Errorf("bind file input indirect [%s] is not pointer", iValue.Type().String())
	}
	iType := iValue.Type()
	for i := 0; i < iType.NumField(); i++ {
		fValue := iValue.Field(i)
		if !fValue.CanSet() {
			continue
		}
		fType := iType.Field(i)
		for _, name := range []string{fType.Name, fType.Tag.Get("form")} {
			if files, ok := filesMap[name]; ok {
				if len(files) == 0 {
					continue
				}
				if fType.Type == reflect.TypeOf((*multipart.FileHeader)(nil)) {
					fValue.Set(reflect.ValueOf(files[0]))
				} else {
					fValue.Set(reflect.ValueOf(files))
				}
				break
			}
		}
	}
	return nil
}

func NewFileBinder(baseBinder echo.Binder) echo.Binder {
	return &fileBinder{
		base: baseBinder,
	}
}
