# tfplan-checks
tfplan-checks



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

    //status, result:= aws.CloudTrail(pr.Type, pr.Values)
	//if status == "ok"{
	//	color.Green(result)
	//}else{
	//	color.Yellow(result)
	//}

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