/*
Copyright 2018  yu.yue@huawei.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api 

import (

        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/runtime"

)


//Config holds the information needed to build connect to huawei cloud
type Config struct{

        Preferences    Preferences   `json:"preferences"`
        AuthInfos      map[string]*AuthInfo `json:"Users"`
        Contexts       map[string]*Context  `json:"context"`
        CurrentContext String `json:"current-context"`
        Extensions     map[string]runtime.Object `json:"extensions,omitempty"`

}

//this 
type Preferences struct{

	Extensions map[string]runtime.Object `json:"extension,omitempty"`
}


//authinfo contains the information that describes identity information
type AuthInfo struct {

	Token      string `json:"Token,omitempty"`
	AK         string `json:"AK,omitempty"`
	SK         string `json:"SK,omitempty"`
	Extensions map[string]runtime.Object `json:"extensions,omitempty"`

}


//context is a tuple of references to the hwcloud
type Context struct{

        AuthInfo string `json:"user"`
        Extensions map[string]runtime.Object 'json:"extensions,omitempty"'

}


func NewConfig() *Config {

	return &Config{
		Preferences: *NewPreferences(),
		AuthInfos:   make(map[string]*AuthInfo),
		Contexts:    make(map[string]*Context),
		Extensions:  make(map[string]runtime.Object),
	}
}

//create an empty context
func NewContext() *Context {
	return &Context{Extensions: make(map[string]runtime.Object)}
}


//New Authinfo is a convenience function that returns a new AuthInfo
func NewAuthInfo() *AuthInfo{

	return &AuthInfo{
		Token: "",
		AK: "",
		SK: "",
		Extensions: make(map[string]runtime.Object)
	}
}

//create new preferences object
func NewPreferences() *Preferences{
	return &Preferences{
		Extensions: make(map[string]runtime.Object)
	}
}
