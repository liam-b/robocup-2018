package main

const LOG_LEVEL = 7

var LOOPING = true
const LOOP_SPEED = 35
const SENSOR_INIT_DELAY = 200
const START_LOOP_DELAY = 300
const END_DELAY = 200
const LOG_MULTI_SAME_GROUPING = true

const NONE = "none"
const BLACK = "black"
const WHITE = "white"
const SILVER = "silver"
const GREEN = "green"

const LEFT = 0
const RIGHT = 1
const BOTH = 2
const FAST = 1
const SLOW = -1

const ENABLED_PIXEL = 0x00
const BATTERY_PIXEL = 0x01
const SCOPE_PIXEL = 0x02
const BEHAVIOUR_PIXEL = 0x04

const COLOR_RIGHT_PIXEL = 0x1a
const COLOR_LEFT_PIXEL = 0x1b

var COLOR_BLACK = [3]int{0, 0, 0}
var COLOR_WHITE = [3]int{120, 120, 120}

var COLOR_RED = [3]int{150, 0, 0}
var COLOR_GREEN = [3]int{0, 120, 0}
var COLOR_BLUE = [3]int{0, 0, 150}

var COLOR_YELLOW = [3]int{250, 200, 0}
var COLOR_PURPLE = [3]int{150, 0, 150}
var COLOR_CYAN = [3]int{0, 150, 150}