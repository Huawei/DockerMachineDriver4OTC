package common

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	HttpClient      *http.Client
	Endpoint        string
	Version         string
}

var defaultTransport *http.Transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

func (client *Client) Init(version, endpoint, accessKeyId, accessKeySecret string) {
	client.AccessKeyId = accessKeyId
	client.AccessKeySecret = accessKeySecret
	tlsConfig := tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	defaultTransport.TLSClientConfig = &tlsConfig
	client.HttpClient = &http.Client{Transport: defaultTransport}
	client.Endpoint = endpoint
	client.Version = version
}
