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


package cmd

import ( 

        "fmt"
        "os"

        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/flag"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/auth"
        cmdconfig "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/config"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/resource"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/rollout"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/set"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/templates"
        cmdutil "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/util"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/i18n"

        "github.com/golang/glog"
        "github.com/spf13/cobra"

)

func NewDefaultHWctlCommand() *cobra.Command{
	return NewHWctlCommand(cmdutil.NewFactory(nil), os.Stdin, os.Stdout, os.Stderr)
}

func NewHWctlCommand(f cmdutil.Factory, in io.Reader, out, err io.Writer) *cobra.Command {
        
        //Parent command to which all subcommands are added
        cmds := &cobra.Command{
                Use: "hwctl",
                Short: i18n.T("hwctl controls the huawei public cloud resources and services"),
                Long: templates.LongDesc(`
        hwctl controls the huawei public cloud resources and services,

        Find more informations at:
               https://wwww.huaweicloud.com`),
                Run: runHelp,
        }
        
        f.BindFlags(cmds.PresistentFlags())
        f.BindExternalFlags(cmds.PersistentFlags())

        //Send nil for the getLanguageFn() results in using 
        //the LANG environment variable
        i18n.LoadTranslation("hwctl",nil)

        //warning the command contains "_" seperators
        cmds.SetGlobalNormalizatoinFunc(flag.WarnWordSepNormalizerFunc)

        groups := templates.CommandGroups{
                {
                        Message: "Basic Commands:",
                        Commands: []*cobra.Command{
                                NewCmdCreate(f, out, err),
                                NewCmdGet(f, out, err),
                                NewCmdDelete(f, out, err),

                        },

                }
        }
        groups.Add(cmds)

        filters := []string{"options"}

        templates.ActsAsRootCommands(cmds, filters, groups...)

        cmds.AddCommand(cmdconfig.NewCmdConfig(f, clientcmd.NewDefaultPathOptions(),out, err))
        cmds.AddCommand(NewCmdPlugin(f, in, out, err))
        cmds.AddCommand(NewCmdVersion(f,out))
        cmds.AddCommand(NewCmdOptions(f,out))


        return cmds

)


func runHelp(cmd *corba.Command, args []string){

	cmd.Help()
}






