package proc

import (
	"fmt"
	"github.com/Advanced-Memory-Analytics/proc-stat/_test"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPSEF(test *testing.T) {
	procs, _, err := PSEF("")
	if err != nil {
		test.Fatal(err)
	}

	for i, proc := range procs {
		process, err := os.FindProcess(proc.Pid)
		if err != nil {
			test.Errorf("Case: %d, Failed to find process with Pid: %d, Error: %v", i, proc.Pid, err)
		}

		if process.Pid != proc.Pid {
			test.Errorf("Case: %d, Actual PID: %d, Expected PID: %d", i, process.Pid, proc.Pid)
		}
	}
}

func TestPSEFWithNameFlag(test *testing.T) {
	testDir := _test.GetTestDir()
	cmd1 := exec.Command("/bin/bash", testDir+"/tester.sh", "-uuid b329c3cb-872b-4762-a2b9-ba0d401907d6 -name guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on")
	cmd2 := exec.Command("/bin/bash", testDir+"/tester.sh", "-name test2")
	foundName1 := false
	foundName2 := false

	go cmd1.Run()
	go cmd2.Run()

	defer cmd1.Wait()
	defer cmd2.Wait()

	procs, _, err := PSEF("")
	if err != nil {
		test.Fatal(err)
	}

	for i, proc := range procs {
		process, err := os.FindProcess(proc.Pid)
		if err != nil {
			test.Errorf("Case: %d, Failed to find process with Pid: %d, Error: %v", i, proc.Pid, err)
		}

		if process.Pid != proc.Pid {
			test.Errorf("Case: %d, Actual PID: %d, Expected PID: %d", i, process.Pid, proc.Pid)
		}

		if proc.Args["name"] == "guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on" {
			foundName1 = true
		}

		if proc.Args["name"] == "test2" {
			foundName2 = true
		}

	}

	if !foundName1 || !foundName2 {
		test.Fatal("Failed to find names.")
	}
}

func TestParseArgs(test *testing.T) {
	proc := &Proc{}
	proc.Args = make(map[string]string, 0)

	args := strings.Split("-uuid b329c3cb-872b-4762-a2b9-ba0d401907d6 -name guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on -S -object {\"qom-type\":\"secret\",\"id\":\"masterKey0\",\"format\":\"raw\",\"file\":\"/var/lib/libvirt/qemu/domain-b329c3cb-872b-4762-a2b9-ba0d401907d6/master-key.aes\"} -machine pc-i440fx-rhel7.6.0,usb=off,dump-guest-core=off,mem-merge=off -accel kvm -cpu host,migratable=on,hv-time=on,kvm-pv-eoi=on,hv-relaxed=on,hv-vapic=on,hv-spinlocks=0x2000,hv-vpindex=on,hv-runtime=on,hv-synic=on,hv-stimer=on,hv-tlbflush=on,hv-ipi=on,l3-cache=on,host-cache-info=off,check=no -m size=209715200k,slots=240,maxmem=8589934592k -overcommit mem-lock=off -smp 32,maxcpus=240,sockets=240,dies=1,cores=1,threads=1 -object {\"qom-type\":\"memory-backend-file\",\"id\":\"ram-node0\",\"mem-path\":\"/dev/hugepages/libvirt/qemu/b329c3cb-872b-4762-a2b9-ba0d401907d6\",\"share\":true,\"prealloc\":true,\"size\":214748364800} -numa node,nodeid=0,cpus=0-239,memdev=ram-node0 -smbios type=1,manufacturer=Nutanix,product=AHV,serial=B329C3CB-872B-4762-A2B9-BA0D401907D6 -device vmgenid,guid=2cc242ca-1bb3-4dc1-b9e3-95add7c8d408,id=vmgenid0 -no-user-config -nodefaults -chardev socket,id=charmonitor,fd=39,server=on,wait=off -mon chardev=charmonitor,id=monitor,mode=control -rtc base=localtime,clock=vm,driftfix=slew -global kvm-pit.lost_tick_policy=delay -no-shutdown -boot menu=off,strict=on -device ich9-usb-ehci1,id=usb,bus=pci.0,addr=0x5.0x7 -device ich9-usb-uhci1,masterbus=usb.0,firstport=0,bus=pci.0,multifunction=on,addr=0x5 -device ich9-usb-uhci2,masterbus=usb.0,firstport=2,bus=pci.0,addr=0x5.0x1 -device ich9-usb-uhci3,masterbus=usb.0,firstport=4,bus=pci.0,addr=0x5.0x2 -device ide-cd,bus=ide.0,unit=0,id=ide0-0-0,bootindex=1,werror=report,rerror=report -netdev tap,fd=42,id=hostua-742dbf17-241a-4772-a2d6-a324e9fbdd97,vhost=on,vhostfd=43 -device virtio-net-pci,rx_queue_size=256,netdev=hostua-742dbf17-241a-4772-a2d6-a324e9fbdd97,id=ua-742dbf17-241a-4772-a2d6-a324e9fbdd97,mac=50:6b:8d:eb:80:ea,bootindex=4,bus=pci.0,addr=0x3 -chardev socket,id=charserial0,fd=38,server=on,wait=off -device isa-serial,chardev=charserial0,id=serial0,index=1,iobase=760,irq=3 -device usb-tablet,id=input0,bus=usb.0,port=1 -audiodev {\"id\":\"audio1\",\"driver\":\"none\"} -vnc 127.0.0.1:1,audiodev=audio1 -device VGA,id=video0,vgamem_mb=16,bus=pci.0,addr=0x2 -device virtio-balloon-pci,id=balloon0,deflate-on-oom=on,bus=pci.0,addr=0x6 -device vmcoreinfo -object dbus-vmstate,addr=unix:path=/run/dbus/system_bus_socket,id=dbus-vmstate1,id-list=b329c3cb-872b-4762-a2b9-ba0d401907d6 -sandbox on,obsolete=deny,elevateprivileges=deny,spawn=deny,resourcecontrol=deny -msg timestamp=on -chardev socket,id=frodo0,fd=3 -device vhost-user-scsi-pci,chardev=frodo0,id=frodo-scsi0,num_queues=32,max_sectors=2048,bus=pci.0,addr=0x4,bootindex=2,boot_tpgt=0", " ")

	proc.parse(args, nil)

	if proc.Args["uuid"] != "b329c3cb-872b-4762-a2b9-ba0d401907d6" {
		test.Errorf("Key mismatch.  Expected: %s, Actual: %s", "b329c3cb-872b-4762-a2b9-ba0d401907d6", proc.Args["uuid"])
	}

	if proc.Args["name"] != "guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on" {
		test.Errorf("Key mismatch.  Expected: %s, Actual: %s", "guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on", proc.Args["name"])
	}

}

func TestPSEFWithFilter(test *testing.T) {
	testDir := _test.GetTestDir()
	cmd1 := exec.Command("/bin/bash", testDir+"/tester.sh", "-uuid b329c3cb-872b-4762-a2b9-ba0d401907d6 -name guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on -junk -moreJunk thisshouldnotshowup")
	cmd2 := exec.Command("/bin/bash", testDir+"/tester.sh", "-name test2 -junk -moreJunk thisshouldnotshowup")

	expectedLongName := "guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on"
	expectedShortName := "test2"
	expectedUUID := "b329c3cb-872b-4762-a2b9-ba0d401907d6"

	go cmd1.Run()
	go cmd2.Run()

	defer cmd1.Wait()
	defer cmd2.Wait()

	procs, _, err := PSEF("bash", "name", "uuid", "test.sh")
	if err != nil {
		test.Fatal(err)
	}

	// Remove testing script if used.
	for i, proc := range procs {
		if proc.Args["./test.sh"] != "" {
			procs = append(procs[:i], procs[i+1:]...)
		}
	}

	if len(procs) != 2 {
		for _, proc := range procs {
			fmt.Printf("%v\n", proc)
		}
		test.Fatalf("Failed to filter processes")
	}

	for _, proc := range procs {
		uuid := proc.Args["uuid"]
		name := proc.Args["name"]

		if name != expectedLongName {
			if name == expectedShortName {
				if uuid != "" {
					test.Errorf("Shouldn't have UUID for: %s", name)
				}
			} else {
				test.Errorf("Expected name: %s, Actual name: %s", expectedShortName, name)
			}
		} else {
			if uuid != expectedUUID {
				test.Errorf("Expected UUID: %s, Actual UUID: %s", expectedUUID, uuid)
			}
		}
	}
}

func TestChecksum(test *testing.T) {
    testDir := _test.GetTestDir()
    cmd1 := exec.Command("/bin/bash", testDir+"/tester.sh", "-uuid b329c3cb-872b-4762-a2b9-ba0d401907d6 -name guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on -junk -moreJunk thisshouldnotshowup")
    cmd2 := exec.Command("/bin/bash", testDir+"/tester.sh", "-name test2 -junk -moreJunk thisshouldnotshowup")
    cmd3 := exec.Command("/bin/bash", testDir+"/tester.sh", "-uuid b329c3cb-872b-4762-a2b9-ba0d401907d6 -name guest=b329c3cb-872b-4762-a2b9-ba0d401907d6,debug-threads=on -junk -moreJunk thisshouldnotshowup")

    cmd1.Start()
    cmd2.Start()

    _, checksum1, err := PSEF("bash")
    if err != nil {
        test.Fatal(err)
    }

    _, checksum2, err := PSEF("bash")
    if err != nil {
        test.Fatal(err)
    }

    if checksum1 != checksum2 {
        test.Errorf("Checksums not equal when should be:  %d, %d", checksum1, checksum2)
    }

    cmd1.Process.Kill()

    _, checksum3, err := PSEF("bash", "name")
    if checksum1 == checksum3 {
        test.Errorf("Checksums are equal when shouldn't be: %d, %d", checksum1, checksum3)
    }

    cmd3.Start()

    _, checksum4, err := PSEF("bash", "name")
    if checksum1 == checksum4 {
        test.Errorf("Checksums are equal when shouldn't be: %d, %d", checksum1, checksum3)
    }


}
