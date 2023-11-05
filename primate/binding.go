package primate

import "syscall/js"
import "encoding/json"

type t_request func(Request) interface{} 
type t_response func(js.Value, []js.Value) interface{}
type t_view func(string, map[string]interface{}) interface{}
type t_redirect func(string, map[string]interface{}) interface{}

type URL struct {
  Href string
  Origin string
  Protocol string
  Username string
  Password string
  Host string
  Hostname string
  Port string
  Pathname string
  Search string
  SearchParams map[string]interface{}
  Hash string
}

type Request struct {
  Url URL
  View t_view
  Redirect t_redirect
}

func make_url(request js.Value) URL {
  url := request.Get("url")
  search_params := make(map[string]interface{})
  json.Unmarshal([]byte(request.Get("search_params").String()), &search_params)

  return URL{
    url.Get("href").String(),
    url.Get("origin").String(),
    url.Get("protocol").String(),
    url.Get("username").String(),
    url.Get("password").String(),
    url.Get("host").String(),
    url.Get("hostname").String(),
    url.Get("port").String(),
    url.Get("pathname").String(),
    url.Get("search").String(),
    search_params,
    url.Get("hash").String(),
  }
}

func make_view(request js.Value) t_view {
  view := request.Get("view")

  return func(component string, props map[string]interface{}) interface{} {
    return view.Invoke(component, props);
  }
}

func make_redirect(request js.Value) t_redirect {
  redirect := request.Get("redirect")

  return func(location string, options map[string]interface{}) interface{} {
    return redirect.Invoke(location, options);
  }
}
 
func MakeRequest(route t_request) t_response {
  return func(this js.Value, args[] js.Value) interface{} {
    request := args[0];
    go_request := Request{
      make_url(request),
      make_view(request),
      make_redirect(request),
    }

    return route(go_request)
  }
}
