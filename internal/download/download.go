package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/devmoded/mp-cli/internal/output"
)

type ProgressReader struct {
	Reader     io.Reader
	Output     output.Output
	Total      int64
	Downloaded int64
}

func (pr *ProgressReader) Read(buf []byte) (int, error) {
	n, err := pr.Reader.Read(buf)
	pr.Downloaded += int64(n)

	pr.Output.Progress(output.Message{
		Event:           "download",
		PercentProgress: float64(pr.Downloaded) / float64(pr.Total) * 100,
		BytesProgress:   pr.Downloaded,
		Total:           pr.Total,
	})
	return n, err
}

func getFilename(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return path.Base(u.Path), nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func Download(o output.Output, url string, path string, filename string) {
	r, err := http.Get(url)
	if err != nil {
		o.Println(output.Message{Event: "error", Message: err.Error()})
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		o.Println(output.Message{Event: "error", Message: r.Status})
	}

	if filename == "" {
		filename, err = getFilename(url)
		if err != nil {
			o.Println(output.Message{Event: "error", Message: err.Error()})
		}
	}

	path, err = filepath.Abs(path)
	if err != nil {
		o.Println(output.Message{Event: "error", Message: err.Error()})
	}
	path = filepath.Join(path, filename)
	if fileExists(path) {
		o.Println(output.Message{Event: "error", Message: path + " is exists"})
		return
	}

	f, err := os.Create(path)
	if err != nil {
		o.Println(output.Message{Event: "error", Message: err.Error()})
	}
	defer f.Close()

	pr := &ProgressReader{
		Reader: r.Body,
		Output: o,
		Total:  r.ContentLength,
	}

	o.Println(output.Message{
		Event:    "download",
		Message:  "start downloading " + filename + " in " + path,
		FilePath: path,
	})
	_, err = io.Copy(f, pr)
	if err != nil {
		o.Println(output.Message{Event: "error", Message: err.Error()})
	}
	fmt.Println()
}
