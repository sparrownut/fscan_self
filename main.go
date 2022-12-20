package main

import (
	"github.com/shadow1ng/fscan/Plugins"
	"github.com/shadow1ng/fscan/common"
	"github.com/shadow1ng/fscan/utils"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	//fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
	utils.Printminfo("扫描结束,耗时: %s\n", t)
}
