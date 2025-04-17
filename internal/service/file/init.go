package file

import (
	"io"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/internal/repository/file"
	"loan-service/pkg/helper"
	"mime/multipart"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type IFileService interface {
	IsExist(id string) bool
	Find(id string) *Model
	Save(file *multipart.FileHeader, fType constant.FileType, pathPrefix, requestedBy string) (*Model, error)
}

type LocalFile struct {
	fileRepo file.IFileRepository
}

func New(fileRepo file.IFileRepository) IFileService {
	return &LocalFile{fileRepo}
}

func (f *LocalFile) Find(id string) *Model {
	file := f.fileRepo.Get(id)

	if file == nil {
		return nil
	}

	return &Model{
		ID:           file.ID,
		OriginalName: file.OriginalName,
		Path:         file.Path,
		MimeType:     file.MimeType,
		Type:         file.Type,
	}
}

func (f *LocalFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func (f *LocalFile) Save(file *multipart.FileHeader, fType constant.FileType, pathPrefix, requestedBy string) (m *Model, err error) {
	mimeType, err := helper.GetMimeType(file)

	if err != nil {
		log.Errorf("Error getting mime type: %s", err.Error())
		return
	}

	fileID := uuid.New().String()

	_, ext := helper.SplitLast(file.Filename, ".")

	p := filepath.Join("file", pathPrefix, fileID+"."+ext)

	if err = f.write(file, p); err != nil {
		log.Errorf("Error writing to local disk: %s", err.Error())
		return
	}

	newFile := &entity.File{
		ID:           fileID,
		OriginalName: file.Filename,
		Path:         p,
		MimeType:     mimeType,
		Type:         fType,
	}
	if err = f.fileRepo.Create(newFile); err != nil {
		log.Errorf("Error creating file: %s", err.Error())
		return
	}

	return &Model{
		ID:           newFile.ID,
		OriginalName: newFile.OriginalName,
		Path:         newFile.Path,
		MimeType:     newFile.MimeType,
		Type:         newFile.Type,
	}, nil
}

func (f *LocalFile) write(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
