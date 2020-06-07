package multierror

import (
	"errors"
	"go/types"
	"io"
	"net"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiError_Error(t *testing.T) {
	require.Equal(t, "multiple errors:\n\t* a\n\t* b\n", (&SyncSlice{Slice: Slice{errors.New("a"), errors.New("b")}}).Error())
}

func TestMultiError_As(t *testing.T) {
	err := &SyncSlice{Slice: Slice{types.Error{Msg: "unit-test"}, &net.OpError{}}}

	var typesErr types.Error
	require.Equal(t, true, errors.As(err, &typesErr))
	require.Equal(t, typesErr, err.Slice[0])

	var urlErr *url.Error
	require.Equal(t, false, errors.As(err, &urlErr))
}

func TestMultiError_Is(t *testing.T) {
	require.Equal(t, true, errors.Is(&SyncSlice{Slice: Slice{io.ErrClosedPipe, io.ErrUnexpectedEOF}}, io.ErrClosedPipe))
	require.Equal(t, false, errors.Is(&SyncSlice{Slice: Slice{io.ErrClosedPipe, io.ErrUnexpectedEOF}}, io.EOF))
}
