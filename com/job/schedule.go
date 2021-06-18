package job

// Update 更新Job列表
func (jqeue *JobQeue) Update() {}

// InitJob 初始化job
func InitJob() (jqeue *JobQeue) {
	jqeue = &JobQeue{}
	jqeue.Update()
	return jqeue
}

// ProcessJob 分发任务
func ProcessJob(job Job, result chan Result) {

}

// ProcessResult 结果处理
func ProcessResult(result chan Result) {}

// Monitor 循环遍历job list，将 将要执行的任务加入队列，准备执行
func (jq *JobQeue) Monitor(job chan Job) {}

func ShouldUpdateJobQeue() {
	shouldUpdate <- 1
}

// Schedule 总体调试
func Schedule() {
	jq := InitJob()
	jobsChannel := make(chan Job, 10)
	resultChannel := make(chan Result, 3)
	go jq.Monitor(jobsChannel)
	for {
		select {
		case <-shouldUpdate:
			jq.Update()
		case job := <-jobsChannel:
			go ProcessJob(job, resultChannel)
			go ProcessResult(resultChannel)
		}
	}
}
