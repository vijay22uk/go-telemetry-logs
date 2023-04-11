# go-telemetry-logs
sample minimal go with telemetry [ Metrics, Traces &amp; logs ]

## Running trace Collector
```docker run -d -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:latest```



```jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))```

```url = 