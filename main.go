package main

import (
	"math"
	"flag"
	"log"
	"os"
)

var myLog = log.New(os.Stdout, "", 0)

func main() {

	delay := flag.Float64("delay", 40, "Max propagation delay, unit is second")
	rate := flag.Float64("rate", 1.0/60, "Block rate, unit is blocks/second")
	threshold := flag.Float64("threshold", 0.49, "The security threshold is the minimal hashing power(percentage) that the attacker must acquire in order to disrupt the protocol’s operation")

	flag.Parse()

	if *delay < 0 || *rate < 0 || *threshold < 0 {
		flag.Usage()
		myLog.Fatalf("keep parameters above zero! delay:%v rate:%v threshold:%v", *delay, *rate, *threshold)
	}

	if *threshold > 1.0 {
		flag.Usage()
		myLog.Fatalf("keep threshold under one! threshold:%v", *threshold)
	}

	antiConeSize(*delay, *rate, *threshold)

}

func antiConeSize(_delay, _rate, _threshold float64) {

	factor := 2 * _delay * _rate
	if factor > 1000 {
		myLog.Fatalf("keep factor:%v = 2 * _delay:%v * _rate:%v under 1000!", factor, _delay, _rate)
	}
	coef := 1 / (math.Pow(math.E, factor) - 1)

	level := 1 - 2*_threshold
	myLog.Printf("_delay:%v _rate:%v _threshold:%v level:%v factor:%v coef:%v\n\n", _delay, _rate, _threshold, level, factor, coef)

	sum := 0.0

	outLen := 10
	kQueue := make([]float64, 0)

	end := 1000

	k := -1
	for kk := 1; kk < end; kk++ {

		sum = 0.0
		for j := kk + 1; j < kk+100+int(factor); j++ {
			xx := 1.0
			for jj := 1; jj <= j; jj++ {
				xx *= factor / float64(jj)
			}
			sum += xx
		}
		sum *= coef

		if k < 0 {
			if sum < level {
				for i := 0; i < len(kQueue); i++ {
					leftBound := outLen
					if kk <= leftBound {
						leftBound = kk - 1
					}
					myLog.Printf("kk=%v sum=%v", kk-(leftBound-i), kQueue[i])
				}
				myLog.Printf("\n[MIN]kk=%v sum=%v\n\n", kk, sum)
				k = kk
				end = kk + 10 + 1
			}

			kQueue = append(kQueue, sum)
			if kk > outLen {
				kQueue = kQueue[1:]
			}

		} else {
			myLog.Printf("kk=%v sum=%v", kk, sum)
		}
	}
}
