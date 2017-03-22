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
package ecsModules

import (
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
)

/* The request params need by createCloudServer cunction Start */

// Struct CreateCloudServerReq
type CreateCloudServerReq struct {
	modules.BaseDataStruct
}

func (createCloudServerReq *CreateCloudServerReq) Init(createCloudServer CreateCloudServer) {
	createCloudServerReq.InitBase()
	createCloudServerReq.MapBodyContent["server"] = createCloudServer.MapBodyContent
}

// Struct CreateCloudServer
type CreateCloudServer struct {
	modules.BaseDataStruct
}

/*func (createCloudServer *CreateCloudServer) Init(imageRef, flavorRef, name, vpcid string,
nics []Nics, root_volume RootVolume, adminpwd, key_name, availability_zone string, sgs []SecGrp) {*/
func (createCloudServer *CreateCloudServer) Init(imageRef, flavorRef, name, vpcid string,
	nics []Nics, root_volume RootVolume, availability_zone string) {
	createCloudServer.InitBase()
	createCloudServer.MapBodyContent["imageRef"] = imageRef
	createCloudServer.MapBodyContent["flavorRef"] = flavorRef
	createCloudServer.MapBodyContent["name"] = name
	createCloudServer.MapBodyContent["vpcid"] = vpcid

	var nicsList []interface{}
	for _, row := range nics {
		nicsList = append(nicsList, row.MapBodyContent)
	}
	createCloudServer.MapBodyContent["nics"] = nicsList

	createCloudServer.MapBodyContent["root_volume"] = root_volume.MapBodyContent
	createCloudServer.MapBodyContent["availability_zone"] = availability_zone

	/*var sgList []interface{}
	for _, row := range sgs {
		sgList = append(sgList, row.MapBodyContent)
	}
	createCloudServer.MapBodyContent["security_groups"] = sgList
	createCloudServer.MapBodyContent["adminPass"] = adminpwd
	createCloudServer.MapBodyContent["key_name"] = key_name*/
}

func (createCloudServer *CreateCloudServer) SetPersonality(personality []Personality) {
	var perList []interface{}
	for _, row := range personality {
		perList = append(perList, row.MapBodyContent)
	}
	createCloudServer.MapBodyContent["personality"] = perList
}

func (createCloudServer *CreateCloudServer) SetUser_data(user_data string) {
	createCloudServer.MapBodyContent["user_data"] = user_data
}

func (createCloudServer *CreateCloudServer) SetAdminPass(adminPass string) {
	createCloudServer.MapBodyContent["adminPass"] = adminPass
}

func (createCloudServer *CreateCloudServer) SetKey_name(key_name string) {
	createCloudServer.MapBodyContent["key_name"] = key_name
}

func (createCloudServer *CreateCloudServer) SetCount(count int) {
	createCloudServer.MapBodyContent["count"] = count
}

func (createCloudServer *CreateCloudServer) SetData_volumes(data_volumes []DataVolume) {
	var dataList []interface{}
	for _, row := range data_volumes {
		dataList = append(dataList, row.MapBodyContent)
	}
	createCloudServer.MapBodyContent["data_volumes"] = dataList
}

func (createCloudServer *CreateCloudServer) SetSecurity_groups(security_groups []SecurityGroup) {
	var securityList []interface{}
	for _, row := range security_groups {
		securityList = append(securityList, row.MapBodyContent)
	}
	createCloudServer.MapBodyContent["security_groups"] = securityList
}

func (createCloudServer *CreateCloudServer) SetExtendparam(extendparam Extendparam) {
	createCloudServer.MapBodyContent["extendparam"] = extendparam.MapBodyContent
}

// Struct Nics
type Nics struct {
	modules.BaseDataStruct
}

func (nics *Nics) Init(subnet_id string) {
	nics.InitBase()
	nics.MapBodyContent["subnet_id"] = subnet_id
}

func (nics *Nics) SetIpAddress(ip_address string) {
	nics.MapBodyContent["ip_address"] = ip_address
}

type SecGrp struct {
	modules.BaseDataStruct
}

func (secGrp *SecGrp) Init(secgrp_id string) {
	secGrp.InitBase()
	secGrp.MapBodyContent["id"] = secgrp_id
}

// Struct RootVolume
type RootVolume struct {
	modules.BaseDataStruct
}

func (root_volume *RootVolume) Init(volumetype string) {
	root_volume.InitBase()
	root_volume.MapBodyContent["volumetype"] = volumetype
}

func (root_volume *RootVolume) SetSize(size int) {
	root_volume.MapBodyContent["size"] = size
}

// Struct Personality
type Personality struct {
	modules.BaseDataStruct
}

func (personality *Personality) Init(path, contexts string) {
	personality.InitBase()
	personality.MapBodyContent["path"] = path
	personality.MapBodyContent["contents"] = contexts
}

// Struct Publicip
type Publicip struct {
	modules.BaseDataStruct
}

func (publicip *Publicip) Init() {
	publicip.InitBase()
}

func (publicip *Publicip) SetId(id string) {
	publicip.MapBodyContent["id"] = id
}

func (publicip *Publicip) SetEip(eip Eip) {
	publicip.MapBodyContent["eip"] = eip.MapBodyContent
}

// Struct Eip
type Eip struct {
	modules.BaseDataStruct
}

func (eip *Eip) Init(iptype string, bandwidth Bandwidth) {
	eip.InitBase()
	eip.MapBodyContent["iptype"] = iptype
	eip.MapBodyContent["bandwidth"] = bandwidth.MapBodyContent
}

// Struct Bandwidth
type Bandwidth struct {
	modules.BaseDataStruct
}

func (bandwidth *Bandwidth) Init(size int, share_type, chargemode string) {
	bandwidth.InitBase()
	bandwidth.MapBodyContent["size"] = size
	bandwidth.MapBodyContent["share_type"] = share_type
	bandwidth.MapBodyContent["chargemode"] = chargemode
}

// Struct DataVolume
type DataVolume struct {
	modules.BaseDataStruct
}

func (dataVolume *DataVolume) Init(volumetype string, size int) {
	dataVolume.InitBase()
	dataVolume.MapBodyContent["volumetype"] = volumetype
	dataVolume.MapBodyContent["size"] = size
}

// Struct SecurityGroup
type SecurityGroup struct {
	modules.BaseDataStruct
}

func (securityGroup *SecurityGroup) Init() {
	securityGroup.InitBase()
}

func (securityGroup *SecurityGroup) SetId(id string) {
	securityGroup.MapBodyContent["id"] = id
}

// Struct Extendparam
type Extendparam struct {
	modules.BaseDataStruct
}

func (extendparam *Extendparam) Init() {
	extendparam.InitBase()
}

func (extendparam *Extendparam) SetRegionID(regionID string) {
	extendparam.MapBodyContent["regionID"] = regionID
}

/* The request params need by createCloudServer End */

/* The request Params need by deleteCloudServer Start */

// Struct DeleteCloudServerReq
type DeleteCloudServerReq struct {
	modules.BaseDataStruct
}

func (deleteCloudServerReq *DeleteCloudServerReq) Init(servers []ServerId, delete_publicip, delete_volume bool) {
	deleteCloudServerReq.InitBase()
	var serversList []interface{}
	for _, row := range servers {
		serversList = append(serversList, row.MapBodyContent)
	}
	deleteCloudServerReq.MapBodyContent["servers"] = serversList

	deleteCloudServerReq.MapBodyContent["delete_publicip"] = delete_publicip
	deleteCloudServerReq.MapBodyContent["delete_volume"] = delete_volume
}

// Struct ServerId
type ServerId struct {
	modules.BaseDataStruct
}

func (serverId *ServerId) Init(id string) {
	serverId.InitBase()
	serverId.MapBodyContent["id"] = id
}

/* The request Params need by deleteCloudServer End */
