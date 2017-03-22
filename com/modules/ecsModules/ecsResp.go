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

/*
	The response of createCloudServer
*/
type CreateCloudServerResp struct {
	ResponseCode int
	modules.ErrorInfo
	Job_id string `json:"job_id"`
}

/*
	The response of deleteCloudServer
*/
type DeleteCloudServerResp struct {
	ResponseCode int
	modules.ErrorInfo
	Job_id string `json:"job_id"`
}

/*
	The response of showEcsJob
*/
type ShowEcsJobResp struct {
	ResponseCode int
	modules.ErrorInfo
	Job_id      string   `json:"job_id"`
	Status      string   `json:"status"`
	Entities    Entities `json:"entities"`
	Job_type    string   `json:"job_type"`
	Begin_time  string   `json:"begin_time"`
	End_time    string   `json:"end_time"`
	Error_code  string   `json:"error_code"`
	Fail_reason string   `json:"fail_reason"`
}

// Entities of showEcsJobResp
type Entities struct {
	Sub_jobs_total int      `json:"sub_jobs_total"`
	Server_id      string   `json:"server_id"`
	SubJobs        []SubJob `json:"sub_jobs"`
}

// SubJob of Entities
type SubJob struct {
	Status      string      `json:"status"`
	Entities    JobEntities `json:"entities"`
	Job_id      string      `json:"job_id"`
	Job_type    string      `json:"job_type"`
	Begin_time  string      `json:"begin_time"`
	End_time    string      `json:"end_time"`
	Error_code  string      `json:"error_code"`
	Fail_reason string      `json:"fail_reason"`
}

// JobEntities of SubJob
type JobEntities struct {
	Server_id string `json:"server_id"`
	Nic_id    string `json:"nic_id"`
}

/*
	The response of listCloudServerFlavorsExt
*/
type ListCloudServerFlavorsExtResp struct {
	ResponseCode int
	modules.ErrorInfo
	Flavors []FlavorListDetail `json:"flavors"`
}

// FlavorListDetail of ListCloudServerFlavorsExtResp
type FlavorListDetail struct {
	Id             string       `json:"id"`
	Name           string       `json:"name"`
	Vcpus          string       `json:"vcpus"`
	Ram            int          `json:"ram"`
	Disk           string       `json:"disk"`
	Swap           string       `json:"swap"`
	Ephemeral      int          `json:"OS-FLV-EXT-DATA:ephemeral"`
	Disabled       bool         `json:"OS-FLV-DISABLED:disabled"`
	Rxtx_factor    float32      `json:"rxtx_factor"`
	Rxtx_quota     string       `json:"rxtx_quota"`
	Rxtx_cap       string       `json:"rxtx_cap"`
	Is_public      bool         `json:"os-flavor-access:is_public"`
	Links          []Links      `json:"links"`
	Os_extra_specs OsExtraSpecs `json:"os_extra_specs"`
}

// Links of FlavorListDetail
type Links struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}

// OsExtraSpecs of FlavorListDetail
type OsExtraSpecs struct {
	Performancetype string `json:"ecs:performancetype"`
	availablezone   string
	resource_type   string
	// availablezone   string `json:"ecs:availablezone"`
	// resource_type   string `json:"resource_type"`
}
