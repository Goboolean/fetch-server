package etcdutil_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"

	etcdutil "github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func Contains[T any](list []T, target T) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, target) {
			return true
		}
	}
	return false
}


func Test_GroupBy(t *testing.T) {

	type args struct {
		list   map[string]string
		prefix string
	}

	tests := []struct {
		name    string
		args    args
		want    []map[string]string
		wantLen int
	} {
		{
			name: "Worker",
			args: args{
				list: map[string]string{
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "waiting",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886": "",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/platform": "kis",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/status": "active",
				},
				prefix: "worker",
			},
			want: []map[string]string{
				{
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "waiting",
				},
				{
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886": "",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/platform": "kis",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/status": "active",
				},
			},
			wantLen: 2,
		},
		{
			name: "Product",
			args: args{
				list: map[string]string{
					"product/test.goboolean.kor": "",
					"product/test.goboolean.kor/platform": "kis",
					"product/test.goboolean.kor/symbol": "goboolean",
					"product/test.goboolean.kor/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					"product/test.goboolean.kor/status": "onsubscribe",
					"product/test.goboolean.eng": "",
					"product/test.goboolean.eng/platform": "kis",
					"product/test.goboolean.eng/symbol": "goboolean",
					"product/test.goboolean.eng/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					"product/test.goboolean.eng/status": "onsubscribe",
				},
				prefix: "product",
			},
			want: []map[string]string{
				{
					"product/test.goboolean.kor": "",
					"product/test.goboolean.kor/platform": "kis",
					"product/test.goboolean.kor/symbol": "goboolean",
					"product/test.goboolean.kor/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					"product/test.goboolean.kor/status": "onsubscribe",
				},
				{
					"product/test.goboolean.eng": "",
					"product/test.goboolean.eng/platform": "kis",
					"product/test.goboolean.eng/symbol": "goboolean",
					"product/test.goboolean.eng/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					"product/test.goboolean.eng/status": "onsubscribe",
				},
			},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := etcdutil.GroupBy(tt.args.list, tt.args.prefix)
			assert.Equal(t, tt.wantLen, len(got))
			for _, v := range got {
				assert.True(t, Contains(tt.want, v))
			}
		})
	}
}


var cases []struct {
	name string
	str map[string]string
	model interface{}
	data interface{}
} = []struct{
	name string
	str map[string]string
	model interface{}
	data interface{}
}{
	{
		name: "Worker",
		str: map[string]string{
			"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
			"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
			"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "active",
		},
		model: struct{
			ID       string `etcd:"id"`
			Platform string `etcd:"platform"`
			Status   string `etcd:"status"`
		} {},
		data: struct{
			ID       string `etcd:"id"`
			Platform string `etcd:"platform"`
			Status   string `etcd:"status"`
		} {
			ID: "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			Platform: "kis",
			Status: "active",
		},
	},
	{
		name: "Product",
		str: map[string]string{
			"product/test.goboolean.kor": "",
			"product/test.goboolean.kor/platform": "kis",
			"product/test.goboolean.kor/symbol": "goboolean",
			"product/test.goboolean.kor/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			"product/test.goboolean.kor/status": "onsubscribe",
		},
		model: struct{
			ID       string `etcd:"id"`
			Platform string `etcd:"platform"`
			Symbol   string `etcd:"symbol"`
			Worker 	 string `etcd:"worker"`
			Status   string `etcd:"status"`
		} {},
		data: struct{
			ID       string `etcd:"id"`
			Platform string `etcd:"platform"`
			Symbol   string `etcd:"symbol"`
			Worker 	 string `etcd:"worker"`
			Status   string `etcd:"status"`
		} {
			ID: "test.goboolean.kor",
			Platform: "kis",
			Symbol: "goboolean",
			Worker: "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
			Status: "onsubscribe",
		},
	},
	{
		name: "Nested Struct",
		str: map[string]string{
			"nested/mulmuri.dev": "",
			"nested/mulmuri.dev/detail": "",
			"nested/mulmuri.dev/detail/name": "goboolean",
			"nested/mulmuri.dev/detail/age": "1",
		},
		model: struct{
			ID string `etcd:"id"`
			Detail struct{
				Name string `etcd:"name"`
				Age int `etcd:"age"`
			} `etcd:"detail"`
		} {},
		data: struct{
			ID string `etcd:"id"`
			Detail struct{
				Name string `etcd:"name"`
				Age int `etcd:"age"`
			} `etcd:"detail"`
		} {
			ID: "mulmuri.dev",
			Detail: struct{
				Name string `etcd:"name"`
				Age int `etcd:"age"`
			} {
				Name: "goboolean",
				Age: 1,
			},
		},
	},
}

func deepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buf).Decode(dst)
}


func Test_Marshal(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			str, err := etcdutil.Mmarshal(tt.data)
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tt.str, str))
		})
	}
}


func Test_Unmarshal(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var input interface{}
			err := deepCopy(tt.data, &input)
			assert.NoError(t, err)

			err = etcdutil.Unmarshal(tt.str, tt.model)
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tt.data, input))
		})
	}
}


func TestCreateUuid(t *testing.T) {
	fmt.Println(uuid.New().String())
}