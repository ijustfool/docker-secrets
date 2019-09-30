package secrets_test

import (
	"testing"

	"github.com/ijustfool/docker-secrets"
)

const (
	secretDir = "test-data"
	userVal   = "root"
	passVal   = "myPass"
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
	checkKey(t, dockerSecrets, "user", userVal)
	checkKey(t, dockerSecrets, "pass", passVal)
}

func checkKey(t *testing.T, dockerSecrets *secrets.DockerSecrets, key, expectedValue string) {
	value, err := dockerSecrets.Get(key)
	if err != nil {
		t.Errorf("dockerSecrets.Get(\"%s\"): %s", key, err.Error())
		return
	}
	if expectedValue != value {
		t.Errorf("dockerSecrets.Get(\"%s\") = `%s`, expected: `%s`", key, value, expectedValue)
	}

}
