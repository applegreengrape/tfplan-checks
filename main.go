package main

import (
	"io/ioutil"
	"encoding/json"
	"github.com/fatih/color"

	//"tfplan/checks/custom"
	"tfplan/checks/aws"
	"tfplan/checks/data"
)

func main() {
	plan, _ := ioutil.ReadFile("/Users/pingzhou.liu/Documents/terrtest/cloudtrail/cloudtrail.json")

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

	status, result:= aws.CloudTrail(pr.Type, pr.Values)
	if status == "ok"{
		color.Green(result)
	}else{
		color.Yellow(result)
	}
	
}