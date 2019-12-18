package main

import (
	"math"
	"flag"
	"log"
	"os"
)

var myLog = log.New(os.Stdout, "", 0)

func main() {

	delay := flag.Float64("delay", 15, "Max propagation delay, unit is second")
	rate := flag.Float64("rate", 1.0/30, "Block rate, unit is blocks/second")
	level := flag.Float64("level", 0.01, "Security level, the probability of an honest block being marked red")

	flag.Parse()

	if *delay < 0 || *rate < 0 || *level < 0 {
		flag.Usage()
		myLog.Fatalf("keep parameters above zero! delay:%v rate:%v level:%v", *delay, *rate, *level)
	}

	if *level > 1.0 {
		flag.Usage()
		myLog.Fatalf("keep level under one! level:%v", *level)
	}

	antiConeSize(*delay, *rate, *level)

}

func antiConeSize(_delay, _rate, _security float64) int {
	expect := 2 * _delay * _rate
	if expect > 999 {
		myLog.Fatalf("keep expect:%v = 2 * _delay:%v * _rate:%v under 1000!", expect, _delay, _rate)
	}

	coef := 1 / math.Pow(math.E, expect)

	myLog.Printf("_delay:%v _rate:%v  _security:%v expect:%v coef:%v\n\n", _delay, _rate, _security, expect, coef)

	end := 100

	sum := 1.0

	for k := 0; k < end; k++ {
		part := 1.0
		for j := 1; j <= k; j++ {
			part *= expect / float64(j)
		}
		sum -= part * coef
		myLog.Printf("k=%v sum=%v", k, sum)
		if sum < _security {
			return k
		}
	}
	return 0
}
