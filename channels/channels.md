### Routines

* Problem => Blocking calls waste time & resources!

* Routine is executor of instructions in the go program
* Main routine (high priority) (entry + exit) (default)
* Child routines (low priority)

* spin a new child routine: Use keyword go
> go someTask()
this someTask is executed in a child routine (a non blocking work!)

* Concurrency is not parallelism!

### Channels

* Used to communicate between go routines
* Two way messaging device
* They are typed in nature! like string channels, int channels

* syntax
> c := make(chan string)

* Sending data
channel <- 5 (value pushed in the channel)

* Receiving data
channelVariable <- channel (reading value in a variable from channel)

fmt.Println(<- channel)

### Function literal => Anonymous function
