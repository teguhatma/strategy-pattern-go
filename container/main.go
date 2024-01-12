package container

import (
	"flag"
	"fmt"
	"time"

	handlestatus "github.com/teguhatma/strategy-pattern-go/task/handle-status"
	othercron "github.com/teguhatma/strategy-pattern-go/task/other-cron"
)

type CronTaskInterface interface {
	Execute() error
}

func ExecuteJob() {
	var (
		name string
		task CronTaskInterface
	)

	flag.StringVar(&name, "name", "handle-status", "Type of activity to run")
	flag.Parse()

	fmt.Println("Name: ", name)

	switch name {
	case handlestatus.Name:
		task = handlestatus.GetHandleStatus()
	case othercron.Name:
		task = othercron.GetOtherCron()
	default:
	}

	startTime := time.Now()

	if err := task.Execute(); err != nil {
		panic(err)
	}

	endTime := time.Now()

	fmt.Println("Start Time: ", startTime)
	fmt.Println("End Time: ", endTime)
	fmt.Println("Duration: ", endTime.Sub(startTime))
}
