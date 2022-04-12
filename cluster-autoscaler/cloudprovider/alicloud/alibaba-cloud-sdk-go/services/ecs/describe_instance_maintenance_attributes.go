package ecs

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

// DescribeInstanceMaintenanceAttributes invokes the ecs.DescribeInstanceMaintenanceAttributes API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemaintenanceattributes.html
func (client *Client) DescribeInstanceMaintenanceAttributes(request *DescribeInstanceMaintenanceAttributesRequest) (response *DescribeInstanceMaintenanceAttributesResponse, err error) {
	response = CreateDescribeInstanceMaintenanceAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceMaintenanceAttributesWithChan invokes the ecs.DescribeInstanceMaintenanceAttributes API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemaintenanceattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMaintenanceAttributesWithChan(request *DescribeInstanceMaintenanceAttributesRequest) (<-chan *DescribeInstanceMaintenanceAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceMaintenanceAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceMaintenanceAttributes(request)
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

// DescribeInstanceMaintenanceAttributesWithCallback invokes the ecs.DescribeInstanceMaintenanceAttributes API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemaintenanceattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMaintenanceAttributesWithCallback(request *DescribeInstanceMaintenanceAttributesRequest, callback func(response *DescribeInstanceMaintenanceAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceMaintenanceAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceMaintenanceAttributes(request)
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

// DescribeInstanceMaintenanceAttributesRequest is the request struct for api DescribeInstanceMaintenanceAttributes
type DescribeInstanceMaintenanceAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// DescribeInstanceMaintenanceAttributesResponse is the response struct for api DescribeInstanceMaintenanceAttributes
type DescribeInstanceMaintenanceAttributesResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	TotalCount            int                   `json:"TotalCount" xml:"TotalCount"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	MaintenanceAttributes MaintenanceAttributes `json:"MaintenanceAttributes" xml:"MaintenanceAttributes"`
}

// CreateDescribeInstanceMaintenanceAttributesRequest creates a request to invoke DescribeInstanceMaintenanceAttributes API
func CreateDescribeInstanceMaintenanceAttributesRequest() (request *DescribeInstanceMaintenanceAttributesRequest) {
	request = &DescribeInstanceMaintenanceAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceMaintenanceAttributes", "ecs", "openAPI")
	return
}

// CreateDescribeInstanceMaintenanceAttributesResponse creates a response to parse from DescribeInstanceMaintenanceAttributes response
func CreateDescribeInstanceMaintenanceAttributesResponse() (response *DescribeInstanceMaintenanceAttributesResponse) {
	response = &DescribeInstanceMaintenanceAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
