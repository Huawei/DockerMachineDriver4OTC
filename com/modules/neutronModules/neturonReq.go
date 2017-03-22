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
package neutronModules

import (
	"encoding/json"
	"fmt"
)

// Body parameters
type CreateSecurityGroupRuleReq struct {
	bodyContent    string
	mapBodyContent map[string]map[string]interface{}
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) Init(security_group_id, direction, ethertype string)
 * @brief		Initialize
 * @param[in]	security_group_id: Specifies the ID of the security group.
 * @param[in]	direction: Specifies the direction of access control.
 * @param[in]	ethertype: Specifies the protocol used by IP addresses.
 * @param[out]
 * @return
 */
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) Init(security_group_id, direction, ethertype string) {
	mapBody := make(map[string]interface{})
	mapBody["security_group_id"] = security_group_id
	mapBody["direction"] = direction
	mapBody["ethertype"] = ethertype
	mapResult := make(map[string]interface{})
	mapResult["security_group_rule"] = mapBody
	bodyResult, _ := json.Marshal(mapResult)
	createSecurityGroupRuleReq.bodyContent = string(bodyResult)
	json.Unmarshal([]byte(bodyResult), &createSecurityGroupRuleReq.mapBodyContent)
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetProtocol(protocol string)
 * @brief		Specifies the protocol type.
 * @param[in]	protocol: Specifies the protocol type.
						  If the parameter is left empty, the security group supports all types of protocols.
						  The value can be icmp, tcp, udp, or others.
 * @param[out]
 * @return
*/
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetProtocol(protocol string) {
	json.Unmarshal([]byte(createSecurityGroupRuleReq.bodyContent), &createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.mapBodyContent["security_group_rule"]["protocol"] = protocol
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetPortRangeMin(port_range_min int)
 * @brief		Specifies the start port.
 * @param[in]	port_range_min: Specifies the start port.
								The value ranges from -1 to 65,535.
								The value must be less than or equal to the value of port_range_max.
								An empty value indicates all ports. If protocol is icmp,
								the value range is determined by the A.2 ICMP-Port Range Relationship Table.
 * @param[out]
 * @return
*/
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetPortRangeMin(port_range_min int) {
	json.Unmarshal([]byte(createSecurityGroupRuleReq.bodyContent), &createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.mapBodyContent["security_group_rule"]["port_range_min"] = fmt.Sprintf("%d", port_range_min)
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetPortRangeMax(port_range_max int)
 * @brief		Specifies the end port.
 * @param[in]	port_range_max: Specifies the end port.
								The value ranges from -1 to 65,535.
								The value must be greater than or equal to the value of port_range_min.
								An empty value indicates all ports. If protocol is icmp,
								the value range is determined by the A.2 ICMP-Port Range Relationship Table.
 * @param[out]
 * @return
*/
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetPortRangeMax(port_range_max int) {
	json.Unmarshal([]byte(createSecurityGroupRuleReq.bodyContent), &createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.mapBodyContent["security_group_rule"]["port_range_max"] = fmt.Sprintf("%d", port_range_max)
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetRemoteIPPrefix(remote_ip_prefix string)
 * @brief		Specifies the remote IP address.
 * @param[in]	remote_ip_prefix: Specifies the remote IP address.
								  If the access control direction is set to egress,
								  the parameter specifies the source IP address.
								  If the access control direction is set to ingress,
								  the parameter specifies the destination IP address.
								  The parameter is exclusive with parameter remote_group_id.
								  The value can be in the CIDR format or IP addresses.
 * @param[out]
 * @return
*/
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetRemoteIPPrefix(remote_ip_prefix string) {
	json.Unmarshal([]byte(createSecurityGroupRuleReq.bodyContent), &createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.mapBodyContent["security_group_rule"]["remote_ip_prefix"] = remote_ip_prefix
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetRemoteGroupID(remote_group_id string)
 * @brief		Specifies the ID of the peer security group.
 * @param[in]	remote_group_id: Specifies the ID of the peer security group.
								 The value is exclusive with parameter remote_ip_prefix.
 * @param[out]
 * @return
*/
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) SetRemoteGroupID(remote_group_id string) {
	json.Unmarshal([]byte(createSecurityGroupRuleReq.bodyContent), &createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.mapBodyContent["security_group_rule"]["remote_group_id"] = remote_group_id
}

/*
 * @fn			func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) GetBodyContent() string
 * @brief		Return the string of json format
 * @param[in]
 * @param[out]
 * @return		string
 */
func (createSecurityGroupRuleReq *CreateSecurityGroupRuleReq) GetBodyContent() string {
	bodyResult, _ := json.Marshal(createSecurityGroupRuleReq.mapBodyContent)
	createSecurityGroupRuleReq.bodyContent = string(bodyResult)
	return createSecurityGroupRuleReq.bodyContent
}
