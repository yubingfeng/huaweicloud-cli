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

package app

import (
        "os"
        
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd"
        cmdutil "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/cmd/util"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/logs"

)

func Run() error {
        logs.InitLogs()
        defer logs.FlushLogs()

        cmd := cmd.NewHWctlCommand(cmdutil.NewFactor(nil), os.Stdin, os.Stdout, os.Stderr)
        return cmd.Execute()


}