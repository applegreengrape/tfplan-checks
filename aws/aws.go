// export the subpackage to the module
// https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
package aws

import (
	"tfplan/checks/aws/cloudtrail"
	"tfplan/checks/aws/codebuild"
	"tfplan/checks/aws/s3"
	"tfplan/checks/aws/ec2"
	"tfplan/checks/data"
)

var s string
var r string

func CloudTrail(typ []interface{}, value []interface{}){
	//print out advice 
	//s, r:= cloudtrail.CloudTrail_01()

	for i:= range typ {
		if typ[i] == "aws_cloudtrail"{
			cloudtrail.CloudTrail_02(value[i])
			//return s,r 
		}
	}
	//return s,r 
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

func CodeBuild(plan data.CodeBuild){
	// CodeBuild_01
	codebuild.CodeBuild_01(plan)
	// CodeBuild_02
	codebuild.CodeBuild_02(plan)
}

func EC2(plan data.EC2)  {
	// [EC2.3] Attached EBS volumes should be encrypted at-rest
	ec2.EC2_03(plan)	
}

