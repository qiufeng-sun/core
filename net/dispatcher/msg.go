package dispatcher

import (
	"strings"

	"util/logs"

	"core/net/dispatcher/pb"
)

var _ = logs.Debug

// temp def // to do
type Frame struct {
	*pb.PbFrame
	DstUrl string
}

//// @return srv, addr, chk, id // to do
//func (this *Frame) Url2Part() (string, string, string, string) {
//	ss := strings.Split(this.Url, "|")
//	srv, addr, chk, id := ss[0], ss[1], ss[2], ss[3]
//	return srv, addr, chk, id
//}

// @return srv, addr, chk, id, ok
func Url2Part(url string) (srv, addr, chk, id string, ok bool) {
	ss := strings.Split(url, "|")
	if len(ss) != 4 {
		ok = false
		return
	}
	srv, addr, chk, id, ok = ss[0], ss[1], ss[2], ss[3], true
	return
}
