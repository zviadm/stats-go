package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	reg := newRegistry()

	c1, err := reg.DefineCounter("stats/test/c1", WithTags("tag"))
	require.NoError(t, err)
	c1.V(KV{"tag": "t1"}).Count(1)
	c1.V(KV{"tag": "t2"}).Count(2)
	c1.V(KV{"tag": "t3"}).Count(3)
	c1.V(KV{"tag": "t3"}).Count(1)

	export := reg.Export()
	require.Len(t, export, 1)
	require.EqualValues(t, CounterType, export["stats/test/c1"].Type)
	c1v := export["stats/test/c1"].F64s
	require.Len(t, c1v, 3)
	require.EqualValues(t, 4, c1v["t3"])

	g1, err := reg.DefineGauge("stats/test/g1")
	require.NoError(t, err)
	g1.V(nil).Set(1)
	g1.V(nil).Add(1)
	require.EqualValues(t, 2, g1.V(nil).Get())

	export = reg.Export()
	require.Len(t, export, 2)
	g1v := export["stats/test/g1"].F64s
	require.Len(t, g1v, 1)
	require.EqualValues(t, 2, g1v[""])
}
