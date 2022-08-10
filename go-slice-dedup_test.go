package goslicededup

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeDuplicateString(t *testing.T) {
	list := []string{"a", "b", "a", "c"}
	listUnique := DeDuplicate(list)
	assert.Equal(t, 1, strings.Count(strings.Join(listUnique, ""), "a"), "should find 1 'a' in joined & deduplicated list of abac")
}

func TestDeDuplicateInt(t *testing.T) {
	list := []int{1, 2, 1, 3}
	listUnique := DeDuplicate(list)
	assert.Equal(t, []int{1, 2, 3}, listUnique, "should find 1,2,3 from deduplicated list of 1,2,1,3")
}
