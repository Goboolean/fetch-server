package etcdutil_test

import (
	"reflect"
	"testing"

	etcdutil "github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd/util"
	"github.com/stretchr/testify/assert"
)




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
