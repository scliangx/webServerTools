package logger

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const DefaultLogPath = "/var/agent/" // 默认输出日志文件路径
const DefaultLogName = "default.log"

type LogConfigs struct {
	LogLevel          string // 日志打印级别 debug  info  warning  error
	LogFormat         string // 输出日志格式	logfmt, json
	LogPath           string // 输出日志文件路径
	LogFileName       string // 输出日志文件名称
	LogFileMaxSize    int    // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    // 日志保留时间，单位: 天 (day)
	LogCompress       bool   // 是否压缩日志
	LogStdout         bool   // 是否输出到控制台
}

// InitLogger 初始化 log
func InitLogger(conf LogConfigs) error {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}

	// init log channel
	logCh = make(chan []byte, 2048)

	writeSyncer, err := getLogWriter(conf) // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getEncoder(conf)          // 获取日志输出编码
	level, ok := logLevel[conf.LogLevel] // 日志打印级别
	if !ok {
		level = logLevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller()) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(logger)
	//zap.L().Debug("")
	//zap.S().Debugf("")
	return nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder(conf LogConfigs) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeTime = customTimeEncoder            // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if conf.LogFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(conf LogConfigs) (zapcore.WriteSyncer, error) {

	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(conf.LogPath); !exist {
		if conf.LogPath == "" {
			conf.LogPath = DefaultLogPath
		}
		if conf.LogFileName == "" {
			conf.LogFileName = DefaultLogName
		}
		if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
			conf.LogPath = DefaultLogPath
			if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFileName), // 日志文件路径
		MaxSize:    conf.LogFileMaxSize,                           // 单个日志文件最大多少 mb
		MaxBackups: conf.LogFileMaxBackups,                        // 日志备份数量
		MaxAge:     conf.LogMaxAge,                                // 日志最长保留时间
		Compress:   conf.LogCompress,                              // 是否压缩日志
	}

	// 日志同时输出到控制台、日志文件和grpc log中
	var writeSyncers []zapcore.WriteSyncer
	writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	writeSyncers = append(writeSyncers, zapcore.AddSync(lumberJackLogger))

	grpcWrite := zapcore.AddSync(&GrpcLogChan{})
	writeSyncers = append(writeSyncers, grpcWrite)
	return zapcore.NewMultiWriteSyncer(writeSyncers...), nil

	// if conf.LogStdout {
	// 	// 日志同时输出到控制台和日志文件中
	// 	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	// } else {
	// 	// 日志只输出到日志文件
	// 	return zapcore.AddSync(lumberJackLogger), nil
	// }
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

type GrpcLogChan struct {
}

var logCh chan []byte

func (w *GrpcLogChan) Write(p []byte) (n int, err error) {
	d := make([]byte, len(p))
	copy(d, p)
	select {
	case logCh <- d:
	default:
	}

	return
}

func (w *GrpcLogChan) Sync() error {
	return nil
}

func SendLogToServerTask(ctx context.Context, grpcConn *grpc.ClientConn) error {
  // log server client
	//client := pb.NewServicesClient(grpcConn)

	for {
		select {
		case <-ctx.Done():
			zap.S().Warn("send log to server task ctx done, Bye...")
			return nil
		case v := <-logCh:
      /*
			if grpcConn.GetState() == connectivity.Ready {
				msgHdr := pb.ReqHeader{
					Time:         common.GetCurrUnixTimestamp(),
					HostId:       common.HostAgent.HostId,
					AgentId:      common.HostAgent.AgentId,
					AgentVersion: common.HostAgent.AgentVersion,
					AgentToken:   common.HostAgent.AgentToken,
				}

				msg := pb.ReportLogReq{
					ReqHdr:  &msgHdr,
					LogData: string(v),
				}
				client.ReportLog(ctx, &msg)
			}
      */
		}
	}
  return nil

}
