package vpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "time"

	"github.com/docker/machine/libmachine/log"
	"github.com/huawei/DockerMachineDriver4OTC/otcgo/common"
	"github.com/huawei/DockerMachineDriver4OTC/otcgo/signer"
)

const (
	VPCDefaultEndpoint = "https://10.175.38.51"
	VPCVersion         = "v1"
	VPCName            = "vpcs"
	BasicDateFormat    = "20060102T150405Z"
)

type Client struct {
	common.Client
	moduleName string
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{
		moduleName: VPCName,
	}
	client.Init(VPCVersion, VPCDefaultEndpoint, accessKeyId, accessKeySecret)
	return client
}

func (client *Client) Do(region, method, uri string, args interface{}) (respbytes []byte, err error) {
	var b []byte
	var errl error
	var req *http.Request

	//initial signer for request signature
	signer := signer.Signer{
		AccessKey: client.AccessKeyId,
		SecretKey: client.AccessKeySecret,
		Region:    region,
		Service:   "ec2",
	}

	if args != nil {
		b, errl = json.Marshal(args)
		if errl != nil {
			return nil, errl
		}
	}
	//compose request URL
	requestURL := client.Endpoint + "/" + client.Version + uri
	//requestURL := client.Endpoint + "/v2/cloudimages"  //check images list
	//requestURL := client.Endpoint + "/v1/15eae18081ba40fabd76979bdbf35d0e/cloudservers/flavors"  //check flavors list
	if args != nil {
		req, errl = http.NewRequest(method, requestURL, bytes.NewReader(b))
	} else {
		req, errl = http.NewRequest(method, requestURL, nil)
	}
	if errl != nil {
		return nil, errl
	}

	//signature http request
	//req.Header.Add("date", time.Now().Format(BasicDateFormat))
	if err := signer.Sign(req); err != nil {
		log.Debugf("sign error: %v", err)
		return nil, err
	}

	//check request details for debuging
	log.Debugf("request authorization header after signature: %v", req.Header.Get("authorization"))
	log.Debugf("request is: %v", req)

	//issue request
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		log.Debugf("http request error: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debugf("error of response body after I/O reading is: %v", err)
		return nil, err
	}

	log.Debugf("response body after io read: %v", string(body))
	//err code handling
	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode <= 599 {
		log.Debugf("response status code for request is: %d", statusCode)
		return nil, fmt.Errorf("Error code: %d", statusCode)
	}
	log.Debugf("status code is: %d", statusCode)

	return body, nil
}
