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
package modules

// Http response code
const (
	HttpOK                         = 200
	HttpBadRequest                 = 400
	HttpUnauthorized               = 401
	HttpForbidden                  = 403
	HttpNotFound                   = 404
	HttpMethodNotAllowed           = 405
	HttpNotAcceptable              = 406
	HttpProxyAuthenticationRequied = 407
	HttpRequestTimeout             = 408
	HttpConflict                   = 409
	HttpInternalServerError        = 500
	HttpNotImplemented             = 501
	HttpBadGateway                 = 502
	HttpServiceUnavailable         = 503
	HttpGatewayTimeout             = 504
)

// Http method
const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_DELETE = "DELETE"
	HTTP_PUT    = "PUT"
	HTTP_HEAD   = "HEAD"
	HTTP_PATCH  = "PATCH"
)

// Scheme type
const (
	HTTP  = "http"
	HTTPS = "https"
)

// Request content type
const (
	ApplicationJson            = "application/json"
	ApplicationOpenstackJson20 = "application/openstack-images-v2.0-json-patch"
	ApplicationOpenstackJson21 = "application/openstack-images-v2.1-json-json-patch"
)

// Authorization type
const (
	V4_AUTH = iota
	Token_AUTH
)

// Request header
const (
	SDK_DATE_HEADER             = "X-Sdk-Date"
	ACCEPT_HEADER               = "accept"
	ACCEPT_AUTHORIZATION_HEADER = "Authorization"
	CONTENT_LENGTH_HEADER       = "content-length"
	CONTENT_TYPE_HEADER         = "content-type"
	HOST_HEADER                 = "Host"
	SDK_HMAC_SHA256             = "SDK-HMAC-SHA256"
	SIGNATURE                   = "Signature"
	SDK_REQUEST                 = "sdk_request"
	SIGNED_HEADERS              = "SignedHeaders"
	CREDENTIAL                  = "Credential"
	SDK_AUTHORIZATION_HEADER    = "Authorization"
	SDK_X_AUTH_TOKEN            = "X-Auth-Token"
	X_PROJECT_ID                = "X-Project-Id"
)
