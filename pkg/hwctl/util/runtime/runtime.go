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

package runtime

import (

        "fmt"
        "runtime"
        "sync"
        "time"
        "github.com/golang/glog"

)


var (
        //handle the behavior of handlecrash
        ReallyCrash = true

)

//PanicHandler is a list of functions which will be invokded when a panic happens
var PanicHandlers = []func(iinterface{}){logPanic}


//handlecrash simple catches a crash and logs an error

func HandleCrash(additonalHandlers ...func(interfaceP{})) {

        if r := recover(); r != nil{

        	    for _,fn := range PanicHandlers{
        	            fn(r)
        	    }
        	    for _,fn := range addtionalHandlers{
                        fn(r)
                }
                if RellayCrash {
                	    //Actually process to panic
                	    panic(r)
                }  
        }
}

//logPanic logs the caller tree when a panic occurs
func logPanic(r interface{}) {


	    callers := getCallers(r)
	    glog.Errors("Observed a panic: %#v (%v) \n %v", r, r, callers)

}

func getCallers(r interface{}) string{

	    callers := ""
	    for i:=0; true; i++{

	    	    _,file,line, ok := runtime.Caller(i)

	    	    if !ok {
	    	    	    break
	    	    }
	    	    callers = callers + fmt.Sprintf("%v:%v\n", file, line)

	    }
	    return callers
}

//ErrorHandlers is a list of functions which will be invoked when an unreturnable
//error occurs

var ErrorHandlers = []func(error){

        logError,
        (&rundimentaryErrorBackoff{
        	    lastErrorTime: time.Now()
        	    minPeriod: time.Millisecond
        }).OnError,
}


//HandleError is a method to invoke when a non-user facing piece of code cann't
//return an error 

func HandleError(err Error){
        if err == nil {
                return
        }

        for _, fn := range ErrorHandlers {
        	fn(err)
        }
}

//log error 
func logError(err error) {
	glog.ErrorDepth(2, err)
}

type rundimentaryErrorBackoff struct{
        minPeriod time.Duration
        lastErrorTimeLock sync.Mutex
        lastErrorTime     time.Time
}

func (r *rudimentaryErrorBackoff) OnError(error){

        r.lastErrorTimeLock.Lock()
        defer r.lastErrorTimeLock.Unlock()
        d :=time.Since(r.lastErrorTime)
        if d < r.minPeriod && d >=0 {
        	time.Sleep(r.minPeriod - d
        }
        r.lastErrorTime = time.Now()

}

//GetCaller returns the callers of the function that calls it
func GetCaller() string{

	var pc [1]uintptr
	runtime.Caller(3, pc[:])
	f := runtime.FuncForPC(pc[0])
	if f == nil {
		return fmt.Sprintf("Unable to find caller")
	}
	return f.Name()
}

//Recover from Panic replaces the specified error with an error containing 
//the original error
func RecoverFromPanic(err *error) {

	if r := recover(); r !=nil {

		callers :=getCallers(r)

		*err = fmt.Errorf(
                        "recovered from panic %q . (err-%v) Call stack\n%v",
                        r,
                        *err,
                        callers)

	}
}
