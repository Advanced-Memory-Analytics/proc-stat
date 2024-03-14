package prc

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
<<<<<<< HEAD
0
func PSEF() ([]*Proc, map[int][]*Proc, error) {
=======

func PSEF(cmd string, argFilters ...string) ([]*Proc, error) {
>>>>>>> main
	processes := make([]*Proc, 0)

	ps := exec.Command("ps", "-ef")
	out, err := ps.Output()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get output of command: %v with error: %v", ps, err)
	}

	lines := strings.Split(string(out), "\n")

	children := make(map[int][]*Proc)
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

		children[proc.Ppid] = append(children[proc.Ppid], proc)
	}

	return processes, children, nil
}

func MonitorPIDChanges(interval time.Duration, stopChan <-chan struct{}){ //chan ensures synchronization
	prevProcs := make(map[string]int)

    	for {
        	select {
        	case <-stopChan:
            		fmt.Println("Stopping process monitoring")
            		return
        	default:
            		currentProcs, _, err := PSEF()
            		if err != nil {
                		fmt.Printf("Error fetching processes: %v\n", err)
                		return
            		}

            		// Use map of command names to PIDs for current processes
            		currentProcsMap := make(map[string]int)
            		for _, process := range currentProcs {
                		currentPID := prevProcs[process.Cmd]
                		currentProcsMap[process.Cmd] = process.Pid

                		if currentPID == 0 {
                    			fmt.Printf("New application detected: %s with PID %d\n", process.Cmd, process.Pid)
            			}else if currentPID != process.Pid{
                                        fmt.Printf("Application %s changed PID from %d to %d\n", process.Cmd, currentPID, process.Pid)
                                }
			}

            // Check for terminated applications
            		for cmd, pid := range prevProcs {
                		if _, exists := currentProcsMap[cmd]; !exists {
                    			fmt.Printf("Application terminated: %s with PID %d\n", cmd, pid)
				}
            		}

            // Update prevProcs to reflect the current state
            		prevProcs = currentProcsMap
            		time.Sleep(interval)
        	}
    	}	
}
