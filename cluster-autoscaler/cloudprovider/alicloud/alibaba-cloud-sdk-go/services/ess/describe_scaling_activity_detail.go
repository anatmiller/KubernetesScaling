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

// DescribeScalingActivityDetail invokes the ess.DescribeScalingActivityDetail API synchronously
// api document: https://help.aliyun.com/api/ess/describescalingactivitydetail.html
func (client *Client) DescribeScalingActivityDetail(request *DescribeScalingActivityDetailRequest) (response *DescribeScalingActivityDetailResponse, err error) {
	response = CreateDescribeScalingActivityDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingActivityDetailWithChan invokes the ess.DescribeScalingActivityDetail API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingactivitydetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingActivityDetailWithChan(request *DescribeScalingActivityDetailRequest) (<-chan *DescribeScalingActivityDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingActivityDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingActivityDetail(request)
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

// DescribeScalingActivityDetailWithCallback invokes the ess.DescribeScalingActivityDetail API asynchronously
// api document: https://help.aliyun.com/api/ess/describescalingactivitydetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingActivityDetailWithCallback(request *DescribeScalingActivityDetailRequest, callback func(response *DescribeScalingActivityDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingActivityDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingActivityDetail(request)
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

// DescribeScalingActivityDetailRequest is the request struct for api DescribeScalingActivityDetail
type DescribeScalingActivityDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScalingActivityId    string           `position:"Query" name:"ScalingActivityId"`
}

// DescribeScalingActivityDetailResponse is the response struct for api DescribeScalingActivityDetail
type DescribeScalingActivityDetailResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	Detail            string `json:"Detail" xml:"Detail"`
}

// CreateDescribeScalingActivityDetailRequest creates a request to invoke DescribeScalingActivityDetail API
func CreateDescribeScalingActivityDetailRequest() (request *DescribeScalingActivityDetailRequest) {
	request = &DescribeScalingActivityDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingActivityDetail", "ess", "openAPI")
	return
}

// CreateDescribeScalingActivityDetailResponse creates a response to parse from DescribeScalingActivityDetail response
func CreateDescribeScalingActivityDetailResponse() (response *DescribeScalingActivityDetailResponse) {
	response = &DescribeScalingActivityDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
