# Race Conditions - Concurrency and Channels

### Code Review

[Race Condition](../example1/example1.go) ([Go Playground](http://play.golang.org/p/Pvq-Dp7HnG))

[Atomic Increments](../example2/example2.go) ([Go Playground](http://play.golang.org/p/4CaNFVZaXn))

[Atomic Store/Load](../example3/example3.go) ([Go Playground](http://play.golang.org/p/vYVpq_l3gw))

[Mutex](../example4/example4.go) ([Go Playground](http://play.golang.org/p/At-ytL7Om_))

### Exercise 1
Review the provided program that will be used for future exercises. Program uses a function type, closures and goroutines to calculate Fibonacci numbers. This program contains a Data Race.
[Starter Program](exercise.go)  ([Go Playground](http://play.golang.org/p/Qh4MFV--hV))

### Exercise 2
From exercise 1, fix the data race using the sync package.

___
[![GoingGo Training](../../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)