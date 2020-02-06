package nps

import (
	"fmt"
	"sort"
	"strings"
	"tool"
)

type NetProcess struct {
	PID    string
	Name   string
	Proto  string
	Listen string
	Count  int
}

func RadNetProcessLocal() (nets []NetProcess) {

	nets = make([]NetProcess, 0)
	netKv := make(map[string]NetProcess)

	output := tool.CallShell("netstat -tunlp")

	for i, process := range strings.Split(output, "\n") {
		if i == 0 || i == 1 {
			continue
		}
		if process == "" {
			continue
		}
		//
		item := make([]string, 0)
		for _, it := range strings.Split(process, " ") {
			if it == "" || it == "0" {
				continue
			}
			item = append(item, it)
		}
		if item[len(item)-1] == "-" {
			continue
		}
		//
		var name, pid string
		tmp := strings.Split(item[len(item)-1], "/")
		for i := 0; i < len(tmp); i++ {
			if i != 0 {
				name = tmp[i]
				//if i == len(tmp)-1 {
				//	name += tmp[i]
				//} else {
				//	name += tmp[i] + "/"
				//}
			}
			pid = tmp[0]
		}
		if name == "" || pid == "" {
			continue
		}
		//
		obj := NetProcess{
			Name:   strings.TrimSpace(name),
			PID:    pid,
			Proto:  item[0],
			Listen: item[1],
		}
		//
		index := obj.Name + obj.Proto
		if this, has := netKv[index]; !has {
			obj.Count = 1
			netKv[index] = obj
		} else {
			this.Count++
			this.Listen += "," + obj.Listen
			netKv[index] = this
		}
	}

	for _, it := range netKv {
		it.Listen = strings.ReplaceAll(it.Listen, "0.0.0.0", "::")
		it.Listen = strings.ReplaceAll(it.Listen, ":::", "")
		it.Listen = strings.ReplaceAll(it.Listen, "127.0.0.1", "local")
		it.Name = fmt.Sprintf("%s*%d", it.Name, it.Count)
		//it.Name = fmt.Sprintf("%d*%s", it.Count, it.Name)
		nets = append(nets, it)
	}

	sort.Slice(nets, func(i, j int) bool {
		I := strings.ToUpper(string(nets[i].Name[0]))
		J := strings.ToUpper(string(nets[j].Name[0]))
		if I == "." {
			I = strings.ToUpper(string(nets[i].Name[2]))
		}
		if J == "." {
			J = strings.ToUpper(string(nets[j].Name[2]))
		}
		if I < J {
			return true
		}
		return false
	})
	return
}
