// export the subpackage to the module
// https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
package aws

import (
	"tfplan/checks/aws/cloudtrail"
	"tfplan/checks/aws/s3"
)

var s string
var r string

func CloudTrail(typ []interface{}, value []interface{}) (string, string){
	//print out advice 
	s, r:= cloudtrail.CloudTrail_01()

	for i:= range typ {
		if typ[i] == "aws_cloudtrail"{
			s, r:= cloudtrail.CloudTrail_02(value[i])
			return s,r 
		}
	}
	return s,r 
}



func S3(typ []interface{}, value []interface{}) (string, string){
	// !aws_s3_bucket_public_access_block
	for i:= range typ {
		if typ[i] == "aws_s3_bucket_public_access_block"{
			s, r:= s3.S3_01(value[i])
			return s,r 
		}
	}
	return s,r 
}

