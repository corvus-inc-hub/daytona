// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package docker

import (
	"testing"

	"github.com/daytonaio/runner/cmd/runner/config"
	"github.com/daytonaio/runner/pkg/api/dto"
)

func TestGetContainerHostConfigSetsSandboxSharedMemory(t *testing.T) {
	t.Setenv("DAYTONA_API_URL", "http://localhost")
	t.Setenv("DAYTONA_RUNNER_TOKEN", "test-token")
	t.Setenv("RUNNER_DOMAIN", "127.0.0.1")
	if _, err := config.GetConfig(); err != nil {
		t.Fatalf("initialize runner config: %v", err)
	}

	client := &DockerClient{
		daemonPath:             "/usr/local/bin/daytona-daemon",
		resourceLimitsDisabled: true,
	}

	hostConfig, err := client.getContainerHostConfig(dto.CreateSandboxDTO{}, nil, nil)
	if err != nil {
		t.Fatalf("getContainerHostConfig returned an error: %v", err)
	}

	if hostConfig.ShmSize != sandboxShmSizeBytes {
		t.Fatalf("shared memory size = %d, want %d", hostConfig.ShmSize, sandboxShmSizeBytes)
	}
}
