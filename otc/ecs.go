package otc

import (
	"crypto/md5"
	"crypto/rand"
	_ "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	_ "os"
	"strings"
	"time"

	_ "github.com/huawei/DockerMachineDriver4OTC/otcgo/ecs"
	_ "github.com/huawei/DockerMachineDriver4OTC/otcgo/ims"
	_ "github.com/huawei/DockerMachineDriver4OTC/otcgo/vpc"

	"github.com/huawei/DockerMachineDriver4OTC/com/client"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/ecsModules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/imsModules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/novaModules"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules/vpcModules"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/docker/machine/libmachine/state"
)

const (
	driverName     = "otc"
	defaultSSHUser = "root"
	dockerPort     = 2376

	serviceName          = "ecs"
	secGrpListNumPerPage = 200
	imgListNumPerPage    = 200
	subnetListNumPerPage = 200

	errorMandatoryEnvOrOption string = "%s must be specified either using the environment variable %s or the CLI option %s"
	errorBothOptions          string = "Both %s and %s must be specified"
)

var user_data_temp string = `#cloud-config
write_files:
  - path: /home/linux/.ssh/authorized_keys
    permissions: '0777'
    content: `

type Driver struct {
	*drivers.BaseDriver
	ServiceEndpoint string
	InstanceId      string
	AccessKey       string
	SecretKey       string
	VpcId           string
	PublicKey       []byte
	TenantId        string
	ImageId         string
	Region          string
	AvailableZone   string
	FlavorId        string
	RootVolType     string
	RootVolSize     int
	ElasticIpType   string
	BandwidthSize   int
	BandwidthType   string
	AdminPass       string
	KeyName         string
	JobId           string

	//network
	ElasticIPId       string
	ElasticIP         string
	SecurityGroupName string
	SecurityGroupId   string
	//Nic               ecs.NicDesc
	SubnetId string
	PortId   string

	//individual business service client
	/*ecsclient *ecs.Client
	imsclient *ims.Client
	vpcclient *vpc.Client*/
	PrivateIPAddress string
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:   "otc-elasticip-type",
			Usage:  "Elastic Ip Type",
			Value:  "5_bgp",
			EnvVar: "ELASTICIP_TYPE",
		},
		mcnflag.StringFlag{
			Name:   "otc-bandwidth-type",
			Usage:  "Bandwidth Type",
			Value:  "PER",
			EnvVar: "BANDWIDTH_TYPE",
		},
		mcnflag.IntFlag{
			Name:   "otc-bandwidth-size",
			Usage:  "Bandwidth Size",
			Value:  10,
			EnvVar: "BANDWIDTH_SIZE",
		},
		mcnflag.StringFlag{
			Name:   "otc-service-endpoint",
			Usage:  "Service Endpoint",
			Value:  "",
			EnvVar: "SERVICE_ENDPOINT",
		},
		mcnflag.StringFlag{
			Name:   "otc-access-key-id",
			Usage:  "Access Key ID",
			Value:  "",
			EnvVar: "ACCESS_KEY_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-access-key-secret",
			Usage:  "Access key Secret",
			Value:  "",
			EnvVar: "ACCESS_KEY_SECRET",
		},
		mcnflag.StringFlag{
			Name:   "otc-security-group",
			Usage:  "VPC Security Group",
			Value:  "",
			EnvVar: "SECURITY_GROUP",
		},
		mcnflag.StringFlag{
			Name:   "otc-image-id",
			Usage:  "Machine Image",
			Value:  "",
			EnvVar: "IMAGE_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-available-zone",
			Usage:  "Availabe Zone",
			Value:  "",
			EnvVar: "AVAILABLE_ZONE",
		},
		mcnflag.StringFlag{
			Name:   "otc-region",
			Usage:  "Region",
			Value:  "",
			EnvVar: "REGION",
		},
		mcnflag.StringFlag{
			Name:   "otc-flavor-id",
			Usage:  "Flavor ID",
			Value:  "",
			EnvVar: "FLAVOR_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-vpc-id",
			Usage:  "Vpc ID",
			Value:  "",
			EnvVar: "VPC_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-subnet-id",
			Usage:  "Subnet ID",
			Value:  "",
			EnvVar: "SUBNET_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-root-volume-type",
			Usage:  "Root Volume Type",
			Value:  "SATA",
			EnvVar: "ROOT_VOLUME_TYPE",
		},
		mcnflag.IntFlag{
			Name:   "otc-root-volume-size",
			Usage:  "Root Volume Size",
			Value:  40,
			EnvVar: "ROOT_VOLUME_SIZE",
		},
		mcnflag.StringFlag{
			Name:   "otc-tenant-id",
			Usage:  "Tenant ID",
			Value:  "",
			EnvVar: "TENANT_ID",
		},
		mcnflag.StringFlag{
			Name:   "otc-admin-password",
			Usage:  "Instance's Root Login Password",
			Value:  "",
			EnvVar: "ADMIN_PWD",
		},
		mcnflag.StringFlag{
			Name:   "otc-ssh-user",
			Usage:  "Instance's optional ssh user",
			Value:  "",
			EnvVar: "SSH_USER",
		},
	}
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.ServiceEndpoint = flags.String("otc-service-endpoint")
	d.SecurityGroupName = flags.String("otc-security-group")
	d.AccessKey = flags.String("otc-access-key-id")
	d.SecretKey = flags.String("otc-access-key-secret")
	d.ImageId = flags.String("otc-image-id")
	d.AvailableZone = flags.String("otc-available-zone")
	d.FlavorId = flags.String("otc-flavor-id")
	d.SubnetId = flags.String("otc-subnet-id")
	d.RootVolType = flags.String("otc-root-volume-type")
	d.RootVolSize = flags.Int("otc-root-volume-size")
	d.TenantId = flags.String("otc-tenant-id")
	d.Region = flags.String("otc-region")
	d.VpcId = flags.String("otc-vpc-id")
	d.AdminPass = flags.String("otc-admin-password")
	d.BandwidthSize = flags.Int("otc-bandwidth-size")
	d.BandwidthType = flags.String("otc-bandwidth-type")
	d.ElasticIpType = flags.String("otc-elasticip-type")
	d.SSHUser = flags.String("otc-ssh-user")
	//fmt.Printf("test for region: %s\n", d.Region)
	return d.checkConfig()
}

func (d *Driver) checkConfig() error {
	if d.ServiceEndpoint == "" {
		return fmt.Errorf(errorMandatoryEnvOrOption, "Service Endpoint", "SERVICE_ENDPOINT", "--otc-service-endpoint")
	}
	if d.Region == "" {
		return fmt.Errorf(errorMandatoryEnvOrOption, "Region", "REGION", "otc-region")
	}
	if d.TenantId == "" {
		return fmt.Errorf(errorMandatoryEnvOrOption, "Tenant ID", "TENANT_ID", "--otc-tenant-id")
	}
	if (d.AccessKey == "" && d.SecretKey == "") || (d.AccessKey != "" && d.SecretKey == "") || (d.AccessKey == "" && d.SecretKey != "") {
		return fmt.Errorf(errorBothOptions, d.AccessKey, d.SecretKey)
	}
	return nil
}

func NewDriver(hostName, storePath string) drivers.Driver {
	id := generateId()
	return &Driver{
		InstanceId: id,
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
			SSHUser:     defaultSSHUser,
		}}
}

func generateId() string {
	rb := make([]byte, 10)
	_, err := rand.Read(rb)
	if err != nil {
		log.Errorf("Unable to generate id: %s", err)
	}

	h := md5.New()
	io.WriteString(h, string(rb))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (d *Driver) initClient() *client.Client {
	var clientConf modules.ClientConfiguration

	clientConf.Endpoint = strings.Split(d.ServiceEndpoint, "https://ecs.")[1]
	clientConf.ServiceName = serviceName
	clientConf.Region = d.Region

	gClient := client.InitV4Client(d.AccessKey, d.SecretKey, d.TenantId, clientConf)

	return gClient
}

func (d *Driver) Create() error {
	var err error

	if err := d.checkPrereqs(); err != nil {
		return err
	}

	log.Infof("%s | Creating key pair for instance ...", d.MachineName)
	if err := d.createKeyPair(); err != nil {
		return fmt.Errorf("%s | Failed to create key pair: %v", d.MachineName, err)
	}

	//var base64string = "I2Nsb3VkLWNvbmZpZwp3cml0ZV9maWxlczoKICAtIHBhdGg6IC9ob21lL2xpbnV4L3Rlc3QKICAgIHBlcm1pc3Npb25zOiAnMDc3NycKICAgIGNvbnRlbnQ6IFlvdXIgVmFsdWU="
	//templatebytes, err := base64.StdEncoding.DecodeString(base64string)
	//if err != nil {
	//	log.Infof("%s | failed to decode base64 string", d.MachineName)
	//}
	//log.Infof("%s | templatebytes: %v", d.MachineName, templatebytes)
	/*log.Infof("%s | const user data temp: %v", d.MachineName, []byte(user_data_temp))

	testfile, err := os.Create("test")
	if err != nil {
		log.Infof("%s | failed to create test file", d.MachineName)
	}

	user_data_input := user_data_temp + string(d.PublicKey)
	//if _, err := testfile.Write(templatebytes); err != nil {
	if _, err := testfile.Write([]byte(user_data_input)); err != nil {
		log.Infof("%s | failed to write file test", d.MachineName)
	}
	for {
	}*/

	log.Infof("%s | Configuring security groups for instance ...", d.MachineName)
	if err := d.configureSecurityGroup(d.SecurityGroupName); err != nil {
		return fmt.Errorf("%s | Failed to configure instance security groups: %v", d.MachineName, err)
	}

	log.Infof("%s | Configuring instance flavor ...", d.MachineName)
	if err := d.configureFlavor(d.FlavorId); err != nil {
		return fmt.Errorf("%s | Failed to configure instance flavor: %v", d.MachineName, err)
	}

	log.Infof("%s | Configure instance image ...", d.MachineName)
	if err := d.configureImage(d.ImageId); err != nil {
		return fmt.Errorf("%s | Failed to configure instance image: %v", d.MachineName, err)
	}

	log.Infof("%s | Configuring subnet for instance ...", d.MachineName)
	if err := d.configureSubnet(d.SubnetId); err != nil {
		return fmt.Errorf("%s | Failed to configure instance subnet: %v ", d.MachineName, err)
	}

	log.Infof("%s | Create SSH key pair ...", d.MachineName)
	if err := d.instanceKeyPairCreate(); err != nil {
		return fmt.Errorf("%s | Failed to crate ssh key pair: %v", d.MachineName, err)
	}

	log.Infof("%s | Issuing request to create instance ...", d.MachineName)
	createResp, err := d.createInstance()
	if err != nil || modules.HttpOK != createResp.ResponseCode {
		pClient := d.initClient()
		deleteKeyPairResp := pClient.DeleteKeyPair(d.KeyName)
		if deleteKeyPairResp.ResponseCode != modules.HttpOK {
			return fmt.Errorf("%s | Failed to delete key pair %s: %v", d.MachineName, d.InstanceId, deleteKeyPairResp.ErrorInfo.Description)
		}
		return fmt.Errorf("%s | Failed to create instance: %v, response code is %d", d.MachineName, err, createResp.ResponseCode)
	}
	d.JobId = createResp.Job_id
	log.Infof("%s | job for creating instance: %s and reponse code: %d and err: %v", d.MachineName, d.JobId, createResp.ResponseCode, createResp.ErrorInfo)

	//check status of creating instance
	log.Infof("%s | Checking instance status...", d.MachineName)
	if err := d.checkJobStatus(d.JobId); err != nil {
		return fmt.Errorf("%s | Failed to check instance", d.MachineName)
	}
	//configure network for elastic cloud server
	log.Infof("%s | Configure network for elastic colud server ...", d.MachineName)
	if err := d.configureNetwork(); err != nil {
		return fmt.Errorf("%s | Failed to configure instance network: %v", d.MachineName, err)
	}

	return err
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

func (d *Driver) GetIP() (string, error) {
	return d.ElasticIP, nil
}

func (d *Driver) Start() error {
	log.Infof("%s | Start instance...", d.MachineName)

	pClient := d.initClient()

	startResp := pClient.StartServer(d.InstanceId)
	if startResp.ResponseCode != modules.HttpOK {
		return fmt.Errorf("%s | Failed to start instance: %v", d.MachineName, startResp.ErrorInfo.Description)
	}

	return nil
}

func (d *Driver) Stop() error {
	log.Infof("%s | Stop instance...", d.MachineName)

	pClient := d.initClient()

	stopResp := pClient.StopServer(d.InstanceId)
	if stopResp.ResponseCode != modules.HttpOK {
		return fmt.Errorf("%s | Failed to stop instance: %v", d.MachineName, stopResp.ErrorInfo.Description)
	}

	return nil
}

func (d *Driver) DriverName() string {
	return driverName
}

func (d *Driver) GetURL() (string, error) {
	return fmt.Sprintf("tcp://%s:%d", d.ElasticIP, dockerPort), nil
}

func (d *Driver) GetState() (state.State, error) {
	log.Infof("%s | Checking instance status...", d.MachineName)

	pClient := d.initClient()

	stateResp := pClient.ShowServer(d.InstanceId)
	if stateResp.ResponseCode != modules.HttpOK {
		err := fmt.Errorf("%s | Failed to get instance %s: state %v", d.MachineName, d.InstanceId, stateResp.ErrorInfo.Description)
		log.Error(err)
		return state.None, err
	}

	// FIXME: PrivateIPAddress is not stored to config.json for now
	// because docker-machine did not get any update from drivers
	for _, paddrs := range stateResp.Server.Addresses {
		for _, paddr := range paddrs {
			if paddr.Addr != "" {
				split := strings.Split(paddr.Addr, " ")
				d.PrivateIPAddress = split[0]
			}
		}
	}

	switch stateResp.Server.Status {
	case "ACTIVE":
		return state.Running, nil
	case "SHUTOFF":
		return state.Stopped, nil
	case "ERROR":
		return state.Error, nil
	}

	return state.None, nil
}

func (d *Driver) Kill() error {
	return d.Stop()
}

func (d *Driver) Remove() error {
	log.Infof("%s | Removing instance...", d.MachineName)

	/*pClient := d.initClient()
	bandwidthListResp := pClient.ListBandwidths(20, "")
	log.Infof("bandwidth list: %v", bandwidthListResp.Bandwidths)*/
	var server ecsModules.ServerId
	server.Init(d.InstanceId)

	var serverList []ecsModules.ServerId
	serverList = append(serverList, server)

	var deleteServerReq ecsModules.DeleteCloudServerReq
	deleteServerReq.Init(serverList, false, false)

	pClient := d.initClient()

	deleteServerResp := pClient.DeleteCloudServer(deleteServerReq)
	if deleteServerResp.ResponseCode != modules.HttpOK {
		return fmt.Errorf("%s | Failed to delete instance %s: %v", d.MachineName, d.InstanceId, deleteServerResp.ErrorInfo.Description)
	}

	//check status of creating instance
	log.Infof("%s | Checking instance status...", d.MachineName)
	if err := d.checkJobStatus(deleteServerResp.Job_id); err != nil {
		return fmt.Errorf("%s | Failed to delete instance %s: %v", d.MachineName, d.InstanceId, err)
	}

	if d.ElasticIPId != "" {
		log.Infof("%s | Removing elastic ip...", d.MachineName)
		pClient = d.initClient()
		deleteElasticIpResp := pClient.DeletePublicIp(d.ElasticIPId)
		if deleteElasticIpResp.ResponseCode != modules.HttpOK {
			return fmt.Errorf("%s | Failed to remove elastic ip: %v", d.MachineName, deleteElasticIpResp.ErrorInfo.Description)
		}
	}

	if d.KeyName != "" {
		log.Infof("%s | Removing instance key pair(%s)...", d.MachineName, d.KeyName)
		pClient = d.initClient()
		deleteKeyPairResp := pClient.DeleteKeyPair(d.KeyName)
		if deleteKeyPairResp.ResponseCode != modules.HttpOK {
			return fmt.Errorf("%s | Failed to delete key pair %s: %v", d.MachineName, d.InstanceId, deleteKeyPairResp.ErrorInfo.Description)
		}
	}

	return nil
}

func (d *Driver) Restart() error {
	log.Infof("%s | Restarting instance...", d.MachineName)

	pClient := d.initClient()

	rebootServerResp := pClient.RebootServer(d.InstanceId, "SOFT")
	if rebootServerResp.ResponseCode != modules.HttpOK {
		return fmt.Errorf("%s | Failed to restart instance %s: %v", d.MachineName, d.InstanceId, rebootServerResp.ErrorInfo.Description)
	}
	return nil
}

func (d *Driver) instanceKeyPairCreate() error {
	var keypairCreate novaModules.KeypairCreate
	var createKeypairReq novaModules.CreateKeypairReq

	log.Debugf("%s | Create ssh key pair...", d.MachineName)

	pClient := d.initClient()

	keypairCreate.Init(d.MachineName)
	keypairCreate.SetPublic_key(string(d.PublicKey))

	createKeypairReq.Init(keypairCreate)

	kp := pClient.CreateKeypair(createKeypairReq)
	if kp.Keypair.Name != "" {
		log.Debugf("%s | Success to create key pair: %s", d.MachineName, kp.Keypair.Name)
		d.KeyName = kp.Keypair.Name
		return nil
	} else if kp.ResponseCode == 409 {
		return fmt.Errorf("Keypair(%s) existed, please remove it from the console, err %v", d.MachineName, kp.ErrorInfo)
	} else {
		return fmt.Errorf("unknown error, status code: %d, err: %v", kp.ResponseCode, kp.ErrorInfo)
	}

}

func (d *Driver) createInstance() (ecsModules.CreateCloudServerResp, error) {
	var nics ecsModules.Nics
	nics.Init(d.SubnetId)

	var nicsList []ecsModules.Nics
	nicsList = append(nicsList, nics)

	/*var secGrp ecsModules.SecGrp
	secGrp.Init(d.SecurityGroupId)

	var sgList []ecsModules.SecGrp
	sgList = append(sgList, secGrp)*/
	var secGrp ecsModules.SecurityGroup
	secGrp.Init()
	secGrp.SetId(d.SecurityGroupId)

	var sgList []ecsModules.SecurityGroup
	sgList = append(sgList, secGrp)

	var rootVol ecsModules.RootVolume

	rootVol.Init(d.RootVolType)
	rootVol.SetSize(d.RootVolSize)

	var instanceDesc ecsModules.CreateCloudServer
	//instanceDesc.Init(d.ImageId, d.FlavorId, d.GetMachineName(), d.VpcId, nicsList, rootVol, d.AdminPass, d.KeyName, d.AvailableZone, sgList)
	instanceDesc.Init(d.ImageId, d.FlavorId, d.GetMachineName(), d.VpcId, nicsList, rootVol, d.AvailableZone)
	instanceDesc.SetKey_name(d.KeyName)
	instanceDesc.SetAdminPass(d.AdminPass)
	instanceDesc.SetSecurity_groups(sgList)

	/*log.Debugf("%s | SSH User: %s", d.MachineName, d.SSHUser)
	if d.SSHUser != "" {
		var perItem ecsModules.Personality
		path := "/home/" + d.SSHUser + "/.ssh/authorized_keys"
		perItem.Init(path, base64.StdEncoding.EncodeToString([]byte(d.PublicKey)))
		//encoding := base64.StdEncoding.EncodeToString([]byte(d.PublicKey))
		//perItem.Init(path, encoding)
		//decoding, _ := base64.StdEncoding.DecodeString(encoding)
		//log.Infof("decoding: %s", string(decoding))

		var perList []ecsModules.Personality
		perList = append(perList, perItem)
		instanceDesc.SetPersonality(perList)
	}
	user_data_input := user_data_temp + string(d.PublicKey)
	instanceDesc.SetUser_data(base64.StdEncoding.EncodeToString([]byte(user_data_input)))*/

	var instanceCreateReq ecsModules.CreateCloudServerReq
	instanceCreateReq.Init(instanceDesc)

	pClient := d.initClient()

	instanceCreateResp := pClient.CreateCloudServer(instanceCreateReq)

	return instanceCreateResp, nil
}

func (d *Driver) checkJobStatus(jobid string) error {
	pClient := d.initClient()

	for {
		ecsJobStatusResp := pClient.ShowEcsJob(jobid)
		if len(ecsJobStatusResp.Entities.SubJobs) > 0 {
			d.InstanceId = ecsJobStatusResp.Entities.SubJobs[0].Entities.Server_id
		}

		if ecsJobStatusResp.Status == "SUCCESS" {
			log.Debugf("%s | job return value are: %v and returned instance id is: %s", d.MachineName, ecsJobStatusResp, d.InstanceId)
			break
		}
		if ecsJobStatusResp.Status == "FAIL" {
			log.Debugf("%s | Failed to check instance status: %v", d.MachineName, ecsJobStatusResp.ErrorInfo.Description)
			return fmt.Errorf("%s | Failed to check instance status: %v", d.MachineName, ecsJobStatusResp.ErrorInfo.Description)
		}
		log.Debugf("job status: %s", ecsJobStatusResp.Status)
		time.Sleep(60 * time.Second)
	}

	return nil
}

func (d *Driver) configureNetwork() error {
	log.Debugf("%s | Allocate elastic ip ...", d.MachineName)

	elasticIpCreate := &vpcModules.PublicipCreate{}
	elasticIpCreate.SetType(d.ElasticIpType)

	bandwidthCreate := &vpcModules.BandwidthCreate{}
	bandwidthCreate.Init(d.MachineName, d.BandwidthSize, d.BandwidthType)

	createElasticIp := &vpcModules.CreatePublicIpReq{}
	createElasticIp.Init(elasticIpCreate, bandwidthCreate)

	pClient := d.initClient()
	createElasticIpResp := pClient.CreatePublicIp(createElasticIp)

	if createElasticIpResp.Id != "" {
		log.Debugf("%s | Success to allocate elastic ip %s and status is %s", d.MachineName, createElasticIpResp.Public_ip_address, createElasticIpResp.Status)
		d.ElasticIPId = createElasticIpResp.Id
		d.ElasticIP = createElasticIpResp.Public_ip_address
	}

	log.Debugf("%s | Query nic port id ...", d.MachineName)
	//check elastic server instance's interface
	pClient = d.initClient()
	interfaceList := pClient.ListInterfaces(d.InstanceId)

	for _, iface := range interfaceList.InterfaceAttachments {
		if iface.Port_id != "" {
			log.Debugf("%s | Found instance network interface port %s", d.MachineName, iface.Port_id)
			d.PortId = iface.Port_id
			break
		}
	}

	log.Debugf("%s | Waiting for elastic ip to be ready ...", d.MachineName)
	//waiting for elastic ip to be ready
	time.Sleep(60 * time.Second)

	log.Debugf("%s | Associate elastic ip to instance nic port ...", d.MachineName)

	elasticIpUpdate := &vpcModules.PublicipUpdate{}
	elasticIpUpdate.SetPortId(d.PortId)
	updateElasticIpReq := &vpcModules.UpdatePublicIpReq{}
	updateElasticIpReq.SetPublicip(elasticIpUpdate)

	pClient = d.initClient()
	updateElasticIpResp := pClient.UpdatePublicIp(d.ElasticIPId, updateElasticIpReq)
	if updateElasticIpResp.ErrorInfo.Description != "" {
		return fmt.Errorf("%s", updateElasticIpResp.ErrorInfo.Description)
	}
	log.Debugf("%s | Success to associate elastic ip to instance nic port ...", d.MachineName)

	return nil
}

func (d *Driver) configureSubnet(subnetId string) error {
	log.Debugf("%s | Checking instance subnet list on cloud ...", d.MachineName)

	pClient := d.initClient()

	subnetList := pClient.ListSubnets(subnetListNumPerPage, "", d.VpcId)
	log.Debugf("subnets:%v response code:%d error info:%v", subnetList.Subnets, subnetList.ResponseCode, subnetList.ErrorInfo)

	for _, subnet := range subnetList.Subnets {
		if subnetId == subnet.Id {
			log.Debugf("%s | Found existing subnet %s", d.MachineName, subnet.Id)
			d.SubnetId = subnet.Id
			d.AvailableZone = subnet.Availability_zone
			log.Debugf("%s | Subnet id is: %s : %s and available zone is: %s", d.MachineName, subnet.Id, d.SubnetId, d.AvailableZone)
			break
		}
	}

	return nil
}

func (d *Driver) configureImage(imageId string) error {
	log.Debugf("%s | Checking instance image list on cloud ...", d.MachineName)

	pClient := d.initClient()

	listImgsReq := &imsModules.ListCloudImagesReqEx{}
	listImgsReq.Init()
	//listImgsReq.SetLimit(imgListNumPerPage)

	imagesList := pClient.ListCloudImages(listImgsReq)
	// buggy log.Debugf("%v", imagesList.Images)

	for _, image := range imagesList.Images {
		if imageId == image.Id {
			log.Debugf("%s | Found existing image %s", d.MachineName, image.Id)
			d.ImageId = image.Id
			log.Debugf("%s | Image id is: %s and name is: %s", d.MachineName, d.ImageId, image.Name)
			break
		}
	}

	return nil
}

func (d *Driver) configureFlavor(flavorId string) error {
	log.Debugf("%s | Checking instance flavor list on cloud ...", d.MachineName)

	pClient := d.initClient()

	flavorList := pClient.ListCloudServerFlavorsExt()
	log.Debugf("%v", flavorList.Flavors)
	if flavorList.ErrorInfo.Code != "" {
		log.Infof("%s | Failed to get flavor", d.MachineName)
		log.Debugf("Error code: %s and description: %s", flavorList.ErrorInfo.Code, flavorList.ErrorInfo.Description)
		return fmt.Errorf("Error code: %s and description: %s", flavorList.ErrorInfo.Code, flavorList.ErrorInfo.Description)
	}
	for _, flavor := range flavorList.Flavors {
		if flavorId == flavor.Id {
			log.Debugf("%s | Found existing flavor spec %s", d.MachineName, flavor.Id)
			d.FlavorId = flavor.Id
			log.Debugf("%s | Flavor id is: %s", d.MachineName, d.FlavorId)
			break
		}
	}

	return nil
}

func (d *Driver) configureSecurityGroup(groupName string) error {
	log.Debugf("%s | Configure security groups", d.MachineName)

	pClient := d.initClient()

	listSecGrpRsp := pClient.ListSecurityGroups(secGrpListNumPerPage, "", "")
	for _, grp := range listSecGrpRsp.SecurityGroups {
		if grp.Name == groupName {
			log.Debugf("%s | Found existing security group (%s) in %s", d.MachineName, groupName, d.VpcId)
			d.SecurityGroupId = grp.Id
			log.Debugf("%s | Security group id is: %s", d.MachineName, d.SecurityGroupId)
			break
		}
	}
	return nil
}

func (d *Driver) PreCreateCheck() error {
	return d.checkPrereqs()
}

func (d *Driver) checkPrereqs() error {
	return nil
}

/*func (d *Driver) getECSClient() *ecs.Client {
	if d.ecsclient == nil {
		client := ecs.NewClient(d.AccessKey, d.SecretKey)
		d.ecsclient = client
	}
	return d.ecsclient
}

func (d *Driver) getVPCClient() *vpc.Client {
	if d.vpcclient == nil {
		client := vpc.NewClient(d.AccessKey, d.SecretKey)
		d.vpcclient = client
	}
	return d.vpcclient
}

func (d *Driver) getIMSClient() *ims.Client {
	if d.imsclient == nil {
		client := ims.NewClient(d.AccessKey, d.SecretKey)
		d.imsclient = client
	}
	return d.imsclient
}*/

func (d *Driver) createKeyPair() error {
	log.Debugf("%s | SSH key path:%s", d.MachineName, d.GetSSHKeyPath())

	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return err
	}

	publicKey, err := ioutil.ReadFile(d.GetSSHKeyPath() + ".pub")
	if err != nil {
		return err
	}

	d.PublicKey = publicKey
	return nil
}
