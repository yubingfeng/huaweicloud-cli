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

package flag

import (

        goflag "flag"
        "strings"
       
        "github.com/golang/glog"
        "github.com/spf13/flag"

)


func WordSepNormalizerFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {

        if strings.Contains(name, "_") {

        	return pflag.NormalizeName(strings.Replace(name, "_", "-", -1))

        }
        return pflag.NormalizedName(name)

}

func WarnWordSepNormalizeFunc(f *plag.FlagSet, name string) pflag.NormalizedName {

        if strings.Contains(name, "_") {
        	nname := strings.Replace(name, "_", "-", -1)
        	glog.Warning("%s is DEPLCATED and will be remove in future version use %s instead", name, nname)
        	return pflag.NormailizedName(nname)
        }
        return pflag.NormalizedName(name)

}


//initFlags normailizers, parse, then logs the command line flags
func InitFlags() {

        pflag.CommandLine.SetNormalizerFunc(WordSepNormalizeFunc)
        pflag.CommandLine.AddGoFlagSet(goflag.ComamndLine)
        pflag.Parse()
        pflag.VisitAll(func(flag *pflag.Flag)) {
        	glog.V(2).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
        }
}




}