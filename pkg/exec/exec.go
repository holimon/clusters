package exec

import (
	"context"
	"os/exec"
	"time"
)

const NetworkExecmdTimeout = 10 * time.Second
const ExecmdIntervalTime = 1 * time.Second
const ExecmdRetryInterval = 5 * time.Second

func Execmd(cmd string) (out string, err error) {
	outbuf, err := exec.Command("sh", "-c", cmd).Output()
	out = string(outbuf)
	return
}

func ExecmdWithTimeout(cmd string, timeout time.Duration) (out string, err error) {
	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	outbuf, err := exec.CommandContext(ctxt, "bash", "-c", cmd).Output()
	out = string(outbuf)
	return
}

func ExeSysctlcmd(option string, service string, timeout time.Duration) (out string, err error) {
	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	outbuf, err := exec.CommandContext(ctxt, "systemctl", option, service).Output()
	out = string(outbuf)
	return
}
