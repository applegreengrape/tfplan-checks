package cloudtrail

import "github.com/fatih/color"

//[CloudTrail.1] CloudTrail should be enabled and configured with at least one multi-Region trail
func CloudTrail_01(){
	r:=`
	💡[AWS Best Practice][CloudTrail.1] CloudTrail should be enabled and configured with at least one multi-Region trail
		- aws_cloudtrail not found
	`
	color.Green(r)
}

//[CloudTrail.2] CloudTrail should have encryption at-rest enabled
func CloudTrail_02(value interface{}){
	kms_key_id:= value.(map[string]interface{})["kms_key_id"]

	if kms_key_id == nil {
		r:=`
	❌[AWS Best Practice][CloudTrail.2] CloudTrail should have encryption at-rest enabled:
		 - kms_key_id: not_null
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
		`
		color.Yellow(r)
	}else{
		r:=`
	✅[AWS Best Practice][CloudTrail.2] CloudTrail should have encryption at-rest enabled
		`
		color.Green(r)
	}
}
