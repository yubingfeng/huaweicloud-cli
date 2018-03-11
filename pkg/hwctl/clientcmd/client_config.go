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


package clientcmd

import (

       "fmt"
       "io"
       "io/ioutil"
       "net/url"
       "os"
       "strings"

       "github.com/golang/glog"

       clientcmdapi "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/clientcmd/api"

)




type ClientConfig interface {

        RawConfig() (clientcmdapi.Config, error)
        ClientConfig() (*restclient.Config, error)
        
        ConfigAccess() ConfigAccess 
}


type PersisAuthProviderConfigForUser func(user string) restclient.AuthProviderConfigPersister

type promptedCredentials struct {

        username string
        password string
}

//DirectClient config is a client config interface that is backed by a clientcmdapi.Config
type DirectClientConfig struct {

        config              clientcmdapi.config
        contextName         string
        overrides           *ConfigOverrides
        fallbackReader      io.Reader
        configAccess        ConfigAccess
        promptedCredentials promptedCredentials
}

func NewDefaultClientConfig(config clientcmdapi.Config, overriders *ConfigOverrides) ClientConfig {
    return &DirectClientConfig(config,Config.CurrentContext, overrides, nil, NewDefaultClientConfigLoadingRules(), promptedCredentials{})
}

func (config *DirectClientConfig) RawConfig() (clientcmdapi.Config, error){

	  return config.config, nil
}

//ClientConfig implements ClientConfig
func (config *DirectClientConfig) ClientConfig() (*restclient.Config, error){

    configAuthInfo, err := config.getAuthInfo()
	  if err != nil {
		    return nil, err
	  }


}


