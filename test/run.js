import { Path } from "rcompat/fs";
import { serve } from "rcompat/http";
import { make_request, make_response, env } from "@primate/binding/go";

env();

const cwd = await new Path(import.meta.url);
const { directory } = cwd;
const file = await directory.join("route.wasm").arrayBuffer();
const typedArray = new Uint8Array(file);

const go = new globalThis.Go();
await WebAssembly.instantiate(typedArray, {
  ...go.importObject,
  env: {}}).then(async result => {
    go.run(result.instance);
  serve(request => {
    const get = globalThis.Get;
    const response = make_response(get(make_request({ url: new URL(request.url) })));
    console.log(response);
    //const $response = typeof response === "object" ? JSON.stringify(response) : response;
    return new Response("11");
  }, {host: "0.0.0.0", port: 6161});
});
