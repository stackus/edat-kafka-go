#edat-kafka-go - Kafka for edat

## Installation

    go get -u github.com/stackus/edat-kafka-go

## Usage Example

    import (
        "github.com/stackus/edat-kafka-go"
        "github.com/stackus/edat/msg"
    )

    // Create a consumer and use it in a message subscriber
    consumer := edatkafkago.NewConsumer(brokers, groupID)
    subscriber := msg.NewSubscriber(consumer)

    // Create a producer and use it in a message publisher
    producer := edatkafkago.NewProducer(brokers)
    publisher := msg.NewPublisher(producer)

## Prerequisites

Go 1.15

## Features

- Message Consumer `NewConsumer(brokers, groupID, ...options)`
- Message Producer `NewProducer(brokers, ...options)`

## TODOs

- Documentation
- Tests, tests, and more tests

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT
