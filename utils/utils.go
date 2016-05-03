// utils.go

package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
)

const CONFIG_DEFAULT_LOCATION = "/etc/gitawsm"
const CONFIG_FILE = "config.json"
const BRANCHES_FILE = "branches.json"

func getConfigUserLocation(userHome string) string {
	return fmt.Sprintf("%s/.gitawsm", userHome)
}

func Print(message string) {
	fmt.Fprintln(os.Stdout, message)
}

func Matches(str string, pattern string) bool {
	// TODO do some preprocessing on the pattern to transform it from bash regex to re2
	r := regexp.MustCompile(fmt.Sprintf("^%s$", pattern))
	match := r.MatchString(str)
	return match
}

func dirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true, nil
	} else if os.IsNotExist(err) || !stat.IsDir() {
		return false, nil
	}
	return false, err
}

func getConfDir() string {
	// first check user home
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	path := getConfigUserLocation(user.HomeDir)
	exists, err := dirExists(path)
	if err != nil {
		Print(fmt.Sprintf("Error looking up directory in user home : %s", err.Error()))
	}

	if exists {
		return path
	}

	path = CONFIG_DEFAULT_LOCATION
	exists, err = dirExists(path)
	if err != nil {
		Print(fmt.Sprintf("Error looking up default directory : %s", err.Error()))
	}

	if exists {
		return path
	}

	log.Fatal("Could not find gitawsm configuration at any of the possible locations. Check installation.")
	return ""
}

func ReadConfig() map[string][]string {
	config := make(map[string][]string)

	// read the configuration for gitawsm
	confDirPath := getConfDir()
	configFile, err := os.Open(fmt.Sprintf("%s/%s", confDirPath, CONFIG_FILE))
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not find gitawsm configuration file in the config folder at %q. Check installation.", confDirPath))
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Unable to read configuration.", err)
	}

	return config
}

func ReadBranches() map[string][]string {
	branches := make(map[string][]string)

	// read branches file
	confDirPath := getConfDir()
	branchesFilePath := fmt.Sprintf("%s/%s", confDirPath, BRANCHES_FILE)
	branchesFile, err := os.Open(branchesFilePath)
	if err != nil {
		// branches file does exists
		return branches
	}
	defer branchesFile.Close()

	decoder := json.NewDecoder(branchesFile)
	err = decoder.Decode(&branches)
	if err != nil {
		log.Fatal("Unable to read branches.", err)
	}

	if branches == nil {
		branches = make(map[string][]string)
	}
	return branches
}

func WriteBranches(b map[string][]string) error {
	// open branches file in write mode
	confDirPath := getConfDir()
	branchesFilePath := fmt.Sprintf("%s/%s", confDirPath, BRANCHES_FILE)
	branchesFile, err := os.OpenFile(branchesFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Unable to open branches file for writing.")
	}

	j, jerr := json.MarshalIndent(b, "", "\t")
	if jerr != nil {
		log.Fatal("Unable to marshal branches to json.")
	}

	_, werr := branchesFile.Write(j)
	if werr != nil {
		log.Fatal("Unable to write branches to file.")
	}

	return nil
}

func GetWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory!")
	}
	return wd
}
