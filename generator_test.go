package idgenerator_test

import (
	"fmt"
	"testing"
	"time"

	idgenerator "github.com/ecumenos-social/id-generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func removeDuplicates[T string | uint64 | idgenerator.ID](strSlice []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func TestGenerator_Generate(t *testing.T) {
	var (
		top int64 = 10
		low int64 = 10
	)
	g, err := idgenerator.New(&idgenerator.NodeID{
		Top: top,
		Low: low,
	})
	require.NoError(t, err)

	num := 100_000
	ids := make([]idgenerator.ID, num)
	start := time.Now()
	for i := 0; i < num; i++ {
		id := g.Generate()
		assert.Equal(t, top, id.NodeID().Top)
		assert.Equal(t, low, id.NodeID().Low)
		ids[i] = id
	}

	assert.Equal(t, len(ids), len(removeDuplicates(ids)))
	fmt.Println(time.Since(start))
}
