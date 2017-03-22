package vpc

import (
	"encoding/json"

	"github.com/docker/machine/libmachine/log"
)

type PubIP struct {
	Type string `json:"type"`
}

type BWDesc struct {
	Name    string `json:"name"`
	Size    uint32 `json:"size"`
	ShrType string `json:"share_type"`
	ChgMode string `json:"charge_mode"`
}

type EIPAllocArg struct {
	EIPDesc   PubIP  `json:"publicip"`
	BandWidth BWDesc `json:"bandwidth"`
}

type EIP struct {
	Id         string `json:"id"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Addr       string `json:"public_ip_address"`
	TenantId   string `json:"tenant_id"`
	CreateTime string `json:"create_time"`
	BWSize     uint32 `json:"bandwidth_size"`
}

type EipResp struct {
	Eip EIP `json:"publicip"`
}

func (client *Client) AllocateEIP(region, projectId string, arg *EIPAllocArg) (eIp *EipResp, err error) {
	var eIpResp EipResp

	//compose uri string
	uri := "/" + projectId + "/publicips"

	respbytes, err := client.Do(region, "POST", uri, arg)
	if err != nil {
		return nil, err
	}

	//unmarshal reponse body into elastic ip
	if err := json.Unmarshal(respbytes, &eIpResp); err != nil {
		log.Debugf("json unmarshal error is: %v", err)
		return nil, err
	}

	log.Debugf("elastic ip is: %v", eIpResp)

	return &eIpResp, nil
}

type PortDesc struct {
	PortId string `json:"port_id"`
}

type EIPAssocArg struct {
	Port PortDesc `json:"publicip"`
}

func (client *Client) AssociateEIP(region, projectId, elasticIPId string, arg *EIPAssocArg) error {
	//compose uri string
	uri := "/" + projectId + "/publicips/" + elasticIPId

	_, err := client.Do(region, "PUT", uri, arg)
	if err != nil {
		return err
	}

	return nil
}
