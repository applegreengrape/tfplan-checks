package main

import (
	"encoding/json"
	"flag"
	//"fmt"
	"io/ioutil"

	//"tfplan/checks/custom"
	"tfplan/checks/aws"
	"tfplan/checks/data"
)

func main() {
	var tfplan string
	var module string
	flag.StringVar(&tfplan, "tfplan", "./test-json/ec2.json", "path to the terraform json planfile")
	flag.StringVar(&module , "module", "ec2", "module you want to check")

	flag.Parse()

	plan, _ := ioutil.ReadFile(tfplan)

	if module == "cloudtrail"{
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
		aws.CloudTrail(pr.Type, pr.Values)
	}

	if module == "codebuild"{
		var pcodebuild data.CodeBuild
		json.Unmarshal(plan, &pcodebuild)
		aws.CodeBuild(pcodebuild)
	}

	if module == "ec2" {
		var pec2 data.EC2
		json.Unmarshal(plan, &pec2)
		aws.EC2(pec2)	
	} 
}