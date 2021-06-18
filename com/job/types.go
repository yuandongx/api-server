package job

type Job struct {
	Id          string
	SshUser     string
	SshPassword string
	SshPort     int
	PingLines   []Ping
}

type Ping struct {
	Id          string
	LineNo      string
	Destination string
	Result
}

type Result struct {
	Id        string
	StartTime string
	EndTime   string
	Count     int
	Received  int
	Loss      int
	Time      string
	Rtt       string
}

type Trigger struct {
	Id       string
	Interval string //任务间隔
	Ttype    int    //触发类型
}

type JobQeue struct {
	Jobs []Job
}
