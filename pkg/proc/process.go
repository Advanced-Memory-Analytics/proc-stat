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
	Args  map[string]any
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
        proc.Args = make(map[string]any, 0)
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
		args := cols[8:]

        skip := false
        for i, arg := range args {
            if skip {
                continue
            }
            if strings.Contains(arg,"-") || strings.Contains(arg,"--") {
                arg = strings.TrimPrefix(arg, "-")
                arg = strings.TrimPrefix(arg, "-")

                if strings.Contains(arg, "=") {
                    split := strings.Split(arg, "=")
                    proc.Args[split[0]] = split[1]
                    continue
                }

                if i < len(args) - 1 {
                    if !strings.Contains(args[i + 1], "-") && !strings.Contains(args[i + 1], "--") {
                        proc.Args[arg] = args[i + 1]
                        skip = true
                        continue
                    }
                }
                proc.Args[arg] = true
                skip = false
            } else {
                proc.Args[arg] = true
            }
        }

		processes = append(processes, proc)
	}

	return processes, nil
}
