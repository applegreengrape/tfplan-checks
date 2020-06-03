package cloudtrail

//[CloudTrail.1] CloudTrail should be enabled and configured with at least one multi-Region trail
func CloudTrail_01() (string, string) {
	str:=`
	💡[AWS Best Practice][CloudTrail.1] CloudTrail should be enabled and configured with at least one multi-Region trail
		- aws_cloudtrail not found
	`
	return "info", str
}

//[CloudTrail.2] CloudTrail should have encryption at-rest enabled
func CloudTrail_02(value interface{}) (string, string){
	kms_key_id:= value.(map[string]interface{})["kms_key_id"]

	if kms_key_id == nil {
		str:=`
	❌[AWS Best Practice][CloudTrail.2] CloudTrail should have encryption at-rest enabled:
		 - kms_key_id: not_null
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
		`
		return "warning", str
	}else{
		str:=`
	✅[AWS Best Practice][CloudTrail.2] CloudTrail should have encryption at-rest enabled
		`
		return "ok", str
	}
}
