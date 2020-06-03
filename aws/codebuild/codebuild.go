package codebuild

import (
	"fmt"
	"tfplan/checks/data"
	"github.com/fatih/color"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
	   if str == a {
		  return true
	   }
	}
	return false
 }

//[CodeBuild.1] CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
func CodeBuild_01(plan data.CodeBuild) {
	var authList []string
	for i:= range plan.PlannedValues.RootModule.ChildModules{
		for n:= range plan.PlannedValues.RootModule.ChildModules[i].Resources{
			if plan.PlannedValues.RootModule.ChildModules[i].Resources[n].Type == "aws_codebuild_project" {
				for l:=range plan.PlannedValues.RootModule.ChildModules[i].Resources[n].Values.Source{
					auth:= plan.PlannedValues.RootModule.ChildModules[i].Resources[n].Values.Source[l].Auth
					for _, a:= range auth{
						auth:= fmt.Sprintf("%v", a)
						authList = append(authList, auth)
					}
				}
			}
		}
	}
	if !contains(authList,"OAuth"){
		r:=`
	❌[AWS Best Practice][CodeBuild.1] CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
		- auth = OAuth
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
		`
		color.Yellow(r)
	}else{
		r:=`
	✅[AWS Best Practice][CodeBuild.1] CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
		`
		color.Green(r)
	}
}

//[CodeBuild.2] CodeBuild project environment variables should not contain clear text credentials
func CodeBuild_02(plan data.CodeBuild){
	s:= "okay"
	for _, m:= range plan.PlannedValues.RootModule.ChildModules{
		for _,r:= range m.Resources{
			if r.Type == "aws_codebuild_project"{
				for _,e:= range r.Values.Environment {
					for _, v:=range e.EnvironmentVariable{
						if (v.Name == "AWS_ACCESS_KEY_ID" && v.Type == "PLAINTEXT"){
							r:= `
	❌[AWS Best Practice][CodeBuild.2] CodeBuild project environment variables should not contain clear text credentials
		 - Found AWS_ACCESS_KEY_ID in PLANTEXT
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
							`
						color.Yellow(r)
						s = "not okay"
						}
						if (v.Name == "AWS_SECRET_ACCESS_KEY" && v.Type == "PLAINTEXT"){
							r:= `
	❌[AWS Best Practice][CodeBuild.2] CodeBuild project environment variables should not contain clear text credentials
		 - Found AWS_SECRET_ACCESS_KEY in PLANTEXT
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
							`
						color.Yellow(r)
						s = "not okay"
						}
					}
				}
			}
		}
	}
	if s == "okay"{
		r:=`
	✅[AWS Best Practice][CodeBuild.2] CodeBuild project environment variables should not contain clear text credentials
		- No credentials AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY in Plantext in environment variables
		`
	color.Green(r)
	}

}