# Trade Executor

## Aproach

For the approach I've try to define the data flow of a trade. First a trade must enter in the system. After that, the system must be 
capable of open a socket to the API for the data retrieval and at the same time, be able to acept new trades. For that reason for every new trade a goroutine is spawned which will listen to the websocket (in another goroutine) and execute trades.


```

```


## Dificulties

The biggest dificult I had was with the unit testing as the engine package is highly coupled which made mock the websocket really hard.

## Best part

Best part was to work with real financial data in a web socket manner.

# Next steps

I'd like to rethink the approach of the system, althoug the data flow makes sense, the actual code is quite coupled.