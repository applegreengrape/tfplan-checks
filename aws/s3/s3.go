package s3

// [S3.1] S3 Block Public Access setting should be enabled
// This control checks whether the following Amazon S3 public access block settings are configured at the account level:
// ignorePublicAcls: true
// blockPublicPolicy: true
// blockPublicAcls: true
// restrictPublicBuckets: true

func S3_01(value interface{}) (string, string){
	ignore_public_acls:= value.(map[string]interface{})["ignore_public_acls"]
	block_public_policy:= value.(map[string]interface{})["block_public_policy"]
	block_public_acls:= value.(map[string]interface{})["block_public_acls"]
	restrict_public_buckets:= value.(map[string]interface{})["restrict_public_buckets"]

	if (ignore_public_acls == "true" && block_public_policy == "true" && block_public_acls == "true" && restrict_public_buckets == "true"){
		return "info", "✅[S3.1] S3 Block Public Access setting should be enabled"
	}else{
		str:=`
	❌[AWS Best Practice][S3.1] S3 Block Public Access Check Failed. Please Update your s3 module to set:
		 - ignore_public_acls: true
		 - block_public_policy: true
		 - block_public_acls: true
		 - restrict_public_buckets: true
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
		`
		return "warning", str
	}

}