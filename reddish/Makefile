build:
	go build -o ./bin/reddish . 

run: build
	./bin/reddish

clean:
	rm -rf bin

redis-cli:
	redis-cli -h 127.0.0.1 -p 6380

benchmark:
	redis-benchmark -t SET,GET -q -h 127.0.0.1 -p 6380

profile:
	go tool pprof -http :8081 http://localhost:6060/debug/pprof/profile\?seconds\=30

