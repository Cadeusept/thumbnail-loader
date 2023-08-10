package utils

import "testing"

func TestRemoveDuplicateStr(t *testing.T) {
	s := []string{"1", "2", "3", "1"}

	s = RemoveDuplicateStr(s)
	if s[0] != "1" || s[1] != "2" || s[2] != "3" {
		t.Errorf("Not expected behaviour: wanted [%d %d %d] . got %s", 1, 2, 3, s)
	}
}
