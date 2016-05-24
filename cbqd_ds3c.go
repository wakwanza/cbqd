package cbqd

import (
	mn "github.com/wakwanza/minio-go"
	"log"
	"os"
	"path/filepath"
)

type AWS struct {
}

type GCS struct {
}

type ALT struct {
}

func (a AWS) CloudSend(s AccessCreds, cobject string, cpath string) error {
	s3client, err := mn.New("s3.amazonaws.com", s.Dkey, s.Dpass, true)
	if err != nil {
		return err
	}
	cbucket := os.Getenv("CBQD_S3_BCKT")
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err := s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/gzip")
	if err != nil {
		return err
	}
	log.Println("Uploaded ", n, " successfully.")
	return nil
}

func (a GCS) CloudSend(s AccessCreds, cobject string, cpath string) error {
	s3client, err := mn.New("storage.googleapis.com", s.Dkey, s.Dpass, true)
	if err != nil {
		return err
	}
	cbucket := os.Getenv("CBQD_S3_BCKT")
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err = s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/gzip")
	if err != nil {
		return err
	}
	log.Println("Uploaded ", n, " successfully.")
	return nil
}

func (a ALT) CloudSend(s AccessCreds, cobject string, cpath string) error {
	clocation := os.Getenv("CBQD_S3_HOST")
	if clocation == "" {
		return S3_BUCKET_ERROR_3
	}
	s3client, err := mn.New(clocation, s.Dkey, s.Dpass, false)
	if err != nil {
		return err
	}
	cbucket := os.Getenv("CBQD_S3_BCKT")
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err = s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/gzip")
	if err != nil {
		return err
	}
	log.Println("Uploaded ", n, " successfully.")
	return nil
}
