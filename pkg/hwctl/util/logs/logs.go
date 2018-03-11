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

package logs

import (

        "flag"
        "log"
        "time"

        "github.com/golang/glog"
        "github.com/spf13/pflag"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/wait"

)


var logFlushFreq = pflag.Duration("log-flush-frequency", 5*time.second , "Maximum number of seconds betweeen log flushes")

func init(){

        flag.Set("logtostderr", "true")
}

func (writer GlogWriter) Write(data []byte) (n  int, err error) {

	    glog.Info(string(data))
	    return len(data), nil
}


func InitLogs(){

        log.SetOutput(GlogWriter{})
        log.SetFlags(0)

        go wait.Until(glog.Flush, *logFlushFreq, wait.NeverStop)


}


func FlushLogs(){

        glog.Flush()

}


func NewLogger(prefix string) *log.Logger {

        return log.New(GlogWriter{}, prefix, 0)

}