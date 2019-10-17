package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(cmdArgs []string) ([]byte, error) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), []string{"LANG=C", "LC_ALL=C"}...)
	out, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf(`Shit hit the fan! "%v" (%+v)`, strings.Join(cmdArgs, " "), err)
	}
	return out, err
}
