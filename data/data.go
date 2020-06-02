package data

type WithOutMoudle struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	PlannedValues    struct {
		RootModule struct {
			Resources []interface {
			} `json:"resources"`
		} `json:"root_module"`
	} `json:"planned_values"`
	ResourceChanges []interface {
	} `json:"resource_changes"`
	PriorState interface {
	} `json:"prior_state"`
	Configuration interface {
	} `json:"configuration"`
}

type WithMoudle struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	PlannedValues    struct {
		RootModule struct {
			ChildModules []struct {
				Resources []interface {
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

type Result struct {
	Type []interface{}
	Values  []interface{}
}
