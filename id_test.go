package idgenerator_test

import (
	"bytes"
	"testing"

	idgenerator "github.com/ecumenos-social/id-generator"
	"github.com/stretchr/testify/require"
)

func TestInt64(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.Int64()

	pID := idgenerator.ParseInt64(i)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	mi := int64(1116766490855473152)
	pID = idgenerator.ParseInt64(mi)
	if pID.Int64() != mi {
		t.Fatalf("pID %v != mi %v", pID.Int64(), mi)
	}
}

func TestString(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	si := oID.String()

	pID, err := idgenerator.ParseString(si)
	require.NoError(t, err)

	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `1116766490855473152`
	_, err = idgenerator.ParseString(ms)
	require.NoError(t, err)

	ms = `1112316766490855473152`
	_, err = idgenerator.ParseString(ms)
	if err == nil {
		t.Fatalf("no error parsing %s", ms)
	}
}

func TestBase2(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.Base2()

	pID, err := idgenerator.ParseBase2(i)
	require.NoError(t, err)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `111101111111101110110101100101001000000000000000000000000000`
	_, err = idgenerator.ParseBase2(ms)
	require.NoError(t, err)

	ms = `1112316766490855473152`
	_, err = idgenerator.ParseBase2(ms)
	if err == nil {
		t.Fatalf("no error parsing %s", ms)
	}
}

func TestBase32(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	for i := 0; i < 100; i++ {
		sf := node.Generate()
		b32i := sf.Base32()
		psf, err := idgenerator.ParseBase32([]byte(b32i))
		require.NoError(t, err)
		if sf != psf {
			t.Fatal("Parsed does not match String.")
		}
	}
}

func TestBase36(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.Base36()

	pID, err := idgenerator.ParseBase36(i)
	require.NoError(t, err)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `8hgmw4blvlkw`
	_, err = idgenerator.ParseBase36(ms)
	require.NoError(t, err)

	ms = `68h5gmw443blv2lk1w`
	_, err = idgenerator.ParseBase36(ms)
	if err == nil {
		t.Fatalf("no error parsing, %s", err)
	}
}

func TestBase58(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		sf := node.Generate()
		b58 := sf.Base58()
		psf, err := idgenerator.ParseBase58([]byte(b58))
		require.NoError(t, err)
		if sf != psf {
			t.Fatal("Parsed does not match String.")
		}
	}
}

func TestBase64(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.Base64()

	pID, err := idgenerator.ParseBase64(i)
	require.NoError(t, err)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `MTExNjgxOTQ5NDY2MDk5NzEyMA==`
	_, err = idgenerator.ParseBase64(ms)
	require.NoError(t, err)

	ms = `MTExNjgxOTQ5NDY2MDk5NzEyMA`
	_, err = idgenerator.ParseBase64(ms)
	if err == nil {
		t.Fatalf("no error parsing, %s", err)
	}
}

func TestBytes(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.Bytes()

	pID, err := idgenerator.ParseBytes(i)
	require.NoError(t, err)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := []byte{0x31, 0x31, 0x31, 0x36, 0x38, 0x32, 0x31, 0x36, 0x37, 0x39, 0x35, 0x37, 0x30, 0x34, 0x31, 0x39, 0x37, 0x31, 0x32}
	_, err = idgenerator.ParseBytes(ms)
	require.NoError(t, err)

	ms = []byte{0xFF, 0xFF, 0xFF, 0x31, 0x31, 0x31, 0x36, 0x38, 0x32, 0x31, 0x36, 0x37, 0x39, 0x35, 0x37, 0x30, 0x34, 0x31, 0x39, 0x37, 0x31, 0x32}
	_, err = idgenerator.ParseBytes(ms)
	if err == nil {
		t.Fatalf("no error parsing, %#v", err)
	}
}

func TestIntBytes(t *testing.T) {
	node, err := idgenerator.NewNode(0)
	require.NoError(t, err)

	oID := node.Generate()
	i := oID.IntBytes()

	pID := idgenerator.ParseIntBytes(i)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := [8]uint8{0xf, 0x7f, 0xc0, 0xfc, 0x2f, 0x80, 0x0, 0x0}
	mi := int64(1116823421972381696)
	pID = idgenerator.ParseIntBytes(ms)
	if pID.Int64() != mi {
		t.Fatalf("pID %v != mi %v", pID.Int64(), mi)
	}
}

func TestMarshalJSON(t *testing.T) {
	id := idgenerator.ID(13587)
	expected := "\"13587\""

	bytes, err := id.MarshalJSON()
	require.NoError(t, err)

	if string(bytes) != expected {
		t.Fatalf("Got %s, expected %s", string(bytes), expected)
	}
}

func TestMarshalsIntBytes(t *testing.T) {
	id := idgenerator.ID(13587).IntBytes()
	expected := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x35, 0x13}
	if !bytes.Equal(id[:], expected) {
		t.Fatalf("Expected ID to be encoded as %v, got %v", expected, id)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tt := []struct {
		json        string
		expectedID  idgenerator.ID
		expectedErr error
	}{
		{`"13587"`, 13587, nil},
		{`1`, 0, idgenerator.JSONSyntaxError{[]byte(`1`)}},
		{`"invalid`, 0, idgenerator.JSONSyntaxError{[]byte(`"invalid`)}},
	}

	for _, tc := range tt {
		var id idgenerator.ID
		err := id.UnmarshalJSON([]byte(tc.json))
		if (tc.expectedErr != nil && err != nil) && (err.Error() != tc.expectedErr.Error()) {
			t.Fatalf("Expected to get error '%s' decoding JSON, but got '%s'", tc.expectedErr, err)
		}
		if (tc.expectedErr != nil && err == nil) || (tc.expectedErr == nil && err != nil) {
			t.Fatalf("Expected to get error '%s' decoding JSON, but got '%s'", tc.expectedErr, err)
		}

		if id != tc.expectedID {
			t.Fatalf("Expected to get ID '%s' decoding JSON, but got '%s'", tc.expectedID, id)
		}
	}
}

func BenchmarkParseBase32(b *testing.B) {
	node, _ := idgenerator.NewNode(10)
	sf := node.Generate()
	b32i := sf.Base32()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		idgenerator.ParseBase32([]byte(b32i))
	}
}

func BenchmarkBase32(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)
	sf := node.Generate()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sf.Base32()
	}
}
func BenchmarkParseBase58(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)

	sf := node.Generate()
	b58 := sf.Base58()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		idgenerator.ParseBase58([]byte(b58))
	}
}
func BenchmarkBase58(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)

	sf := node.Generate()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sf.Base58()
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)

	id := node.Generate()
	bytes, err := id.MarshalJSON()
	require.NoError(b, err)

	var id2 idgenerator.ID

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = id2.UnmarshalJSON(bytes)
		require.NoError(b, err)
	}
}

func BenchmarkMarshal(b *testing.B) {
	node, err := idgenerator.NewNode(10)
	require.NoError(b, err)
	id := node.Generate()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err = id.MarshalJSON()
		require.NoError(b, err)
	}
}

func TestParseBase32(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    idgenerator.ID
		wantErr bool
	}{
		{
			name:    "ok",
			arg:     "b8wjm1zroyyyy",
			want:    1427970479175499776,
			wantErr: false,
		},
		{
			name:    "capital case is invalid encoding",
			arg:     "B8WJM1ZROYYYY",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "l is not allowed",
			arg:     "b8wjm1zroyyyl",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "v is not allowed",
			arg:     "b8wjm1zroyyyv",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "2 is not allowed",
			arg:     "b8wjm1zroyyy2",
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idgenerator.ParseBase32([]byte(tt.arg))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBase32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBase32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBase58(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    idgenerator.ID
		wantErr bool
	}{
		{
			name:    "ok",
			arg:     "4jgmnx8Js8A",
			want:    1428076403798048768,
			wantErr: false,
		},
		{
			name:    "0 not allowed",
			arg:     "0jgmnx8Js8A",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "I not allowed",
			arg:     "Ijgmnx8Js8A",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "O not allowed",
			arg:     "Ojgmnx8Js8A",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "l not allowed",
			arg:     "ljgmnx8Js8A",
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idgenerator.ParseBase58([]byte(tt.arg))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBase58() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBase58() got = %v, want %v", got, tt.want)
			}
		})
	}
}
