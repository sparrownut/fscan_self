package InfoExpand

import (
	"github.com/shadow1ng/fscan/utils"
	"github.com/zhshch2002/goreq"
)

type Hostinfo struct {
	Domain   string
	Ip       string
	Csegment string
}

func domainexpand(host Hostinfo) []Hostinfo {
	Httpclient := goreq.NewClient(goreq.WithRandomUA())
	//通过域名扩张资产
	var hostlist []Hostinfo
	fofaSearchApi := "https://fofa.info/api/v1/search/all?email=om2bg0tjsptuh7weetboc2t7vvzk@open_wechat&key=fff73dae777ddbcf971bb5e85f2ae3aa&size=10000&qbase64="
	//inp = base64.b64encode(bytes(('domain=' + domain).encode())).decode()
	req := goreq.Get(fofaSearchApi + "").SetClient(Httpclient) //base64 req
	requestText := goreq.Do(req)
	utils.Printsuc(requestText.Text)
	return hostlist
}
