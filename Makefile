RTP=0.1

run-server:
	go run ./cmd/rtp-server/main.go -rtp=${RTP}

DEFAULT-GOAL: run-server

build:
	mkdir -p bin && \
	go build -o ./bin/rtpapi-server ./cmd/rtpapi-server/main.go

docker-up:
	docker build . -t rtpapi-img && \
	docker run -d -p "64333:64333" rtpapi-img -rtp=${RTP}
