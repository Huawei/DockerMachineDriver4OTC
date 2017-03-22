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
``--otc-api-endpoint``|The custom API endpoint.| |
``--otc-description`` | The description of instance.| |
 ``-otcs-disk-size``| The data disk size for /var/lib/docker (in GB)||
 ``-otcs-disk-category``|The category of data disk, the valid values could be `cloud` (default), `cloud_efficiency` or `cloud_ssd`.|| 
``--otc-image-id``| The image ID of the instance to use, default is the Ubuntu server 14.04 64bits provided by system||
``--otc-io-optimized``| The I/O optimized instance type, the valid values could be `none` (default) or `optimized`||
``--otc-instance-type``| The instance type to run.  Default: `ecs.t1.small`||
``--otc-internet-max-bandwidth``| Maxium bandwidth for Internet access (in Mbps), default 1||
``--otc-private-address-only``| Use the private IP address only||
``--otc-region``| The region to use when launching the instance. Default: `cn-hangzhou`||
``--otc-route-cidr``| The CIDR to use configure the route entry for the instance in VPC. Sample: 192.168.200.0/24||
``--otc-security-group``| Security group name. Default: `docker-machine`||
``--otc-slb-id``|SLB id for instance association||
``--otc-ssh-password``| SSH password for created virtual machine. Default is random generated.||
``--otc-system-disk-category``|System disk category for instance||
``--otc-tag``| Tag for the instance.||
``--otc-vpc-id``| Your VPC ID to launch the instance in. (required for VPC network only)||
``--otc-vswitch-id``| Your VSwitch ID to launch the instance with. (required for VPC network only)||
``--otc-available-zone``| The availabilty zone to launch the instance||

## Environment variables and default values:

| CLI option                          | Environment variable        | Default          |
|-------------------------------------|-----------------------------|------------------|
| **`--otc-access-key-id`**     | `ACCESS_KEY_ID`         | -                |
| **`--otc-access-key-key`**    | `ACCESS_KEY_SECRET`     | -                |
| `--otc-api-endpoint`          | `API_ENDPOINT`          | -                |
| `--otc-description`           | `DESCRIPTION`           | -                |
| `--otc-disk-size`             | `DISK_SIZE`             | -                |
| `--otc-disk-category`         | `DISK_CATEGORY`         | -                |
| `--otc-image-id`              | `IMAGE_ID`              | -                |
| `--otc-aliyunecs-io-optimized`| `IO_OPTIMIZED`          | `none`           |
| `--otc-instance-type`         | `INSTANCE_TYPE`         | `ecs.t1.small`   |
| `--otc-internet-max-bandwidth`| `INTERNET_MAX_BANDWIDTH`| `1`              |
| `--otc-private-address-only`  | `PRIVATE_ADDR_ONLY`     | `false`          |
| `--otc-region`                | `REGION`                | `cn-hangzhou`    |
| `--otc-route-cidr`            | `ROUTE_CIDR`            | -                |
| `--otc-security-group`        | `SECURITY_GROUP`        | -                |
| `--otc-slb-id`                | `SLB_ID`                | -                |
| `--otc-ssh-password`          | `SSH_PASSWORD`          | Random generated |
| `--otc-tag`                   | `TAGS`                  | -                |
| `--otc-vpc-id`                | `VPC_ID`                | -                |
| `--otc-vswitch-id`            | `VSWITCH_ID`            | -                |
| `--otc-available-zone`                  | `ZONE`                  | -                |

Each environment variable may be overloaded by its option equivalent at runtime.

## Hacking
### Get the sources
```bash
go get github.com/huawei/DockerMachineDriver4OTC
cd $GOPATH/src/github.com/huawei/DockerMachineDriver4OTC
```
### Test the driver
To test the driver, make sure your current build directory has been added into ```$PATH``` so that docker-machine can find it:
```
export PATH=$GOPATH/src/github.com/huawei/DockerMachineDriver4OTC:$PATH
```

## Related links

- **Docker Machine**: https://docs.docker.com/machine/
- **Contribute**: https://github.com/huawei/DockerMachineDriver4OTC
- **Report bugs**: https://github.com/huawei/DockerMachineDriver4OTC/issues

## License
Apache 2.0

#test command
./bin/docker-machine -D create -d otc --otc-access-key-id=DFNQWPE4JSXA6BQEOEBY --otc-access-key-secret=1DUumFBPMcE5AIcO6olMNvjmaOA76k0MLTAOfAyM --otc-security-group=default --otc-tenant-id=15eae18081ba40fabd76979bdbf35d0e --otc-region=southchina --otc-vpc-id=ecbd1d70-8c7e-4bdf-bb2c-b3e2b7f7e15b --otc-flavor-id=103 --otc-image-id=627a1223-2ca3-46a7-8d5f-7aef22c74ee6 --otc-subnet-id=a81eee33-c0c3-445f-988a-248ee426fd8d test
=======
# DockerMachineDriver4OTC
