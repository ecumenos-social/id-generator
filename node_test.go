package idgenerator_test

import (
	"testing"

	idgenerator "github.com/ecumenos-social/id-generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	_, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	if _, err := idgenerator.NewNode(50000000000000); err == nil {
		t.Fatalf("no error creating NewNode, %s", err)
	}
}

func TestID_Node(t *testing.T) {
	var node int64 = 110
	n, err := idgenerator.NewNode(node)
	require.NoError(t, err)
	id := n.Generate()

	assert.Equal(t, node, id.Node())
}

func TestGenerateDuplicateID(t *testing.T) {
	node, err := idgenerator.NewNode(1)
	require.NoError(t, err)

	var x, y idgenerator.ID
	for i := 0; i < 100_000; i++ {
		y = node.Generate()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

func TestRace(t *testing.T) {
	node, err := idgenerator.NewNode(1)
	require.NoError(t, err)

	go func() {
		for i := 0; i < 1_000_000_000; i++ {
			idgenerator.NewNode(1)
		}
	}()

	for i := 0; i < 4000; i++ {
		node.Generate()
	}
}

func TestPrintAll(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	id := node.Generate()

	t.Logf("Int64    : %#v", id.Int64())
	t.Logf("String   : %#v", id.String())
	t.Logf("Base2    : %#v", id.Base2())
	t.Logf("Base32   : %#v", id.Base32())
	t.Logf("Base36   : %#v", id.Base36())
	t.Logf("Base58   : %#v", id.Base58())
	t.Logf("Base64   : %#v", id.Base64())
	t.Logf("Bytes    : %#v", id.Bytes())
	t.Logf("IntBytes : %#v", id.IntBytes())
}

func BenchmarkGenerate(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}
