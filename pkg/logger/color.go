package logger

// Colors.
type Color string

const (
	reset       Color = "\033[0m"
	red         Color = "\033[31m"
	green       Color = "\033[32m"
	yellow      Color = "\033[33m"
	blue        Color = "\033[34m"
	magenta     Color = "\033[35m" // 紫色
	cyan        Color = "\033[36m" // 蓝色
	white       Color = "\033[37m"
	blueBold    Color = "\033[34;1m"
	magentaBold Color = "\033[35;1m"
	redBold     Color = "\033[31;1m"
	yellowBold  Color = "\033[33;1m"
)

const (
	debugColor = green
	infoColor  = white
	warnColor  = cyan
	errorColor = redBold
	panicColor = yellowBold
	fatalColor = magentaBold
)

func (l Level) Color() Color {
	switch l {
	case DebugLevel:
		return debugColor
	case InfoLevel:
		return infoColor
	case WarnLevel:
		return warnColor
	case ErrorLevel:
		return errorColor
	case PanicLevel:
		return panicColor
	case FatalLevel:
		return fatalColor
	default:
		return infoColor
	}
}
