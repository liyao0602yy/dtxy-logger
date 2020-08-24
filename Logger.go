package Logger

import (
	"fmt"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
	"os"
	"time"
)


var Log * glog.Logger

func Init() *glog.Logger {
	var LoggerConfig map[string]interface{}
 	if _, v := gutil.MapPossibleItemByKey(gcfg.Instance().GetMap("."), "Logger"); v != nil {
		LoggerConfig = gconv.Map(v)
	} else {

	}
	if Log == nil {
		Log = glog.New()
		logPath := LoggerConfig["path"].(string)
		Log.SetPath(logPath)

		Log.SetLevelStr(LoggerConfig["loglevel"].(string))
		enableClear, succeed := LoggerConfig["enableClear"].(bool)
		if succeed && enableClear{
			_, err := gcron.Add("0 0 1 1 * ?", func() {
				nowTime :=  time.Now()
				clearMonth := (nowTime.Month() - 2 + 12) % 12
				isNotJiewei := int((nowTime.Month() - 2 + 12) / 12)
				yearOffset := 0
				if isNotJiewei == 0 {
					yearOffset = -1
				} else {
				}
				clearYear := (nowTime.Year() + yearOffset)

				monthFileName := fmt.Sprintf("%04d-%02d", clearYear, clearMonth)

				for i := 1; i <= 31; i++ {
					perDayFile := logPath + fmt.Sprintf("%s-%02d.log", monthFileName, i)
					_, err := os.Stat(perDayFile)    //os.Stat获取文件信息

					if err == nil {
						_ = os.Remove(perDayFile)
					} else {
						if os.IsExist(err) {
							os.Remove(logPath + "/" + perDayFile)
						}
					}
				}
			}, "clearFileTask")
			if err != nil {
				Log.Error(err)
			}
		}

		/*
		loggerIns.SetConfigWithMap(g.Map{
			"file":		"temp",
			"path":     "./log",
			"level":    "all",
			"stdout":   false,
			"StStatus": 0,
		})*/
	}



	return Log
}
