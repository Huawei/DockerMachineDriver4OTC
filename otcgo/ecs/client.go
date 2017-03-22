package ecs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/machine/libmachine/log"
	"github.com/huawei/DockerMachineDriver4OTC/otcgo/common"
	"github.com/huawei/DockerMachineDriver4OTC/otcgo/signer"
)

const (
	ECSDefaultEndpoint = "https://10.175.38.51"
	ECSVersion         = "v1"
	ECSName            = "cloudservers"
)

type Client struct {
	common.Client
	moduleName string
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{
		moduleName: ECSName,
	}
	client.Init(ECSVersion, ECSDefaultEndpoint, accessKeyId, accessKeySecret)
	return client
}

func (client *Client) Do(version, region, method, uri string, args interface{}) (respbytes []byte, err error) {
	var b []byte
	var errl error
	var req *http.Request
	var requestURL string

	//create signer for request signature
	signer := signer.Signer{
		AccessKey: client.AccessKeyId,
		SecretKey: client.AccessKeySecret,
		Region:    region,
		Service:   "ec2",
	}

	//setup request
	if args != nil {
		b, errl = json.Marshal(args)
		if errl != nil {
			return nil, errl
		}
	}

	//Generate request URL
	if version != "" && version != client.Version {
		requestURL = client.Endpoint + "/" + version + uri
	} else {
		requestURL = client.Endpoint + "/" + client.Version + uri
	}
	if args != nil {
		req, errl = http.NewRequest(method, requestURL, bytes.NewReader(b))
	} else {
		req, errl = http.NewRequest(method, requestURL, nil)
	}
	if errl != nil {
		return nil, errl
	}

	req.Header.Set("Content-Type", "application/json")

	//signature http request
	if err := signer.Sign(req); err != nil {
		return nil, err
	}

	log.Debugf("request is: %v", req)
	//issue request
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//log.Debugf("response body are: %v and status code is: %s", body, resp.StatusCode)
	log.Debugf("response body are: %s and status code is: %s", string(body), resp.StatusCode)
	//err code handling
	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode <= 599 {
		return nil, fmt.Errorf("Error code: %d", statusCode)
	}

	return body, nil
}
