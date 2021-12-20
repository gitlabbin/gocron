//go:build windows
// +build windows

package utils

import (
	"errors"
	"os/exec"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"bufio"
	"github.com/armon/circbuf"
	"io"
)

type Result struct {
	output string
	err    error
}

// 执行shell命令，可设置执行超时时间
func ExecShell(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("cmd", "/C", command)
	// 隐藏cmd窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	var resultChan chan Result = make(chan Result)
	go func() {
		output, err := cmd.CombinedOutput()
		resultChan <- Result{string(output), err}
	}()
	select {
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(cmd.Process.Pid)).Run()
			cmd.Process.Kill()
		}

		// Block waiting for command to exit, be stopped, or be killed
		finalStatus := <-resultChan
		log.Errorf("end with %v", finalStatus.err)
		return "", errors.New("timeout killed")
	case result := <-resultChan:
		return ConvertEncoding(result.output), result.err
	}
}

func ExecShellPipe(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("cmd", "/C", command)
	// 隐藏cmd窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	var resultChan chan Result = make(chan Result)
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
		buf, _ := circbuf.NewBuffer(2000)
		for scanner.Scan() {
			msg := scanner.Text()
			buf.Write([]byte(msg + "\n"))
			log.Infof("cmd output: %s\n", msg)
		}

		err = cmd.Wait()
		resultChan <- Result{string(buf.Bytes()), err}
		log.Infof("shell Routine done.... %v", err)
	}()
	select {
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(cmd.Process.Pid)).Run()
			cmd.Process.Kill()
		}

		// Block waiting for command to exit, be stopped, or be killed
		finalStatus := <-resultChan
		log.Errorf("end with %v", finalStatus.err)
		return "", errors.New("timeout killed")
	case result := <-resultChan:
		return ConvertEncoding(result.output), result.err
	}
}

func ConvertEncoding(outputGBK string) string {
	// windows平台编码为gbk，需转换为utf8才能入库
	outputUTF8, ok := GBK2UTF8(outputGBK)
	if ok {
		return outputUTF8
	}

	return outputGBK
}
