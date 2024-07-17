package client

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	tests := []struct {
		name          string
		spec          []byte
		err           *plugin.TestConnError
		clientBuilder func() (plugin.Client, error)
	}{
		{
			name: "ok",
			spec: []byte(`{"addresses": ["test"]}`),
			err:  nil,
		},
		{
			name: "invalid spec",
			spec: []byte(`{"addresses": 12}`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "invalid spec JSON",
			spec: []byte(`{"addresses"`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "unreachable",
			spec: []byte(`{"addresses": ["test"]}`),
			err:  &plugin.TestConnError{Code: codeUnreachable},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errUnreachable
			},
		},
		{
			name: "unauthorized",
			spec: []byte(`{"addresses": ["test"]}`),
			err:  &plugin.TestConnError{Code: codeUnauthorized},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errUnauthorized
			},
		},
		{
			name: "connection failed with other error",
			spec: []byte(`{"addresses": ["test"]}`),
			err:  &plugin.TestConnError{Code: codeConnectionFailed},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errors.New("test")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.clientBuilder == nil {
				tt.clientBuilder = func() (plugin.Client, error) {
					return &Client{}, nil
				}
			}
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				sp := &Spec{}
				if err := json.Unmarshal(specBytes, sp); err != nil {
					return nil, errInvalidSpec
				}
				sp.SetDefaults()
				if err := sp.Validate(); err != nil {
					return nil, errInvalidSpec
				}
				return tt.clientBuilder()
			})
			err := tester(context.Background(), zerolog.Nop(), tt.spec)
			if tt.err == nil {
				require.NoError(t, err)
				return
			}
			var expErr *plugin.TestConnError
			require.ErrorAs(t, err, &expErr)
			require.Equal(t, tt.err.Code, err.(*plugin.TestConnError).Code)
		})
	}
}
