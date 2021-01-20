// +build !GO_TEST

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	defaultConfigFile = "./.GO_TEST_IGNORE.json"
)

func main() {
	err := work()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// 打印帮助信息
func help(exe string) {
	fmt.Printf("Usage   : %s submit\n", exe)
	fmt.Printf("          %s add <file-path>\n", exe)
	fmt.Printf("          %s del <file-path>\n", exe)
	fmt.Printf("          %s set-flag <build-flag>\n", exe)
	fmt.Println()
	fmt.Printf("Example : %s add main.go\n", exe)
	fmt.Printf("          %s add file_modify.go\n", exe)
	fmt.Printf("          %s set-flag !GO_TEST\n", exe)
	fmt.Printf("          %s submit\n", exe)
}

// 处理函数
func work() error {
	// 加载配置
	config := NewIgnoreConfig()
	config.Load(path.Join(".", defaultConfigFile))

	// 参数检查
	if len(os.Args) == 1 {
		help(os.Args[0])
	} else if os.Args[1] == "add" && len(os.Args) > 2 {
		for _, file := range os.Args[2:] {
			if !strings.HasSuffix(file, ".go") {
				fmt.Fprintf(os.Stderr, "skip '%s' not a go file\n", file)
				continue
			}
			fmt.Fprintf(os.Stderr, "add '%s'\n", file)
			config.AddIgnore(file)
		}
		return config.Write(defaultConfigFile)
	} else if os.Args[1] == "del" && len(os.Args) > 2 {
		for _, file := range os.Args[2:] {
			if !strings.HasSuffix(file, ".go") {
				fmt.Fprintf(os.Stderr, "skip '%s' not a go file\n", file)
				continue
			}
			fmt.Fprintf(os.Stderr, "del '%s'\n", file)
			config.DelIgnore(file)
		}
		return config.Write(defaultConfigFile)
	} else if os.Args[1] == "set-flag" && len(os.Args) == 3 {
		config.BuildFlag = os.Args[2]
		return config.Write(defaultConfigFile)
	} else if os.Args[1] == "submit" && len(os.Args) == 2 {
		// 遍历源码
		return scanDir(".", createFileHandler(config))
	} else {
		help(os.Args[0])
	}
	return nil
}

// 递归扫描文件夹
func scanDir(dir string, handler func(filename string) error) error {
	dir = path.Clean(dir)
	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range fileList {
		if file.IsDir() {
			subdir := path.Join(dir, file.Name())
			err = scanDir(subdir, handler)
			if err != nil {
				return err
			}
		} else {
			filename := path.Join(dir, file.Name())
			err = handler(filename)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 创建文件处理器
func createFileHandler(config *IgnoreConfig) func(string) error {
	return func(filename string) error {
		flag := config.BuildFlag
		if !strings.HasSuffix(filename, ".go") {
			// fmt.Printf("Skip '%s' : not a go file.\n", filename)
		} else if config.ShouldIgnore(filename) {
			fmt.Printf("Add '%s' into '%s'\n", flag, filename)
			return AddBuildFlagToFile(filename, flag)
		} else {
			fmt.Printf("Remove '%s' from '%s'\n", flag, filename)
			return DelBuildFlagFromFile(filename, flag)
		}
		return nil
	}
}
