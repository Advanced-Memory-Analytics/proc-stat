package proc

import (
	"fmt"
	"hash/fnv"
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
	Args  map[string]string
}

func PSEF(cmd string, argFilters ...string) ([]*Proc, uint64, error) {
	processes := make([]*Proc, 0)
	checksum := uint64(0)
	h := fnv.New32a()

	ps := exec.Command("ps", "-ef")
	out, err := ps.Output()
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to get output of command: %v with error: %v", ps, err)
	}

	lines := strings.Split(string(out), "\n")

	for _, process := range lines {
		if len(process) == 0 {
			continue
		}

		proc := &Proc{}
		proc.Args = make(map[string]string, 0)
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

		if !strings.Contains(proc.Cmd, cmd) {
			continue
		}

		proc.parse(args, argFilters)

		processes = append(processes, proc)

		h.Write([]byte(cols[1] + proc.Stime))
		checksum += uint64(h.Sum32())
	}

	return processes, checksum, nil
}
