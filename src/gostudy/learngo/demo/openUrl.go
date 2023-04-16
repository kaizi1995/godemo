package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// 不同平台启动指令不同
var commands = map[string]string{
	"windows": "explorer",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func main() {
	Open("https://baidu.com")
}
