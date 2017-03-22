package ims

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
	IMSDefaultEndpoint = "https://10.175.38.51"
	IMSVersion         = "v2"
	IMSName            = "cloudimages"
)

type Client struct {
	common.Client
	moduleName string
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{
		moduleName: IMSName,
	}
	client.Init(IMSVersion, IMSDefaultEndpoint, accessKeyId, accessKeySecret)
	return client
}

func (client *Client) Do(region, method, uri string, args interface{}) (respbytes []byte, err error) {
	var b []byte
	var errl error
	var req *http.Request

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
	requestURL := client.Endpoint + "/" + client.Version + uri
	if args != nil {
		req, errl = http.NewRequest(method, requestURL, bytes.NewReader(b))
	} else {
		req, errl = http.NewRequest(method, requestURL, nil)
	}
	if errl != nil {
		return nil, errl
	}

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

	log.Debugf("response body are: %s and status code is: %s", string(body), resp.StatusCode)

	//err code handling
	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode <= 599 {
		return nil, fmt.Errorf("Error Code: %d", statusCode)
	}

	return body, nil
}
