package logging

import (
	"time"

	"github.com/pure-project/purelog"
)

var Logger *purelog.Logger
var Config *purelog.Config

func init() {
	Config = purelog.NewConfig().
		SetStdout(true).
		SetCaller(true).
		SetFlush(100 * time.Millisecond)
	Config.SetLevel(purelog.LevelInfo)
	Logger = purelog.New(Config)
}
