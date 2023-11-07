/// {{{ start of primate wrapper, prefix
package main
import "syscall/js"
// }}} end

import "github.com/primatejs/go/primate";

func Get(request Request) interface{} {
  //return request.View("1test", map[string]interface{}{
  //  "test": 1234,
  //});
  //return map[string]interface{}{
  //  "test": 1234,
  //};
//  return request.Body.Get("foo")
  return primate.View("test", map[string]interface{}{
    "foo": "bar",
    "bar": []string{"baz"},
    "baz": []interface{}{
      map[string]interface{}{"HO": "HI"},
    },
  });
}

// {{{ start primate wrapper, postfix
func main() {
  c := make(chan bool)
  js.Global().Set("Get", js.FuncOf(make_request(Get)))
  <-c
}
// }}} end
