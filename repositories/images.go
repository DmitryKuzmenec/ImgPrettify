package repositories

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"

	"github.com/DmitryKuzmenec/ImgPrettify/config"
	log "github.com/sirupsen/logrus"
)

// Repo
type Repo struct {
	conf *config.Config
}

// NewImgRepo
func NewImgRepo(conf *config.Config) *Repo {
	return &Repo{
		conf: conf,
	}
}

// Pretty
func (r *Repo) Pretty(src multipart.File) (*os.File, error) {
	dst, err := ioutil.TempFile("", "*")
	if err != nil {
		log.Errorf("[repositories] ioutil.TempFile: %s", err)
		return nil, err
	}
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	convertedFileName := dst.Name() + "_new"
	cmd := exec.Command("convert", "-fuzz", "20%", "-trim", dst.Name(), convertedFileName)
	var out bytes.Buffer
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		log.Errorf("[repositories] exec.Command: %s (%s)", err, out.String())
		return nil, err
	}
	file, err := os.Open(convertedFileName)
	if err != nil {
		log.Errorf("[repositories] os.Open: %s", err)
		return nil, err
	}
	defer file.Close()
	return file, nil
}
