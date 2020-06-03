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