package othercron

import "fmt"

const Name string = "other-cron"

type OtherCron struct{}

func (oc *OtherCron) Execute() error {
	fmt.Println("Run a cron")

	return nil
}

func New() *OtherCron {
	return &OtherCron{}
}
