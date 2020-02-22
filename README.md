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

## Getting involved

Check the [CONTRIBUTING guide](.github/CONTRIBUTING.md)

