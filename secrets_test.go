package secrets_test

import (
	"testing"

	"github.com/ijustfool/docker-secrets"
)

const (
	secretDir = "test-data"
	userVal   = "root"
	passVal   = "password"
)

func TestNewDockerSecrets(t *testing.T) {
	_, err := secrets.NewDockerSecrets(secretDir)
	if err != nil {
		t.Errorf("NewDockerSecrets(): %v", err)
		return
	}
}

func TestDockerSecrets_GetAll(t *testing.T) {
	dockerSecrets, _ := secrets.NewDockerSecrets(secretDir)
	allSecrets := dockerSecrets.GetAll()
	if _, ok := allSecrets["user"]; !ok {
		t.Errorf("allSecrets[\"user\"] not exsist, expected: `%v`", userVal)
	}
	if _, ok := allSecrets["pass"]; !ok {
		t.Errorf("allSecrets[\"pass\"] not exsist, expected: `%v`", passVal)
	}
}

func TestDockerSecrets_Get(t *testing.T) {
	dockerSecrets, _ := secrets.NewDockerSecrets(secretDir)
	val, err := dockerSecrets.Get("user")
	if err != nil {
		t.Errorf("dockerSecrets.Get(\"user\"): %v", err)
		return
	}
	if val != userVal {
		t.Errorf("dockerSecrets.Get(\"user\") = `%v`, expected: `%v`", val, userVal)
	}
}
