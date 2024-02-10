JS_PKG_MANAGER=yarn

go:
	go run main.go

js:
	$(JS_PKG_MANAGER) start

ts:
	npx ts-node index.ts