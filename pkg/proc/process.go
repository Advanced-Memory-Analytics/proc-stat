package proc

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Proc struct {
	uid   string
	pid   int
	ppid  int
	c     int
	stime string
	tty   string
	time  string
	cmd   string
	args  string
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
		if len(process) == 0 {
			continue
		}

		proc := &Proc{}
		var cols []string

		process = strings.Join(strings.Fields(strings.TrimSpace(process)), " ")

		cols = strings.Split(process, " ")

		proc.uid = cols[0]
		proc.pid, _ = strconv.Atoi(cols[1])
		proc.ppid, _ = strconv.Atoi(cols[2])
		proc.c, _ = strconv.Atoi(cols[3])
		proc.stime = cols[4]
		proc.tty = cols[5]
		proc.time = cols[6]
		proc.cmd = cols[7]
		proc.args = strings.Join(cols[8:], " ")

		if strings.Contains(proc.args, "-name") {
			indexOfArg := strings.Index(proc.args, "-name")
			indexOfName := indexOfArg + len("-name ")
			endOfName := strings.Index(proc.args[indexOfName:], " ")
			if endOfName == -1 {
				proc.name = proc.args[indexOfName:]
			} else {
				proc.name = proc.args[indexOfName : indexOfName+endOfName]
			}
		}

		processes = append(processes, proc)
	}

	return processes, nil
}

func filter(orig []string, filter string) (result []string) {
    for _, line := range orig {
        if strings.Contains(line, filter) {
            result = append(result, line)
        }
    }
    return result
}
