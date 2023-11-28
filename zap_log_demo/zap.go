package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// basic
var log *zap.Logger

// level
var log1 *zap.SugaredLogger

// leve2
var log2 *zap.SugaredLogger

//添加日志切割  go get -u github.com/natefinch/lumberjack

func InitLoger() {
	log, _ = zap.NewProduction()
	//level
	log1 = log.Sugar()
	//level2
	//_, _ = os.Create("./lumber.log")
	//添加日志切割
	lumber := &lumberjack.Logger{
		Filename:   "./lumber.log",
		MaxSize:    1,     //文件最大大小 以M为单位
		MaxBackups: 1,     //所保留的文件个数
		MaxAge:     30,    //保留天数
		Compress:   false, //是否压缩分割的文件
	}
	write := zapcore.AddSync(lumber)
	config := zap.NewProductionEncoderConfig()
	// config.EncodeTime = zapcore.ISO8601TimeEncoder //默认时间格式
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") //自定义时间格式
	// config.EncodeLevel = zapcore.CapitalColorLevelEncoder //大写
	encode := zapcore.NewJSONEncoder(config)
	core := zapcore.NewCore(encode, write, zapcore.DebugLevel)
	coreLog := zap.New(core)
	// coreLog := zap.New(core, zap.AddCaller())//添加调用的函数
	log2 = coreLog.Sugar()
}
func main() {
	InitLoger()
	defer log.Sync()
	defer log1.Sync()
	defer log2.Sync()
	httpGet("http://www.baidu.com")
	httpGet("http://www.google.com")
}
func httpGet(url string) {
	log1.Debug("sug:tring fetch url:", url)
	log2.Debug("sug:tring fetch url:", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Error(
			"fetch url err", zap.String("url", url), zap.Error(err),
		)
		//level
		log1.Error("sug:error fetch ing")
		//level2
		log2.Error("sug:error fetch ing")
	} else {
		log.Info("fetch rul success", zap.Int("statuscode", resp.StatusCode), zap.String("status", resp.Status), zap.String("url", url))
		//level
		log1.Info("sug:fetch url succes")
		//level2
		log2.Info("sug:fetch url succes")
	}
	defer resp.Body.Close()

}
