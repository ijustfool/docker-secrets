package secrets

import (
	"fmt"
	"io/ioutil"
	"os"
)

// DockerSecrets contains secrets
type DockerSecrets struct {
	secretsDir string
	secrets    map[string]string
}

// NewDockerSecrets creates an instance of DockerSecrets
// The secretsDir argument has a default value of: /run/secrets
func NewDockerSecrets(secretsDir string) (*DockerSecrets, error) {
	if secretsDir == "" {
		secretsDir = "/run/secrets"
	}
	dockerSecrets := &DockerSecrets{secretsDir: secretsDir}
	err := dockerSecrets.readAll()
	return dockerSecrets, err
}

// GetAll returns all secrets from memory
func (ds *DockerSecrets) GetAll() map[string]string {
	return ds.secrets
}

// Get returns one secret by secretName
func (ds *DockerSecrets) Get(secretName string) (string, error) {
	if _, ok := ds.secrets[secretName]; !ok {
		return "", fmt.Errorf("secret not exsist: %s", secretName)
	}
	return ds.secrets[secretName], nil
}

// Reads all secrets on the specified path in the secretsDir
func (ds *DockerSecrets) readAll() error {
	secretsDir := ds.getDir()
	err := isDir(secretsDir)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(secretsDir)
	if err != nil {
		return err
	}

	ds.secrets = make(map[string]string)
	for _, file := range files {
		err := ds.read(file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Reads a secret from file
func (ds *DockerSecrets) read(file string) error {
	buf, err := ioutil.ReadFile(ds.getDir() + "/" + file)
	if err != nil {
		return err
	}
	ds.secrets[file] = string(buf)
	return nil
}

// Returns the secretsDir
func (ds *DockerSecrets) getDir() string {
	return ds.secretsDir
}

// Checks if the given path is a directory
func isDir(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("is not a directory: %s", path)
	}
	return nil
}
