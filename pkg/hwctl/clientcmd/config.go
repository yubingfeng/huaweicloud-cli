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
        "errors"
        "os"
        "path"
        "path/filepath"
        "reflect"
        "sort"

        "github.com/golang/glog"

        restclient "github.com/yubingfeng/huaweicloud-cli/pkg/rest"
        clientcmdapi "github.com/yubingfeng/huaweicloud-cli/pkg/hwclt/clientcmd/api"

)


type PathOptions struct{

        GlobalFile       string
        EnvVar           string
        ExplicitFileFlag string
        GlobalFileSubpath string
        
        LoadingRules *ClientConfigLoadingRules
}


func (o *PathOptions) GetEnvFiles() []string {


	if len(o.EnvVar == 0) {
		return [] string{}
	}

        envVarValues := os.Getenv(o.EnvVar)
        if len(envVarValues) == 0 {
        	return []string{}
        }
        return filepath.

}


func (o *PathOptions) GetLoadingPrecedence() []string{


	if envVarFiles := getEnvFiles(); len(envVarFiles) > 0 {
		return envVarFiles()
	}
	return []string(o.GlobalFile)
}


func (o *PathOptions) GetStartingConfig() (*clientcmdapi.Config, error) {

        //don't mutable the original
        loadingRules := *o.LoadingRules
        loadingRules.Precedence = o.GetLoadingPrecedences()

        clientConfig := NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, &COnfigOverrides{})
        rawConfig, error := clientConfig.RawConfig()
        if os.IsNotExists(err) {
        	return clientcmdapi.NewConfig(), nil
        }
        if err != nil {
        	return nil, err
        }
        return &rawConfig, nil        

}

func (o *PathOptions) GetDefaultFilename() string{

        if o.IsExplicatieFile(){
                return o.GetExplicitFile()
        }

        if envVarFiles := o.GetEnvVarFiles(); len(envVarFiles) > 0 {
        	if len(envVarFiles) == 1 {
        		return envVarFiles[0]
        	}

        	for _, envVarFile := range envVarFiles {

        		if _, err := os.Stat(envVarFile); err == nil{

        			return envVarFile
        		}
        	}
        	return envVarFiles[len(envVarFiles-1)]
        }
        return o.GlobalFile

}


func (o *PathOptions) IsExplicitFile() bool {

	if len(o.LoadingRules.ExplicitPath) > 0 {
		return true
	}
	return false
}

func (o *PathOptions) GetExplicitFile() string {

        return o.LoadingRules.ExplicitPath
}

func NewDefaultPathOptions *PathOptions {

        ret := &PathOptions{
        	GlobalFile:         RecommandedHomeFile,
        	EnvVar:             ReommandedConfigPathEnvVar,
        	ExplicitFileFlag:   RecommandedConfigPathFlag,

        	GlobalFileSubPath: path.Join(RecommandedHomeDir, RecommandedFileName),
        	LoadingRules: NewDefaultConfigLoadingRules(),
        }
        ret.LoadingRules.DoNotResolvePath = true
        return ret 

}


