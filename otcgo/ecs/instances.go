package ecs

import (
	"encoding/json"

	"github.com/docker/machine/libmachine/log"
)

type NicDesc struct {
	SubnetId string `json:"subnet_id"`
	IpAddr   string `json:"ip_address"`
}

type RootVolDesc struct {
	VolType string `json:"volumetype"`
	Size    uint32 `json:"size"`
}

type SecGrp struct {
	Id string `json:"id"`
}

type CreateInstanceAttribute struct {
	ImageId       string      `json:"imageRef"`
	FlavorId      string      `json:"flavorRef"`
	InstanceName  string      `json:"name"`
	VpcId         string      `json:"vpcid"`
	AvailableZone string      `json:"availability_zone"`
	Nic           []NicDesc   `json:"nics"`
	RootVol       RootVolDesc `json:"root_volume"`
	SecGrps       []SecGrp    `json:"security_groups"`
	AdminPass     string      `json:"adminPass"`
	KeyName       string      `json:"key_name"`
	UserData      []byte      `json:"user_data"`
}

type CreateInstanceArgs struct {
	ServerAttr CreateInstanceAttribute `json:"server"`
}

type CreateInstanceResponse struct {
}

type JobPara struct {
	ProjectId string
	JobId     string
}

//{"status":"SUCCESS","entities":{"sub_jobs_total":1,"sub_jobs":[{"status":"SUCCESS","entities":{"server_id":"d89f8ca7-e54b-4d00-b539-d2621707f1a3"},"job_id":"4010b39b5404b277015404b72e8a0090","job_type":"createSingleServer","begin_time":"2016-04-11T09:46:34.505Z","end_time":"2016-04-11T09:50:16.819Z","error_code":null,"fail_reason":null}]},"job_id":"4010b39b5404b277015404b71f150084","job_type":"createServer","begin_time":"2016-04-11T09:46:30.547Z","end_time":"2016-04-11T09:50:34.756Z","error_code":null,"fail_reason":null}
type Server struct {
	Id string `json:"server_id"`
}

type SubJob struct {
	Status     string `json:"status"`
	Entities   Server `json:"entities"`
	JobId      string `json:"job_id"`
	JobType    string `json:"job_type"`
	BeginTime  string `json:"begin_time"`
	EndTime    string `json:"end_time"`
	ErrorCode  string `json:"error_code"`
	FailReason string `json:"fail_reason"`
	Message    string `json:"message"`
	Code       string `json:"code"`
}

type Entity struct {
	SubJobsTotal uint32   `json:"sub_jobs_total"`
	SubJobs      []SubJob `json:"sub_jobs"`
}

type JobRespDesc struct {
	Status       string `json:"status"`
	ServerEntity Entity `json:"entities"`
	JobId        string `json:"job_id"`
	JobType      string `json:"job_type"`
	BeginTime    string `json:"begin_time"`
	EndTime      string `json:"end_time"`
	ErrorCode    string `json:"error_code"`
	FailReason   string `json:"fail_reason"`
	Message      string `json:"message"`
	Code         string `json:"code"`
}

type NormalResp struct {
	Job_Id string `json:"job_id"`
}

type ErrBody struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type AbnormalResp struct {
	Error ErrBody `json:"error"`
}

func (client *Client) CreateInstance(region, projectId string, args *CreateInstanceArgs) (retVal interface{}, err error) {
	var normalResp NormalResp

	//compose uri string
	uri := "/" + projectId + "/cloudservers"
	respbytes, err := client.Do("", region, "POST", uri, args)
	if err != nil {
		return "", err
	}

	//unmarshal response body data into job response
	if err := json.Unmarshal(respbytes, &normalResp); err != nil {
		return nil, err
	}

	return normalResp.Job_Id, nil
}

type Svr struct {
	Id string `json:"id"`
}

type DeleteInstanceArg struct {
	Servers  []Svr `json:"servers"`
	DelPubIp bool  `json:"delete_publicip"`
	DelVol   bool  `json:"delete_volume"`
}

func (client *Client) DeleteInstance(region, projectId string, arg *DeleteInstanceArg) (retVal interface{}, err error) {
	var normalResp NormalResp

	//compose uri string
	uri := "/" + projectId + "/cloudservers/delete"
	respbytes, err := client.Do("", region, "POST", uri, arg)
	if err != nil {
		return "", err
	}

	//unmarshal response body data into job response
	if err := json.Unmarshal(respbytes, &normalResp); err != nil {
		return nil, err
	}
	return normalResp.Job_Id, nil
}

func (client *Client) CheckInstanceStatus(region string, args *JobPara) (jobResp *JobRespDesc, err error) {
	var jobRespDesc JobRespDesc

	//compose uri string
	uri := "/" + args.ProjectId + "/jobs/" + args.JobId
	respbytes, err := client.Do("", region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal response body data into keypair
	if err := json.Unmarshal(respbytes, &jobRespDesc); err != nil {
		return nil, err
	}

	return &jobRespDesc, nil
}

type OS_Ext_Spec struct {
	PerfType  string `json:"ecs:performancetype"`
	AvailZone string `json:"ecs:availablezone"`
}

type LinkDesc struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}

type FlavorList struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Vcpus         string      `json:"vcpus"`
	Ram           uint32      `sjon:"ram"`
	Disk          string      `json:"disk"`
	Swap          string      `json:"swap"`
	Flavor_OS_Ext uint32      `json:"OS-FLV-EXT-DATA:ephemeral"`
	Flavor_OS_Dis bool        `json:"OS-FLV-DISABLED:disabled"`
	RxTx_Factor   float32     `json:"rxtx_factor"`
	RxTx_Quota    string      `json:"rxtx_quota"`
	RxTx_Cap      string      `json:"rxtx_cap"`
	Flavor_OS_Acc bool        `json:"os-flavor-access:is_public"`
	Links         []LinkDesc  `json:"links"`
	Os_Ext_Specs  OS_Ext_Spec `json:"os_extra_specs"`
}

type FlavorResponse struct {
	Flavors []FlavorList `json:"flavors"`
}

func (client *Client) ListFlavors(region, projectId string) (flvList *FlavorResponse, err error) {
	var flavorList FlavorResponse

	//compose uri string
	uri := "/" + projectId + "/cloudservers/flavors"

	respbytes, err := client.Do("", region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal response body data into flavor list
	if err := json.Unmarshal(respbytes, &flavorList); err != nil {
		return nil, err
	}

	log.Debugf("flavor list are: %v", flavorList)
	return &flavorList, nil
}

type FixIP struct {
	SubnetId string `json:"subnet_id"`
	IpAddr   string `json:"ip_address"`
}

type Interface struct {
	PortState string  `json:"port_state"`
	FixIPs    []FixIP `json:"fixed_ips"`
	NetId     string  `json:"net_id"`
	PortId    string  `json:"port_id"`
	MacAddr   string  `json:"mac_addr"`
}

type InterfaceResponse struct {
	InterfaceList []Interface `json:"interfaceAttachments"`
}

func (client *Client) QueryInterfaces(region, projectId, instanceId string) (ifList *InterfaceResponse, err error) {
	var interfaceList InterfaceResponse

	//compose uri string
	uri := "/" + projectId + "/servers/" + instanceId + "/os-interface"

	respbytes, err := client.Do("v2", region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal reponse body data into interface list
	if err := json.Unmarshal(respbytes, &interfaceList); err != nil {
		return nil, err
	}

	log.Debugf("interface list are: %v", interfaceList)

	return &interfaceList, nil
}

type KeyPair struct {
	Name   string `json:"name"`
	PubKey string `json:"public_key"`
}

type KeyPairAttr struct {
	KP KeyPair `json:"keypair"`
}

type KeyPairDesc struct {
	FP         string `json:"fingerprint"`
	Name       string `json:"name"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	UserId     string `json:"user_id"`
}

type KeyPairResponse struct {
	KPr KeyPairDesc `json:"keypair"`
}

func (client *Client) CreateKeyPair(region, projectId string, arg *KeyPairAttr) (kp *KeyPairResponse, err error) {
	var keyPair KeyPairResponse

	//compose uri string
	uri := "/" + projectId + "/os-keypairs"

	respbytes, err := client.Do("v2", region, "POST", uri, arg)
	if err != nil {
		return nil, err
	}

	//unmarshal response body data into keypair
	if err := json.Unmarshal(respbytes, &keyPair); err != nil {
		return nil, err
	}

	log.Debugf("key pair is: %v", keyPair)

	return &keyPair, nil
}

func (client *Client) DeleteKeyPair(region, projectId, kpName string) error {
	//compose uri string
	uri := "/" + projectId + "/os-keypairs/" + kpName

	_, err := client.Do("v2", region, "DELETE", uri, nil)
	if err != nil {
		return err
	}

	return nil
}

type InstanceStartAttr struct {
	StartCmd string `json:"os-start"`
}

func (client *Client) StartInstance(region, projectId, instanceId string, arg *InstanceStartAttr) error {
	//compose uri string
	uri := "/" + projectId + "/servers/" + instanceId + "/action"

	_, err := client.Do("v2", region, "POST", uri, arg)
	if err != nil {
		return err
	}

	return nil
}

type InstanceStopAttr struct {
	StopCmd string `json:"os-stop"`
}

func (client *Client) StopInstance(region, projectId, instanceId string, arg *InstanceStopAttr) error {
	//compose uri string
	uri := "/" + projectId + "/servers/" + instanceId + "/action"

	_, err := client.Do("v2", region, "POST", uri, arg)
	if err != nil {
		return err
	}

	return nil
}

type CmdType struct {
	Type string `json:"type"`
}

type InstanceRestartAttr struct {
	RebootCmd CmdType `json:"reboot"`
}

func (client *Client) RestartInstance(region, projectId, instanceId string, arg *InstanceRestartAttr) error {
	//compose uri string
	uri := "/" + projectId + "/servers/" + instanceId + "/action"

	_, err := client.Do("v2", region, "POST", uri, arg)
	if err != nil {
		return err
	}

	return nil
}

type EmbServer struct {
	Status string `json:"status"`
}

type StateResp struct {
	Server EmbServer `json:"server"`
}

func (client *Client) GetInstanceState(region, projectId, instanceId string) (sr *StateResp, err error) {
	var stateResp StateResp

	//compose uri string
	uri := "/" + projectId + "/servers/" + instanceId

	respbytes, err := client.Do("v2", region, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	//unmarshal reponse body into state response
	if err := json.Unmarshal(respbytes, &stateResp); err != nil {
		return nil, err
	}

	log.Debugf("response state is: %v", stateResp)

	return &stateResp, nil
}
