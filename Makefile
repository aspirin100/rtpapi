RTP=0

run-server:
	go run ./cmd/rtp-server/main.go -rtp=${RTP}

DEFAULT-GOAL: run-server