package bufferpool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCap(t *testing.T) {
	require := require.New(t)

	for size := 1; size <= 8; size++ {
		buf := Get(size)
		require.Equal(8, cap(buf))
		Put(buf)
	}
}

func TestLen(t *testing.T) {
	require := require.New(t)

	for size := 1; size <= 8; size++ {
		buf := Get(size)
		require.Equal(size, len(buf))
		Put(buf)
	}
}

func TestBufReuse(t *testing.T) {
	require := require.New(t)

	buf1 := Get(16)
	buf1[0] = 1
	require.Equal(byte(1), buf1[0])
	Put(buf1)

	buf2 := Get(16)
	require.Equal(buf1, buf2)
	require.Equal(byte(1), buf2[0])
	buf2[1] = 2
	require.Equal(byte(2), buf1[1])
	Put(buf2)
}
