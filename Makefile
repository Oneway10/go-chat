ProjectPath = ..\go-chat\biz

gen:
	go mod tidy

build:
	go build .

run-dev: build
	./chat

run: build
	./chat &

cp-server:
	scp -r ..\go-chat\biz server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\cmd server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\common server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\conf server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\config server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\dal server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\idl server0:/home/ubuntu/run/go-chat
	scp -r ..\go-chat\.gitignore .hz go.mod go.sum main.go Makefile router.go router_gen.go server0:/home/ubuntu/run/go-chat

hz-new:
	hz new -idl ./idl/user.thrift

hz-update:
	hz update -idl ./idl/user.thrift
