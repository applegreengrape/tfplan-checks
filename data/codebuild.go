package data

type CodeBuild struct {
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
						Artifacts []struct {
							ArtifactIdentifier   interface{} `json:"artifact_identifier"`
							EncryptionDisabled   bool        `json:"encryption_disabled"`
							Location             interface{} `json:"location"`
							Name                 interface{} `json:"name"`
							NamespaceType        interface{} `json:"namespace_type"`
							OverrideArtifactName bool        `json:"override_artifact_name"`
							Packaging            interface{} `json:"packaging"`
							Path                 interface{} `json:"path"`
							Type                 string      `json:"type"`
						} `json:"artifacts"`
						BadgeEnabled bool `json:"badge_enabled"`
						BuildTimeout int  `json:"build_timeout"`
						Cache        []struct {
							Location interface{} `json:"location"`
							Modes    interface{} `json:"modes"`
							Type     string      `json:"type"`
						} `json:"cache"`
						Environment []struct {
							Certificate         interface{} `json:"certificate"`
							ComputeType         string      `json:"compute_type"`
							EnvironmentVariable []struct {
								Name  string `json:"name"`
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"environment_variable"`
							Image                    string        `json:"image"`
							ImagePullCredentialsType string        `json:"image_pull_credentials_type"`
							PrivilegedMode           bool          `json:"privileged_mode"`
							RegistryCredential       []interface{} `json:"registry_credential"`
							Type                     string        `json:"type"`
						} `json:"environment"`
						LogsConfig         []interface{} `json:"logs_config"`
						Name               string        `json:"name"`
						QueuedTimeout      int           `json:"queued_timeout"`
						SecondaryArtifacts []interface{} `json:"secondary_artifacts"`
						SecondarySources   []interface{} `json:"secondary_sources"`
						Source             []struct {
							Auth                []interface{} `json:"auth"`
							Buildspec           string        `json:"buildspec"`
							GitCloneDepth       interface{}   `json:"git_clone_depth"`
							GitSubmodulesConfig []interface{} `json:"git_submodules_config"`
							InsecureSsl         interface{}   `json:"insecure_ssl"`
							Location            string        `json:"location"`
							ReportBuildStatus   bool          `json:"report_build_status"`
							Type                string        `json:"type"`
						} `json:"source"`
						SourceVersion interface{} `json:"source_version"`
						Tags          struct {
							Name      string `json:"Name"`
							Namespace string `json:"Namespace"`
							Stage     string `json:"Stage"`
						} `json:"tags"`
						VpcConfig []interface{} `json:"vpc_config"`
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
	