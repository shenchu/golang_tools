package task

import "sync/atomic"

// 任务优雅退出工具

type TaskManage interface {
	Add()
	Done()
	CanExit() bool
	Exiting() bool
	Exit() bool
	Num() int32
}

type DefaultTaskManage struct {
	count int32
}

func NewDefaultTaskManager() DefaultTaskManage {
	return DefaultTaskManage{}
}

func (this *DefaultTaskManage) Add() {
	atomic.AddInt32(&this.count, 1)
}

func (this *DefaultTaskManage) Done() {
	atomic.AddInt32(&this.count, -1)
}

func (this *DefaultTaskManage) CanExit() bool {
	return (atomic.LoadInt32(&this.count) & 0x3FFFFFFF) == 0
}

func (this *DefaultTaskManage) Exiting() bool {
	return (atomic.LoadInt32(&this.count) & 0x40000000) == 0x40000000
}

func (this *DefaultTaskManage) Exit() {
	for {
		oldValue := atomic.LoadInt32(&this.count)
		newValue := oldValue | 0x40000000
		if atomic.CompareAndSwapInt32(&this.count, oldValue, newValue) {
			break
		}
	}
}

func (this *DefaultTaskManage) Num() int32 {
	return atomic.LoadInt32(&this.count) & 0x3FFFFFFF
}

