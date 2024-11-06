package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ndy/realworld-gin/internal/config"
	"os"
	"time"
)

var Log *zap.Logger

// InitLogger 함수는 로거를 초기화합니다.
func InitLogger() {
	// 커스텀 인코더 설정
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",  // 타임스탬프 키
		LevelKey:       "level",      // 레벨 키
		NameKey:        "logger",     // 로거 이름 키
		CallerKey:      "caller",     // 호출 위치 키
		MessageKey:     "msg",        // 메시지 키
		StacktraceKey:  "stacktrace", // 스택트레이스 키
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 레벨을 대문자로 출력
		EncodeTime:     customTimeEncoder,             // 커스텀 시간 포맷
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 파일명 및 라인 출력
		EncodeDuration: zapcore.StringDurationEncoder, // 지속 시간 포맷
	}

	// 콘솔 인코더를 사용하는 코어 생성
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 콘솔 출력용 인코더
		zapcore.AddSync(os.Stdout),               // 출력 대상 설정
		config.ZapConfig.Level,                   // 최소 로그 레벨 설정
	)

	// 로거 생성
	Log = zap.New(core, zap.AddCaller())
}

// customTimeEncoder 함수는 시간을 원하는 형식으로 출력합니다.
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05")) // 원하는 시간 형식
}

// Sync 함수는 로그를 출력합니다.
func Sync() {
	_ = Log.Sync()
}
