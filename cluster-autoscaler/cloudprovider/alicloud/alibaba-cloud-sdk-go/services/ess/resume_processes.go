package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/requests"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud/alibaba-cloud-sdk-go/sdk/responses"
)

// ResumeProcesses invokes the ess.ResumeProcesses API synchronously
// api document: https://help.aliyun.com/api/ess/resumeprocesses.html
func (client *Client) ResumeProcesses(request *ResumeProcessesRequest) (response *ResumeProcessesResponse, err error) {
	response = CreateResumeProcessesResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeProcessesWithChan invokes the ess.ResumeProcesses API asynchronously
// api document: https://help.aliyun.com/api/ess/resumeprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeProcessesWithChan(request *ResumeProcessesRequest) (<-chan *ResumeProcessesResponse, <-chan error) {
	responseChan := make(chan *ResumeProcessesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeProcesses(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ResumeProcessesWithCallback invokes the ess.ResumeProcesses API asynchronously
// api document: https://help.aliyun.com/api/ess/resumeprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeProcessesWithCallback(request *ResumeProcessesRequest, callback func(response *ResumeProcessesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeProcessesResponse
		var err error
		defer close(result)
		response, err = client.ResumeProcesses(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ResumeProcessesRequest is the request struct for api ResumeProcesses
type ResumeProcessesRequest struct {
	*requests.RpcRequest
	Process              *[]string        `position:"Query" name:"Process"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ResumeProcessesResponse is the response struct for api ResumeProcesses
type ResumeProcessesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResumeProcessesRequest creates a request to invoke ResumeProcesses API
func CreateResumeProcessesRequest() (request *ResumeProcessesRequest) {
	request = &ResumeProcessesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ResumeProcesses", "ess", "openAPI")
	return
}

// CreateResumeProcessesResponse creates a response to parse from ResumeProcesses response
func CreateResumeProcessesResponse() (response *ResumeProcessesResponse) {
	response = &ResumeProcessesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
