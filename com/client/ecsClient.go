/*
 *Copyright 2015 Huawei Technologies Co., Ltd. All rights reserved.
 *	   eSDK is licensed under the Apache License, Version 2.0 (the "License");
 *	   you may not use this file except in compliance with the License.
 *	   You may obtain a copy of the License at
 *
 *	       http://www.apache.org/licenses/LICENSE-2.0
 *
 *
 *	   Unless required by applicable law or agreed to in writing, software
 *	   distributed under the License is distributed on an "AS IS" BASIS,
 *	   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *	   See the License for the specific language governing permissions and
 *	   limitations under the License.
 */
package client

import (
	"encoding/json"
	"github.com/docker/machine/libmachine/log"
	"github.com/huawei/DockerMachineDriver4OTC/com/cpp_sdk_core"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/ecsModules"
)

func (client *Client) CreateCloudServer(createCloudServerReq ecsModules.CreateCloudServerReq) ecsModules.CreateCloudServerResp {
	createCloudServerResp := ecsModules.CreateCloudServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/cloudservers"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = createCloudServerReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		createCloudServerResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &createCloudServerResp)
	} else {
		createCloudServerResp.ResponseCode = Result.ResponseCode
		log.Errorf("CreateCloudServer failed with: %s", Result.RespMessage)
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		createCloudServerResp.ErrorInfo = Error
	}

	return createCloudServerResp
}

func (client *Client) DeleteCloudServer(deleteCloudServerReq ecsModules.DeleteCloudServerReq) ecsModules.DeleteCloudServerResp {
	deleteCloudServerResp := ecsModules.DeleteCloudServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/cloudservers/delete"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = deleteCloudServerReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		deleteCloudServerResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &deleteCloudServerResp)
	} else {
		deleteCloudServerResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		deleteCloudServerResp.ErrorInfo = Error
	}

	return deleteCloudServerResp
}

func (client *Client) ShowEcsJob(job_id string) ecsModules.ShowEcsJobResp {
	showEcsJobResp := ecsModules.ShowEcsJobResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/jobs/" + job_id
	client.RequestParam.Method = modules.HTTP_GET
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		showEcsJobResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &showEcsJobResp)
		log.Debugf("get ECS job status: %v", Result.RespMessage)
	} else {
		showEcsJobResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		showEcsJobResp.ErrorInfo = Error
		log.Debugf("get ECS job status failed: %v", Result.RespMessage)
	}

	return showEcsJobResp
}

func (client *Client) ListCloudServerFlavorsExt() ecsModules.ListCloudServerFlavorsExtResp {
	listCloudServerFlavorsExtResp := ecsModules.ListCloudServerFlavorsExtResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/cloudservers/flavors"
	client.RequestParam.Method = modules.HTTP_GET
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		listCloudServerFlavorsExtResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &listCloudServerFlavorsExtResp)
	} else {
		listCloudServerFlavorsExtResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		listCloudServerFlavorsExtResp.ErrorInfo = Error
	}

	return listCloudServerFlavorsExtResp
}
