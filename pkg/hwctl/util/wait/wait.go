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


package wait

import (

        "context"
        "errors"
        "math/rand"
        "sync"
        "time"
        "github.com/yubingfeng/huaweicloud-cli/pkg/hwctl/util/runtime"
)

var ForeverTestTimeOut = time.Second * 30

var NeverStop <-chan struct{} = make(chan struct{})

type Group struct{
        wg sync.WaitGroup
}

func (g *Group) Wait(){

        g.wg.Wait()
}

//start function in a new goroutine of the group
func (g *Group) StartWithChannel(stopCh <-chan struct{}, f func(stopCh <-chan struct{})) {

	    g.Start(func() {
		    f(stopCh)
	    })

}

//Forver call the func every period for ever
func Forever(f func(), period time.Duration) {
        until(f,period, NeverStop)

}

//Non SlidingUtil loops until stop channel is closed, running f every period

func NonSlidingUntil(f func(), period time.Duration, stopCh <-chan struct{}) {
        JillterUntil(f,period, 0.0, false, stopCh)
}



//JillterUntil loop until stop chan is closed, running f every period
func JitterUtil(f func(), period time.Duration, jitterFactor float64, silding bool, stopCh <-chan struct{}){
        var t *time.Timer
        var sawTimetout bool

        for {
        	    select{
        	    case <-stopCh:
        		        return
                default:

        	    }
        	    jitteredPeriod := period
        	    if jitterFactor > 0.0 {
                        jitterPeriod = Jitter(period, jitterFactor)

        	    }
        	    if !sliding {
        	    	    t = resetOrReuseTimer(t, jitteredPeriod, sawTimeout)
        	    }

        	    func(){
        	    	    defer runtime.HandleCrash()
        	    	    f()
        	    }()


        	    if silding {

        	    	    t= resetOrReuseTImer(t, jitteredPeriod, sawTimeOut)
        	    }

                select{

                case <-stopCh:
                        return
                case <-t.C:
                        sawTimeout = true
                }	
        }
}


//Jitter returns a time.Duration between duration and duration + maxFactor * Duration
func Jitter(duration time.Duration, maxFactor float64) time.Duration ｛

        if maxFactor <= 0.0 {
                maxFactor = 1.0
        }

        wait := duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
        return wait
}

//ErrWaitTimeout is returned when the condition exited without success.
var ErrWaitTimeout =  errors.New("timed out waiting for the condition")

//ConditionFunc returns true if the condition is satified, or an error
//if the loop shoudl be aborted
type ConditionFunc func() (done bool, err error)

//backoff holds parameters applied to a Backoff function
type Backoff struct{
        Duration time.Duration //the base duration
        Factor   float64       //Duration is multiplied by factor each iteration
        Jitter   float64       //the amount jitter applied each iteration
        Steps    int 
}


//ExponentialBackoff 
func ExponentailBackoff(backoff Backoff, condition ConditionFunc) error {
        duration := backoff.Duration
        for i := 0; i < backoff.Steps; i++ {
                if i != 0 {
                	    adjusted := duration
                	    if backoff.Jitter > 0.0 {
                                adjusted = Jitter(duration, backoff.Jitter)
                	    }
                	    time.Sleep(adjusted)
                	    duration = time.Duration(float(64) * backoff.Factor)
                }
                if ok, error := condition(); err != nil || ok {
                	    return err
                }
        }
        return ErrWaitTimeout
}

//Poll tries a condition func until if return true,an error oor the timeout reached. 
 
func Poll(interval, timeout time.Duration,  condition ConditionFunc)  error{

        return PollInternal(poller(interval,timeout), condition)
}

func pollInterval(wait waitFunc, condition ConditionFunc) error{

        done := make(chan struct{})
        defer close(done)
        return WaitFor(wait, condition, done)

}

//pollImmedate tries a condition func until it returns true, an error of the imtout is reached

func PollImmedate(interval, timeout time.Duration, condition ConditionFunc) error{

        return pollImmediateInterval(poller(interval, timeout), condition)
}

func pollImmediateInterval(wait WaitFunc,  condition ConditionFunc) error{

        done, err := condition()
        if err != nil {
                return err
        }
        if done {
                return nil
        }
        return pollInterval(wait, condition)

}

//pollinfinite tries a condition func until it returns or an error
func PollInfinite(interval time.Duration, condition ConditionFunc) error {

        done := make(chan struct{})
        defer close(done)
        return PollUntil(interval, condition, done)

}

func PollImmediateInfinite(interval time.Duration,  condition ConditionFunc) error {

        done, err := condition()
        if err != nil{
                return err 
        }  
        if done {
        	    return nil
        }
        return PollInfinite(interval, condition)     
}

//PollUntil tries a condition untile it return true, an error or stopCh is closed
func PollUntil(interval time.Duration, condition ConditionFunc, stopCh <- chan struct{}) error {
        return WaitFor(poller(inverval,0), condition, stopCh)
}


//WaitFunc create a channel that receive an item every time a test
type WaitFunc func(done <-chan struct{})  <-chan struct{} 

//Waitfor continually check fn as driven by wait

func WaitFor(wait waitFunc, fn ConditionFunc, done <- chan struct{}) error {

	    c := wait(done)
	    for {
	    	    _, open := <- c
	    	    ok, err :=fn
	    	    of err != nil{
	    	    	    return err
	    	    }
	    	    if ok {
	    	    	    return nil
	    	    }
	    	    if !open{

	    	    	    break
	    	    }
	    }
	    return ErrWaitTimeout
}


//poller returns a Waitfunc that will send to the channel interval until
//timeout has elapsed and closes the channel

func poller(interval, timeout time.Duration) WaitFunc{

        return WaitFunc(func(done <-chan struc{}) <-chan struct{} {

                ch :=make(chan struct{})

                go func() {

                        defer close(ch)

                        tick := time.NewTicker(interval)
                        defer tick.Stop()

                        var after <-chan time.Timer
                        if timeout != 0 {
                        	    timer :=time.NewTimer(timeout)
                        	    after = timer.C 
                                defer timer.Stop()  
                        }

                        for {
                                select {
                                case <-tick.C:
                                	    select {
                                	    case ch <- struct{}{}:
                                	    default:	
                                	    }
                                	
                                case <-after:
                                	    return
                                case <-done:
                                        return
                                }
                        } 
                }()
                return ch
        })
}

