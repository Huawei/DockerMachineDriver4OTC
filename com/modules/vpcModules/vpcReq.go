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
package vpcModules

import (
	"encoding/json"
)

// The request of creating a VPC
type CreateVpcReq struct {
	bodyContent string
}

/*
 * @fn			func (createVpcReq *CreateVpcReq) Init(name, cidr string)
 * @brief		Initialize
 * @param[in]	name: Specifies the name of the VPC.
 * @param[in]	cidr: Specifies the range of available subnets in the VPC.
 * @param[out]
 * @return
 */
func (createVpcReq *CreateVpcReq) Init(name, cidr string) {
	mapBody := make(map[string]interface{})
	mapBody["name"] = name
	mapBody["cidr"] = cidr
	mapResult := make(map[string]interface{})
	mapResult["vpc"] = mapBody
	bodyResult, _ := json.Marshal(mapResult)
	createVpcReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createVpcReq *CreateVpcReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (createVpcReq *CreateVpcReq) GetBodyContent() string {
	return createVpcReq.bodyContent
}

// The request of creating a security group
type CreateSecurityGroupReq struct {
	bodyContent string
}

/*
 * @fn			func (createSecurityGroupReq *CreateSecurityGroupReq) Init(name, vpc_id string)
 * @brief		Initialize
 * @param[in]	name: Specifies the name of the security group.
 * @param[in]	vpc_id: Specifies the resource ID of the VPC to which the security group belongs.
 * @param[out]
 * @return
 */
func (createSecurityGroupReq *CreateSecurityGroupReq) Init(name, vpc_id string) {
	mapBody := make(map[string]interface{})
	mapBody["name"] = name
	mapBody["vpc_id"] = vpc_id
	mapResult := make(map[string]interface{})
	mapResult["security_group"] = mapBody
	bodyResult, _ := json.Marshal(mapResult)
	createSecurityGroupReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createSecurityGroupReq *CreateSecurityGroupReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (createSecurityGroupReq *CreateSecurityGroupReq) GetBodyContent() string {
	return createSecurityGroupReq.bodyContent
}

// The request of creating a subnet
type CreateSubnetReq struct {
	bodyContent    string
	mapBodyContent map[string]map[string]interface{}
}

/*
 * @fn			func (createSubnetReq *CreateSubnetReq) Init(name, cidr, gateway_ip, availability_zone, vpc_id string)
 * @brief		Initialize
 * @param[in]	name: Specifies the name of the subnet.
 * @param[in]	cidr: Specifies the network segment on which the subnet resides.
 * @param[in]	gateway_ip: Specifies the gateway of the subnet.
 * @param[in]	availability_zone: Specifies the ID of the availability zone (AZ) to which the subnet belongs.
 * @param[out]
 * @return
 */
func (createSubnetReq *CreateSubnetReq) Init(name, cidr, gateway_ip, availability_zone, vpc_id string) {
	mapBody := make(map[string]interface{})
	mapBody["name"] = name
	mapBody["cidr"] = cidr
	mapBody["gateway_ip"] = gateway_ip
	mapBody["availability_zone"] = availability_zone
	mapBody["vpc_id"] = vpc_id
	mapResult := make(map[string]interface{})
	mapResult["subnet"] = mapBody
	bodyResult, _ := json.Marshal(mapResult)
	createSubnetReq.bodyContent = string(bodyResult)
	json.Unmarshal([]byte(bodyResult), &createSubnetReq.mapBodyContent)
}

/*
 * @fn			func (createSubnetReq *CreateSubnetReq) SetDhcpEnable(dhcp_enable bool)
 * @brief		Specifies whether the DHCP function is enabled for the subnet.
 * @param[in]	dhcp_enable: Specifies whether the DHCP function is enabled for the subnet.
							 The value can be true or false.
							 If this parameter is left blank, it is set to true by default.
 * @param[out]
 * @return
*/
func (createSubnetReq *CreateSubnetReq) SetDhcpEnable(dhcp_enable bool) {
	json.Unmarshal([]byte(createSubnetReq.bodyContent), &createSubnetReq.mapBodyContent)
	createSubnetReq.mapBodyContent["subnet"]["dhcp_enable"] = dhcp_enable
	bodyResult, _ := json.Marshal(createSubnetReq.mapBodyContent)
	createSubnetReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createSubnetReq *CreateSubnetReq) SetPrimaryDNS(primary_dns string)
 * @brief		Specifies the primary IP address of the DNS server on the subnet.
 * @param[in]	primary_dns: Specifies the primary IP address of the DNS server on the subnet.
							 The value must be a valid IP address.
 * @param[out]
 * @return
*/
func (createSubnetReq *CreateSubnetReq) SetPrimaryDNS(primary_dns string) {
	json.Unmarshal([]byte(createSubnetReq.bodyContent), &createSubnetReq.mapBodyContent)
	createSubnetReq.mapBodyContent["subnet"]["primary_dns"] = primary_dns
	bodyResult, _ := json.Marshal(createSubnetReq.mapBodyContent)
	createSubnetReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createSubnetReq *CreateSubnetReq) SetSecondaryDNS(secondary_dns string)
 * @brief		Specifies the secondary IP address of the DNS server on the subnet.
 * @param[in]	secondary_dns: Specifies the secondary IP address of the DNS server on the subnet.
							   The value must be a valid IP address.
 * @param[out]
 * @return
*/
func (createSubnetReq *CreateSubnetReq) SetSecondaryDNS(secondary_dns string) {
	json.Unmarshal([]byte(createSubnetReq.bodyContent), &createSubnetReq.mapBodyContent)
	createSubnetReq.mapBodyContent["subnet"]["secondary_dns"] = secondary_dns
	bodyResult, _ := json.Marshal(createSubnetReq.mapBodyContent)
	createSubnetReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createSubnetReq *CreateSubnetReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (createSubnetReq *CreateSubnetReq) GetBodyContent() string {
	bodyResult, _ := json.Marshal(createSubnetReq.mapBodyContent)
	createSubnetReq.bodyContent = string(bodyResult)
	return createSubnetReq.bodyContent
}

// The request of updating bandwidth information
type UpdateBandwidthReq struct {
	// name string `json:"name"`
	name string /*
		Specifies the name of the bandwidth.
		At least one in parameter name or parameter size must be set.
		The value is a string of 1 to 64 characters that contain digits,
		letters, underscores (_), and hyphens (-). If the value is null
		or is left empty, the name of the bandwidth is not changed.
	*/
	// size int `json:"size"` /*
	size int /*
		Specifies the bandwidth capacity.
		At least one in the parameter size or parameter name must be set.
		The value ranges from 1 to 300 Mbit/s. If the parameter is not included
		or the value is 0, the size of the bandwidth is not changed.
	*/
}

/*
 * @fn			func (updateBandwidthReq *UpdateBandwidthReq) SetName(name string)
 * @brief		Specifies the name of the bandwidth.
 * @param[in]	name: The name of the bandwidth.
 * @param[out]
 * @return
 */
func (updateBandwidthReq *UpdateBandwidthReq) SetName(name string) {
	updateBandwidthReq.name = name
}

/*
 * @fn			func (updateBandwidthReq *UpdateBandwidthReq) SetSize(size int)
 * @brief		Specifies the bandwidth capacity.
 * @param[in]	size: The bandwidth capacity.
 * @param[out]
 * @return
 */
func (updateBandwidthReq *UpdateBandwidthReq) SetSize(size int) {
	updateBandwidthReq.size = size
}

/*
 * @fn			func (updateBandwidthReq *UpdateBandwidthReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (updateBandwidthReq *UpdateBandwidthReq) GetBodyContent() string {
	mapBody := make(map[string]interface{})
	if len(updateBandwidthReq.name) > 0 {
		mapBody["name"] = updateBandwidthReq.name
	}

	if updateBandwidthReq.size < 0 {
		updateBandwidthReq.size = 0
	}
	mapBody["size"] = updateBandwidthReq.size
	mapResult := make(map[string]interface{})
	mapResult["bandwidth"] = mapBody
	bodyResult, _ := json.Marshal(mapResult)
	return string(bodyResult)
}

// Specifies the elastic IP address objects.
type PublicipCreate struct {
	strType string
}

/*
 * @fn			func (publicipCreate *PublicipCreate) SetType(Type string)
 * @brief		Specifies the elastic IP address type.
 * @param[in]	Type: Specifies the elastic IP address type.
					  The value must be a type supported by the system.
					  The value is 5_bgp.
 * @param[out]
 * @return
*/
func (publicipCreate *PublicipCreate) SetType(Type string) {
	publicipCreate.strType = Type
}

// Specifies the bandwidth objects.
type BandwidthCreate struct {
	name        string
	size        int
	share_type  string
	charge_mode string
}

/*
 * @fn			func (publicipCreate *PublicipCreate) SetType(Type string)
 * @brief		Initialize
 * @param[in]	name: Specifies the name of the bandwidth.
 * @param[in]	size: Specifies the bandwidth capacity.
 * @param[in]	share_type: The value is PER, indicating that the bandwidth is exclusive.
 * @param[out]
 * @return
 */
func (bandwidthCreate *BandwidthCreate) Init(name string, size int, share_type string) {
	bandwidthCreate.name = name
	bandwidthCreate.size = size
	bandwidthCreate.share_type = share_type
}

/*
 * @fn			func (bandwidthCreate *BandwidthCreate) SetChargeMode(charge_mode string)
 * @brief		Set the charging mode
 * @param[in]	charge_mode: The value is traffic, indicating the charging is based on traffic.
 * @param[out]
 * @return
 */
func (bandwidthCreate *BandwidthCreate) SetChargeMode(charge_mode string) {
	bandwidthCreate.charge_mode = charge_mode
}

// The request of applying for an elastic ip address
type CreatePublicIpReq struct {
	bodyContent string
}

/*
 * @fn			func (createPublicIpReq *CreatePublicIpReq) Init(publicipCreate *PublicipCreate, bandwidthCreate *BandwidthCreate)
 * @brief		Initialize
 * @param[in]	publicipCreate: Specifies the elastic IP address objects.
 * @param[in]	bandwidthCreate: Specifies the bandwidth objects.
 * @param[out]
 * @return
 */
func (createPublicIpReq *CreatePublicIpReq) Init(publicipCreate *PublicipCreate, bandwidthCreate *BandwidthCreate) {
	mapPublicip := make(map[string]interface{})
	mapPublicip["type"] = publicipCreate.strType
	mapBandwidth := make(map[string]interface{})
	mapBandwidth["name"] = bandwidthCreate.name
	mapBandwidth["size"] = bandwidthCreate.size
	mapBandwidth["share_type"] = bandwidthCreate.share_type
	if len(bandwidthCreate.charge_mode) > 0 {
		mapBandwidth["charge_mode"] = bandwidthCreate.charge_mode
	}
	mapResult := make(map[string]interface{})
	mapResult["publicip"] = mapPublicip
	mapResult["bandwidth"] = mapBandwidth
	bodyResult, _ := json.Marshal(mapResult)
	createPublicIpReq.bodyContent = string(bodyResult)
}

/*
 * @fn			func (createPublicIpReq *CreatePublicIpReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (createPublicIpReq *CreatePublicIpReq) GetBodyContent() string {
	return createPublicIpReq.bodyContent
}

// Specifies the elastic IP address objects.
type PublicipUpdate struct {
	bodyContent string
}

/*
 * @fn			func (publicipUpdate *PublicipUpdate) SetPortId(port_id string)
 * @brief		Specifies the ID of the VM NIC.
 * @param[in]	port_id: Specifies the ID of the VM NIC.
						 The parameter must be an existing NIC ID.
						If this parameter is not specified, this command unbinds
						the elastic IP address from the NIC.
 * @param[out]
 * @return
*/
func (publicipUpdate *PublicipUpdate) SetPortId(port_id string) {
	mapBody := make(map[string]interface{})
	mapBody["port_id"] = port_id
	mapResult := make(map[string]interface{})
	mapResult["publicip"] = mapBody
	body, _ := json.Marshal(mapResult)
	publicipUpdate.bodyContent = string(body)
}

// The request of updating elastic ip address information
type UpdatePublicIpReq struct {
	bodyContent string
}

/*
 * @fn			func (updatePublicIpReq *UpdatePublicIpReq) SetPublicip(publicipUpdate *PublicipUpdate)
 * @brief		Specifies the elastic IP address objects.
 * @param[in]	publicipUpdate: Specifies the elastic IP address objects.
 * @param[out]
 * @return
 */
func (updatePublicIpReq *UpdatePublicIpReq) SetPublicip(publicipUpdate *PublicipUpdate) {
	updatePublicIpReq.bodyContent = publicipUpdate.bodyContent
}

/*
 * @fn			func (updatePublicIpReq *UpdatePublicIpReq) GetBodyContent() string
 * @brief		Get the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (updatePublicIpReq *UpdatePublicIpReq) GetBodyContent() string {
	return updatePublicIpReq.bodyContent
}
