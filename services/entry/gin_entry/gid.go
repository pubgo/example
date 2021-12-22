package gin_entry

import (
	"github.com/pubgo/example/services/entry/gin_entry/handler"
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/entry/ginEntry"
)

func GetEntry() entry.Entry {
	var ent = ginEntry.New("gid1")
	ent.Description("gid1 generate")
	ent.Register(handler.NewId())
	return ent
}
