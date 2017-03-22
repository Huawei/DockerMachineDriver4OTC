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
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/imsModules"
)

/*
 * @fn			func (client *Client) ListCloudImages(listCloudImagesReqEx
				*imsModules.ListCloudImagesReqEx) *imsModules.ListCloudImagesResp
 * @brief		This interface is used to query images using search criteria and to display the images in a list.
 * @param[in] 	listCloudImagesReqEx: URI parameters
 * @param[out]
 * @return		*imsModules.ListCloudImagesResp
*/
func (client *Client) ListCloudImages(listCloudImagesReqEx *imsModules.ListCloudImagesReqEx) *imsModules.ListCloudImagesResp {
	ListCloudImagesResp := &imsModules.ListCloudImagesResp{}
	client.RequestParam.Url = client.RequestParam.Endpoint + "/v2/cloudimages" + listCloudImagesReqEx.GetAbsolutePath()
	client.RequestParam.Method = "GET"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ListCloudImagesResp.ResponseCode = modules.HttpOK
		var data modules.ImageInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ListCloudImagesResp.Images = data.Images
	} else {
		ListCloudImagesResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ListCloudImagesResp.ErrorInfo = Error
	}

	return ListCloudImagesResp
}
