package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"

	"os/exec"
)

// The filename of the configuration file.
const confFilename = "fastmailsync.json"

type ConfigurationError struct {
	Message string
	Err     error
}

func (ce *ConfigurationError) Error() string {
	return fmt.Sprintf("%s: %v", ce.Message, ce.Err)
}

func NewConfigurationError(message string, err error) *ConfigurationError {
	return &ConfigurationError{
		Message: message,
		Err:     err,
	}
}

type Configuration struct {
	AccountId string
	Username  string
	Password  string
	MailDir   string
}

func ReadConfiguration() Configuration {
	configHomePath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln("Cannot read user config directory:", err)
	}

	confDir := path.Join(configHomePath, "fastmailsync")
	ensurePath(confDir)

	confFile := path.Join(confDir, confFilename)
	file, err := os.Open(confFile)
	if err != nil {
		return Configuration{}
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln(err)
	}
	configuration.Password, err = readPassword(configuration.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return configuration
}

func ensurePath(path string) {
	info, err := os.Stat(path)
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
			os.MkdirAll(path, 0700)
			return
		} else {
			log.Fatalln("Could not create directory for configuration:", err)
		}
	}
	if !info.IsDir() {
		log.Fatalln("Configuration path does not point to a directory.")
	}
}

func createDefaultFile(path string) error {
	return os.WriteFile(path, []byte{}, 0600)
}

func readPassword(passEntry string) (string, error) {
	cmd := exec.Command("pass", passEntry)
	password, err := cmd.Output()
	if err != nil {
		return "", NewConfigurationError("Could not read password with pass", err)
	}
	return string(password), nil
}
