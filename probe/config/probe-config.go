package config

const Version = "v1.0.1"

//为了系统稳定进行脱敏处理
var WsUrl = "xxxx"
var AgentUrl = "xxxx"
var ApiUrl = "xxxx"

type Config struct {
	ServerId string  `json:"serverId"`
	Warning  Warning `json:"warning"`
	Monitor  Monitor `json:"monitor"`
}

type Warning struct {
	Switch int         `json:"switch"`
	Rule   WarningRule `json:"rule"`
}

type WarningRule struct {
	Cpu WarningRuleCpu `json:"cpu"`
	Mem WarningRuleMem `json:"mem"`
}

type WarningRuleCpu struct {
	Switch   int `json:"switch"`
	Value    int `json:"value"`
	Duration int `json:"duration"`
	Count    int `json:"count"`
}

type WarningRuleMem struct {
	Switch   int `json:"switch"`
	Value    int `json:"value"`
	Duration int `json:"duration"`
	Count    int `json:"count"`
}

type Monitor struct {
	Rule MonitorRule `json:"rule"`
}

type MonitorRule struct {
	Process []string `json:"process"`
}
