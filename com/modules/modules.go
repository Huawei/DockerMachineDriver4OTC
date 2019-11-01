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
package modules

func IsHttpOk(RespCode int) (ok bool) {
	if (RespCode >= 200) && (RespCode <= 299) {
		ok = true
	} else {
		ok = false
	}
	return ok
}

// Error message
type ErrorInfo struct {
	Code        string `json:"code"`
	Description string `json:"message"`
}

// 执行结果
type Result struct {
	Err              error             // Error
	ResponseCode     int               // Response code
	RespMessage      string            // Response message
	HeaderCollection map[string]string // Response Header
}

// Configuration
type ClientConfiguration struct {
	Endpoint    string // Server address
	ServiceName string // Service name
	Region      string // Region
}

// Request paramters
type RequestParam struct {
	AK                 string // Access key
	SK                 string // Security key
	Token              string // Token
	Endpoint           string // Server address
	ServiceName        string // Service name
	Region             string // Region
	BodyContent        string // The string of body
	Method             string // Method
	RequestContentType string // Request content type
	AuthType           int    // Authorization mode
	Url                string // Http url
	Subproject_ID      string // Subproject ID
}

/*
 * @fn			func InitRequestParam(AK, SK, Uri, ServiceName, Region, BodyContent,
				Method, RequestContentType string) (requestParam *RequestParam)
 * @brief		Initialize
 * @param[in]	AK: Access key
 * @param[in]	SK: Security key
 * @param[in]	Token: Token value
 * @param[in]	Endpoint: Service name
 * @param[in]	ServiceName: Service name
 * @param[in]	Region: Region
 * @param[in]	BodyContent: The string of body
 * @param[in]	Method: Method
 * @param[in]	RequestContentType: Request content type
 * @param[in]	AuthType: Authrization mode
 * @param[out]
 * @return		requestParam
*/

/*
func InitRequestParam(Endpoint, ServiceName, Region, BodyContent, Method, RequestContentType string, AuthType int) (requestParam *RequestParam) {
	requestParam = &RequestParam{"", "", "", Endpoint, ServiceName, Region, BodyContent, Method, RequestContentType, AuthType}
	return requestParam
}

*/
