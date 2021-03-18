# SQL Rest server


## Preqreq
export SQL_USERNAME=example 
export SQL_PASSWORD=examplePAssword

```cmd
curl -v "localhost:8080/pricing" \
       -X POST \
       -d '{"three":1000,"six":2000,"ten":3000}'
```
