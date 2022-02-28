package runner

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

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
	r.getClusterConfig()
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

func (r *K6Runner) getClusterConfig() {
	b, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(b))
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Pod example-xxxxx not found in default namespace\n")
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(config.String())
	fmt.Printf("%+v\n", config)
	countPrint := 1
	for countPrint < 5 {
		countPrint++
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for idx, pod := range pods.Items {
			fmt.Println(idx, pod.Name)
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("testkube").Get(context.TODO(), "testkube-api-server-876676fd5-fgxm5", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod testkube-api-server not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found testkube-api-server pod in default namespace\n")
		}

		time.Sleep(10 * time.Second)
	}
}
