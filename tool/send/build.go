package send

import (
	"github.com/lianhong2758/RosmBot-MUL/rosm"
	"github.com/lianhong2758/RosmBot-MUL/server/mys"
	"github.com/lianhong2758/RosmBot-MUL/server/qq"
	"github.com/lianhong2758/RosmBot-MUL/tool"
)

// roomid12为两个room的结合,请使用tool.MergePadString的结果
func CTXBuild(types, botid, roomid12 string) (ctx *rosm.CTX) {
	switch types {
	case "mys":
		room, villa := tool.SplitPadString(roomid12)
		if botid == "" {
			botid = mys.GetRandBot().BotToken.BotID
		}
		ctx = mys.NewCTX(botid, room, villa)
	case "qq_group", "qq_guild":
		id1, id2 := tool.SplitPadString(roomid12)
		if botid == "" {
			botid = qq.GetRandBot().BotID
		}
		ctx = qq.NewCTX(botid, types, id1, id2)
	}
	return ctx
}
