package main

import "fmt"
import "strconv"
import "strings"
import "time"

const TEXT_PURPLE = "\x1b[35m"
const TEXT_BOLD = "\x1b[1m"
const TEXT_BLUE = "\x1b[34m"
const TEXT_CYAN = "\x1b[36m"
const TEXT_GREEN = "\x1b[32m"
const TEXT_YELLOW = "\x1b[33m"
const TEXT_RED = "\x1b[31m"
const TEXT_BLACK = "\x1b[30m"
const TEXT_END = "\x1b[0m"
const TEXT_WHITE = ""

var counter int

func printLog(flag string, difference string, color string, name string, symbol string, method []string, text string) {
  fmt.Println(TEXT_BLACK + difference + " " + pad(strconv.Itoa(counter), 5) + " " + TEXT_BOLD + "(" + flag + ")" + TEXT_END + " " + TEXT_BOLD + color + symbol + " " + strings.ToUpper(name) + TEXT_END + " " + TEXT_PURPLE + strings.Join(method, "") + TEXT_END + " " + text)
  counter += 1
}

func pad(str string, plength int) string {
  for i := len(str); i < plength; i++ {
    str = "0" + str
  }
  return str
}

type Logger struct {
  flag string
  level int
  methodString []string
  startTime time.Time
  calledOnce bool
}

func (logger Logger) New(initialMethod string) Logger {
  fmt.Println("  ____   ___  ____   ___   ____ _   _ ____    ____   ___  _  ___  \r\n |  _ \\ / _ \\| __ ) / _ \\ / ___| | | |  _ \\  |___ \\ / _ \\/ |( _ ) \r\n | |_) | | | |  _ \\| | | | |   | | | | |_) |   __) | | | | |/ _ \\ \r\n |  _ <| |_| | |_) | |_| | |___| |_| |  __/   / __/| |_| | | (_) |\r\n |_| \\_\\\\___/|____/ \\___/ \\____|\\___/|_|     |_____|\\___/|_|\\___/ \r\n          ")
  logger.startTime = time.Now()
  fmt.Println(TEXT_BOLD + "TRACE " + TEXT_GREEN + "DEBUG " + TEXT_BLUE + "INFO " + TEXT_CYAN + "NOTICE " + /*TEXT_GREEN + "SUCCESS " + */TEXT_YELLOW + "WARN " + TEXT_RED + "ERROR " + "FATAL " + TEXT_END + TEXT_PURPLE + "method " + TEXT_BLACK + "meta " + TEXT_BOLD + "flag" + TEXT_END)
  if initialMethod != "" {
    logger.inc(initialMethod[1:])
  }
  if logger.level > 6 { logger.trace("logger started at level " + logger.value(strconv.Itoa(logger.level))) }
  logger.calledOnce = false
  return logger
}

func (logger *Logger) timeDifference() string {
  return pad(strconv.Itoa(int(time.Since(logger.startTime).Minutes())), 2) + ":" + pad(strconv.Itoa(int(time.Since(logger.startTime).Seconds()) % 60), 2)
}

func (logger *Logger) inc(method string) {
  logger.methodString = append(logger.methodString, method)
}

func (logger *Logger) dec() {
  logger.methodString = logger.methodString[:len(logger.methodString) - 1]
}

func (logger *Logger) rep(method string) {
  logger.dec()
  logger.inc(method)
}

func (logger *Logger) set(methods string) {
  logger.methodString = strings.Split(methods[1:], ":")
}

func (logger *Logger) once(method string) {
  logger.inc(method)
  logger.calledOnce = true
}

func (logger *Logger) trace(text string) {
  if logger.level > 6 { printLog(logger.flag, logger.timeDifference(), TEXT_WHITE, "trace", "-", logger.methodString, text) }
  logger.handleOnceCall()
}

func (logger *Logger) debug(text string) {
  if logger.level > 5 { printLog(logger.flag, logger.timeDifference(), TEXT_GREEN, "debug", "➤", logger.methodString, text) }
  logger.handleOnceCall()
}

func (logger *Logger) info(text string) {
  if logger.level > 4 { printLog(logger.flag, logger.timeDifference(), TEXT_BLUE, "info ", "ℹ", logger.methodString, text) }
  logger.handleOnceCall()
}

// func (logger *Logger) success(text string) {
//   if logger.level > 4 { printLog(logger.flag, logger.timeDifference(), TEXT_GREEN, "success", "✓", logger.methodString, text) }
//   logger.handleOnceCall()
// }

func (logger *Logger) notice(text string) {
  if logger.level > 3 {
    // fmt.Println(TEXT_CYAN + "                 ________" + TEXT_END)
    printLog(logger.flag, logger.timeDifference(), TEXT_CYAN, "notice", "!", logger.methodString, text)
    // fmt.Println(TEXT_CYAN + "                 ‾‾‾‾‾‾‾‾" + TEXT_END)
  }
  logger.handleOnceCall()
}

func (logger *Logger) warn(text string) {
  if logger.level > 2 { printLog(logger.flag, logger.timeDifference(), TEXT_YELLOW, "warn ", "⚠", logger.methodString, text) }
  // go bot.speaker.warn()
  logger.handleOnceCall()
}

func (logger *Logger) error(text string) {
  if logger.level > 1 { printLog(logger.flag, logger.timeDifference(), TEXT_RED, "error", "×", logger.methodString, text) }
  // go bot.speaker.error()
  logger.handleOnceCall()
}

func (logger *Logger) fatal(text string) {
  // fmt.Println(TEXT_RED + "                ________" + TEXT_END)
  if logger.level > 0 { printLog(logger.flag, logger.timeDifference(), TEXT_RED, "fatal", "☢", logger.methodString, text) }
  // fmt.Println(TEXT_RED + "                ‾‾‾‾‾‾‾‾" + TEXT_END)
  // go bot.speaker.fatal()
  logger.handleOnceCall()
}

func (logger Logger) value(text string) string {
  return TEXT_CYAN + text + TEXT_END
}

func (logger *Logger) handleOnceCall() {
  if logger.calledOnce {
    logger.calledOnce = false
    logger.dec()
  }
}
