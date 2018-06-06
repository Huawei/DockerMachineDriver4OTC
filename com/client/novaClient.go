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
	"fmt"
	"github.com/huawei/DockerMachineDriver4OTC/com/cpp_sdk_core"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/novaModules"
)

func (client *Client) ListInterfaces(server_id string) novaModules.ListInterfacesResp {
	listInterfacesResp := novaModules.ListInterfacesResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id + "/os-interface"
	client.RequestParam.Method = modules.HTTP_GET
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		listInterfacesResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &listInterfacesResp)
	} else {
		listInterfacesResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		listInterfacesResp.ErrorInfo = Error
	}

	return listInterfacesResp
}

func (client *Client) CreateKeypair(createKeypairReq novaModules.CreateKeypairReq) novaModules.CreateKeypairResp {
	createKeypairResp := novaModules.CreateKeypairResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/os-keypairs"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = createKeypairReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		createKeypairResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &createKeypairResp)
	} else {
		createKeypairResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		createKeypairResp.ErrorInfo = Error
	}

	return createKeypairResp
}

func (client *Client) DeleteKeyPair(keypair_name string) novaModules.DeleteKeyPairResp {
	deleteKeyPairResp := novaModules.DeleteKeyPairResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/os-keypairs/" + keypair_name
	client.RequestParam.Method = modules.HTTP_DELETE
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		deleteKeyPairResp.ResponseCode = modules.HttpOK
	} else {
		deleteKeyPairResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		deleteKeyPairResp.ErrorInfo = Error
	}

	return deleteKeyPairResp
}

func (client *Client) StartServer(server_id string) novaModules.StartServerResp {
	startServer := novaModules.StartServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id + "/action"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = "{ \"os-start\": {} }"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		startServer.ResponseCode = modules.HttpOK
	} else {
		startServer.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		startServer.ErrorInfo = Error
	}

	return startServer
}

func (client *Client) StopServer(server_id string) novaModules.StopServerResp {
	stopServer := novaModules.StopServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id + "/action"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = "{ \"os-stop\": {} }"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		stopServer.ResponseCode = modules.HttpOK
	} else {
		stopServer.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		stopServer.ErrorInfo = Error
	}

	return stopServer
}

func (client *Client) RebootServer(server_id, rebootType string) novaModules.RebootServerResp {
	rebootServer := novaModules.RebootServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id + "/action"
	client.RequestParam.Method = modules.HTTP_POST
	client.RequestParam.RequestContentType = modules.ApplicationJson
	bodyMap := make(map[string]map[string]string)
	typeMap := make(map[string]string)
	typeMap["type"] = rebootType
	bodyMap["reboot"] = typeMap
	bodyContent, _ := json.Marshal(bodyMap)
	strBody := string(bodyContent)
	client.RequestParam.BodyContent = strBody

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		rebootServer.ResponseCode = modules.HttpOK
	} else {
		rebootServer.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		rebootServer.ErrorInfo = Error
	}

	return rebootServer
}

func (client *Client) DeleteAnServer(server_id string) novaModules.DeleteAnServerResp {
	deleteAnServerResp := novaModules.DeleteAnServerResp{}

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id
	client.RequestParam.Method = modules.HTTP_DELETE
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		deleteAnServerResp.ResponseCode = modules.HttpOK
	} else {
		deleteAnServerResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		deleteAnServerResp.ErrorInfo = Error
	}

	return deleteAnServerResp
}

func (client *Client) ShowServer(server_id string) novaModules.ShowServerResp {
	showServerResp := novaModules.ShowServerResp{}
	showServerResp.Server.Power_state = -1
	showServerResp.Server.EvsOpts = -1
	showServerResp.Server.NumaOpts = -1
	showServerResp.Server.Progress = -1

	client.RequestParam.Url = "https://ecs." + client.RequestParam.Endpoint + "/v2/" + client.TenantID + "/servers/" + server_id
	client.RequestParam.Method = modules.HTTP_GET
	client.RequestParam.RequestContentType = modules.ApplicationJson
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		showServerResp.ResponseCode = modules.HttpOK
		json.Unmarshal([]byte(Result.RespMessage), &showServerResp)
		fmt.Println(Result.RespMessage)
	} else {
		showServerResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		showServerResp.ErrorInfo = Error
	}

	return showServerResp
}
