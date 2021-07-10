// +build !windows

package utils

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io"
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

func ExecShellPipe(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	resultChan := make(chan Result)
	go func() {
		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalf("could not get stderr pipe: %v", err)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatalf("could not get stdout pipe: %v", err)
		}
		// start the command after having set up the pipe
		if err := cmd.Start(); err != nil {
			log.Fatalf("could not start cmd: %v", err)
		}

		merged := io.MultiReader(stderr, stdout)
		scanner := bufio.NewScanner(merged)
		var msg string
		for scanner.Scan() {
			msg = scanner.Text()
			log.Infof("cmd output: %s\n", msg)
		}

		err = cmd.Wait()
		resultChan <- Result{string(msg), err}
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
