package cis

import (
	"fmt"
	"tfplan/checks/cis/cloudtrail"
)

func main()  {
	fmt.Println(cloudtrail.Checks())
}