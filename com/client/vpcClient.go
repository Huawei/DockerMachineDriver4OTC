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
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/vpcModules"
)

/*
 * @fn			func (client *Client) DeleteVpc(vpc_id string) *vpcModules.DeleteVpcResp
 * @brief		This interface is used to delete a VPC.
 * @param[in] 	vpc_id: Specifies the VPC ID, which uniquely identifies the VPC.
 * @param[out]
 * @return		*vpcModules.DeleteVpcResp
 */
func (client *Client) DeleteVpc(vpc_id string) *vpcModules.DeleteVpcResp {
	DeleteVpcResp := &vpcModules.DeleteVpcResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/vpcs/" + vpc_id

	client.RequestParam.Method = "DELETE"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		DeleteVpcResp.ResponseCode = modules.HttpOK
	} else {
		DeleteVpcResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		DeleteVpcResp.ErrorInfo = Error
	}

	return DeleteVpcResp
}

/*
 * @fn			func (client *Client) CreateVpc(createVpcReq vpcModules.CreateVpcReq) *vpcModules.CreateVpcResp
 * @brief		This interface is used to create a VPC.
 * @param[in] 	createVpcReq: Body parameters
 * @param[out]
 * @return		*vpcModules.CreateVpcResp
 */
func (client *Client) CreateVpc(createVpcReq vpcModules.CreateVpcReq) *vpcModules.CreateVpcResp {
	CreateVpcResp := &vpcModules.CreateVpcResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/vpcs"
	client.RequestParam.Method = "POST"
	client.RequestParam.BodyContent = createVpcReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		CreateVpcResp.ResponseCode = modules.HttpOK
		var data modules.VpcInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		CreateVpcResp.Vpc = data.Vpc
	} else {
		CreateVpcResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		CreateVpcResp.ErrorInfo = Error
	}

	return CreateVpcResp
}

/*
 * @fn			func (client *Client) ShowVpc(vpc_id string) *vpcModules.ShowVpcResp
 * @brief		This interface is used to query details about a VPC.
 * @param[in] 	vpc_id: Specifies the VPC ID, which uniquely identifies the VPC.
 * @param[out]
 * @return		*vpcModules.ShowVpcResp
 */
func (client *Client) ShowVpc(vpc_id string) *vpcModules.ShowVpcResp {
	ShowVpcResp := &vpcModules.ShowVpcResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/vpcs/" + vpc_id

	client.RequestParam.Method = "GET"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ShowVpcResp.ResponseCode = modules.HttpOK
		var data modules.VpcInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ShowVpcResp.Vpc = data.Vpc
	} else {
		ShowVpcResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ShowVpcResp.ErrorInfo = Error
	}

	return ShowVpcResp
}

/*
 * @fn			func (client *Client) CreateSubnet(createSubnetReq *vpcModules.CreateSubnetReq)
				*vpcModules.CreateSubnetResp
 * @brief		This interface is used to create a subnet.
 * @param[in] 	createSubnetReq: URI parameters
 * @param[out]
 * @return		*vpcModules.CreateSubnetResp
*/
func (client *Client) CreateSubnet(createSubnetReq *vpcModules.CreateSubnetReq) *vpcModules.CreateSubnetResp {
	CreateSubnetResp := &vpcModules.CreateSubnetResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/subnets"

	client.RequestParam.Method = "POST"
	client.RequestParam.BodyContent = createSubnetReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		CreateSubnetResp.ResponseCode = modules.HttpOK
		var data modules.SubnetInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		CreateSubnetResp.Subnet = data.Subnet
	} else {
		CreateSubnetResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		CreateSubnetResp.ErrorInfo = Error
	}

	return CreateSubnetResp
}

/*
 * @fn			func (client *Client) ShowSubnet(subnet_id string) *vpcModules.ShowSubnetResp
 * @brief		This interface is used to query details about a subnet.
 * @param[in] 	subnet_id:Specifies the subnet ID, which uniquely identifies the subnet.
 * @param[out]
 * @return		*vpcModules.ShowSubnetResp
 */
func (client *Client) ShowSubnet(subnet_id string) *vpcModules.ShowSubnetResp {
	ShowSubnetResp := &vpcModules.ShowSubnetResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/subnets/" + subnet_id

	client.RequestParam.BodyContent = ""
	client.RequestParam.Method = "GET"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ShowSubnetResp.ResponseCode = modules.HttpOK
		var data modules.SubnetInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ShowSubnetResp.Subnet = data.Subnet
	} else {
		ShowSubnetResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ShowSubnetResp.ErrorInfo = Error
	}

	return ShowSubnetResp
}

func (client *Client) ListSubnets(limit int, marker, vpc_id string) *vpcModules.ListSubnetsResp {
	ListSubnetsResp := &vpcModules.ListSubnetsResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/subnets" + "?limit=" + fmt.Sprintf("%d", limit) + "&marker=" + marker + "&vpc_id=" + vpc_id

	client.RequestParam.BodyContent = ""
	client.RequestParam.Method = "GET"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ListSubnetsResp.ResponseCode = modules.HttpOK
		var data modules.SubnetsInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ListSubnetsResp.Subnets = data.Subnets
	} else {
		ListSubnetsResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ListSubnetsResp.ErrorInfo = Error
	}

	return ListSubnetsResp
}

/*
 * @fn			func (client *Client) DeleteSubnet(vpc_id, subnet_id string) *vpcModules.DeleteSubnetResp
 * @brief		This interface is used to delete a subnet.
 * @param[in] 	vpc_id: Specifies the ID of the subnet VPC.
 * @param[in]	subnet_id:Specifies the subnet ID, which uniquely identifies the subnet.
 * @param[out]
 * @return		*vpcModules.DeleteSubnetResp
 */
func (client *Client) DeleteSubnet(vpc_id, subnet_id string) *vpcModules.DeleteSubnetResp {
	DeleteSubnetResp := &vpcModules.DeleteSubnetResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/vpcs/" + vpc_id + "/subnets/" + subnet_id

	client.RequestParam.BodyContent = ""
	client.RequestParam.Method = "DELETE"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		DeleteSubnetResp.ResponseCode = modules.HttpOK
	} else {
		DeleteSubnetResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		DeleteSubnetResp.ErrorInfo = Error
	}

	return DeleteSubnetResp
}

/*
 * @fn			func (client *Client) UpdateBandwidth(bandwidth_id string,
				updateBandwidthReq *vpcModules.UpdateBandwidthReq) *vpcModules.UpdateBandwidthResp
 * @brief		This interface is used to update information about a bandwidth.
 * @param[in] 	bandwidth_id: Specifies the bandwidth ID, which uniquely identifies the bandwidth.
 * @param[in]	updateBandwidthReq: Body parameters
 * @param[out]
 * @return		*vpcModules.UpdateBandwidthResp
*/
func (client *Client) UpdateBandwidth(bandwidth_id string, updateBandwidthReq *vpcModules.UpdateBandwidthReq) *vpcModules.UpdateBandwidthResp {
	UpdateBandwidthResp := &vpcModules.UpdateBandwidthResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/bandwidths/" + bandwidth_id

	client.RequestParam.BodyContent = updateBandwidthReq.GetBodyContent()
	client.RequestParam.Method = "PUT"

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		UpdateBandwidthResp.ResponseCode = modules.HttpOK
		var data modules.BandwidthInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		UpdateBandwidthResp.Bandwidth = data.Bandwidth
	} else {
		UpdateBandwidthResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		UpdateBandwidthResp.ErrorInfo = Error
	}

	return UpdateBandwidthResp
}

/*
 * @fn			func (client *Client) ListBandwidths(limit int, marker string)
				*vpcModules.ListBandwidthsResp
 * @brief		This interface is used to query bandwidths using search criteria
				and to display the bandwidths in a list.
 * @param[in] 	limit:Specifies the number of records returned on each page.
					  The value ranges from 0 to intmax.
 * @param[in]	marker:Specifies the resource ID of the pagination query. If the
					   parameter is left empty, only resources on the first page are queried.
 * @param[out]
 * @return		*vpcModules.ListBandwidthsResp
*/
func (client *Client) ListBandwidths(limit int, marker string) *vpcModules.ListBandwidthsResp {
	ListBandwidthsResp := &vpcModules.ListBandwidthsResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/bandwidths" + "?limit=" + fmt.Sprintf("%d", limit) + "&marker=" + marker

	client.RequestParam.Method = "GET"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ListBandwidthsResp.ResponseCode = modules.HttpOK
		var data modules.BandwidthsInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ListBandwidthsResp.Bandwidths = data.Bandwidths
	} else {
		ListBandwidthsResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ListBandwidthsResp.ErrorInfo = Error
	}

	return ListBandwidthsResp
}

/*
 * @fn			func (client *Client) CreateSecurityGroup(createSecurityGroupReq
				*vpcModules.CreateSecurityGroupReq) *vpcModules.CreateSecurityGroupResp
 * @brief		This interface is used to create a security group.
 * @param[in] 	createSecurityGroupReq: Body parameters
 * @param[out]
 * @return		*vpcModules.CreateSecurityGroupResp
*/
func (client *Client) CreateSecurityGroup(createSecurityGroupReq *vpcModules.CreateSecurityGroupReq) *vpcModules.CreateSecurityGroupResp {
	CreateSecurityGroupResp := &vpcModules.CreateSecurityGroupResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/security-groups"

	client.RequestParam.Method = "POST"
	client.RequestParam.BodyContent = createSecurityGroupReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		CreateSecurityGroupResp.ResponseCode = modules.HttpOK
		var data modules.SecurityGroupInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		CreateSecurityGroupResp.SecurityGroup = data.SecurityGroup
	} else {
		CreateSecurityGroupResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		CreateSecurityGroupResp.ErrorInfo = Error
	}

	return CreateSecurityGroupResp
}

/*
 * @fn			func (client *Client) ListSecurityGroups(limit int, marker,
				vpc_id string) *vpcModules.ListSecurityGroupsResp
 * @brief		This interface is used to query security groups using search
				criteria and to display the security groups in a list.
 * @param[in] 	limit: Specifies the number of records returned on each page.
					   The value ranges from 0 to intmax.
 * @param[in]	marker: Specifies the resource ID of the pagination query.
						If the parameter is left empty, only resources on
						the first page are queried.
 * @param[in]	vpc_id: Specifies the VPC ID used as the query filter.
 * @param[out]
 * @return		*vpcModules.ListSecurityGroupsResp
*/
func (client *Client) ListSecurityGroups(limit int, marker, vpc_id string) *vpcModules.ListSecurityGroupsResp {
	ListSecurityGroupsResp := &vpcModules.ListSecurityGroupsResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/security-groups" + "?limit=" + fmt.Sprintf("%d", limit) + "&marker=" + marker + "&vpc_id=" + vpc_id

	client.RequestParam.Method = "GET"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		ListSecurityGroupsResp.ResponseCode = modules.HttpOK
		var data modules.SecurityGroupsInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		ListSecurityGroupsResp.SecurityGroups = data.SecurityGroups
	} else {
		ListSecurityGroupsResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		ListSecurityGroupsResp.ErrorInfo = Error
	}

	return ListSecurityGroupsResp
}

/*
 * @fn			func (client *Client) CreatePublicIp(createPublicipReq
				*vpcModules.CreatePublicIpReq) *vpcModules.CreatePublicIpResp
 * @brief		This interface is used to apply for an elastic IP address.
 * @param[in] 	createPublicipReq: Body parameters
 * @param[out]
 * @return		*vpcModules.CreatePublicIpResp
*/
func (client *Client) CreatePublicIp(createPublicipReq *vpcModules.CreatePublicIpReq) *vpcModules.CreatePublicIpResp {
	CreatePublicIpResp := &vpcModules.CreatePublicIpResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/publicips"

	client.RequestParam.Method = "POST"
	client.RequestParam.BodyContent = createPublicipReq.GetBodyContent()
	fmt.Println(client.RequestParam.BodyContent)
	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	fmt.Println(Result)
	if modules.IsHttpOk(Result.ResponseCode) {
		CreatePublicIpResp.ResponseCode = modules.HttpOK
		var data modules.PublicipCreateDataInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		CreatePublicIpResp.PublicipCreateData = data.PublicipCreateData
	} else {
		CreatePublicIpResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		CreatePublicIpResp.ErrorInfo = Error
	}

	return CreatePublicIpResp
}

/*
 * @fn			func (client *Client) UpdatePublicIp(publicip_id string,
				updatePublicIpReq *vpcModules.UpdatePublicIpReq) *vpcModules.UpdatePublicIpResp
 * @brief		This interface is used to update information about an elastic IP address,
				that is, binding an elastic IP address to a NIC or unbinding an
				elastic IP address from a NIC.
 * @param[in] 	publicip_id: Specifies the ID of the elastic IP address, which
							 uniquely identifies the elastic IP address.
 * @param[out]
 * @return		*vpcModules.UpdatePublicIpResp
*/
func (client *Client) UpdatePublicIp(publicip_id string, updatePublicIpReq *vpcModules.UpdatePublicIpReq) *vpcModules.UpdatePublicIpResp {
	UpdatePublicIpResp := &vpcModules.UpdatePublicIpResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/publicips/" + publicip_id

	client.RequestParam.Method = "PUT"
	client.RequestParam.BodyContent = updatePublicIpReq.GetBodyContent()

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		UpdatePublicIpResp.ResponseCode = modules.HttpOK
		var data modules.PublicipUpdateDataInfo
		json.Unmarshal([]byte(Result.RespMessage), &data)
		UpdatePublicIpResp.PublicipUpdateData = data.PublicipUpdateData
	} else {
		UpdatePublicIpResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		UpdatePublicIpResp.ErrorInfo = Error
	}

	return UpdatePublicIpResp
}

/*
 * @fn			func (client *Client) DeletePublicIp(publicip_id string) *vpcModules.DeletePublicIpResp
 * @brief		This interface is used to delete an elastic IP address.
 * @param[in] 	publicip_id: Specifies the ID of the elastic IP address, which
							 uniquely identifies the elastic IP address.
 * @param[out]
 * @return		*vpcModules.DeletePublicIpResp
*/
func (client *Client) DeletePublicIp(publicip_id string) *vpcModules.DeletePublicIpResp {
	DeletePublicIpResp := &vpcModules.DeletePublicIpResp{}
	client.RequestParam.Url = "https://vpc." +  client.RequestParam.Endpoint + "/v1/" + client.TenantID + "/publicips/" + publicip_id

	client.RequestParam.Method = "DELETE"
	client.RequestParam.BodyContent = ""

	var Result *modules.Result = nil
	Result = cpp_sdk_core.SendRequest(client.RequestParam)
	if modules.IsHttpOk(Result.ResponseCode) {
		DeletePublicIpResp.ResponseCode = modules.HttpOK
	} else {
		DeletePublicIpResp.ResponseCode = Result.ResponseCode
		var Error modules.ErrorInfo
		json.Unmarshal([]byte(Result.RespMessage), &Error)
		DeletePublicIpResp.ErrorInfo = Error
	}

	return DeletePublicIpResp
}
