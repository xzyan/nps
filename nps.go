package main

import (
	"fmt"
	"net/http"
	"nps"
	"os"
)

const format = " %2d  %5s  %s  %4s  %15s\n"

func main() {
	if len(os.Args) == 1 {
		data := nps.RadNetProcessLocal()
		fmt.Printf(format, 0, "PID", "PROCESS        ", "PROTO", "LISTEN")
		for i, it := range data {
			F := 15 - len(it.Name)
			for i := 0; i < F; i++ {
				it.Name += " "
			}
			fmt.Printf(format, i+1, it.PID, it.Name, it.Protocol, it.Listen)
		}
		return
	}
	switch os.Args[1] {
	case "hs":
		println("http listen:2345")
		e := http.ListenAndServe(":2345", http.FileServer(http.Dir(".")))
		println(e)
	}
}
