.PHONY: wire

wire:
	cd ./example/cmd/server && wire

.PHONY: run

run:
	cd ./example/cmd/server && go run . -conf /Users/welong/Project/egret/example/configs/config.yaml
