package main
import (
        "fmt"
        "github.com/Advanced-Memory-Analytics/proc-stat/pkg/proc"
)

func main() {
          _, children, err := proc.PSEF()
          if err != nil{
                  fmt.Println("Error:", err)
                  return
          }
          //parentChildMap := proc.FindChildren(procs)

	//fmt.Printf("procceses:", procs)
        for ppid, children := range children {
                fmt.Printf("Parent PID: %d\n", ppid)
        for _, child := range children {
                fmt.Printf("\tChild PID: %d, CMD: %s\n", child.Pid, child.Cmd)
        }
 }
}

