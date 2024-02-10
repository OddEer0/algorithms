JS_PKG_MANAGER=yarn

go:
	go run ./golang/main.go

js:
	$(JS_PKG_MANAGER) start

ts:
	npx ts-node ./typescript/main.ts