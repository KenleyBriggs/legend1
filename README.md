# SQL Rest server


## Preqreq
export SQL_USERNAME=example 
export SQL_PASSWORD=examplePAssword
export SQL_SERVER="briggs-roofing-pricing.database.windows.net"
export SQL_PORT=1433
export SQL_DATABASE="Legend"

```cmd
curl -v "localhost:8080/pricing" \
       -X POST \
       -d '{"three":1000,"six":2000,"ten":3000}'
```




```cmd
curl -v "localhost:8080/insertrow" \
       -X POST \
       -d '{"prodName":"bobs burgers","three":1000,"six":2000,"ten":3000}'
```
