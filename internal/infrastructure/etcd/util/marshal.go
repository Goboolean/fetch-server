package etcdutil

import (
	"strings"
)

// GroupBy groups the list by the distinquisher that appears just after the prefix.
// It assums that list is sorted by key.
func GroupBy(list map[string]string, prefix string) []map[string]string {
	var m = make(map[string]map[string]string)
	var result []map[string]string

	for k, v := range list {
		p := strings.Split(strings.TrimPrefix(k, prefix), "/")[1]
		if _, ok := m[p]; !ok {
			m[p] = make(map[string]string)
		}
		m[p][k] = v
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}


func Mmarshal(i Model) (map[string]string, error) {
	return nil, nil
}

func Unmarshal(str map[string]string, i Model) error {
	return nil
}