package proc

import (
	"fmt"
	"os/exec"
	"strings"
)

type Proc struct {
	uid   string
	pid   int64
	ppid  int64
	c     int64
	stime string
	tty   string
	time  string
	cmd   string
    args []string
	name  string
}

func Processes(filter ...string) ([]*Proc, error) {
	processes := make([]*Proc, 0)

	ps := exec.Command("ps", "-ef")
	out, err := ps.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed to get output of command: %v with error: %v", ps, err)
	}

	lines := strings.Split(string(out), "\n")

	for _, process := range lines {
        //proc := &Proc{}
        var cols []string
        args := ""
        process = strings.Join(strings.Fields(strings.TrimSpace(process)), " ")
        argIndex := strings.Index(process, "-")
        if argIndex != -1 {
            cols = strings.Split(process[:argIndex], " ")
            args = process[argIndex:]
        } else {
            cols = strings.Split(process, " ")
        }
        println(process)
        for _, col := range cols {
          println(col)
        }
	}

	return processes, nil
}
