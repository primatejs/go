/// {{{ start of primate wrapper, prefix
package main
import "syscall/js"
// }}} end

import "github.com/primatejs/go/primate";

func Get(request primate.Request) interface{} {
  //return request.View("1test", map[string]interface{}{
  //  "test": 1234,
  //});
  //return map[string]interface{}{
  //  "test": 1234,
  //};
  return primate.View("test", map[string]interface{}{});
}

// {{{ start primate wrapper, postfix
func main() {
  c := make(chan bool)
  js.Global().Set("Get", js.FuncOf(primate.MakeRequest(Get)))
  <-c
}
// }}} end
