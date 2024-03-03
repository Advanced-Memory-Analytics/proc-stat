package proc

import (
    "github.com/Advanced-Memory-Analytics/proc-stat/_test"
    "os"
    "os/exec"
    "testing"
)

func TestPSEF(test *testing.T) {
	procs, err := PSEF()
	if err != nil {
		test.Fatal(err)
	}

	for i, proc := range procs {
		process, err := os.FindProcess(proc.pid)
		if err != nil {
			test.Errorf("Case: %d, Failed to find process with pid: %d, Error: %v", i, proc.pid, err)
		}

		if process.Pid != proc.pid {
			test.Errorf("Case: %d, Actual PID: %d, Expected PID: %d", i, process.Pid, proc.pid)
		}
	}
}

func TestPSEFWithNameFlag(test *testing.T) {
	testDir := _test.GetTestDir()
	cmd1 := exec.Command("/bin/bash", testDir+"/tester.sh", "-name test1")
	cmd2 := exec.Command("/bin/bash", testDir+"/tester.sh", "-name test2")
	foundName1 := false
	foundName2 := false

	go cmd1.Run()
	go cmd2.Run()

	procs, err := PSEF()
	if err != nil {
		test.Fatal(err)
	}

	for i, proc := range procs {
		process, err := os.FindProcess(proc.pid)
		if err != nil {
			test.Errorf("Case: %d, Failed to find process with pid: %d, Error: %v", i, proc.pid, err)
		}

		if process.Pid != proc.pid {
			test.Errorf("Case: %d, Actual PID: %d, Expected PID: %d", i, process.Pid, proc.pid)
		}

		if proc.name == "test1" {
			foundName1 = true
		}

		if proc.name == "test2" {
			foundName2 = true
		}

	}

	if !foundName1 && !foundName2 {
		test.Fatal("Failed to find names.")
	}
}
