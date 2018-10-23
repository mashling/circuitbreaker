# Circuit Breaker Activity Gateway

## Install

To install run the following commands:
```
flogo create -f flogo.json
cd CircuitBreakerActivity
flogo install github.com/mashling/circuitbreaker
flogo install github.com/mashling/httpactivity
flogo install github.com/TIBCOSoftware/flogo-contrib/activity/rest
flogo install github.com/TIBCOSoftware/flogo-contrib/activity/actreply
flogo install github.com/TIBCOSoftware/flogo-contrib/activity/log
flogo build
```

## Testing

Run:
```
CircuitBreakerGateway
```

Scenario 1: Target service in working state
Open another terminal and run:
```
go run testserver.go -server
```

Then open another terminal and run:
```
curl http://localhost:9096/pets/4
```

You should then see something like:
```
{"category":{"id":0,"name":"string"},"id":1,"name":"sally","photoUrls":["string"],"status":"available","tags":[{"id":0,"name":"string"}]}
```

Scenario 2: Circuit breaker gets tripped
Simulate a service failure by stopping the target service
and then run the following command 6 times:
```
curl http://localhost:9096/pets/4
```

You should see the below response 5 times:
```
 "connection failure"
```

Followed by:
```
"circuit breaker tripped"
```

The circuit breaker is now in the tripped state.
Start the gateway target service back up and wait at least one minute.

Repeat Scenario 1 and you will get response as expected.
The circuit breaker is no longer in the tripped state, and the target service is working.
