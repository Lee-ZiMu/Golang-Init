/**
 * @Description: 日志
 * @Author Lee
 * @Date 2023/12/13 13:18
 **/

package logger

import (
	"github.com/robfig/cron/v3"
	viperInit "github.com/spf13/viper"
	"hy_heymate/common/file"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init() {
	once.Do(func() {
		// 初始化 core
		core := zapcore.NewTee(
			zapcore.NewCore(getEncoder(), getLogWriter("info.log"), zapcore.DebugLevel),
			zapcore.NewCore(getEncoder(), getLogWriter("error.log"), zapcore.ErrorLevel),
		)

		// 初始化 Logger
		logger = zap.New(core,
			zap.AddCaller(),      // 调用文件和行号，内部使用 runtime.Caller
			zap.AddCallerSkip(1), // 封装了一层，调用文件去除一层(runtime.Caller(1))
			// zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
		)

		// 将自定义的 logger 替换为全局的 logger
		// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
		zap.ReplaceGlobals(logger)
	})
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",   // 时间对应的key名
		LevelKey:       "level",  // 日志级别对应的key名
		NameKey:        "logger", // logger名对应的key名
		CallerKey:      "caller", // 调用者对应的key名
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",                      // 日志内容对应的key名，此参数必须不为空，否则日志主体不处理
		StacktraceKey:  "stacktrace",                   // 栈追踪的key名
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志编码级别，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	// 获取配置文件
	logDir, err := filepath.Abs(viperInit.GetString("log.path"))
	if err != nil {
		panic(err)
	}

	// 创建日志文件夹
	if err := file.MkdirAll(logDir); err != nil {
		panic(err)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, filename),
		MaxAge:     viperInit.GetInt("log.rotateDays"),
		MaxSize:    viperInit.GetInt("log.rotateSize"),
		MaxBackups: viperInit.GetInt("log.backups"),
		LocalTime:  true,
	}

	c := cron.New()
	c.AddFunc("0 0 * * ?", func() { lumberJackLogger.Rotate() })
	// c.AddFunc("0/1 * * * ?", func() { lumberJackLogger.Rotate() })
	c.Start()

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateTime))
}

func Fatal(msg string, err error) {
	logger.Fatal(msg, zap.Error(err))
}

func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func ErrorE(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	zap.S().Debugf(template, args...)
}
