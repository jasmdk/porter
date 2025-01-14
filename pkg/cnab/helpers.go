package cnab

import (
	"io/ioutil"
	"testing"

	"github.com/cnabio/cnab-go/bundle"
	"github.com/stretchr/testify/require"
)

func ReadTestBundle(t *testing.T, path string) ExtendedBundle {
	bunD, err := ioutil.ReadFile(path)
	require.NoError(t, err, "ReadFile failed for %s", path)

	bun, err := bundle.Unmarshal(bunD)
	require.NoError(t, err, "Unmarshal failed for bundle at %s", path)

	return NewBundle(*bun)
}
