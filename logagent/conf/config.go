package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	//TaillogConf `ini:"taillog"`
	EtcdConf `ini:"etcd"`

	ESConf `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	ChanMax int    `ini:"chan_max"`
	Topic   string `ini:"topic"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

// -- unused ---
type TaillogConf struct {
	FileName string `ini:"filename"`
}

type ESConf struct {
	Address      string `ini:"address"`
	ChanSize     int    `ini:"chansize"`
	NumsGorutine int    `ini:"numsGorutine"`
}
