// export the subpackage to the module
// https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
package aws

import (
	"tfplan/checks/aws/s3"
)

// checks for all s3 resources
func S3(typ []interface{}, value []interface{}) (string, string){
	var s string
	var r string
	for i:= range typ {
		if typ[i] == "aws_s3_bucket_public_access_block"{
			s, r:= s3.S3_01(value[i])
			return s,r 
		}
		if typ[i] == "aws_s3_bucket"{
			
		}
	}
	return s,r 
}
