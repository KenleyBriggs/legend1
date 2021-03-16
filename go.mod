module github.com/briggs/test

go 1.16

//replace github.com/briggs/myfunc => ../myfunc

//require github.com/briggs/myfunc v0.0.0-00010101000000-000000000000

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/common v0.19.0
	go.uber.org/zap v1.16.0
	knative.dev/pkg v0.0.0-20210315160101-6a33a1ab29ac
)
