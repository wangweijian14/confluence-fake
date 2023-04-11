package importserv

import (
	"fmt"
	"runtime"
)

type jobResult struct {
	ret interface{}
	err error
}

type job func() (interface{}, error)

type concurrentHelper struct {
	concurrentLimit int
	jobs            []job
}

func NewConcurrentHelper(concurrentLimit int) *concurrentHelper {
	return &concurrentHelper{
		concurrentLimit: concurrentLimit,
		jobs:            []job{},
	}
}

func (c *concurrentHelper) Add(j job) {
	c.jobs = append(c.jobs, j)
}

// 返回两个列表，长度和 Jobs 相同，顺序与 Job 添加顺序相同
// 如果 job 返回错误，则 results 列表中对应的值为 nil， errors 列表中的值为对应的错误
func (c *concurrentHelper) Run() (results []interface{}, errors []error) {
	count := len(c.jobs)
	//
	limiterChan := make(chan bool, c.concurrentLimit)
	jobResultsChan := make([]chan *jobResult, count)
	defer func() {
		close(limiterChan)
		for i := 0; i < count; i++ {
			close(jobResultsChan[i])
		}
	}()
	//
	for i := 0; i < count; i++ {
		// 初始化一个对应位置的缓存长度为1的结果 channel
		jobResultsChan[i] = make(chan *jobResult, 1)
		go func(j job, jobResultChan chan *jobResult, limiterChan chan bool) {
			defer func() {
				if r := recover(); r != nil {
					buf := make([]byte, 1<<18)
					n := runtime.Stack(buf, false)
					fmt.Println("%v, STACK: %s", r, buf[0:n])
					// 出现 panic 恢复时
					// 将结果写入 channel，避免最后读取时阻塞
					jobResultChan <- &jobResult{ret: nil, err: fmt.Errorf("panic")}
					// 释放掉占用
					<-limiterChan
				}
			}()
			//
			limiterChan <- true
			ret, err := j()
			jobResultChan <- &jobResult{ret: ret, err: err}
			<-limiterChan
		}(c.jobs[i], jobResultsChan[i], limiterChan)
	}
	// 遍历 channel 读取结果
	for i := 0; i < count; i++ {
		jobResult := <-jobResultsChan[i]
		results = append(results, jobResult.ret)
		errors = append(errors, jobResult.err)
	}
	//
	return results, errors
}
