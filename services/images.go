package services

import (
	"mime/multipart"
	"os"

	"github.com/DmitryKuzmenec/ImgPrettify/repositories"
	log "github.com/sirupsen/logrus"
)

// Svc
type Svc struct {
	repo *repositories.Repo
}

// NewImgSvc
func NewImgSvc(repo *repositories.Repo) *Svc {
	return &Svc{
		repo: repo,
	}
}

// Pretty
func (s *Svc) Pretty(file *multipart.FileHeader) (*os.File, error) {
	src, err := file.Open()
	if err != nil {
		log.Errorf("[service] file.Open: %s", err)
		return nil, err
	}
	defer src.Close()

	dst, err := s.repo.Pretty(src)
	if err != nil {
		log.Errorf("[service] repo.Pretty: %s", err)
		return nil, err
	}
	return dst, nil
}
