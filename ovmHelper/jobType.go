package ovmhelper

import "fmt"

// JobError - interface for a failed job
type JobError struct {
	Message string `json:"message"`
}

func (j *JobError) Error() string {
	return fmt.Sprintf("error, %v", j.Message)
}

func (j *JobResponse) succeed() bool {
	if j.JobRunState != "SUCCESS" {
		return false
	}
	return true
}

// JobResponse - Job response interface
type JobResponse struct {
	ID                           *ID       `json:"id"`
	Done                         bool      `json:"done,omitempty"`
	ResultID                     *ID       `json:"resultId,omitempty"`
	ResourceGroupIds             *[]ID     `json:"resourceGroupIds,omitempty"`
	SummaryDone                  bool      `json:"summaryDone,omitempty"`
	JobGroup                     bool      `json:"jobGroup,omitempty"`
	JobRunState                  string    `json:"jobRunState,omitempty"`
	JobSummaryState              string    `json:"jobSummaryState,omitempty"`
	AbortedByUser                string    `json:"abortedByUser,omitempty"`
	ExtraInfo                    string    `json:"extraInfo,omitempty"`
	Name                         string    `json:"name,omitempty"`
	Description                  string    `json:"description,omitempty"`
	Locked                       bool      `json:"locked,omitempty"`
	ReadOnly                     bool      `json:"readOnly,omitempty"`
	Generation                   int       `json:"generation,omitempty"`
	ProgressMessage              string    `json:"progressMessage,omitempty"`
	LatestSummaryProgressMessage string    `json:"latestSummaryProgressMessage,omitempty"`
	StartTime                    int64     `json:"startTime,omitempty"`
	EndTime                      int64     `json:"endTime,omitempty"`
	ParentJobID                  *ID       `json:"parentJobId,omitempty"`
	ChildJobIds                  *[]ID     `json:"childJobIds,omitempty"`
	Error                        *JobError `json:"error,omitempty"`
	User                         string    `json:"user,omitempty"`
	WsErrorCode                  string    `json:"wsErrorCode,omitempty"`
}
