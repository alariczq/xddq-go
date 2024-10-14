.PHONY: proto
proto:
	protoc -I=proto --go_out=paths=source_relative:./pkg/game/protocol/pb proto/*.proto