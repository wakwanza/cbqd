package cbqd

import (
	mn "github.com/wakwanza/minio-go"
	"os"
	"path/filepath"
)

var (
	cbucket   = os.Getenv("CBQD_S3_BCKT")
	clocation = os.Getenv("CBQD_S3_HOST")
	cregion   = os.Getenv("CBQD_S3_REGION")
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
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err := s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/x-gzip")
	if err != nil {
		return err
	}
	log.Info("Uploaded ", n, " successfully.")
	return nil
}

func (a GCS) CloudSend(s AccessCreds, cobject string, cpath string) error {
	s3client, err := mn.New("storage.googleapis.com", s.Dkey, s.Dpass, true)
	if err != nil {
		return err
	}
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err = s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/x-gzip")
	if err != nil {
		return err
	}
	log.Info("Uploaded ", n, " successfully.")
	return nil
}

func (a ALT) CloudSend(s AccessCreds, cobject string, cpath string) error {
	if clocation == "" {
		return S3_BUCKET_ERROR_3
	}
	s3client, err := mn.New(clocation, s.Dkey, s.Dpass, false)
	if err != nil {
		return err
	}
	if cbucket == "" {
		return S3_BUCKET_ERROR_1
	}
	if err = s3client.BucketExists(cbucket); err != nil {
		return S3_BUCKET_ERROR_2
	}

	fpath := filepath.Join(cpath, cobject)

	n, err := s3client.FPutObject(cbucket, cobject, fpath, "application/x-gzip")
	if err != nil {
		return err
	}
	log.Info("Uploaded ", n, " successfully.")
	return nil
}
