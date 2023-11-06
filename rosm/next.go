package rosm

import (
	"time"

	"github.com/sirupsen/logrus"
)

var nextList = map[EventType]map[int]*Matcher{}

// 获取下一事件
func (ctx *CTX) GetNext(types EventType, SetBlock bool, rs ...Rule) (chan *CTX, func()) {
	next := make(chan *CTX, 1)
	ids := int(0xfffffff & time.Now().Unix())
	m := &Matcher{block: SetBlock, rules: rs, nestchan: next}
	if nextList[types] != nil {
		nextList[types][ids] = m
	} else {
		nextList[types] = map[int]*Matcher{ids: m}
	}
	return next, func() {
		close(next)
		delete(nextList[types], ids)
	}
}

func (ctx *CTX) sendNext(types EventType) (block bool) {
	if len(nextList) == 0 || nextList[types] == nil {
		return false
	}
	logrus.Debug("[next]匹配事件type", types)
	for _, v := range nextList[types] {
		if v.RulePass(ctx) {
			v.nestchan <- ctx
			if v.block {
				return true
			}
		}
	}
	return false
}
