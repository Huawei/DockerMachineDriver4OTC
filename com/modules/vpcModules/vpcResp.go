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
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
)

// The response of creating a VPC
type CreateVpcResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.Vpc
}

// The response of deleting a vpc
type DeleteVpcResp struct {
	ResponseCode int
	modules.ErrorInfo
}

// The response of querying VPC details
type ShowVpcResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.Vpc
}

// The response of creating a subnet
type CreateSubnetResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.Subnet
}

// The response of querying subnet details
type ShowSubnetResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.Subnet
}

type ListSubnetsResp struct {
	ResponseCode int
	modules.ErrorInfo
	Subnets []modules.Subnet
}

// The response of deleting a subnet
type DeleteSubnetResp struct {
	ResponseCode int
	modules.ErrorInfo
}

// The response of updating bandwidth information
type UpdateBandwidthResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.Bandwidth
}

// The response of querying bandwidths
type ListBandwidthsResp struct {
	ResponseCode int
	modules.ErrorInfo
	Bandwidths []modules.Bandwidth
}

// The response of creating a security group
type CreateSecurityGroupResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.SecurityGroup
}

// The response of querying security groups
type ListSecurityGroupsResp struct {
	ResponseCode int
	modules.ErrorInfo
	SecurityGroups []modules.SecurityGroup
}

// The response of applying for an elastic IP address
type CreatePublicIpResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.PublicipCreateData
}

// The response of updating elastic IP address information
type UpdatePublicIpResp struct {
	ResponseCode int
	modules.ErrorInfo
	modules.PublicipUpdateData
}

// The response of deleting an elastic IP address
type DeletePublicIpResp struct {
	ResponseCode int
	modules.ErrorInfo
}
