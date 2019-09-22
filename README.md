[![Build Status](https://travis-ci.com/jaztec/golife.svg?branch=master)](https://travis-ci.com/jaztec/golife)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaztec/golife)](https://goreportcard.com/report/github.com/jaztec/golife)
[![License MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://github.com/jaztec/golife/blob/master/LICENSE)
[![GoDoc eurodnsgo](https://godoc.org/github.com/jaztec/golife?status.svg)](https://godoc.org/github.com/jaztec/golife)

# Just some experiment in Go

With the Conway's game of life and Golang. The game itself is just an excuse (no view what so ever)
to experiment with goroutines and calculations. Looking into when a linear calculation would be 
surpassed by multiple routines connected through pipeline channels. Apparently the change comes 
somewhere between 100 and 961 game cells:


```bash
BenchmarkLinearSimulator100_100-8         	     100	    10011850 ns/op	    189915 B/op	     936 allocs/op
BenchmarkLinearSimulator1000_100-8        	      20	   104818889 ns/op	   2480088 B/op	     798 allocs/op
BenchmarkLinearSimulator1000_1-8          	    2000	     1004545 ns/op	     24800 B/op	       7 allocs/op
BenchmarkLinearSimulator10000_100-8       	       1	  1181144229 ns/op	  20147640 B/op	    1348 allocs/op
BenchmarkLinearSimulator1000000_100-8     	       1	183367434139 ns/op	2491420344 B/op	    2223 allocs/op
BenchmarkThreadedSimulator100_100-8       	     200	    10248275 ns/op	    189647 B/op	     931 allocs/op
BenchmarkThreadedSimulator1000_100-8      	      20	    60629836 ns/op	   2483929 B/op	     805 allocs/op
BenchmarkThreadedSimulator1000_1-8        	    2000	      610659 ns/op	     24837 B/op	       8 allocs/op
BenchmarkThreadedSimulator10000_100-8     	       2	   617678912 ns/op	  20374668 B/op	    1358 allocs/op
BenchmarkThreadedSimulator1000000_100-8   	       1	 79860474458 ns/op	2555447512 B/op	    2401 allocs/op
```