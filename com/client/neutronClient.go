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
	"github.com/huawei/DockerMachineDriver4OTC/com/cpp_sdk_core"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/neutronModules"
)

/*
 * @fn			func (client *Client) CreateSecurityGroupRule(createSecurityGroupRuleReq
				*neutronModules.CreateSecurityGroupRuleReq) *neutronModules.CreateSecurityGroupRuleResp
 * @brief		This interface is used to create a security group rule.
 * @param[in] 	createSecurityGroupRuleReq: URI parameters
 * @param[out]
 * @return		*neutronModules.CreateSecurityGroupRuleResp
*/
func (client *Client) CreateSecurityGroupRule(createSecurityGroupRuleReq *neutronModules.CreateSecurityGroupRuleReq) *neutronModules.CreateSecurityGroupRuleResp {
	CreateSecurityGroupRuleResp := &neutronModules.CreateSecurityGroupRuleResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v2.0/security-group-rules"

	client.RequestParam.Method = "POST"
	client.RequestParam.BodyContent = createSecurityGroupRuleReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		CreateSecurityGroupRuleResp.ResponseCode = modules.HttpOK
		var data modules.NeutronSecurityGroupRuleInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		CreateSecurityGroupRuleResp.SecurityGroupRuleCreateInfo = data.SecurityGroupRuleCreateInfo
	} else {
		CreateSecurityGroupRuleResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		CreateSecurityGroupRuleResp.ErrorInfo = Error
	}

	return CreateSecurityGroupRuleResp
}

/*
 * @fn			func (client *Client) ShowSecurityGroupRule(rules_security_groups_id string)
				*neutronModules.ShowSecurityGroupRuleResp
 * @brief		This interface is used to query details about a security group rule.
 * @param[in] 	rules_security_groups_id: Specifies the security group rule ID, which
									  uniquely identifies the security group rule.
 * @param[out]
 * @return		*neutronModules.ShowSecurityGroupRuleResp
*/
func (client *Client) ShowSecurityGroupRule(rules_security_groups_id string) *neutronModules.ShowSecurityGroupRuleResp {
	ShowSecurityGroupRuleResp := &neutronModules.ShowSecurityGroupRuleResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v2.0/security-group-rules/" + rules_security_groups_id

	client.RequestParam.Method = "GET"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ShowSecurityGroupRuleResp.ResponseCode = modules.HttpOK
		var data modules.NeutronSecurityGroupRuleInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ShowSecurityGroupRuleResp.SecurityGroupRuleCreateInfo = data.SecurityGroupRuleCreateInfo
	} else {
		ShowSecurityGroupRuleResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ShowSecurityGroupRuleResp.ErrorInfo = Error
	}

	return ShowSecurityGroupRuleResp
}

/*
 * @fn			func (client *Client) DeleteSecurityGroupRule(rules_security_groups_id string)
				*neutronModules.DeleteSecurityGroupRuleResp
 * @brief		This interface is used to delete a security group rule.
 * @param[in] 	rules_security_groups_id: Specifies the security group rule ID, which
									  uniquely identifies the security group rule.
 * @param[out]
 * @return		*neutronModules.DeleteSecurityGroupRuleResp
*/
func (client *Client) DeleteSecurityGroupRule(rules_security_groups_id string) *neutronModules.DeleteSecurityGroupRuleResp {
	DeleteSecurityGroupRuleResp := &neutronModules.DeleteSecurityGroupRuleResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v2.0/security-group-rules/" + rules_security_groups_id

	client.RequestParam.Method = "DELETE"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		DeleteSecurityGroupRuleResp.ResponseCode = modules.HttpOK
	} else {
		DeleteSecurityGroupRuleResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		DeleteSecurityGroupRuleResp.ErrorInfo = Error
	}

	return DeleteSecurityGroupRuleResp
}
