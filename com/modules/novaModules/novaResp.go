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
package novaModules

import (
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
)

/*
	The response of listInterfaces
*/
type ListInterfacesResp struct {
	ResponseCode int
	modules.ErrorInfo
	InterfaceAttachments []InterfaceAttachmentListInfo `json:"interfaceAttachments"`
}

// InterfaceAttachmentListInfo of ListInterfacesResp
type InterfaceAttachmentListInfo struct {
	Port_state string     `json:"port_state"`
	Fixed_ips  []FixedIps `json:"fixed_ips"`
	Net_id     string     `json:"net_id"`
	Port_id    string     `json:"port_id"`
	Mac_addr   string     `json:"mac_addr"`
}

// FixedIps of InterfaceAttachmentListInfo
type FixedIps struct {
	Subnet_id  string `json:"subnet_id"`
	Ip_address string `json:"ip_address"`
}

/*
	The response of createKeypair
*/
type CreateKeypairResp struct {
	ResponseCode int
	modules.ErrorInfo
	Keypair KeypairCreateDetail `json:"keypair"`
}

// KeypairCreateDetail of CreateKeypairResp
type KeypairCreateDetail struct {
	Fingerprint string `json:"fingerprint"`
	Name        string `json:"name"`
	Public_key  string `json:"public_key"`
	Private_key string `json:"private_key"`
	User_id     string `json:"user_id"`
}

/*
	The response of deleteKeyPair
*/
type DeleteKeyPairResp struct {
	ResponseCode int
	modules.ErrorInfo
}

/*
	The response of startServer
*/
type StartServerResp struct {
	ResponseCode int
	modules.ErrorInfo
}

/*
	The response of stopServer
*/
type StopServerResp struct {
	ResponseCode int
	modules.ErrorInfo
}

/*
	The response of rebootServer
*/
type RebootServerResp struct {
	ResponseCode int
	modules.ErrorInfo
}

/*
	The response of deleteAnServer
*/
type DeleteAnServerResp struct {
	ResponseCode int
	modules.ErrorInfo
}

/*
	The response of showServer
*/
type ShowServerResp struct {
	ResponseCode int
	modules.ErrorInfo
	Server ServerShowDetail `json:"server"`
}

// ServerShowDetail of ShowServerResp
type ServerShowDetail struct {
	AccessIPv4          string                             `json:"accessIPv4"`
	AccessIPv6          string                             `json:"accessIPv6"`
	Id                  string                             `json:"id"`
	Name                string                             `json:"name"`
	Status              string                             `json:"status"`
	Created             string                             `json:"created"`
	Updated             string                             `json:"updated"`
	Flavor              Flavor                             `json:"flavor"`
	Image               string                             `json:"image"`
	Tenant_id           string                             `json:"tenant_id"`
	Key_name            string                             `json:"key_name"`
	User_id             string                             `json:"user_id"`
	SS_Metadata         Metadata                           `json:"metadata"`
	HostId              string                             `json:"hostId"`
	Addresses           map[string][]PrivateAddr           `json:"addresses"`
	Security_groups     []ServerSecurityGroup              `json:"security_groups"`
	Links               []Elinks                           `json:"links"`
	DiskConfig          string                             `json:"OS-DCF:diskConfig"`
	Availability_zone   string                             `json:"OS-EXT-AZ:availability_zone"`
	Service_state       string                             `json:"OS-EXT-SERVICE:service_state"`
	Host                string                             `json:"OS-EXT-SRV-ATTR:host"`
	Hypervisor_hostname string                             `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	Instance_name       string                             `json:"OS-EXT-SRV-ATTR:instance_name"`
	Power_state         int                                `json:"OS-EXT-STS:power_state"`
	Task_state          string                             `json:"OS-EXT-STS:task_state"`
	Vm_state            string                             `json:"OS-EXT-STS:vm_state"`
	Launched_at         string                             `json:"OS-SRV-USG:launched_at"`
	Terminated_at       string                             `json:"OS-SRV-USG:terminated_at"`
	Volumes_attached    []OsExtendedvolumesVolumesAttached `json:"os-extended-volumes:volumes_attached"`
	Tags                []string                           `json:"tags"`
	Config_drive        string                             `json:"config_drive"`
	EvsOpts             int                                `json:"evsOpts"`
	HyperThreadAffinity string                             `json:"hyperThreadAffinity"`
	NumaOpts            int                                `json:"numaOpts"`
	Progress            int                                `json:"progress"`
	VcpuAffinity        []int                              `json:"vcpuAffinity"`
}

// Flavor of ServerShowDetail
type Flavor struct {
	Id    string   `json:"id"`
	Links []Elinks `json:"links"`
}

// Elinks of Flavor and ServerShowDetail
type Elinks struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// Metadata of ServerShowDetail
type Metadata struct {
	Openstack_region_name string `json:"__openstack_region_name"`
	Charging_mode         string `json:"chargiong_mode"`
	Image_name            string `json:"image_name"`
	CloudServiceType      string `json:"metering.cloudServiceType"`
	Image_id              string `json:"metering.image_id"`
	Imagetype             string `json:"metering.imagetype"`
	Resourcespeccode      string `json:"metering.resourcespeccode"`
	Resourcetype          string `json:"metering.resourcetype"`
	Os_bit                string `json:"os_bit"`
	Os_type               string `json:"os_type"`
	Vpc_id                string `json:"vpc_id"`
	Op_svc_lockaction     string `json:"op_svc_lockaction"`
}

// PrivateAddr of ServerShowDetail
type PrivateAddr struct {
	Addr     string `json:"addr"`
	Version  int    `json:"version"`
	Mac_addr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	Type     string `json:"OS-EXT-IPS:type"`
}

// ServerSecurityGroup of ServerShowDetail
type ServerSecurityGroup struct {
	Name string `json:"name"`
}

// OsExtendedvolumesVolumesAttachedm of ServerShowDetail
type OsExtendedvolumesVolumesAttached struct {
	Id string `json:"id"`
}
