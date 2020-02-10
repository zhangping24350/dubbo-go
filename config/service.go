/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"github.com/apache/dubbo-go/common"
)

var (
	conServices = map[string]common.RPCService{} // service name -> service
	proServices = map[string]common.RPCService{} // service name -> service
)

// SetConsumerService is called by init() of implement of RPCService
func SetConsumerService(service common.RPCService) {
	conServices[service.Reference()] = service
}

// SetProviderService is called by init() of implement of RPCService
func SetProviderService(service common.RPCService) {
	proServices[service.Reference()] = service
}

// GetConsumerService ...
func GetConsumerService(name string) common.RPCService {
	return conServices[name]
}

// GetProviderService ...
func GetProviderService(name string) common.RPCService {
	return proServices[name]
}

// GetCallback ...
func GetCallback(name string) func(response common.CallbackResponse) {
	service := GetConsumerService(name)
	if sv, ok := service.(common.AsyncCallbackService); ok {
		return sv.CallBack
	}
	return nil
}
