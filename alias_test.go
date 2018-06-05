package aliastest

import (
	"bytes"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yurishkuro/gogoproto-custom-type/jsontest"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		obj *jsontest.Simple
		out string
	}{
		{&jsontest.Simple{Val: 0x42}, `{"val":"42"}`},
	}
	for _, test := range tests {
		out := new(bytes.Buffer)
		require.NoError(t, new(jsonpb.Marshaler).Marshal(out, test.obj))
		assert.Equal(t, test.out, out.String())
		var val jsontest.Simple
		require.NoError(t, jsonpb.Unmarshal(bytes.NewReader([]byte(test.out)), &val))
		assert.Equal(t, *test.obj, val)
	}
}