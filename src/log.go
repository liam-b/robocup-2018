package main

import "fmt"
import "strconv"
import "strings"
import "time"

const PURPLE = "\x1b[35m"
const BOLD = "\x1b[1m"
const BLUE = "\x1b[34m"
const CYAN = "\x1b[36m"
const GREEN = "\x1b[32m"
const YELLOW = "\x1b[33m"
const RED = "\x1b[31m"
const BLACK = "\x1b[30m"
const END = "\x1b[0m"
const WHITE = ""

var counter int = 0

func printLog(difference string, color string, name string, symbol string, method []string, text string) {
  fmt.Println(BLACK + difference + " " + pad(strconv.Itoa(counter), 3) + END + " " + BOLD + color + symbol + " " + strings.ToUpper(name) + END + " " + PURPLE + strings.Join(method, ":") + END + " " + text)
  counter += 1
}

func pad(str string, plength int) string {
  for i := len(str); i < plength; i++ {
    str = "0" + str
  }
  return str
}

type Logger struct {
  methodString []string
  startTime time.Time
}

func (logger *Logger) init() {
  logger.startTime = time.Now()
  fmt.Println(BOLD + "TRACE " + GREEN + "DEBUG " + BLUE + "INFO " + CYAN + "NOTICE " + GREEN + "SUCCESS " + YELLOW + "YELLOW " + RED + "ERROR " + "FATAL " + END + PURPLE + "method " + BLACK + "timestamp")
}

func (logger *Logger) timeDifference() string {
  return strconv.Itoa(int(time.Since(logger.startTime).Minutes())) + ":" + strconv.Itoa(int(time.Since(logger.startTime).Seconds()))
}

func (logger *Logger) add(method string) {
  logger.methodString = append(logger.methodString, method)
}

func (logger *Logger) remove() {
  logger.methodString = logger.methodString[:len(logger.methodString) - 1]
}

func (logger Logger) trace(text string) {
  printLog(logger.timeDifference(), WHITE, "trace", "-", logger.methodString, text)
}

func (logger Logger) debug(text string) {
  printLog(logger.timeDifference(), GREEN, "debug", "➤", logger.methodString, text)
}

func (logger Logger) info(text string) {
  printLog(logger.timeDifference(), BLUE, "info", "ℹ", logger.methodString, text)
}

func (logger Logger) success(text string) {
  printLog(logger.timeDifference(), GREEN, "success", "✓", logger.methodString, text)
}

func (logger Logger) notice(text string) {
  printLog(logger.timeDifference(), CYAN, "notice", "!", logger.methodString, text)
}

func (logger Logger) warn(text string) {
  printLog(logger.timeDifference(), YELLOW, "warn", "⚠", logger.methodString, text)
}

func (logger Logger) error(text string) {
  printLog(logger.timeDifference(), RED, "error", "×", logger.methodString, text)
}

func (logger Logger) fatal(text string) {
  printLog(logger.timeDifference(), RED, "fatal", "☢", logger.methodString, text)
}
