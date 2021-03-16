# Legned1

## Startup

* From the main directory excute the following command
```cmd
go run . 
```

## How to send data to the server. 

```cmd
curl -v "localhost:8080" \
       -X POST \
       -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
       -H "Ce-Specversion: 1.0" \
       -H "Ce-Type: io.triggermesh.google.bigquery.query" \
       -H "Ce-Source: dev.knative.samples/helloworldsource" \
       -H "Content-Type: application/json" \
       -d '{"three":1000,"six":2000,"ten":3000}'
```
