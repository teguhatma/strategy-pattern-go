package handlestatus

import "fmt"

const Name string = "handle-status"

type HandleStatus struct{}

func (hs *HandleStatus) Execute() error {
	fmt.Println("Handle Status Cron")

	return nil
}

func New() *HandleStatus {
	return &HandleStatus{}
}
