package configmodel

type ConfigSettings struct {
	MQSettings MQSettings `json:"mq"`
}

//MQSettings structure
type MQSettings struct {
	Link string `json:"link"`
}
