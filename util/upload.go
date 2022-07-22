package util

import (
	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
	"github.com/qiniupd/qiniu-go-sdk/x/log.v7"
	"os"
)

func UploadAndRemove(src, dest string) error {
	uploader := operation.NewUploaderV2()
	log.Infof("uploading from: %s to: %s", src, dest)

	if err := uploader.Upload(src, dest); err != nil {
		return err
	}

	log.Infof("remove download %s", src)

	if err := os.RemoveAll(src); err != nil {
		log.Errorf("removing car  %s: %+v", src, err)
	}
	return nil
}
