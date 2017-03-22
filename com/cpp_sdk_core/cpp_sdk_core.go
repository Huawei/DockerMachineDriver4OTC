/*
 *Copyright 2015 Huawei Technologies Co., Ltd. All rights reserved.
 *	   eSDK is licensed under the Apache License, Version 2.0 (the "License");
 *	   you may not use this file except in compliance with the License.
 *	   You may obtain a copy of the License at
 *
 *	       http://www.apache.org/licenses/LICENSE-2.0
 *
 *
 *	   Unless required by applicable law or agreed to in writing, software
 *	   distributed under the License is distributed on an "AS IS" BASIS,
 *	   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *	   See the License for the specific language governing permissions and
 *	   limitations under the License.
 */
package cpp_sdk_core

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/docker/machine/libmachine/log"
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Resquest header collection
type HeaderValueCollection map[string]string

/**
* @brief		hmacSHA256加密
* @param[in]	key:加密的key值
* @param[in]    content加密的字符串
* @return		[]byte
 */
func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

/*
* @brief		http连接设置
* @param[in]
* @param[out]
* @return
 */
func getTransport() *http.Transport {
	dial := func(netw, addr string) (net.Conn, error) {
		con, err := net.DialTimeout(netw, addr, time.Second*time.Duration(30))
		if err != nil {
			return nil, err
		}

		tcp_conn := con.(*net.TCPConn)
		tcp_conn.SetKeepAlive(false)
		return tcp_conn, nil
	}

	return &http.Transport{
		DisableCompression: true,
		Dial:               dial,
		ResponseHeaderTimeout: time.Second * time.Duration(30),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
}

/*
 * @fn			func SendRequest(requestParam *modules.RequestParam) (result *modules.Result)
 * @brief		Send request
 * @param[in] 	requestParam: Request parameters
 * @param[out]
 * @return		result
 */
func SendRequest(requestParam *modules.RequestParam) (result *modules.Result) {
	result = &modules.Result{}
	result.HeaderCollection = make(map[string]string)
	var urlStr string = requestParam.Url
	uriParts, _ := url.Parse(urlStr)
	headerCollection := make(HeaderValueCollection)
	uriHost := uriParts.Host
	pos := strings.IndexByte(uriHost, ':')
	if pos != -1 {
		uriHost = uriHost[:pos]
	}
	// set host header
	headerCollection[strings.ToLower(modules.HOST_HEADER)] = strings.TrimSpace(uriHost)
	t := time.Now()
	simpleTime := fmt.Sprintf("%4d%02d%02d", t.UTC().Year(), t.UTC().Month(), t.UTC().Day())
	longTime := simpleTime
	longTime += fmt.Sprintf("T%02d%02d%02dZ", t.UTC().Hour(), t.UTC().Minute(), t.UTC().Second())

	// set date header
	switch requestParam.AuthType {
	case modules.V4_AUTH:
		headerCollection[strings.ToLower(modules.SDK_DATE_HEADER)] = strings.TrimSpace(longTime)
	case modules.Token_AUTH:
		if len([]rune(requestParam.Token)) > 0 {
			headerCollection[modules.SDK_X_AUTH_TOKEN] = requestParam.Token
		}
	}
	// set type header
	headerCollection[strings.ToLower(modules.CONTENT_TYPE_HEADER)] = strings.TrimSpace(requestParam.RequestContentType)

	// set content length header
	if len([]byte(requestParam.BodyContent)) > 0 {
		lengthBuffer := strconv.Itoa(len([]byte(requestParam.BodyContent)))
		headerCollection[strings.ToLower(modules.CONTENT_LENGTH_HEADER)] = strings.TrimSpace(lengthBuffer)
	}

	switch requestParam.AuthType {
	case modules.V4_AUTH:
		// compute hash on payload if it exists
		hash := sha256.New()
		hash.Write([]byte(requestParam.BodyContent))
		md := hash.Sum(nil)
		payloadHash := hex.EncodeToString(md)
		var canonicalHeadersStr string
		var signedHeadersValue string

		sorted_keys := make([]string, 0)
		for k, _ := range headerCollection {
			sorted_keys = append(sorted_keys, k)
		}
		// sort 'string' key in increasing order
		sort.Strings(sorted_keys)
		for _, key := range sorted_keys {
			canonicalHeadersStr += strings.ToLower(key) + ":" + headerCollection[key] + "\n"
			signedHeadersValue += key + ";"
		}

		signedHeadersValue = signedHeadersValue[:(len([]rune(signedHeadersValue)) - 1)]
		canonicalHeadersStr = canonicalHeadersStr[:(len([]rune(canonicalHeadersStr)) - 1)]

		// append v4 stuff to the canonical request string
		sortedParameters := uriParts.Query()
		var queryString string
		var first bool = true
		if len(sortedParameters) > 0 {
			queryString = "?"
		}
		pos = strings.IndexByte(urlStr, '?')
		var strQueryString string
		if pos != -1 {
			strQueryString = urlStr[pos:]
		}
		if strings.IndexByte(strQueryString, '=') != -1 {
			for key, value := range sortedParameters {
				if !first {
					queryString += "&"
				}
				first = false
				queryString += key + "=" + fmt.Sprintf("%s", value)
			}
			//strQueryString = queryString
		}
		var signingString string
		escapedPath := fmt.Sprintf("%s", uriParts.EscapedPath())
		escapedPath += "/"
		signingString += requestParam.Method + "\n" + escapedPath + "\n"
		if len([]rune(strQueryString)) > 0 && strings.IndexByte(strQueryString, '=') != -1 {
			signingString += strQueryString[1:] + "\n"
		} else if len([]rune(strQueryString)) > 1 {
			signingString += strQueryString[1:] + "=" + "\n"
		} else {
			signingString += "\n"
		}
		for _, key := range sorted_keys {
			signingString += key + ":" + headerCollection[key]
			signingString += "\n"
		}
		signingString += "\n"
		signingString += signedHeadersValue
		signingString += "\n"
		signingString += payloadHash

		// now compute sha256 on that request string
		hashRequest := sha256.New()
		hashRequest.Write([]byte(signingString))
		mdRequest := hashRequest.Sum(nil)
		strRequest := hex.EncodeToString(mdRequest)

		// generate string to sign
		var strGenerate string
		strGenerate += modules.SDK_HMAC_SHA256 + "\n" + longTime + "\n" + simpleTime + "/" + requestParam.Region + "/" + requestParam.ServiceName + "/" + modules.SDK_REQUEST + "\n" + strRequest

		var tempKey string
		tempKey += "SDK" + requestParam.SK
		dateKey := hmacSHA256([]byte(tempKey), simpleTime)
		dateRegionKey := hmacSHA256([]byte(dateKey), requestParam.Region)
		dateRegionServiceKey := hmacSHA256([]byte(dateRegionKey), requestParam.ServiceName)
		signingKey := hmacSHA256([]byte(dateRegionServiceKey), modules.SDK_REQUEST)
		signatureHmac := hmacSHA256([]byte(signingKey), strGenerate)
		finalSignature := hex.EncodeToString(signatureHmac)

		var authString string
		authString += modules.SDK_HMAC_SHA256 + " " + modules.CREDENTIAL + "=" + requestParam.AK + "/" + simpleTime + "/" + requestParam.Region + "/" + requestParam.ServiceName + "/" + modules.SDK_REQUEST + ", " + modules.SIGNED_HEADERS + "=" + signedHeadersValue + ", " + modules.SIGNATURE + "=" + finalSignature
		headerCollection[strings.ToLower(modules.SDK_AUTHORIZATION_HEADER)] = strings.TrimSpace(authString)
	}

	conn := &http.Client{Transport: getTransport()}
	var req *http.Request = nil
	log.Debugf("request body: %v", requestParam.BodyContent)
	if requestParam.BodyContent != "" {
		req, _ = http.NewRequest(requestParam.Method, requestParam.Url, strings.NewReader(requestParam.BodyContent))
	} else {
		req, _ = http.NewRequest(requestParam.Method, requestParam.Url, nil)
	}
	req.Host = uriHost

	for key, _ := range headerCollection {
		req.Header.Set(key, headerCollection[key])
	}
	log.Debugf("request: %v", req)
	resp, err := conn.Do(req)
	if err != nil {
		result.RespMessage = err.Error()
		return result
	}

	arrayToken := resp.Header["X-Subject-Token"]
	var strToken string
	for _, item := range arrayToken {
		strToken += item
	}
	result.HeaderCollection["X-Subject-Token"] = strToken

	defer resp.Body.Close()
	if resp == nil {
		result.Err = errors.New("Failed to send request with response")
		return result
	}

	result.ResponseCode = resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	result.RespMessage = string(body)
	return result
}
