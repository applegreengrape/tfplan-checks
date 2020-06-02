package cis

import (
	"fmt"
	"github.com/applegreengrape/tfplan-checks/cis/cloudtrail"
)

func Status()  {
	fmt.Println(cloudtrail.Checks())
}