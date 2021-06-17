package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/docker/pkg/ioutils"
	"github.com/stretchr/testify/assert"
)

func TestService_ClonePrivateRepository_GitHub(t *testing.T) {
	ensureIntegrationTest(t)

	pat := getRequiredValue(t, "GITHUB_PAT")
	username := getRequiredValue(t, "GITHUB_USERNAME")
	service := NewService()

	dst, err := ioutils.TempDir("", "clone")
	assert.NoError(t, err)
	defer os.RemoveAll(dst)

	repositoryUrl := "https://github.com/portainer/private-test-repository.git"
	err = service.CloneRepository(dst, repositoryUrl, "refs/heads/main", username, pat)
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(dst, "README.md"))
}
