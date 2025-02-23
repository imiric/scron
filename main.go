package main

import (
	"fmt"
	"time"

	scron "github.com/imiric/scron/util"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	for _, cronEntity := range scron.Crontab("crontab") {
		fmt.Println("add spec=", cronEntity.Spec, "command=", cronEntity.Command)
		c.AddFunc(cronEntity.Spec, cronFunc(cronEntity))
	}

	c.Run()
}

func cronFunc(cron scron.Cron) func() {
	return func() {
		out, _ := scron.ShellRun(cron.Command)
		fmt.Println("Run", cron.Command, "at", time.Now(), "output: ", out)
	}
}
