package manifest_test

import (
	"testing"

	"github.com/convox/rack/pkg/helpers"
	"github.com/convox/rack/pkg/manifest"
	"github.com/stretchr/testify/require"
)

func TestManifestLoadClevy(t *testing.T) {
	n := &manifest.Manifest{
		Services: manifest.Services{
			manifest.Service{
				Name: "q-delete-intent",
				Build: manifest.ServiceBuild{
					Manifest: "Dockerfile",
					Path:     "./resources/queues",
				},
				Environment: []string{
					"QUEUE_NAME=delete-intent",
					"POEPOE=giveme-somestuff",
				},
				Health: manifest.ServiceHealth{
					Grace:    5,
					Interval: 5,
					Path:     "/",
					Timeout:  4,
				},
				Scale: manifest.ServiceScale{
					Count:  manifest.ServiceScaleCount{Min: 1, Max: 1},
					Cpu:    256,
					Memory: 512,
				},
				Sticky: true,
			},
		},
	}

	attrs := []string{
		"services",
		"services.q-delete-intent",
		"services.q-delete-intent.build",
		"services.q-delete-intent.environment",
	}

	env := map[string]string{ "INTERPOLATE": "somestuff"}

	n.SetAttributes(attrs)
	n.SetEnv(env)

	// env processing that normally happens as part of load
	require.NoError(t, n.CombineEnv())
	require.NoError(t, n.ValidateEnv())

	m, err := testdataManifest("clevy", env)
	require.NoError(t, err)
	require.Equal(t, n, m)
}

func testdataManifest(name string, env map[string]string) (*manifest.Manifest, error) {
	data, err := helpers.Testdata(name)
	if err != nil {
		return nil, err
	}

	return manifest.Load(data, env)
}
