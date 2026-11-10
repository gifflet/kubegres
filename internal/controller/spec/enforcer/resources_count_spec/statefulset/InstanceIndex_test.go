package statefulset

import (
	"testing"

	"reactive-tech.io/kubegres/internal/controller/states/statefulset"
)

func TestNextInstanceIndex(t *testing.T) {
	deployed := func(indexes ...int32) []statefulset.StatefulSetWrapper {
		var wrappers []statefulset.StatefulSetWrapper
		for _, index := range indexes {
			wrappers = append(wrappers, statefulset.StatefulSetWrapper{InstanceIndex: index})
		}
		return wrappers
	}
	tests := []struct {
		name       string
		lastStatus int32
		deployed   []statefulset.StatefulSetWrapper
		expected   int32
	}{
		{"nothing deployed yet", 0, nil, 1},
		{"status in sync with the cluster", 2, deployed(1, 2), 3},
		{"status behind the cluster (status update was rejected)", 1, deployed(1, 2), 3},
		{"status ahead of the cluster (StatefulSet deleted)", 5, deployed(1, 2), 6},
	}
	for _, test := range tests {
		if got := nextInstanceIndex(test.lastStatus, test.deployed); got != test.expected {
			t.Errorf("%s: expected %d, got %d", test.name, test.expected, got)
		}
	}
}
