package secrets_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ijustfool/docker-secrets"
)

const (
	secretDir = "test-data"
	userVal   = "root"
	passVal   = "myPass"
)

type testSecrets struct {
	User     string
	Password string `mapstructure:"pass"`
}

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

func TestDockerSecrets_Unmarshal(t *testing.T) {
	dockerSecrets, _ := secrets.NewDockerSecrets(secretDir)
	testSecrets := testSecrets{}
	err := dockerSecrets.Unmarshal(&testSecrets)
	if err != nil {
		t.Errorf("dockerSecrets.Unmarshal(): %v", err)
		return
	}
	if testSecrets.User != userVal {
		t.Errorf("testSecrets.User = `%v`, expected: `%v`", testSecrets.User, userVal)
	}
	if testSecrets.Password != passVal {
		t.Errorf("testSecrets.Password = `%v`, expected: `%v`", testSecrets.Password, passVal)
	}
}

func TestReplaceInFile(t *testing.T) {
	b, err := ioutil.ReadFile(secretDir + "/config.json")
	if err != nil {
		t.Errorf("ReplaceInFile(): %v", err)
		return
	}

	b, err = secrets.ReplaceInFile(b)
	if err != nil {
		t.Errorf("ReplaceInFile(): %v", err)
		return
	}

	testSecrets := testSecrets{}
	err = json.Unmarshal(b, &testSecrets)
	if err != nil {
		t.Errorf("ReplaceInFile(): %v", err)
		return
	}

	if testSecrets.User != userVal {
		t.Errorf("testSecrets.User = `%v`, expected: `%v`", testSecrets.User, userVal)
	}
	if testSecrets.Password != passVal {
		t.Errorf("testSecrets.Password = `%v`, expected: `%v`", testSecrets.Password, passVal)
	}
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
