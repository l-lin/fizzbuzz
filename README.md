# fizzbuzz

[![GoDoc](https://godoc.org/l-lin/fizzbuzz?status.svg)](https://pkg.go.dev/github.com/l-lin/fizzbuzz)
![Go](https://github.com/l-lin/fizzbuzz/workflows/Go/badge.svg)

> Fizz-buzz in a web server

## Usage

Download latest binary in the [release page](https://github.com/l-lin/fizzbuzz/releases).

```bash
# Launch web server (default port: 3000)
fizzbuzz serve
# execute fizz-buzz
curl http://localhost:3000/fizz-buzz -d '{"int1": 3, "int2": 5, "limit": 100, "str1": "Fizz", "str2": "Buzz"}'
# use Apache Server Benchmarking tool to simulate multiple requests
echo '{"int1": 3, "int2": 5, "limit": 100, "str1": "Fizz", "str2": "Buzz"}' > /tmp/data.json
ab -n 1000 -T application/json -p /tmp/data.json http://localhost:3000/fizz-buzz
# check most used requests with its parameters
curl http://localhost:3000/requests/stats
```

## Available endpoints

| Endpoints         | Description                                                                                                                                                                                                                                                   | Request body                                                                        |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| `POST /fizz-buzz`          | Perform fizz-buzz<br/>Generate returns a list of strings with numbers from 1 to limit, where:<br/>- all multiples of int1 are replaced by str1<br/>- all multiples of int2 are replaced by str2<br/>- all multiples of int1 and int2 are replaced by str1str2 |  `{"int1": number, "int2": number, "limit": number, "str1": string,"str2": string}` |
| `GET /requests/stats`      | Get most used request with its parameters                                                                                                                                                                                                                     |                                                                                     |

## Testing horizontal scaling

By default, this webapp will store the request stats in memory, which means it's not adequate to use
this webapp in cluster. Indeed, if there are multiple instances of the fizzbuzz webapp, the stats
are not shared between them. Hence, it's not scalable horizontally.

There are several way to mitigate this issue. One way is to use [Apache
Kafka](https://kafka.apache.org) to synchronize the stats between all instances by producing and
consuming Kafka message that will contain the HTTP requests that were received.

The architecture of the sample is the following:

```txt
             +----------+
             |  Nginx   |
             +----+-----+
                  |
     +----------+-+------------+
     |          |              |
     v          v              v
+----+---+ +----+---+     +----+---+
|FizzBuzz| |FizzBuzz| ... |FizzBuzz|
+----+---+ +--------+     +--------+
     ^          ^              ^
     |          |              |
   +-------+----+--+-----------+
   |       |       |
   v       v       v
+-----+ +-----+ +-----+
|Kafka| |Kafka| |Kafka|
+-----+ +-----+ +-----+
   ^       ^       ^
   |       |       |
   +-------+-------+
           |
   +-------+-------+
   |       |       |
   v       v       v
+-----+ +-----+ +-----+
| Zk  | | Zk  | | Zk  |
+-----+ +-----+ +-----+
```

We have:

- a loadbalancer (using Nginx) that will redirect the requests to a fizzbuzz node
- 10 fizzbuzz nodes that produce and consume Kafka message in the topic `fizzbuzz-request-stats`
- 3 kafka nodes in cluster
- 3 zookeeper nodes (just to test it out)

I've put a [docker-compose.yml](./docker-compose.yml) that will spawn the components describe above.

```bash
# use 10 instances of fizzbuzz webapp
docker-compose up --scale fizzbuzz=10

# execute fizz-buzz
curl http://localhost/fizz-buzz -d '{"int1": 3, "int2": 5, "limit": 100, "str1": "Fizz", "str2": "Buzz"}'
# use Apache Server Benchmarking tool to simulate multiple requests
echo '{"int1": 3, "int2": 5, "limit": 100, "str1": "Fizz", "str2": "Buzz"}' > /tmp/data.json
ab -n 1000 -T application/json -p /tmp/data.json http://localhost/fizz-buzz
# check most used requests with its parameters
curl http://localhost/requests/stats
```

### Helpful commands

```bash
# create a kafka topic "fizzbuzz-stats" with 4 partitions and replication factor 2
docker run --rm --net fizzbuzz_default \
  confluentinc/cp-kafka:5.4.1 kafka-topics \
  --create --topic fizzbuzz-request-stats \
  --partitions 4 \
  --replication-factor 2 \
  --if-not-exists \
  --zookeeper zk1:2181
# connect a kafka consumer that listens on topic "fizzbuzz-stats"
docker run -it --rm --name kafka-consumer --net fizzbuzz_default \
  confluentinc/cp-kafkacat:5.4.1 kafkacat \
  -b kafka1:9092,kafka2:9092,kafka3:9092 \
  -t fizzbuzz-request-stats \
  -C
# connect a kafka consumer that publishes on topic "fizzbuzz-stats"
docker run -it --rm --name kafka-producer --net fizzbuzz_default \
  confluentinc/cp-kafkacat:5.4.1 kafkacat \
  -b kafka1:9092,kafka2:9092,kafka3:9092 \
  -t fizzbuzz-request-stats \
  -P
# publish "1 2 Fizz" message to the topic "fizzbuzz-stats"
echo "1 2 Fizz" | docker run -i --rm --name kafka-producer --net fizzbuzz_default \
  confluentinc/cp-kafkacat:5.4.1 kafkacat \
  -b kafka1:9092,kafka2:9092,kafka3:9092 \
  -t fizzbuzz-request-stats \
  -P
# stop and rebuild fizzbuzz app
docker-compose stop fizzbuzz && docker-compose up --build fizzbuzz
```

## Getting involved

Check the [CONTRIBUTING guide](.github/CONTRIBUTING.md)

