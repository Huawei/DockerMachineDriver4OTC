package vpc

import (
	"encoding/json"

	"github.com/docker/machine/libmachine/log"
)

type SubnetList struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Cidr        string `json:"cidr"`
	Gateway_Ip  string `json:"gateway_ip"`
	Dhcp_En     bool   `json:"dhcp_enable"`
	Primary_Dns string `json:"primary_dns"`
	Second_Dns  string `json:"secondary_dns"`
	Avail_Zone  string `json:"availability_zone"`
	Vpc_Id      string `json:"vpc_id"`
	Status      string `json:"status"`
}

type SubnetResponse struct {
	Subnets []SubnetList `json:"subnets"`
}

func (client *Client) ListSubnet(region, projectId, vpcId string) (snList *SubnetResponse, err error) {
	var subnetList SubnetResponse

	//compose uri string
	uri := "/" + projectId + "/subnets?vpc_id=" + vpcId

	respbytes, err := client.Do(region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal response body data into subnet list
	if err := json.Unmarshal(respbytes, &subnetList); err != nil {
		return nil, err
	}

	log.Debugf("subnet list are: %v", subnetList)
	return &subnetList, nil
}
