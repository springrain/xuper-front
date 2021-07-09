package clixchain

import (
	"testing"
)

func TestGroup(t *testing.T) {
	groupA := group{
		groupID:    "xuper",
		admin:      []string{"A", "B"},
		identities: []string{"A", "B", "C"},
	}
	set := make(map[string]bool)
	result := groupA.GetAddrs(set)
	for _, a := range result {
		if a != "A" && a != "B" && a != "C" {
			t.Errorf("group GetAddrs error")
		}
	}
	groupB := group{
		groupID:    "xuper",
		admin:      []string{"A", "B"},
		identities: []string{"D"},
	}
	result = groupB.GetAddrs(set)
	if len(result) != 1 {
		t.Errorf("group GetAddrs error")
	}
}

func TestGroupCache(t *testing.T) {
	gc := groupCache{
		value: []string{"A", "B", "C"},
	}
	gc.put([]string{"D"})
	resultB := gc.get()
	if len(resultB) != 1 {
		t.Errorf("groupCache Get error, result = %v", resultB)
	}
}