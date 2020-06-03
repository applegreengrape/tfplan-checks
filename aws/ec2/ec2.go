package ec2

import (
	"tfplan/checks/data"
	"github.com/fatih/color"
)

//[EC2.3] Attached EBS volumes should be encrypted at-rest
func EC2_03(plan data.EC2)  {
	s:= "okay"
	for _, m:= range plan.PlannedValues.RootModule.ChildModules{
		for _,r:= range m.Resources{
			if len(r.Values.EbsBlockDevice) == 0 {
				s = "not okay"
			}else{
				for _,e :=range r.Values.EbsBlockDevice{
					if e.Encrypted {
						s = "okay"
					}else{
						s = "not okay"
					}
				}
			}
		}
	}

	if s == "okay"{
		r:=`
	✅[AWS Best Practice][EC2.3] Attached EBS volumes should be encrypted at-rest
		`
		color.Green(r)
	}else{
		r:=`
	❌[AWS Best Practice][EC2.3] Attached EBS volumes should be encrypted at-rest
		- EbsBlockDevice encrpyted is not enabled 
	ref: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-fsbp-controls.html
		
		`
		color.Yellow(r)
	}
	
}