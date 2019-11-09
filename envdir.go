package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type EnvVar struct {
	Name  string
	Value string
}

func isAllUpperCase(s string) bool {
	return s == strings.ToUpper(s)
}

func isEnvVariableFile(info os.FileInfo) bool {
	name := info.Name()
	if !info.IsDir() && isAllUpperCase(name) && strings.Contains(name, "_") && info.Size() <= 1024 {
		return true
	}
	return false
}

func GetEnvVarsFromDirectory(pathToEnvDir string) []EnvVar {
	envVars := make([]EnvVar, 0)

	_ = filepath.Walk(pathToEnvDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if isEnvVariableFile(info) {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return nil
			}
			envVars = append(envVars, EnvVar{
				Name:  info.Name(),
				Value: string(content),
			})
		}
		return nil
	})

	return envVars
}

func convertEnvVarsToStrings(envVars []EnvVar) []string {
	envStrings := make([]string, len(envVars), len(envVars))
	for _, envVar := range envVars {
		envStrings = append(envStrings, envVar.Name + "=" + envVar.Value)
	}
	return envStrings
}

type Args struct {
	pathToEnvDir string
	command string
}

func getArgs() *Args {
	args := Args{
		pathToEnvDir: os.Args[1],
		command:      os.Args[2],
	}

	if args.pathToEnvDir == "" {
		log.Fatal("Specify 1st parameter: the path to the folder with environment variables")
	}
	if args.command == "" {
		log.Fatal("Specify 2nd parameter: the command to execute")
	}

	return &args
}

func main() {
	args := getArgs()

	envVars := GetEnvVarsFromDirectory(args.pathToEnvDir)
	cmd := exec.Command(args.command)
	cmd.Env = convertEnvVarsToStrings(envVars)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}