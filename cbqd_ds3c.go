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

type S3C struct {
	StoreURL string
}

func (a S3C) CloudSend(s AccessCreds, cobject string, cpath string) error {

	if a.StoreURL == "alt" && clocation == "" {
		return S3_BUCKET_ERROR_3
	}
	cloud_url := ""
	switch a.StoreURL {
	case "aws":
		cloud_url = "s3.amazonaws.com"
	case "gce":
		cloud_url = "storage.googleapis.com"
	case "alt":
		cloud_url = clocation
	}

	s3client, err := mn.New(cloud_url, s.Dkey, s.Dpass, true)
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
	log.Info("begin storage bucket upload process.")
	n, err0 := s3client.FPutObject(cbucket, cobject, fpath, "application/x-gzip")
	if err0 != nil {
		return err0
	}
	log.Info("uploaded ", n, " successfully.")
	return nil
}
