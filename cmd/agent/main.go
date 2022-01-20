package main

import (
	"os"

	"github.com/minhhoangvn/testkube-executor-k6/pkg/runner"
	"github.com/kubeshop/testkube/pkg/runner/agent"
)

func main() {
	agent.Run(runner.NewRunner(), os.Args)
}
