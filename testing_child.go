package main
import (
        "fmt"
        "github.com/Advanced-Memory-Analytics/proc-stat/pkg/proc"
)

func main() {
          procs, err := proc.PSEF()
          if err != nil{
                  fmt.Println("Error:", err)
                  return
          }
          parentChildMap := proc.FindChildren(procs)

        for ppid, children := range parentChildMap {
                fmt.Printf("Parent PID: %d\n", ppid)
        for _, child := range children {
                fmt.Printf("\tChild PID: %d, CMD: %s\n", child.Pid, child.Cmd)
        }
 }
}

