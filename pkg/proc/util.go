package proc

import (
	"strings"
)

func (this *Proc) parse(args []string, filters []string) {

	skip := false
	for i, arg := range args {
		for _, filter := range filters {
			if strings.Contains(arg, filter) {
				skip = false
				break
			}
			skip = true
		}

		if skip {
			skip = false
			continue
		}
		if strings.Contains(arg, "-") || strings.Contains(arg, "--") {
			arg = strings.TrimPrefix(arg, "-")
			arg = strings.TrimPrefix(arg, "-")

			if strings.Contains(arg, "=") {
				split := strings.Split(arg, "=")
				this.Args[split[0]] = split[1]
				continue
			}

			if i < len(args)-1 {
				if !strings.HasPrefix(args[i+1], "-") && !strings.HasPrefix(args[i+1], "--") {
					this.Args[arg] = args[i+1]
					skip = true
					continue
				}
			}
			this.Args[arg] = "true"
		} else {
			this.Args[arg] = "true"
		}
	}
}
