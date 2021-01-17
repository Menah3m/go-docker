package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main()  {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 隔离 uts ipc pid mount user network
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWNET,
		// 设置容器的UID
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1, //容器的UID
				HostID: 0,      // 宿主机的GID
				Size: 1,
			},
		},
		// 设置容器的GID
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1, // 容器的GID
				HostID: 0,      // 宿主机的GID
				Size: 1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
