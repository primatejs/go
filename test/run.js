import { Path } from "rcompat/fs";
import { serve } from "rcompat/http";
import { make_request, make_response, env } from "@primate/binding/go";

env();

const cwd = await new Path(import.meta.url);
const { directory } = cwd;
const file = await directory.join("route.wasm").arrayBuffer();
const typedArray = new Uint8Array(file);

const getter = {
  get() {
    return {foo: "bar"};
  }
}

const fake_request = request => ({
  url: new URL(request.url),
  body: getter,
  path: getter,
  query: getter,
  cookies: getter,
  headers: getter,
});

const go = new globalThis.Go();
await WebAssembly.instantiate(typedArray, {
  ...go.importObject,
  env: {}}).then(async result => {
    go.run(result.instance);
  serve(request => {
    const get = globalThis.Get;
    const response = make_response(get(make_request(fake_request(request))));
    //const $response = typeof response === "object" ? JSON.stringify(response) : response;
    console.log(response);
    return new Response(response);
  }, {host: "0.0.0.0", port: 6161});
});
