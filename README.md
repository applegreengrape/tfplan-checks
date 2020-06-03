# tfplan-checks
tfplan-checks

```
$ go run main.go --tfplan test-json/cloudtrail.json --module cloudtrail

        ❌[AWS Best Practice][CloudTrail.2] CloudTrail should have encryption at-rest enabled:
                 - kms_key_id: not_null
        ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html

$ go run main.go --tfplan test-json/codebuild.json --module codebuild

        ❌[AWS Best Practice][CodeBuild.1] CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
                - auth = OAuth
        ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html


        ❌[AWS Best Practice][CodeBuild.2] CodeBuild project environment variables should not contain clear text credentials
                 - Found AWS_ACCESS_KEY_ID in PLANTEXT
        ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html

$ go run main.go --tfplan test-json/ec2.json --module ec2

        ❌[AWS Best Practice][EC2.3] Attached EBS volumes should be encrypted at-rest
                - EbsBlockDevice encrpyted is not enabled 
        ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
```