package Solutions

import (
	"reflect"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	tests := []struct {
		name          string
		methods       []string
		args          [][]int
		expected      []interface{}
	}{
		{
			name:    "test01",
			methods: []string{"RecentCounter", "ping", "ping", "ping", "ping"},
			args:    [][]int{nil, {1}, {100}, {3001}, {3002}},
			expected: []interface{}{nil, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj RecentCounter
			var result []interface{}

			for i, method := range tt.methods {
				switch method {
				case "RecentCounter":
					obj = Constructor()
					result = append(result, nil)
				case "ping":
					res := obj.Ping(tt.args[i][0])
					result = append(result, res)
				default:
					t.Errorf("Unknown method: %s", method)
				}
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
