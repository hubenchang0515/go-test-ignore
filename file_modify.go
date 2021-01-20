// +build !GO_TEST

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	stateFinding    = iota // 查找build flag中
	stateFound             // 已经找到了build flag
	stateTerminated        // 已经遇到package，查找结束
)

func AddBuildFlagToFile(file string, flag string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	outputCode := ""                                  // 将要输出的代码
	source := strings.SplitAfter(string(bytes), "\n") // 分行后的源代码
	state := stateFinding
	for _, line := range source {
		if state == stateFound || state == stateTerminated {
			// 已经查找结束，直接往后写
			outputCode += line
		} else if strings.HasPrefix(strings.TrimSpace(line), "package") {
			// 没有找到flag，找到了package关键字，在开头写flag并空一行
			outputCode = fmt.Sprintf("// +build %s\n\n", flag) + outputCode

			// 写源码并更新状态
			outputCode += line
			state = stateTerminated
		} else if CheckMultiBuildFlag(flag, line) {
			// 找到了flag，写源码并更新状态
			outputCode += line
			state = stateFound
		} else {
			// 查找中，什么也没找到，直接往后写
			outputCode += line
		}
	}

	err = ioutil.WriteFile(file, []byte(outputCode), 0644)
	return err
}

func DelBuildFlagFromFile(file string, flag string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	outputCode := ""                                  // 将要输出的代码
	source := strings.SplitAfter(string(bytes), "\n") // 分行后的源代码
	state := stateFinding
	for _, line := range source {
		if state == stateTerminated {
			// 已经查找结束，直接往后写
			outputCode += line
		} else if strings.TrimSpace(line) == "" {
			// 忽略，清除package之前的空行，避免反复添加、删除FLAG导致积累大量空行
		} else if strings.HasPrefix(strings.TrimSpace(line), "package") {
			// 没有找到flag，找到了package关键字，写源码并更新状态
			// 由于之前清除了空行，所以需要在package之前补充一个空行，避免其他的build flag报错
			if outputCode != "" {
				outputCode += "\n" + line
			} else {
				outputCode += line
			}

			state = stateTerminated
		} else if CheckMultiBuildFlag(flag, line) {
			// 找到了flag，删除该flag项并更新状态
			newFlag := RemoveFlag(flag, line)
			if newFlag != "" { // 避免积累空行
				outputCode += newFlag + "\n"
				state = stateFound
			}
		} else {
			// 查找中，什么也没找到，直接往后写
			outputCode += line
		}
	}

	err = ioutil.WriteFile(file, []byte(outputCode), 0644)
	return err
}
