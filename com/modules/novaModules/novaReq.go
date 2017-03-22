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
package novaModules

import (
	"github.com/huawei/DockerMachineDriver4OTC/com/modules"
)

/* The request Params need by createKeypair Start */

// Struct KeypairCreateReq
type CreateKeypairReq struct {
	modules.BaseDataStruct
}

func (createkeypairReq *CreateKeypairReq) Init(keypair KeypairCreate) {
	createkeypairReq.InitBase()
	createkeypairReq.MapBodyContent["keypair"] = keypair.MapBodyContent
}

// Struct KeypairCreate
type KeypairCreate struct {
	modules.BaseDataStruct
}

func (keypairCreate *KeypairCreate) Init(name string) {
	keypairCreate.InitBase()
	keypairCreate.MapBodyContent["name"] = name
}

func (keypairCreate *KeypairCreate) SetPublic_key(public_key string) {
	keypairCreate.MapBodyContent["public_key"] = public_key
}

/* The request Params need by createKeypair End */
