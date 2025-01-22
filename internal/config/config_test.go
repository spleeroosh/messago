//go:build !integration

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("APP_NAME", "messago"); err != nil {
		require.NoError(t, err)
	}

	conf, _ := GetConfig()
	require.Equal(t, "messago", conf.App.Name)

	os.Clearenv()
	conf, _ = GetConfig()
	require.Equal(t, "messago", conf.App.Name)
}
