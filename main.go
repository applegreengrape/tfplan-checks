package main

import (
	"fmt"

	"github.com/hashicorp/terraform/plans/planfile"
	"github.com/hashicorp/terraform/tfdiags"
)


func main(){
	r, err := planfile.Open("tfplan-v12")
	if err != nil {
		fmt.Println(err)
	}

	var diags tfdiags.Diagnostics
	c, diags := r.ReadConfig() 
	if diags != nil {
		fmt.Println(diags)
	}

	p, err := r.ReadPlan() 
	if err != nil {
		fmt.Println(err)
	}

	s, err := r.ReadStateFile() 
	if err != nil {
		fmt.Println(err)
	}

	//tfplan, _ := jsonplan.Marshal(c, p, s, ss)
	fmt.Println(c,p,s)

	fmt.Println(p.Changes.Resources[0].ChangeSrc.Before)
}