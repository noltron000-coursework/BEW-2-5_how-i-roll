const go = new Go();
WebAssembly
	.instantiateStreaming(fetch("../go/main.wasm"), go.importObject)
	.then((result) => {
		go.run(result.instance);
	});
