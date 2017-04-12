package dispatcher

import (
	"strconv"
)

// 消息处理单元
type Unit interface {
	SetId(id int)
	GetId() int
	GetIdStr() string
	AddFrame(f *Frame) bool
}

//
type BaseUnit struct {
	Id int // global id

	//
	Frames chan *Frame // servers' msg(frame)
	CurF   *Frame      // 当前正在处理的服务端发送过来的消息
}

func NewBaseUnit(frameNum int) *BaseUnit {
	return &BaseUnit{Frames: make(chan *Frame, frameNum)}
}

func (this *BaseUnit) SetId(id int) {
	this.Id = id
}

func (this *BaseUnit) GetId() int {
	return this.Id
}

func (this *BaseUnit) GetIdStr() string {
	return strconv.Itoa(this.Id)
}

func (this *BaseUnit) AddFrame(f *Frame) bool {
	select {
	case this.Frames <- f:
		return true
	default:
	}

	return false
}
