package primate

import "syscall/js";
import "encoding/json";
import "fmt";

func Redirect1(location string) interface{} {
  return js.FuncOf(func(this js.Value, args[] js.Value) interface{} {
    return map[string]interface{}{
      "handler": "redirect",
      "location": location,
    };
  });
}

func Redirect(location string, options map[string]interface{}) interface{} {
  return js.FuncOf(func(this js.Value, args[] js.Value) interface{} {
    return map[string]interface{}{
      "handler": "redirect",
      "location": location,
      "options": options,
    };
  });
}

func View1(component string) interface{} {
  return js.FuncOf(func(this js.Value, args[] js.Value) interface{} {
    return map[string]interface{}{
      "handler": "view",
      "component": component,
    };
  });
}

func View(component string, props map[string]interface{}) interface{} {
  serialized, err := json.Marshal(props);
  if err != nil {
    fmt.Println(err.Error());
    return nil;
  }
  stringified_props := string(serialized);

  return js.FuncOf(func(this js.Value, args[] js.Value) interface{} {
    return map[string]interface{}{
      "handler": "view",
      "component": component,
      "props": stringified_props,
    };
  });
}
