package runner

import (
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/log"
	"go.uber.org/zap"
)

func NewRunner() *K6Runner {
	return &K6Runner{Log: log.DefaultLogger}
}

// ExampleRunner for template - change me to some valid runner
type K6Runner struct {
	Log *zap.SugaredLogger
}

func (r *K6Runner) Run(execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	output, err := executor.Run("", "k6", "--help")
	if err != nil {
		r.Log.Errorf("Error occured when running a command %s", err)
		return result.Err(err), nil
	}
	outputString := string(output)
	result.Output = outputString
	return testkube.ExecutionResult{
		Status: testkube.StatusPtr(testkube.SUCCESS_ExecutionStatus),
		Output: outputString,
	}, nil
}
