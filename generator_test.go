package idgenerator_test

import (
	"testing"

	idgenerator "github.com/ecumenos-social/id-generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func removeDuplicates(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func TestGenerator_Generate(t *testing.T) {
	g, err := idgenerator.New(100, 1)
	require.NoError(t, err)

	num := 1_000_000
	ids := make([]string, num)
	for i := 0; i < num; i++ {
		ids[i] = g.Generate().String()
	}

	assert.Equal(t, len(ids), len(removeDuplicates(ids)))
}
