package rds

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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAvailableInstanceClass invokes the rds.DescribeAvailableInstanceClass API synchronously
// api document: https://help.aliyun.com/api/rds/describeavailableinstanceclass.html
func (client *Client) DescribeAvailableInstanceClass(request *DescribeAvailableInstanceClassRequest) (response *DescribeAvailableInstanceClassResponse, err error) {
	response = CreateDescribeAvailableInstanceClassResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableInstanceClassWithChan invokes the rds.DescribeAvailableInstanceClass API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailableinstanceclass.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableInstanceClassWithChan(request *DescribeAvailableInstanceClassRequest) (<-chan *DescribeAvailableInstanceClassResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableInstanceClassResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableInstanceClass(request)
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

// DescribeAvailableInstanceClassWithCallback invokes the rds.DescribeAvailableInstanceClass API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailableinstanceclass.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableInstanceClassWithCallback(request *DescribeAvailableInstanceClassRequest, callback func(response *DescribeAvailableInstanceClassResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableInstanceClassResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableInstanceClass(request)
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

// DescribeAvailableInstanceClassRequest is the request struct for api DescribeAvailableInstanceClass
type DescribeAvailableInstanceClassRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	Engine               string           `position:"Query" name:"Engine"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// DescribeAvailableInstanceClassResponse is the response struct for api DescribeAvailableInstanceClass
type DescribeAvailableInstanceClassResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	AvailableZones []AvailableZone `json:"AvailableZones" xml:"AvailableZones"`
}

// CreateDescribeAvailableInstanceClassRequest creates a request to invoke DescribeAvailableInstanceClass API
func CreateDescribeAvailableInstanceClassRequest() (request *DescribeAvailableInstanceClassRequest) {
	request = &DescribeAvailableInstanceClassRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableInstanceClass", "rds", "openAPI")
	return
}

// CreateDescribeAvailableInstanceClassResponse creates a response to parse from DescribeAvailableInstanceClass response
func CreateDescribeAvailableInstanceClassResponse() (response *DescribeAvailableInstanceClassResponse) {
	response = &DescribeAvailableInstanceClassResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
