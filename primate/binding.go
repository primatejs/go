package primate

import "syscall/js";
import "encoding/json";
import "fmt";

func tryposition(array []map[string]any, position uint8) map[string]any {
  if (len(array) < int(position)) {
    return map[string]any{};
  }
  return array[position];
}

func serialize(data map[string]any) (string, error) {
  serialized, err := json.Marshal(data);
  if err != nil {
    return "", err;
  }
  return string(serialized), nil;
}

func Redirect(location string, rest ...map[string]any) any {
  options := tryposition(rest, 0);

  return js.FuncOf(func(this js.Value, args[] js.Value) any {
    return map[string]any{
      "handler": "redirect",
      "location": location,
      "options": options,
    };
  });
}

func View(component string, rest ...map[string]any) any {
  props, err := serialize(tryposition(rest, 0));
  if err != nil {
    fmt.Println(err.Error());
    return nil;
  }

  options, err := serialize(tryposition(rest, 1));
  if err != nil {
    fmt.Println(err.Error());
    return nil;
  }

  return js.FuncOf(func(this js.Value, args[] js.Value) any {
    return map[string]any{
      "handler": "view",
      "component": component,
      "props": props,
      "options": options,
    };
  });
}
