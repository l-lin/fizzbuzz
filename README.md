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
curl http://localhost:3000/stats
```

## Available endpoints

| Endpoints         | Description                                                                                                                                                                                                                                                   | Request body                                                                        |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| `POST /fizz-buzz`          | Perform fizz-buzz<br/>Generate returns a list of strings with numbers from 1 to limit, where:<br/>- all multiples of int1 are replaced by str1<br/>- all multiples of int2 are replaced by str2<br/>- all multiples of int1 and int2 are replaced by str1str2 |  `{"int1": number, "int2": number, "limit": number, "str1": string,"str2": string}` |
| `GET /requests/stats`      | Get most used request with its parameters                                                                                                                                                                                                                     |                                                                                     |

## Getting involved

Check the [CONTRIBUTING guide](.github/CONTRIBUTING.md)

