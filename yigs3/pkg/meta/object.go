package meta

import (
	. "github.com/opensds/multi-cloud/yigs3/error"
	"github.com/opensds/multi-cloud/yigs3/pkg/helper"
	. "github.com/opensds/multi-cloud/yigs3/pkg/meta/types"
	"github.com/opensds/multi-cloud/yigs3/pkg/redis"
)

const (
	OBJECT_CACHE_PREFIX = "object:"
)

func (m *Meta) GetObject(bucketName string, objectName string, willNeed bool) (object *Object, err error) {
	getObject := func() (o helper.Serializable, err error) {
		helper.Logger.Println(10, "GetObject CacheMiss. bucket:", bucketName, "object:", objectName)
		object, err := m.Client.GetObject(bucketName, objectName, "")
		if err != nil {
			return
		}
		helper.Debugln("GetObject object.Name:", object.Name)
		if object.Name != objectName {
			err = ErrNoSuchKey
			return
		}
		return object, nil
	}

	toObject := func(fields map[string]string) (interface{}, error) {
		o := &Object{}
		return o.Deserialize(fields)
	}

	o, err := m.Cache.Get(redis.ObjectTable, OBJECT_CACHE_PREFIX, bucketName+":"+objectName+":",
		getObject, toObject, willNeed)
	if err != nil {
		return
	}
	object, ok := o.(*Object)
	if !ok {
		err = ErrInternalError
		return
	}
	return object, nil
}

func (m *Meta) GetAllObject(bucketName string, objectName string) (object []*Object, err error) {
	return m.Client.GetAllObject(bucketName, objectName, "")
}

func (m *Meta) GetObjectMap(bucketName, objectName string) (objMap *ObjMap, err error) {
	m.Client.GetObjectMap(bucketName, objectName)
	return
}

func (m *Meta) GetObjectVersion(bucketName, objectName, version string, willNeed bool) (object *Object, err error) {
	getObjectVersion := func() (o helper.Serializable, err error) {
		object, err := m.Client.GetObject(bucketName, objectName, version)
		if err != nil {
			return
		}
		if object.Name != objectName {
			err = ErrNoSuchKey
			return
		}
		return object, nil
	}

	toObject := func(fields map[string]string) (interface{}, error) {
		o := &Object{}
		return o.Deserialize(fields)
	}

	o, err := m.Cache.Get(redis.ObjectTable, OBJECT_CACHE_PREFIX, bucketName+":"+objectName+":"+version,
		getObjectVersion, toObject, willNeed)
	if err != nil {
		return
	}
	object, ok := o.(*Object)
	if !ok {
		err = ErrInternalError
		return
	}
	return object, nil
}

func (m *Meta) PutObject(object *Object, multipart *Multipart, objMap *ObjMap, updateUsage bool) error {
	tx, err := m.Client.NewTrans()
	defer func() {
		if err != nil {
			m.Client.AbortTrans(tx)
		}
	}()

	err = m.Client.PutObject(object, tx)
	if err != nil {
		return err
	}

	if objMap != nil {
		err = m.Client.PutObjectMap(objMap, tx)
		if err != nil {
			return err
		}
	}

	if multipart != nil {
		err = m.Client.DeleteMultipart(multipart, tx)
		if err != nil {
			return err
		}
	}

	if updateUsage {
		err = m.UpdateUsage(object.BucketName, object.Size)
		if err != nil {
			return err
		}
	}
	err = m.Client.CommitTrans(tx)
	return nil
}

func (m *Meta) PutObjectEntry(object *Object) error {
	err := m.Client.PutObject(object, nil)
	return err
}

func (m *Meta) UpdateObjectAcl(object *Object) error {
	err := m.Client.UpdateObjectAcl(object)
	return err
}

func (m *Meta) UpdateObjectAttrs(object *Object) error {
	err := m.Client.UpdateObjectAttrs(object)
	return err
}

func (m *Meta) PutObjMapEntry(objMap *ObjMap) error {
	err := m.Client.PutObjectMap(objMap, nil)
	return err
}

func (m *Meta) DeleteObject(object *Object, DeleteMarker bool, objMap *ObjMap) error {
	tx, err := m.Client.NewTrans()
	defer func() {
		if err != nil {
			m.Client.AbortTrans(tx)
		}
	}()

	err = m.Client.DeleteObject(object, tx)
	if err != nil {
		return err
	}

	if objMap != nil {
		err = m.Client.DeleteObjectMap(objMap, tx)
		if err != nil {
			return err
		}
	}

	if DeleteMarker {
		return nil
	}

	err = m.Client.PutObjectToGarbageCollection(object, tx)
	if err != nil {
		return err
	}

	err = m.UpdateUsage(object.BucketName, -object.Size)
	if err != nil {
		return err
	}
	err = m.Client.CommitTrans(tx)

	return err
}

func (m *Meta) AppendObject(object *Object, isExist bool) error {
	if !isExist {
		return m.Client.PutObject(object, nil)
	}
	return m.Client.UpdateAppendObject(object)
}

//func (m *Meta) DeleteObjectEntry(object *Object) error {
//	err := m.Client.DeleteObject(object, nil)
//	return err
//}

//func (m *Meta) DeleteObjMapEntry(objMap *ObjMap) error {
//	err := m.Client.DeleteObjectMap(objMap, nil)
//	return err
//}
