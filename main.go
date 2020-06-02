package main

import (
	"fmt"
	"github.com/applegreengrape/tfplan-checks/custom"
	"github.com/applegreengrape/tfplan-checks/cis"
	//"tfplan/checks/cis"

)

func main() {
	fmt.Println(custom.Checks_01())
	cis.main()
}