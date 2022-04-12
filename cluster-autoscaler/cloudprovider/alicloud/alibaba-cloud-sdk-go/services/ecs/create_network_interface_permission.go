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

// CreateNetworkInterfacePermission invokes the ecs.CreateNetworkInterfacePermission API synchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterfacepermission.html
func (client *Client) CreateNetworkInterfacePermission(request *CreateNetworkInterfacePermissionRequest) (response *CreateNetworkInterfacePermissionResponse, err error) {
	response = CreateCreateNetworkInterfacePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNetworkInterfacePermissionWithChan invokes the ecs.CreateNetworkInterfacePermission API asynchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterfacepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNetworkInterfacePermissionWithChan(request *CreateNetworkInterfacePermissionRequest) (<-chan *CreateNetworkInterfacePermissionResponse, <-chan error) {
	responseChan := make(chan *CreateNetworkInterfacePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNetworkInterfacePermission(request)
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

// CreateNetworkInterfacePermissionWithCallback invokes the ecs.CreateNetworkInterfacePermission API asynchronously
// api document: https://help.aliyun.com/api/ecs/createnetworkinterfacepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNetworkInterfacePermissionWithCallback(request *CreateNetworkInterfacePermissionRequest, callback func(response *CreateNetworkInterfacePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNetworkInterfacePermissionResponse
		var err error
		defer close(result)
		response, err = client.CreateNetworkInterfacePermission(request)
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

// CreateNetworkInterfacePermissionRequest is the request struct for api CreateNetworkInterfacePermission
type CreateNetworkInterfacePermissionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountId            requests.Integer `position:"Query" name:"AccountId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Permission           string           `position:"Query" name:"Permission"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string           `position:"Query" name:"NetworkInterfaceId"`
}

// CreateNetworkInterfacePermissionResponse is the response struct for api CreateNetworkInterfacePermission
type CreateNetworkInterfacePermissionResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	NetworkInterfacePermission NetworkInterfacePermission `json:"NetworkInterfacePermission" xml:"NetworkInterfacePermission"`
}

// CreateCreateNetworkInterfacePermissionRequest creates a request to invoke CreateNetworkInterfacePermission API
func CreateCreateNetworkInterfacePermissionRequest() (request *CreateNetworkInterfacePermissionRequest) {
	request = &CreateNetworkInterfacePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterfacePermission", "ecs", "openAPI")
	return
}

// CreateCreateNetworkInterfacePermissionResponse creates a response to parse from CreateNetworkInterfacePermission response
func CreateCreateNetworkInterfacePermissionResponse() (response *CreateNetworkInterfacePermissionResponse) {
	response = &CreateNetworkInterfacePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
