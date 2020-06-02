package main

import (
	"io/ioutil"
	"encoding/json"
	"github.com/fatih/color"

	"tfplan/checks/custom"
	"tfplan/checks/aws"
	"tfplan/checks/data"
)

func main() {
	plan, _ := ioutil.ReadFile("tfplan-with-module.json")

	// pwm - plan with module
	var pwm data.WithMoudle
	json.Unmarshal(plan, &pwm)

	var pr data.Result
	for i:= range pwm.PlannedValues.RootModule.ChildModules {
		for n:= range pwm.PlannedValues.RootModule.ChildModules[i].Resources {
			p:= pwm.PlannedValues.RootModule.ChildModules[i].Resources[n].(map[string]interface{})
			pr.Type = append(pr.Type,p["type"])
			pr.Values = append(pr.Values,p["values"])
		}
	}

	status, result:= aws.S3(pr.Type, pr.Values)
	if status == "warning"{
		color.Yellow(result)
	}

	// call a custom checks
	s, r:= custom.S3_Tag_Check_01(pr.Type, pr.Values)
	if s == "okay"{
		color.Green(r)
	}else{
		color.Yellow(r)
	}

	// call another custom checks
	s2, r2:= custom.S3_Tag_Check_02(pr.Type, pr.Values)
	if s2 == "okay"{
		color.Green(r2)
	}else{
		color.Yellow(r2)
	}
	
}