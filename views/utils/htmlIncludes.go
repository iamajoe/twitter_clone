package utils

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

func cssStylesheet(name string) (res string) {
	filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == name {
			p := strings.Replace(path, "views/", "", 1)
			res = "<link rel=\"stylesheet\" href=\"/" + p + "\">"
		}

		return nil
	})

	return res
}

func jsScript(name string) (res string) {
	filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == name {
			p := strings.Replace(path, "views/", "", 1)
			res = "<script type=\"module\" src=\"/" + p + "\" defer></script>"
		}

		return nil
	})

	return res
}

func IncludeCss(name string) templ.ComponentFunc {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) (err error) {
		// not a public file, include it already
		if strings.Index(name, "http") == 0 {
			_, err = io.WriteString(w, name)
			return
		}

		_, err = io.WriteString(w, cssStylesheet(name))
		return
	})
}

func IncludeJs(name string) templ.ComponentFunc {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) (err error) {
		// not a public file, include it already
		if strings.Index(name, "http") == 0 {
			_, err = io.WriteString(w, name)
			return
		}

		_, err = io.WriteString(w, jsScript(name))
		return
	})
}
