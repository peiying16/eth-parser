# eth parser

This project implements a parser interface to keep track of eth addresses transations

Here is how it works

1. Compare chain block number and parsed block number and send a todo number to message queue (A message guarantee to process, and it won't lose block )
1. Scaner receive the number and get transcation from the chain, and save to storage.
1. API reads storage and provides data.

API and the scanner must be able to stateless, so they can scale.

## Diagrams

![1](docs/1.png)

![2](docs/2.png)
