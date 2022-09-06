// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package braketiface provides an interface to enable mocking the Braket service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package braketiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/braket"
)

// BraketAPI provides an interface to enable mocking the
// braket.Braket service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Braket.
//	func myFunc(svc braketiface.BraketAPI) bool {
//	    // Make svc.CancelJob request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := braket.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockBraketClient struct {
//	    braketiface.BraketAPI
//	}
//	func (m *mockBraketClient) CancelJob(input *braket.CancelJobInput) (*braket.CancelJobOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockBraketClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BraketAPI interface {
	CancelJob(*braket.CancelJobInput) (*braket.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *braket.CancelJobInput, ...request.Option) (*braket.CancelJobOutput, error)
	CancelJobRequest(*braket.CancelJobInput) (*request.Request, *braket.CancelJobOutput)

	CancelQuantumTask(*braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error)
	CancelQuantumTaskWithContext(aws.Context, *braket.CancelQuantumTaskInput, ...request.Option) (*braket.CancelQuantumTaskOutput, error)
	CancelQuantumTaskRequest(*braket.CancelQuantumTaskInput) (*request.Request, *braket.CancelQuantumTaskOutput)

	CreateJob(*braket.CreateJobInput) (*braket.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *braket.CreateJobInput, ...request.Option) (*braket.CreateJobOutput, error)
	CreateJobRequest(*braket.CreateJobInput) (*request.Request, *braket.CreateJobOutput)

	CreateQuantumTask(*braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error)
	CreateQuantumTaskWithContext(aws.Context, *braket.CreateQuantumTaskInput, ...request.Option) (*braket.CreateQuantumTaskOutput, error)
	CreateQuantumTaskRequest(*braket.CreateQuantumTaskInput) (*request.Request, *braket.CreateQuantumTaskOutput)

	GetDevice(*braket.GetDeviceInput) (*braket.GetDeviceOutput, error)
	GetDeviceWithContext(aws.Context, *braket.GetDeviceInput, ...request.Option) (*braket.GetDeviceOutput, error)
	GetDeviceRequest(*braket.GetDeviceInput) (*request.Request, *braket.GetDeviceOutput)

	GetJob(*braket.GetJobInput) (*braket.GetJobOutput, error)
	GetJobWithContext(aws.Context, *braket.GetJobInput, ...request.Option) (*braket.GetJobOutput, error)
	GetJobRequest(*braket.GetJobInput) (*request.Request, *braket.GetJobOutput)

	GetQuantumTask(*braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error)
	GetQuantumTaskWithContext(aws.Context, *braket.GetQuantumTaskInput, ...request.Option) (*braket.GetQuantumTaskOutput, error)
	GetQuantumTaskRequest(*braket.GetQuantumTaskInput) (*request.Request, *braket.GetQuantumTaskOutput)

	ListTagsForResource(*braket.ListTagsForResourceInput) (*braket.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *braket.ListTagsForResourceInput, ...request.Option) (*braket.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*braket.ListTagsForResourceInput) (*request.Request, *braket.ListTagsForResourceOutput)

	SearchDevices(*braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error)
	SearchDevicesWithContext(aws.Context, *braket.SearchDevicesInput, ...request.Option) (*braket.SearchDevicesOutput, error)
	SearchDevicesRequest(*braket.SearchDevicesInput) (*request.Request, *braket.SearchDevicesOutput)

	SearchDevicesPages(*braket.SearchDevicesInput, func(*braket.SearchDevicesOutput, bool) bool) error
	SearchDevicesPagesWithContext(aws.Context, *braket.SearchDevicesInput, func(*braket.SearchDevicesOutput, bool) bool, ...request.Option) error

	SearchJobs(*braket.SearchJobsInput) (*braket.SearchJobsOutput, error)
	SearchJobsWithContext(aws.Context, *braket.SearchJobsInput, ...request.Option) (*braket.SearchJobsOutput, error)
	SearchJobsRequest(*braket.SearchJobsInput) (*request.Request, *braket.SearchJobsOutput)

	SearchJobsPages(*braket.SearchJobsInput, func(*braket.SearchJobsOutput, bool) bool) error
	SearchJobsPagesWithContext(aws.Context, *braket.SearchJobsInput, func(*braket.SearchJobsOutput, bool) bool, ...request.Option) error

	SearchQuantumTasks(*braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error)
	SearchQuantumTasksWithContext(aws.Context, *braket.SearchQuantumTasksInput, ...request.Option) (*braket.SearchQuantumTasksOutput, error)
	SearchQuantumTasksRequest(*braket.SearchQuantumTasksInput) (*request.Request, *braket.SearchQuantumTasksOutput)

	SearchQuantumTasksPages(*braket.SearchQuantumTasksInput, func(*braket.SearchQuantumTasksOutput, bool) bool) error
	SearchQuantumTasksPagesWithContext(aws.Context, *braket.SearchQuantumTasksInput, func(*braket.SearchQuantumTasksOutput, bool) bool, ...request.Option) error

	TagResource(*braket.TagResourceInput) (*braket.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *braket.TagResourceInput, ...request.Option) (*braket.TagResourceOutput, error)
	TagResourceRequest(*braket.TagResourceInput) (*request.Request, *braket.TagResourceOutput)

	UntagResource(*braket.UntagResourceInput) (*braket.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *braket.UntagResourceInput, ...request.Option) (*braket.UntagResourceOutput, error)
	UntagResourceRequest(*braket.UntagResourceInput) (*request.Request, *braket.UntagResourceOutput)
}

var _ BraketAPI = (*braket.Braket)(nil)
