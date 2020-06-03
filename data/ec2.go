package data

type EC2 struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	PlannedValues    struct {
		RootModule struct {
			ChildModules []struct {
				Resources []struct {
					Address       string `json:"address"`
					Mode          string `json:"mode"`
					Type          string `json:"type"`
					Name          string `json:"name"`
					Index         int    `json:"index"`
					ProviderName  string `json:"provider_name"`
					SchemaVersion int    `json:"schema_version"`
					Values        struct {
						Ami                 string `json:"ami"`
						CreditSpecification []struct {
							CPUCredits string `json:"cpu_credits"`
						} `json:"credit_specification"`
						DisableAPITermination bool `json:"disable_api_termination"`
						EbsBlockDevice        []struct {
							DeleteOnTermination bool   `json:"delete_on_termination"`
							DeviceName          string `json:"device_name"`
							Encrypted           bool   `json:"encrypted"`
							VolumeSize          int    `json:"volume_size"`
							VolumeType          string `json:"volume_type"`
						} `json:"ebs_block_device"`
						EbsOptimized                      bool        `json:"ebs_optimized"`
						GetPasswordData                   bool        `json:"get_password_data"`
						Hibernation                       interface{} `json:"hibernation"`
						IamInstanceProfile                string      `json:"iam_instance_profile"`
						InstanceInitiatedShutdownBehavior string      `json:"instance_initiated_shutdown_behavior"`
						InstanceType                      string      `json:"instance_type"`
						KeyName                           string      `json:"key_name"`
						Monitoring                        bool        `json:"monitoring"`
						SourceDestCheck                   bool        `json:"source_dest_check"`
						SubnetID                          string      `json:"subnet_id"`
						Tags                              struct {
							Environment string `json:"Environment"`
							Name        string `json:"Name"`
							Terraform   string `json:"Terraform"`
						} `json:"tags"`
						Tenancy        string      `json:"tenancy"`
						Timeouts       interface{} `json:"timeouts"`
						UserData       interface{} `json:"user_data"`
						UserDataBase64 interface{} `json:"user_data_base64"`
						VolumeTags     struct {
							Name string `json:"Name"`
						} `json:"volume_tags"`
						VpcSecurityGroupIds []string `json:"vpc_security_group_ids"`
					} `json:"values"`
				} `json:"resources"`
				Address string `json:"address"`
			} `json:"child_modules"`
		} `json:"root_module"`
	} `json:"planned_values"`
	ResourceChanges []interface {
		} `json:"resource_changes"`
	Configuration interface {
	} `json:"configuration"`
}