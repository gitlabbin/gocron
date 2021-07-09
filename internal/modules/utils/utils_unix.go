// +build !windows

package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os/exec"
	"syscall"
)

type Result struct {
	output string
	err    error
}

// 执行shell命令，可设置执行超时时间
func ExecShell(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	resultChan := make(chan Result)
	go func() {
		output, err := cmd.CombinedOutput()
		resultChan <- Result{string(output), err}
		log.Infof("shell routine down.... %v", err)
	}()

	select {
	case result := <-resultChan:
		return result.output, result.err
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		// Block waiting for command to exit, be stopped, or be killed
		finalStatus := <-resultChan
		log.Errorf("end with %v", finalStatus.err)
		return "", errors.New("timeout killed")
	}
}
