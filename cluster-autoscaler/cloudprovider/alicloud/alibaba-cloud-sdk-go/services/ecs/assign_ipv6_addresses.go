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

// AssignIpv6Addresses invokes the ecs.AssignIpv6Addresses API synchronously
// api document: https://help.aliyun.com/api/ecs/assignipv6addresses.html
func (client *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
	response = CreateAssignIpv6AddressesResponse()
	err = client.DoAction(request, response)
	return
}

// AssignIpv6AddressesWithChan invokes the ecs.AssignIpv6Addresses API asynchronously
// api document: https://help.aliyun.com/api/ecs/assignipv6addresses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignIpv6AddressesWithChan(request *AssignIpv6AddressesRequest) (<-chan *AssignIpv6AddressesResponse, <-chan error) {
	responseChan := make(chan *AssignIpv6AddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignIpv6Addresses(request)
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

// AssignIpv6AddressesWithCallback invokes the ecs.AssignIpv6Addresses API asynchronously
// api document: https://help.aliyun.com/api/ecs/assignipv6addresses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignIpv6AddressesWithCallback(request *AssignIpv6AddressesRequest, callback func(response *AssignIpv6AddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignIpv6AddressesResponse
		var err error
		defer close(result)
		response, err = client.AssignIpv6Addresses(request)
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

// AssignIpv6AddressesRequest is the request struct for api AssignIpv6Addresses
type AssignIpv6AddressesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Ipv6AddressCount     requests.Integer `position:"Query" name:"Ipv6AddressCount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string           `position:"Query" name:"NetworkInterfaceId"`
	Ipv6Address          *[]string        `position:"Query" name:"Ipv6Address"  type:"Repeated"`
}

// AssignIpv6AddressesResponse is the response struct for api AssignIpv6Addresses
type AssignIpv6AddressesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssignIpv6AddressesRequest creates a request to invoke AssignIpv6Addresses API
func CreateAssignIpv6AddressesRequest() (request *AssignIpv6AddressesRequest) {
	request = &AssignIpv6AddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AssignIpv6Addresses", "ecs", "openAPI")
	return
}

// CreateAssignIpv6AddressesResponse creates a response to parse from AssignIpv6Addresses response
func CreateAssignIpv6AddressesResponse() (response *AssignIpv6AddressesResponse) {
	response = &AssignIpv6AddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
