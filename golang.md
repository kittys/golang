# Golang
* 1st article you shuould look at - https://howistart.org/posts/go/1
* example code https://gobyexample.com/
* live package docuement http://godoc.org/
* 

## Basic stuff
* [Markdown syntax](https://help.github.com/articles/markdown-basics/) and [advance](https://help.github.com/articles/github-flavored-markdown/)
* [My Star Repos](https://github.com/stars)
* [Golang Trending Repos](https://github.com/trending?l=go&since=monthly)

## Questions
Lists of question I found during study
* How to iterate through struct? = Use [reflect](http://stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go)

## Sample Code (Open Source Software)
Learning from runnable code is the most best thing you can get. Here's some of it
* Result from Hackathon at [GopherGala.com](http://gopher-gala.challengepost.com/submissions/). Here [winners](http://gophergala.com/blog/gopher/gala/2015/02/03/winners/)
 * [Go Report Card](http://gopher-gala.challengepost.com/submissions/32189-go-report-card)
 * [AppStract](http://gopher-gala.challengepost.com/submissions/32181-appstract)

## Web Development
* main entry point is from [Mark Bates](https://github.com/markbates) from [gothamgo video](https://vimeo.com/115940590). All topics in gothamgo can found [here](https://blog.golang.org/gothamgo)
* Live Reload [Gin](https://github.com/codegangsta/gin), but I prefer [rerun](https://github.com/skelterjohn/rerun)
* Database [sqlx](https://github.com/jmoiron/sqlx), but i prefer [mgo](https://labix.org/mgo)
* Rendering [render](https://github.com/unrolled/render)
* Routing [Negroni](https://github.com/codegangsta/negroni), but i prefer [httpRouter](https://github.com/julienschmidt/httprouter)
* Authentication
* MiddleWare
* Context Sharing
* CORS - Cross Origin Resource Sharing

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
