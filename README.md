# go-test-ignore
在golang项目中自动添加构建约束，从而在覆盖率测试中忽略部分文件

## Usage
```bash
Usage   : go-test-ignore submit                    # 确认生效
          go-test-ignore add <file-path>           # 添加一个需要忽略的文件
          go-test-ignore del <file-path>           # 删除一个需要忽略的文件
          go-test-ignore set-flag <build-flag>     # 设置需要忽略的文件使用的构建约束

Example : go-test-ignore add main.go               # 忽略 main.go
          go-test-ignore add file_modify.go        # 忽略 file_modify.go
          go-test-ignore set-flag '!GO_TEST'       # 需要忽略的文件的构建约束设为 !GO_TEST
          go-test-ignore submit                    # 确认生效
```

## 注意事项
对于在 `package` 的行开头写注释的魔幻代码不生效
``` go
/**/ package main
```