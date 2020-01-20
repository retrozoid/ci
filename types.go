package ci

// Property ...
type Property struct {
	Name      string        `json:"name"`
	Value     string        `json:"value"`
	Inherited bool          `json:"inherited"`
	Type      *PropertyType `json:"type,omitempty"`
}

// Properties ...
type Properties struct {
	Property []*Property `json:"property"`
	Count    int64       `json:"count"`
	Href     string      `json:"href"`
}

// PropertyType ...
type PropertyType struct {
	RawValue string `json:"rawValue"`
}

// TriggeredBy ...
type TriggeredBy struct {
}

// TestOccurrences ...
type TestOccurrences struct {
}

// RunningInfo ...
type RunningInfo struct {
	PercentageComplete    int64  `json:"percentageComplete"`
	ElapsedSeconds        int64  `json:"elapsedSeconds"`
	EstimatedTotalSeconds int64  `json:"estimatedTotalSeconds"`
	CurrentStageText      string `json:"currentStageText"`
	Outdated              bool   `json:"outdated"`
	ProbablyHanging       bool   `json:"probablyHanging"`
}

// Build ...
type Build struct {
	StatusText          string           `json:"statusText"`
	StartEstimate       string           `json:"startEstimate"`
	QueuedDate          *RunningInfo     `json:"running-info"`
	StartDate           string           `json:"startDate"`
	FinishDate          string           `json:"finishDate"`
	Triggered           *TriggeredBy     `json:"triggered"`
	TestOccurrences     *TestOccurrences `json:"testOccurrences"`
	SettingsHash        string           `json:"settingsHash"`
	CurrentSettingsHash string           `json:"currentSettingsHash"`
	ModificationID      string           `json:"modificationId"`
	ChainModificationID string           `json:"chainModificationId"`
	ID                  int64            `json:"id"`
	TaskID              int64            `json:"taskId"`
	BuildTypeID         string           `json:"buildTypeId"`
	BuildTypeInternalID string           `json:"buildTypeInternalId"`
	Number              string           `json:"number"`
	Status              string           `json:"status"`
	State               string           `json:"state"`
	Running             bool             `json:"running"`
	FailedToStart       bool             `json:"failedToStart"`
	Personal            bool             `json:"personal"`
	PercentageComplete  int64            `json:"percentageComplete"`
	BranchName          string           `json:"branchName"`
	DefaultBranch       bool             `json:"defaultBranch"`
	UnspecifiedBranch   bool             `json:"unspecifiedBranch"`
	History             bool             `json:"history"`
	Pinned              bool             `json:"pinned"`
	Href                string           `json:"href"`
	WebURL              string           `json:"webUrl"`
	Locator             string           `json:"locator"`
	Properties          *Properties      `json:"properties"`
}
