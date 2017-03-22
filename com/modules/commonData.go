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

import (
	"encoding/json"
)

// The base data struct
type BaseDataStruct struct {
	MapBodyContent map[string]interface{}
}

func (baseDataStruct *BaseDataStruct) InitBase() {
	baseDataStruct.MapBodyContent = make(map[string]interface{})
}

func (baseDataStruct *BaseDataStruct) GetBodyContent() string {
	body, _ := json.Marshal(baseDataStruct.MapBodyContent)
	return string(body)
}

// Used to parse the Vpc response
type VpcInfo struct {
	Vpc Vpc `json:"vpc"`
}

// Vpc response fields
type Vpc struct {
	Id   string `json:"id"`   // Specifies a resource ID in UUID format.
	Name string `json:"name"` /*
		Specifies the name of the VPC.
		The name must be unique for a tenant.
		The value is a string of 1 to 64 characters that contain letters, digits,
		underscores (_), and hyphens (-).
	*/
	Cidr string `json:"cidr"` /*
		Specifies the range of available subnets in the VPC.
		The value must be in CIDR format, for example, 192.168.0.0/16.
		The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to
		172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.
	*/
	Status string `json:"status"` /*
		Specifies the status of the VPC.
		The value can be CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.
	*/
	NoSecurityGroup bool `json:"noSecurityGroup"`
}

// Used to parse the Subnet response
type SubnetInfo struct {
	Subnet Subnet `json:"subnet"`
}

type SubnetsInfo struct {
	Subnets []Subnet `json:"subnets"`
}

// Subnet response fields
type Subnet struct {
	Id                string `json:"id"`                // Specifies a resource ID in UUID format.
	Name              string `json:"name"`              // Specifies the name of the subnet.
	Cidr              string `json:"cidr"`              // Specifies the network segment of the subnet.
	Gateway_ip        string `json:"gateway_ip"`        // Specifies the gateway address of the subnet.
	Dhcp_enable       bool   `json:"dhcp_enable"`       // Specifies whether the DHCP function is enabled for the subnet.
	Primary_dns       string `json:"primary_dns"`       // Specifies the primary IP address of the DNS server on the subnet.
	Secondary_dns     string `json:"secondary_dns"`     // Specifies the secondary IP address of the DNS server on the subnet.
	Availability_zone string `json:"availability_zone"` // Specifies the ID of the AZ to which the subnet belongs.
	Vpc_id            string `json:"vpc_id"`            // Specifies the ID of the VPC to which the subnet belongs.
	Status            string `json:"status"`            // Specifies the status of the subnet.The value can be ACTIVE, DOWN, BUILD, ERROR, or DELETE.
}

// Publicip fields
type PublicipInfo struct {
	Publicip_id string `json:"publicip_id"` /*
		Specifies the ID of the elastic IP address,
		which uniquely identifies the elastic IP address.*/
	Publicip_address string `json:"publicip_address"` // Specifies the elastic IP address.
	Publicip_type    string `json:"publicip_type"`    // Specifies the type of the elastic IP address.
}

// Bandwidth fields
type Bandwidth struct {
	Name           string         `json:"name"`           // Specifies the bandwidth service.
	Size           int            `json:"size"`           // Specifies the bandwidth capacity.
	Id             string         `json:"id"`             // Specifies the bandwidth ID, which uniquely identifies the bandwidth.
	Share_type     string         `json:"share_type"`     // The value is PER, indicating that the bandwidth is exclusive.
	Publicip_info  []PublicipInfo `json:"publicip_info"`  // Specifies the elastic IP address of the bandwidth.
	Tenant_id      string         `json:"tenant_id"`      // Specifies the tenant ID of the user.
	Bandwidth_type string         `json:"bandwidth_type"` // Specifies the type of the bandwidth.
	Charge_mode    string         `json:"charge_mode"`    /*
		Specifies whether the charging is based on traffic or bandwidth.
		The value can be bandwidth or traffic. If the value is an empty
		character string or no value is returned, the charging is based on bandwidth.
	*/
}

// Used to parse the Bandwidth response
type BandwidthInfo struct {
	Bandwidth Bandwidth `json:"bandwidth"`
}

// Used to parse the Bandwidths response
type BandwidthsInfo struct {
	Bandwidths []Bandwidth `json:"bandwidths"`
}

// SecurityGroupRule  response fields
type SecurityGroupRule struct {
	Id                string `json:"id"`                // Specifies the ID of the security group rule.
	Security_group_id string `json:"security_group_id"` // Specifies the ID of the security group.
	Direction         string `json:"direction"`         /*
		Specifies the direction of access control.
		The value can be egress or ingress.
	*/
	Ethertype string `json:"ethertype"` /* Specifies the protocol used by IP addresses.
	The value can be IPv4 or IPv6.
	*/
	Protocol string `json:"protocol"` /*
		Specifies the protocol type.
		If the parameter is left empty, the security group supports all types of protocols.
		The value can be icmp, tcp, udp, or others.
	*/
	Port_range_min int `json:"port_range_min"` /*
		Specifies the start port.
		The value ranges from -1 to 65,535.
		The value must be less than or equal to the value of port_range_max.
		An empty value indicates all ports. If protocol is icmp, the value range
		 is determined by the A.2 ICMP-Port Range Relationship Table.
	*/
	Port_range_max int `json:"port_range_max"` /*
		Specifies the end port.
		The value ranges from -1 to 65,535.
		The value must be greater than or equal to the value of port_range_min.
		An empty value indicates all ports. If protocol is icmp, the value range
		is determined by the A.2 ICMP-Port Range Relationship Table.
	*/
	Remote_ip_perfix string `json:"remote_ip_prefix"` /*
		Specifies the remote IP address. If the access control direction is set to egress,
		the parameter specifies the source IP address. If the access control direction is
		set to ingress, the parameter specifies the destination IP address.
		The parameter is exclusive with parameter remote_group_id.
		The value can be in the CIDR format or IP addresses.
	*/
	Remote_group_id string `json:"remote_group_id"` /*
		Specifies the ID of the peer security group.
		The value is exclusive with parameter remote_ip_prefix.
	*/
}

// SecurityGroup  response fields
type SecurityGroup struct {
	Name               string              `json:"name"`                 //	Specifies the name of the security group.
	Id                 string              `json:"id"`                   // Specifies the security group ID, which uniquely identifies the security group.
	Vpc_id             string              `json:"vpc_id"`               // Specifies the resource ID of the VPC to which the security group belongs.
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"` /*
		Specifies the default security group rule, which ensures that hosts
		in the security group to communicate with one another.
	*/
}

// Used to parse the SecurityGroup response
type SecurityGroupInfo struct {
	SecurityGroup SecurityGroup `json:"security_group"`
}

// Used to parse the SecurityGroups response
type SecurityGroupsInfo struct {
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

// SecurityGroupRuleCreateInfo  response fields
type SecurityGroupRuleCreateInfo struct {
	Id                string `json:"id"`                // Specifies the ID of the security group rule.
	Tenant_id         string `json:"tenant_id"`         // Specifies the ID of the tenant.
	Security_group_id string `json:"security_group_id"` // Specifies the ID of the security group.
	Direction         string `json:"direction"`         /*
		Specifies the direction of access control.
		The value can be egress or ingress.
	*/
	Ethertype string `json:"ethertype"` /*
		Specifies the protocol used by IP addresses.
		The value can be IPv4 or IPv6.
	*/
	Protocol string `json:"protocol"` /*
		Specifies the protocol type.
		If the parameter is left empty, the security group supports all types of protocols.
		The value can be icmp, tcp, udp, or others.
	*/
	Port_range_min int `json:"port_range_min"` /*
		Specifies the start port.
		The value ranges from -1 to 65,535.
		The value must be less than or equal to the value of port_range_max.
	*/
	Port_range_max int `json:"port_range_max"` /*
		Specifies the end port.
		The value ranges from -1 to 65,535.
		The value must be greater than or equal to the value of port_range_min.
	*/
	Remote_ip_prefix string `json:"remote_ip_prefix"` /*
		Specifies the remote IP address. If the access control direction is set
		to egress, the parameter specifies the source IP address. If the access
		control direction is set to ingress, the parameter specifies the destination IP address.
		The parameter is exclusive with parameter remote_group_id.
		The value can be in the CIDR format or IP addresses.
	*/
	Remote_group_id string `json:"remote_group_id"` /*
		Specifies the ID of the peer security group.
		The value is exclusive with parameter remote_ip_prefix.
	*/
}

// Used to parse the SecurityGroupRuleCreateInfo response
type NeutronSecurityGroupRuleInfo struct {
	SecurityGroupRuleCreateInfo SecurityGroupRuleCreateInfo `json:"security_group_rule"`
}

// PublicipCreateData  response fields
type PublicipCreateData struct {
	Id string `json:"id"` /*
		Specifies the ID of the elastic IP address,
		which uniquely identifies the elastic IP address.
	*/
	Status string `json:"status"` /*
		Specifies the status of the elastic IP address.
		The value can be FREEZED, BIND_ERROR, BINDING,
		PENDING_DELETE, PENDING_CREATE, NOTIFYING, NOTIFY_DELETE,
		PENDING_UPDATE, DOWN, ACTIVE, ELB, or ERROR.
	*/
	Type              string `json:"type"`              // Specifies the type of the elastic IP address.
	Public_ip_address string `json:"public_ip_address"` // Specifies the elastic IP address obtained.
	Tenant_id         string `json:"tenant_id"`         // Specifies the tenant ID of the operator.
	Create_time       string `json:"create_time"`       // Specifies the time applying for the elastic IP address.
	Bandwidth_size    int    `json:"bandwidth_size"`    // Specifies the bandwidth capacity.
}

// Used to parse the PublicipCreateData response
type PublicipCreateDataInfo struct {
	PublicipCreateData PublicipCreateData `json:"publicip"`
}

// PublicipUpdateData  response fields
type PublicipUpdateData struct {
	Id string `json:"id"` /*
		Specifies the ID of the elastic IP address,
		which uniquely identifies the elastic IP address.
	*/
	Status string `json:"status"` /*
		Specifies the status of the elastic IP address.
		The value can be FREEZED, BIND_ERROR, BINDING,
		PENDING_DELETE, PENDING_CREATE, NOTIFYING, NOTIFY_DELETE,
		PENDING_UPDATE, DOWN, ACTIVE, ELB, or ERROR.
	*/
	Type              string `json:"type"`              // Specifies the type of the elastic IP address.
	Public_ip_address string `json:"public_ip_address"` // Specifies the elastic IP address obtained.
	Port_id           string `json:"port_id"`           // Specifies the ID of the VM NIC.
	Tenant_id         string `json:"tenant_id"`         // Specifies the tenant ID of the operator.
	Create_time       string `json:"create_time"`       // Specifies the time applying for the elastic IP address.
	Bandwidth_size    int    `json:"bandwidth_size"`    // Specifies the bandwidth capacity.
}

// Used to parse the PublicipUpdateData response
type PublicipUpdateDataInfo struct {
	PublicipUpdateData PublicipUpdateData `json:"publicip"`
}

// Image response fields
type Image struct {
	File   string `json:"file"`   // Specifies the URL for uploading and downloading the image file.
	Owner  string `json:"owner"`  // Specifies the tenant to which the image belongs.
	Id     string `json:"id"`     // Specifies the image ID.
	Size   int64  `json:"size"`   // This parameter is unavailable currently.
	Self   string `json:"self"`   // Specifies the URL of the image.
	Schema string `json:"schema"` // Specifies the image schema.
	Status string `json:"status"` /*
		Specifies the image status. The value can be active, queued, saving, d
		eleted, or killed. An image can be used only when it is in the active state.
	*/
	Tags       []string `json:"tags"`       // Specifies the image tag.
	Visibility string   `json:"visibility"` /*
		Specifies whether the image can be seen by other
		tenants. The value can be private or public.
	*/
	Name     string `json:"name"`     // Specifies the image name.
	Checksum string `json:"checksum"` // This parameter is unavailable currently.
	Deleted  bool   `json:"deleted"`  /*
		Specifies whether the image has been deleted.
		The value can be true or false.
	*/
	Protected bool `json:"protected"` /*
		Specifies whether the image is protected.
		A protected image cannot be deleted. The value can be true or false.
	*/
	Container_format string `json:"container_format"` // Specifies the container type.
	Min_ram          int    `json:"min_ram"`          // Specifies the minimum memory size (MB) required for running the image.
	Update_at        string `json:"update_at"`        // Specifies the time when the image was updated.
	Os_bit           string `json:"__os_bit"`         /*
		Specifies the system property of the operating system.
		The value can be 32 or 64.
	*/
	Os_version   string `json:"__os_version"`   // Specifies the operating system version.
	Description  string `json:"__description"`  // Provides supplementary information about an image.
	Disk_format  string `json:"disk_format"`    // Specifies the image format. Only the zvhd format is supported.
	IsRegistered string `json:"__isregistered"` /*
		Specifies whether the image has been registered.
		The value can be true or false.
	*/
	Platform string `json:"__platform"` /*
		Specifies the image platform type.
		The value can be Windows, Ubuntu, RedHat, SUSE, CentOS, or Other.
	*/
	Os_type string `json:"__os_type"` /*
		Specifies the operating system type.
		The value can be Linux, Windows, or Other.
	*/
	Min_disk         int    `json:"min_disk"`         // Specifies the minimum disk space (GB) required for running the image.
	Virtual_env_type string `json:"virtual_env_type"` /*
		Specifies the type of the environment where the image is used.
		If the image is used in FusionCompute, the value is FusionCompute.
	*/
	Image_source_type string `json:"__image_source_type"` /*
		Specifies the backend storage used by images.
		The value can be glance, nfs, or uds. The default value is glance.
	*/
	Imagetype string `json:"__imagetype"` /*
		Specifies the image type. The value can be gold,
		which indicates a public image, or private, which indicates a private image.
	*/
	Create_at         string `json:"created_at"`          // Specifies the time when the image was created.
	Virtual_size      int    `json:"virtual_size"`        // This parameter is unavailable currently.
	Deleted_at        string `json:"deleted_at"`          // Specifies the time when the image was deleted.
	Originalimagename string `json:"__originalimagename"` /*
		Specifies the ID of the parent image.
		If the image is a public image or a image that is created using a file,
		this value is left empty.
	*/
	Backup_id string `json:"__backup_id"` /*
		Specifies the backup ID. To create the a image using a backup,
		set the value to the backup ID. Otherwise, this value is left empty.
	*/
	Productcode    string `json:"__productcode"`    // Specifies the ID of the market image product.
	Image_location string `json:"__image_location"` // Specifies the location where the image is stored.
	Image_size     string `json:"__image_size"`     // Specifies the size of the image file (byte).
	Data_origin    string `json:"__data_origin"`    /*
		Specifies the image resource.
		If the image is a public image, this parameter is left empty.
	*/
}

// Used to parse the Images response
type ImageInfo struct {
	Images []Image `json:"images"`
}
