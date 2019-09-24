package main

import (
	"fmt"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	data := map[string]string{}

	data["version"] = exec("git describe --tags --abbrev=0")
	data["buildDate"] = exec("date +%Y-%m-%d %H:%M")
	data["gitCommit"] = exec("git rev-parse --short HEAD")
	data["gitState"] = exec("git diff --quiet || echo 'dirty'")
	data["gitBranch"] = exec("git symbolic-ref -q --short HEAD")
	data["gitSummary"] = exec("git describe --tags --dirty --always")

	out := []string{}
	for k, v := range data {
		out = append(out, fmt.Sprintf("-X 'github.com/NoUseFreak/go-vembed.%s=%s'", k, v))
	}

	fmt.Println(strings.Join(out, " "))
}

func exec(cmd string) string {
	out := script.Exec(cmd)

	if out.ExitStatus() == 0 {
		if d, err := out.String(); err == nil {
			return strings.TrimSpace(d)
		}
	}

	return ""
}
