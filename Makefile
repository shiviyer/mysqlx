all: test

export MYSQLX_TEST_DATASOURCE ?= mysqlx://my_user:my_password@127.0.0.1:33060/?_auth-method=PLAIN&time_zone=UTC

init:
	go get -u golang.org/x/perf/cmd/benchstat
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

install:
	go test -v -i
	go test -v -i -race
	go install -v
	go install -v -race
	go vet .

test: install
	go test -race -tags gofuzzgen
	go test -v -coverprofile=coverage.txt

cover: test
	go tool cover -html=coverage.txt

bench: test
	go test -run=NONE -bench=. -benchtime=3s -count=5 -benchmem | tee bench.txt

check: install
	-gometalinter.v2 --tests --deadline=60s ./...

proto:
	cd internal && go run compile.go

fuzz: test
	go-fuzz-build -func=FuzzUnmarshalDecimal github.com/AlekSi/mysqlx
	go-fuzz -bin=mysqlx-fuzz.zip -workdir=go-fuzz/unmarshalDecimal

seed:
	docker exec -ti mysqlx_mysql_1 sh -c 'mysql < /test_db/mysql/world_x/world_x.sql'
	docker exec -ti mysqlx_mysql_1 mysql -e "GRANT ALL ON world_x.* TO 'my_user'@'%';"

.PHONY: fuzz
