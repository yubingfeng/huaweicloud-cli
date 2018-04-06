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


package config 

import (
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/clientcmd"
        cmdutil "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/util"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwclt/util/i18n"
        
)

//new cmd config

func NewCmdConfig(f cmdutil.Factory, pathOptions *clientcmd.PathOptions, errOut io.Writer) *cobra.Command{

        if len(pathOptions.ExplicitFileFlag) == 0 {
        	pathOtions.ExplicitFileFlag = clientcmd.RecommendedConfigPathFlag
        }

        cmd := &cobra.Command {
        	Use: "config SUBCOMMAND",
        	Short: i18n.T("Modify hwctl config files")
        	Long: templates.LongDesc(`
        		Modify hwctl config files using subcommands like "hwctl config set"
        		`)
        	Run: cmdutil.DefaultSubCommandRun(errOut),

        }

        return cmd

}