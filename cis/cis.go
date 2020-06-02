package cis

import (
	"fmt"
	"github.com/applegreengrape/tfplan-checks/cis/cloudtrail"
)

func main()  {
	fmt.Println(cloudtrail.Checks())
}