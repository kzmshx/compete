package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	u := New(7) // {0}, {1}, {2}, {3}, {4}, {5}, {6}
	assert.False(t, u.IsSameSet(1, 2))
	assert.Equal(t, 1, u.Size(1))

	u.Union(1, 2) // {0}, {1, 2}, {3}, {4}, {5}, {6}
	assert.True(t, u.IsSameSet(1, 2))
	assert.Equal(t, 2, u.Size(2))

	u.Union(2, 3) // {0}, {1, 2, 3}, {4}, {5}, {6}
	assert.True(t, u.IsSameSet(2, 3))
	assert.True(t, u.IsSameSet(1, 3))
	assert.Equal(t, 3, u.Size(2))
	assert.Equal(t, 3, u.Size(3))

	u.Union(4, 5) // {0}, {1, 2, 3}, {4, 5}, {6}
	assert.True(t, u.IsSameSet(4, 5))
	assert.False(t, u.IsSameSet(2, 5))
	assert.Equal(t, 3, u.Size(3))
	assert.Equal(t, 2, u.Size(5))

	u.Union(1, 6) // {0}, {1, 2, 3, 6}, {4, 5}
	assert.True(t, u.IsSameSet(1, 6))
	assert.Equal(t, 4, u.Size(6))
}
