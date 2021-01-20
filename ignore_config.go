package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"sort"
)

type IgnoreConfig struct {
	BuildFlag string   // Build Constraints add to ignored source files
	FileList  []string // ignored source files list

	fileMap map[string]bool
}

func NewIgnoreConfig() *IgnoreConfig {
	return &IgnoreConfig{
		BuildFlag: "!GO_TEST",
		FileList:  make([]string, 0),
		fileMap:   make(map[string]bool),
	}
}

func (config *IgnoreConfig) Load(file string) error {
	file = path.Clean(file)
	byets, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byets, config)
	if err != nil {
		return err
	}

	for _, f := range config.FileList {
		config.fileMap[f] = true
	}

	return nil
}

func (config *IgnoreConfig) Write(file string) error {
	file = path.Clean(file)
	config.FileList = make([]string, 0)
	for f, _ := range config.fileMap {
		config.FileList = append(config.FileList, f)
	}

	sort.Strings(config.FileList)

	bytes, err := json.MarshalIndent(*config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (config *IgnoreConfig) SetBuildFlag(flag string) {
	config.BuildFlag = flag
}

func (config *IgnoreConfig) AddIgnore(file string) {
	file = path.Clean(file)
	config.fileMap[file] = true
}

func (config *IgnoreConfig) DelIgnore(file string) {
	file = path.Clean(file)
	delete(config.fileMap, file)
}

func (config *IgnoreConfig) ShouldIgnore(file string) bool {
	file = path.Clean(file)
	_, ok := config.fileMap[file]
	return ok
}
