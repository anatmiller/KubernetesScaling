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

// ImportSnapshot invokes the ecs.ImportSnapshot API synchronously
// api document: https://help.aliyun.com/api/ecs/importsnapshot.html
func (client *Client) ImportSnapshot(request *ImportSnapshotRequest) (response *ImportSnapshotResponse, err error) {
	response = CreateImportSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// ImportSnapshotWithChan invokes the ecs.ImportSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/importsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportSnapshotWithChan(request *ImportSnapshotRequest) (<-chan *ImportSnapshotResponse, <-chan error) {
	responseChan := make(chan *ImportSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportSnapshot(request)
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

// ImportSnapshotWithCallback invokes the ecs.ImportSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/importsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportSnapshotWithCallback(request *ImportSnapshotRequest, callback func(response *ImportSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportSnapshotResponse
		var err error
		defer close(result)
		response, err = client.ImportSnapshot(request)
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

// ImportSnapshotRequest is the request struct for api ImportSnapshot
type ImportSnapshotRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SnapshotName         string           `position:"Query" name:"SnapshotName"`
	OssObject            string           `position:"Query" name:"OssObject"`
	OssBucket            string           `position:"Query" name:"OssBucket"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string           `position:"Query" name:"RoleName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ImportSnapshotResponse is the response struct for api ImportSnapshot
type ImportSnapshotResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
	SnapshotId string `json:"SnapshotId" xml:"SnapshotId"`
}

// CreateImportSnapshotRequest creates a request to invoke ImportSnapshot API
func CreateImportSnapshotRequest() (request *ImportSnapshotRequest) {
	request = &ImportSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ImportSnapshot", "ecs", "openAPI")
	return
}

// CreateImportSnapshotResponse creates a response to parse from ImportSnapshot response
func CreateImportSnapshotResponse() (response *ImportSnapshotResponse) {
	response = &ImportSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
