package main

import (
	"fmt"
	"nps"
)

const format = " %2d  %5s  %s  %4s  %15s\n"

func main() {
	data := nps.RadNetProcessLocal()
	fmt.Printf(format, 0, "PID", "PROCESS        ", "PROTO", "LISTEN")
	for i, it := range data {
		F := 15 - len(it.Name)
		for i := 0; i < F; i++ {
			it.Name += " "
		}
		fmt.Printf(format, i+1, it.PID, it.Name, it.Protocol, it.Listen)
	}
}
