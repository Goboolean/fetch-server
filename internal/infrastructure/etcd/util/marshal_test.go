package etcdutil_test

import (
	"fmt"
	"reflect"
	"testing"

	etcdutil "github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)



func Test_GroupBy(t *testing.T) {

	type args struct {
		list   map[string]string
		prefix string
	}

	tests := []struct {
		name string
		args args
		want []map[string]string
	} {
		{
			name: "Worker",
			args: args{
				list: map[string]string{
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886": "",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/platform": "kis",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/status": "active",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "waiting",
				},
				prefix: "worker",
			},
			want: []map[string]string{
				{
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886": "",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/platform": "kis",
					"worker/b9992d7b-a926-483a-84f8-bbc05dee7886/status": "active",
				},
				{
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "waiting",
				},
			},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := etcdutil.GroupBy(tt.args.list, tt.args.prefix)
			assert.Equal(t, tt.want, got)
		})
	}
}


func Test_Marshal(t *testing.T) {

	type args struct {
		str  map[string]string
		i    interface{} // struct initialized with default values
		want interface{} // same type of struct with args i with expected values
	}

	tests := []struct {
		name string
		args args
	} {
		{
			name: "Worker",
			args: args{
				str: map[string]string{
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d": "",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/platform": "kis",
					"worker/9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d/status": "active",
				},
				i: struct{
					ID       string `etcd:"id"`
					Platform string `etcd:"platform"`
					Status   string `etcd:"status"`
				} {},
				want: struct{
					ID       string `etcd:"id"`
					Platform string `etcd:"platform"`
					Status   string `etcd:"status"`
				} {
					ID: "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					Platform: "kis",
					Status: "active",
				},
			},
		},
		{
			name: "Product",
			args: args{
				str: map[string]string{
					"product/test.goboolean.kor": "",
					"product/test.goboolean.kor/platform": "kis",
					"product/test.goboolean.kor/symbol": "goboolean",
					"product/test.goboolean.kor/worker": "9cf226f7-4ee8-4a5c-9d2f-6d7c74f6727d",
					"product/test.goboolean.kor/status": "onsubscribe",
				},
				i: struct{
					ID       string `etcd:"id"`
					Platform string `etcd:"platform"`
					Symbol   string `etcd:"symbol"`
					Worker 	 string `etcd:"worker"`
					Status   string `etcd:"status"`
				} {},
				want: struct{
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
		},
		{
			name: "Nested Struct",
			args: args{
				str: map[string]string{
					"nested/mulmuri.dev": "",
					"nested/mulmuri.dev/detail": "",
					"nested/mulmuri.dev/detail/name": "goboolean",
					"nested/mulmuri.dev/detail/age": "1",
				},
				i: struct{
					ID string `etcd:"id"`
					Detail struct{
						Name string `etcd:"name"`
						Age int `etcd:"age"`
					} `etcd:"detail"`
				} {},
				want: struct{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := etcdutil.Marshal(tt.args.str, tt.args.i)
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tt.args.want, tt.args.i))
		})
	}
}


func TestCreateUuid(t *testing.T) {
	fmt.Println(uuid.New().String())
}