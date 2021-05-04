### Publisher

A golang tool to interact with pubsub emulator on your local env

### Install 

```
make install
```

### Usage

Flags definitions : 

```
--host="" // to define your local host for the pubsub emulator
--topic="" // a topic ID
```

List pub sub topics 

```
publisher listTopic --host="127.0.0.1:8085"
```


Create a new topic 

```
publisher createTopic --host="127.0.0.1:8085" --topic="topic"
```

Publish a message to a topic

```
publisher publish --host="127.0.0.1:8085" --topic="topic" --message='{"message":"content"}'
```