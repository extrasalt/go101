package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "host":
		host()
	case "container":
		container()
	default:
		fmt.Println("Usage: barebonescontainer <mode>")
	}
}

func host() {
	pid := os.Getpid()
	fmt.Println("pid", pid)
}

func container() {
	cmd := exec.Command("/proc/self/exe", "host")
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWPID,
	}
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
