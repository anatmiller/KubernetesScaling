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

// ModifyInstanceSpec invokes the ecs.ModifyInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancespec.html
func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (response *ModifyInstanceSpecResponse, err error) {
	response = CreateModifyInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceSpecWithChan invokes the ecs.ModifyInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceSpecWithChan(request *ModifyInstanceSpecRequest) (<-chan *ModifyInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceSpec(request)
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

// ModifyInstanceSpecWithCallback invokes the ecs.ModifyInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceSpecWithCallback(request *ModifyInstanceSpecRequest, callback func(response *ModifyInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceSpec(request)
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

// ModifyInstanceSpecRequest is the request struct for api ModifyInstanceSpec
type ModifyInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                      string           `position:"Query" name:"ClientToken"`
	AllowMigrateAcrossZone           requests.Boolean `position:"Query" name:"AllowMigrateAcrossZone"`
	InternetMaxBandwidthOut          requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskCategory               string           `position:"Query" name:"SystemDisk.Category"`
	InstanceType                     string           `position:"Query" name:"InstanceType"`
	TemporaryEndTime                 string           `position:"Query" name:"Temporary.EndTime"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                     string           `position:"Query" name:"OwnerAccount"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	TemporaryInternetMaxBandwidthOut requests.Integer `position:"Query" name:"Temporary.InternetMaxBandwidthOut"`
	TemporaryStartTime               string           `position:"Query" name:"Temporary.StartTime"`
	Async                            requests.Boolean `position:"Query" name:"Async"`
	InstanceId                       string           `position:"Query" name:"InstanceId"`
	InternetMaxBandwidthIn           requests.Integer `position:"Query" name:"InternetMaxBandwidthIn"`
}

// ModifyInstanceSpecResponse is the response struct for api ModifyInstanceSpec
type ModifyInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceSpecRequest creates a request to invoke ModifyInstanceSpec API
func CreateModifyInstanceSpecRequest() (request *ModifyInstanceSpecRequest) {
	request = &ModifyInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceSpec", "ecs", "openAPI")
	return
}

// CreateModifyInstanceSpecResponse creates a response to parse from ModifyInstanceSpec response
func CreateModifyInstanceSpecResponse() (response *ModifyInstanceSpecResponse) {
	response = &ModifyInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
