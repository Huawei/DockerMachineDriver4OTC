package vpc

import (
	"encoding/json"

	"github.com/docker/machine/libmachine/log"
)

type ListSecurityGroupsAttributeArgs struct {
	TenantId string
	Marker   string
	Limit    uint32
	VpcId    string
}

type sgrule struct {
	Id                string `json:"id"`
	Security_group_id string `json:"security_group_id"`
	Direction         string `json:"direction"`
	Ethertype         string `json:"ethertype"`
	Protocol          string `json:"protocol"`
	Port_range_min    uint32 `json:"port_range_min"`
	Port_range_max    uint32 `json:"port_range_max"`
	Remote_ip_prefix  string `json:"remote_ip_prefix"`
	Remote_group_id   string `json:"remote_group_id"`
}

type SecurityGroup struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	VpcId string   `json:"vpc_id"`
	Rules []sgrule `json:"security_group_rules"`
}

type SecurityGroupsResponse struct {
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

func (client *Client) ListSecurityGroups(region, projectId string, args *ListSecurityGroupsAttributeArgs) (*SecurityGroupsResponse, error) {
	var securityGroups SecurityGroupsResponse
	//compose uri string
	uri := "/" + projectId + "/security-groups?limit=10&vpc_id=" + args.VpcId

	respbytes, err := client.Do(region, "GET", uri, args)
	if err != nil {
		return nil, err
	}

	//unmarshal reponse body data into security group list
	if err := json.Unmarshal(respbytes, &securityGroups); err != nil {
		log.Debugf("json unmarshal error is: %v", err)
		return nil, err
	}

	log.Debugf("security group list are: %v", securityGroups)

	return &securityGroups, nil
}
