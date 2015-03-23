# Golang
Memo during my study on go language

## Basic stuff
[Markdown syntax](https://help.github.com/articles/markdown-basics/) and [advance](https://help.github.com/articles/github-flavored-markdown/)


## Unorder Lists
* [Error Handling](http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html) - gothamgo video in [vimeo](https://vimeo.com/115782573) 
```go
err := json.Unmarshal([]byte(`{"name":"bill"}`), u)
if err != nil {
 switch e := err.(type) {
 case *json.UnmarshalTypeError:
     log.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
 case *json.InvalidUnmarshalError:
   log.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
   default:
     log.Println(err)
  }
}```
