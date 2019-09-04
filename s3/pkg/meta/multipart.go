package meta

import (
	. "github.com/opensds/multi-cloud/s3/pkg/meta/types"
)

func (m *Meta) GetMultipart(bucketName, objectName, uploadId string) (multipart Multipart, err error) {
	return m.Client.GetMultipart(bucketName, objectName, uploadId)
}

func (m *Meta) DeleteMultipart(multipart Multipart) (err error) {
	tx, err := m.Client.NewTrans()
	defer func() {
		if err != nil {
			m.Client.AbortTrans(tx)
		}
	}()
	err = m.Client.DeleteMultipart(&multipart, tx)
	if err != nil {
		return
	}
	var removedSize int64 = 0
	for _, p := range multipart.Parts {
		removedSize += p.Size
	}
	err = m.UpdateUsage(multipart.BucketName, -removedSize)
	if err != nil {
		return
	}
	err = m.Client.CommitTrans(tx)
	return
}

func (m *Meta) PutObjectPart(multipart Multipart, part Part) (err error) {
	tx, err := m.Client.NewTrans()
	defer func() {
		if err != nil {
			m.Client.AbortTrans(tx)
		}
	}()
	err = m.Client.PutObjectPart(&multipart, &part, tx)
	if err != nil {
		return
	}
	var removedSize int64 = 0
	if part, ok := multipart.Parts[part.PartNumber]; ok {
		removedSize += part.Size
	}
	err = m.UpdateUsage(multipart.BucketName, part.Size-removedSize)
	if err != nil {
		return
	}
	err = m.Client.CommitTrans(tx)
	return
}