## Testing

### Code Review

The sample program implements support for accessing a MongoDB database from MongoLab. The program implements two different find calls that return documents that are unmarshaled into user defined types.

[example1.go](../example1/example1.go)

[buoy/buoy.go](../example1/buoy/buoy.go)

[mongodb/mongodb.go](../example1/mongodb/mongodb.go)

### Tests

[Standard tests for testing calls to MongoDB](../example1/tests/example1_test.go)

[Table tests that perform multiple calls with different ids](../example1/tests/example1_table_test.go)

[Benchmarks that test the performance of the MongoDB find](../advanced/tests/example1_bench_test.go)

### Exercise 1
Write new tests for the FindRegion function found in the buoy package. Write a standard test, table test and benchmark.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)