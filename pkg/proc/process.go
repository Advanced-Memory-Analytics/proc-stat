package proc

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Proc struct {
	Uid   string
	Pid   int
	Ppid  int
	C     int
	Stime string
	Tty   string
	Time  string
	Cmd   string
	Args  string
	Name  string
}

func PSEF() ([]*Proc, error) {
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

		proc.Uid = cols[0]
		proc.Pid, _ = strconv.Atoi(cols[1])
		proc.Ppid, _ = strconv.Atoi(cols[2])
		proc.C, _ = strconv.Atoi(cols[3])
		proc.Stime = cols[4]
		proc.Tty = cols[5]
		proc.Time = cols[6]
		proc.Cmd = cols[7]
		proc.Args = strings.Join(cols[8:], " ")

		if strings.Contains(proc.Args, "-Name") {
			indexOfArg := strings.Index(proc.Args, "-Name")
			indexOfName := indexOfArg + len("-Name ")
			endOfName := strings.Index(proc.Args[indexOfName:], " ")
			if endOfName == -1 {
				proc.Name = proc.Args[indexOfName:]
			} else {
				proc.Name = proc.Args[indexOfName : indexOfName+endOfName]
			}
		}

		processes = append(processes, proc)
	}

	return processes, nil
}
