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
	"time"
)

var expiresDate string = ""

// The client
type Client struct {
	TenantID     string                // Tenant ID
	RequestParam *modules.RequestParam // Request parameters
}

// User name
type DomainStruct struct {
	Name string `json:"name"`
}

// User stucture
type UserStruct struct {
	Domain   DomainStruct `json:"domain"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
}

// Password structure
type PasswordStruct struct {
	User UserStruct `json:"user"`
}

// Identity
type IdentityStruct struct {
	Methods  []string       `json:"methods"`
	Password PasswordStruct `json:"password"`
}

// Tenant id
type ProjectStruct struct {
	Id string `json:"id"`
}

// Scope structure
type ScopeStruct struct {
	Domain  DomainStruct  `json:"domain"`
	Project ProjectStruct `json:"project"`
}

// Authration information
type AuthStruct struct {
	Identity IdentityStruct `json:"identity"`
	Scope    ScopeStruct    `json:"scope"`
}

// Authration structure
type AuthStructInfo struct {
	AuthStruct AuthStruct `json:"auth"`
}

type Role struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Domin struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type User struct {
	Domin Domin  `json:"domin"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}

type Project struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Domin Domin  `json:"domin"`
}

type Token struct {
	Expires_at string  `json:"expires_at"`
	Issued_at  string  `json:"issued_at"`
	Project    Project `json:"project"`
	User       User    `json:"user"`
	Roles      []Role  `json:"roles"`
}

type TokenInfo struct {
	Token Token `json:"token"`
}

/*
 * @fn			func InitClient(AK, SK, TenantID, Endpoint string) (client *Client)
 * @brief		Initialize the client
 * @param[in] 	AK: The Access Key
 * @param[in]	SK: The Secret Key
 * @param[in]	TenantID: The tenant ID
 * @param[in]	Endpoint: Server address
 * @param[out]
 * @return		client
 */
func InitV4Client(AK, SK, TenantID string, clientConfiguration modules.ClientConfiguration) (client *Client) {
	RequestParam := &modules.RequestParam{}
	RequestParam.Endpoint = clientConfiguration.Endpoint
	RequestParam.Region = clientConfiguration.Region
	RequestParam.ServiceName = clientConfiguration.ServiceName
	RequestParam.AK = AK
	RequestParam.SK = SK
	RequestParam.Token = ""
	RequestParam.AuthType = modules.V4_AUTH
	client = &Client{TenantID, RequestParam}

	return client
}

/*
 * @fn			func InitClient(AK, SK, TenantID, Endpoint string) (client *Client)
 * @brief		Initialize the client
 * @param[in] 	Token: Token value
 * @param[in]	TenantID: The tenant ID
 * @param[in]	Endpoint: Server address
 * @param[out]
 * @return		client
 */
func InitTokenClient(Token, TenantID string, clientConfiguration modules.ClientConfiguration) (client *Client) {
	RequestParam := &modules.RequestParam{}
	RequestParam.Endpoint = clientConfiguration.Endpoint
	RequestParam.Region = clientConfiguration.Region
	RequestParam.ServiceName = clientConfiguration.ServiceName
	RequestParam.Token = Token
	RequestParam.AuthType = modules.Token_AUTH
	return &Client{TenantID, RequestParam}
}

func GetToken(username, password, tenant_id string, clientConfiguration modules.ClientConfiguration) string {
	if expiresDate != "" {
		t := time.Now()
		localDate := fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%d.%dZ", t.UTC().Year(), t.UTC().Month(), t.UTC().Day(), t.UTC().Hour(), t.UTC().Minute(), t.UTC().Second(), (t.UTC().Nanosecond())/1000)
		if localDate < expiresDate {
			return expiresDate
		}
	}

	RequestParam := &modules.RequestParam{}
	RequestParam.Url = clientConfiguration.Endpoint + "/v3/auth/tokens"
	RequestParam.Region = clientConfiguration.Region
	RequestParam.ServiceName = clientConfiguration.ServiceName
	RequestParam.Method = "POST"
	RequestParam.AuthType = modules.Token_AUTH
	RequestParam.RequestContentType = modules.ApplicationJson
	RequestParam.Token = ""

	jsonStr := `{"auth": {"identity": {"methods": ["password"], "password": {"user": {"domain": {"name": ""}, "name": "", "password": ""}}}, "scope": {"domain": {"name": ""}, "project": {"id": ""}}}}`
	var data AuthStructInfo
	json.Unmarshal([]byte(jsonStr), &data)
	data.AuthStruct.Identity.Password.User.Domain.Name = username
	data.AuthStruct.Identity.Password.User.Name = username
	data.AuthStruct.Identity.Password.User.Password = password
	data.AuthStruct.Scope.Domain.Name = username
	data.AuthStruct.Scope.Project.Id = tenant_id
	body, _ := json.Marshal(data)
	RequestParam.BodyContent = string(body)

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(RequestParam)
	strToken := Result.HeaderCollection["X-Subject-Token"]
	if modules.IsHttpOk(Result.ResponseCode) {
		var token TokenInfo
		json.Unmarshal([]byte(Result.RespMessage), &token)
		expiresDate = token.Token.Expires_at
		return strToken
	} else {
		return ""
	}
}
