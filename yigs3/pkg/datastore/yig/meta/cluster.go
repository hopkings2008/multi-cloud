package meta

import (
	"fmt"

	. "github.com/opensds/multi-cloud/yigs3/pkg/datastore/yig/error"
	"github.com/opensds/multi-cloud/yigs3/pkg/datastore/yig/helper"
	. "github.com/opensds/multi-cloud/yigs3/pkg/datastore/yig/meta/types"
	"github.com/opensds/multi-cloud/yigs3/pkg/datastore/yig/redis"
)

const (
	CLUSTER_CACHE_PREFIX = "cluster:"
)

func (m *Meta) GetCluster(fsid string, pool string) (cluster Cluster, err error) {
	rowKey := fsid + ObjectNameEnding + pool
	getCluster := func() (c helper.Serializable, err error) {
		helper.Logger.Println(10, "GetCluster CacheMiss. fsid:", fsid)
		cl, err := m.Client.GetCluster(fsid, pool)
		c = &cl
		return c, err
	}

	toCluster := func(fields map[string]string) (interface{}, error) {
		c := &Cluster{}
		return c.Deserialize(fields)
	}

	c, err := m.Cache.Get(redis.ClusterTable, CLUSTER_CACHE_PREFIX, rowKey, getCluster, toCluster, true)
	if err != nil {
		helper.Logger.Println(20, fmt.Sprintf("failed to get cluster for fsid: %s, err: %v", fsid, err))
		return
	}
	cluster, ok := c.(Cluster)
	if !ok {
		err = ErrInternalError
		return
	}
	return cluster, nil
}
