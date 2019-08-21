# Docker Machine Driver OTC
Create machines on OTC. You will need an Access Key ID, Secret Access Key and a Region Name. If you want to setup instance on the VPC network, you will need the VPC ID. Please login to the console and select the one where you would like to launch instance.

Create docker instances on OTC:
```bash
docker-machine create -d otc <machine name>
```

## Installation


## Example Usage
eg. Export your credentials
```bash
export ACCESS_KEY_ID = '<Your access key ID>'
export ACCESS_KEY_SECRET = '<Your secret access key>'

docker-machine create -d otc <machine name>
```
or type cmdline
```bash
docker-machine create -d otc --otc-access-key-id=<Your ak> --otc-access-key-secret=<Your sk> --otc-security-group=<Security group name> --otc-tenant-id=<Your tenant ID> --otc-region=<Region name> --otc-vpc-id=<Your VPC ID> --otc-flavor-id=<instance flavor ID> --otc-image-id=<Guest OS image ID> --otc-subnet-id=<Your subnet ID> --otc-admin-password=<instance login password> --otc-available-zone=<available zone name> <machine name>
```

## Options
```bash
docker-machine create -d otc --help
```
Option Name                                          | Description                                         | required 
------------------------------------------------------|----------------------------------------------------|----|
``--otc-access-key-id`` | Your access key ID.  |**yes**|
``--otc-access-key-secret``|Your secret access key.| **yes** |
``--otc-service-endpoint``|The custom API endpoint.| |
``--otc-image-id``| The image ID of the instance to use, default is the Ubuntu server 14.04 64bits provided by system||
``--otc-region``| The region to use when launching the instance||
``--otc-security-group``| Security group name. Default: `docker-machine`||
``--otc-admin-password``| Admin password for created virtual machine. Default is random generated.||
``--otc-vpc-id``| Your VPC ID to launch the instance in. (required for VPC network only)||
``--otc-available-zone``| The availabilty zone to launch the instance||
``--otc-bandwidth-size``|Bandwidth Size for Elastic IP||
``--otc-bandwidth-type``|Bandwidth Type for Elastic IP||
``--otc-elasticip-type``|Your Elastic IP Type||
``--otc-flavor-id``|Flavor for you instance||
``--otc-root-volume-size``|Root Volume Size for your instance||
``--otc-root-volume-type``|Root Volume type for your instance||
``--otc-ssh-user``|Instance's optional ssh user||
``--otc-subnet-id``|Subnet ID for your instance private network (Network ID)||
``--otc-tenant-id``|Tenant ID (Project ID)||


## Environment variables and default values:

| CLI option                          | Environment variable        | Default          |
|-------------------------------------|-----------------------------|------------------|
| **`--otc-access-key-id`**	| `ACCESS_KEY_ID`	| -                |
| **`--otc-access-key-key`**	| `ACCESS_KEY_SECRET`	| -                |
| `--otc-admin-password`	| `ADMIN_PWD`		| -                |
| `--otc-available-zone`	| `AVAILABLE_ZONE`	| -                |
| `--otc-service-endpoint`	| `SERVICE_ENDPOINT`	| -                |
| `--otc-image-id`		| `IMAGE_ID`		| -                |
| `--otc-region`		| `REGION`		| -                |
| `--otc-security-group`	| `SECURITY_GROUP`	| -                |
| `--otc-vpc-id`		| `VPC_ID`		| -                |
| `--otc-bandwidth-size`	| `BANDWIDTH_SIZE`	| `10`             |
| `--otc-bandwidth-type`	| `BANDWIDTH_TYPE`	| `PER`            |
| `--otc-elasticip-type`	| `ELASTICIP_TYPE`	| `5_bgp`          |
| `--otc-flavor-id`		| `FLAVOR_ID`		| -                |
| `--otc-root-volume-size`	| `ROOT_VOLUME_SIZE`	| `40`             |
| `--otc-root-volume-type`	| `ROOT_VOLUME_TYPE`	| `SATA`           |
| `--otc-ssh-user`		| `SSH_USER`		| -                |
| `--otc-subnet-id`		| `SUBNET_ID`		| -                |
| `--otc-tenant-id`		| `TENANT_ID`		| -                |

Each environment variable may be overloaded by its option equivalent at runtime.

## Installing
### Install Go and git
```bash
yum install golang git
```
### Install DockerMachineDriver4OTC
```bash
export GOPATH=<Path to your Go Build Folder>
go get github.com/huawei/DockerMachineDriver4OTC
cd $GOPATH/src/github.com/huawei/DockerMachineDriver4OTC
./build.sh
```
The Driver Binary will be under ./bin/docker-machine-driver.linux-amd64
Rename it and copy it to you local path
```bash
cp ./bin/docker-machine-driver.linux-amd64 /usr/local/bin/docker-machine-driver-otc
```

Alternatively you can also just rename it and add the folder to your ```$PATH```:
```
export PATH=$GOPATH/src/github.com/huawei/DockerMachineDriver4OTC:$PATH
```

### Binary direct download
You can also directly download the Binary from here:
https://dockermachinedriver.obs.eu-de.otc.t-systems.com/docker-machine-driver-otc


## Related links

- **Docker Machine**: https://docs.docker.com/machine/
- **Contribute**: https://github.com/huawei/DockerMachineDriver4OTC
- **Report bugs**: https://github.com/huawei/DockerMachineDriver4OTC/issues

## License
Apache 2.0

## test command
docker-machine -D create -d otc --otc-access-key-id BCE_shortened --otc-access-key-secret 4UR_shortened --otc-available-zone eu-de-01 --otc-bandwidth-size 10 --otc-bandwidth-type PER --otc-elasticip-type 5_bgp --otc-flavor-id normal1 --otc-image-id d6944a41-5ec7-44a4-970e-ce330da390d2 --otc-region eu-de --otc-root-volume-size 40 --otc-root-volume-type SATA --otc-security-group sg-tino --otc-service-endpoint https://ecs.eu-de.otc.t-systems.com --otc-ssh-user ubuntu --otc-subnet-id 91e2f28b-50dc-4a2d-b856-39d9204323e2 --otc-tenant-id 16d53a84a13b49529d2e2c3646691288 --otc-vpc-id 9d5c46ec-b3f0-42a2-9a65-f2d77e124516 test

## debugging
docker-machine --debug [....]

**parameters may vary on different cloud platforms**
