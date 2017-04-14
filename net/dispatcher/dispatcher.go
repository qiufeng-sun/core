// 消息分发器
package dispatcher

import (
	"strconv"
	"sync"

	"util/logs"

	"core/net/dispatcher/pb"
)

var _ = logs.Debug

// 消息分发器
type Dispatcher struct {
	Name  string // 标识
	SrvId string

	Id    int             // global id
	Units map[string]Unit // global id => unit
	sync.Mutex
}

func New(name string, srvId string) *Dispatcher {
	return &Dispatcher{
		Name:  name,
		SrvId: srvId,
		Units: map[string]Unit{},
	}
}

//
func (this *Dispatcher) Register(u Unit) {
	this.Lock()
	defer this.Unlock()

	this.Id++
	url := Url(this.SrvId, this.Id)
	u.Set(this.Id, url)
	this.Units[strconv.Itoa(this.Id)] = u
}

//
func (this *Dispatcher) Unregister(u Unit) {
	this.Lock()
	defer this.Unlock()

	delete(this.Units, u.GetIdStr())
}

//
func (this *Dispatcher) Dispatch(f *pb.PbFrame) {
	for _, url := range f.DstUrls {
		//srv, addr, chk, id, ok := lan.Url2Part(url)
		_, id, ok := Url2Part(url)
		if !ok {
			logs.Warn("invalid dst url in frame! %v:%v", this.Name, url)
			continue
		}

		unit, ok := this.Units[id]
		if !ok {
			// 已经不在线了
			logs.Info("not found unit! maybe offline. %v:%v", this.Name, url)
			continue
		}

		nf := &Frame{PbFrame: f, DstUrl: url}
		unit.AddFrame(nf)
	}
}
