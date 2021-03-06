Movie and TV Metadata Library [![Build Status](https://travis-ci.org/amahi/go-metadata.png?branch=master)](https://travis-ci.org/amahi/go-metadata)
==========
We have created a metadata library that helps us retrieve metadata information and artwork about media files.
![Metadata Flow](docs/img/metadata_flow.jpg)

Here is a high level picture of how the library works:
![Metadata Library Architecture](docs/img/metadata_server.jpg)

And here is the godoc [go-metadata documentation](http://godoc.org/github.com/amahi/go-metadata).

Install
=======
`go get github.com/amahi/go-metadata`

Details
=======
* In our implementation of this library L, the origin server acts as a client to the library
* The main entry point to L is method/function call with a media name(string) and Hint(string) such as "tv" or "movie" whose metadata we return in a structured way
* Format of the metadata returned will be a json structure with details such as cast, director, and url of cover art images that are related to it. 
* The library keeps a transparent cache as part of the implementation
* Whenever L gets a request from the client, it queries a local database (the cache) to find if the requested information is present in cache already
* If it is present, L reads the data from cache and returns it
* Otherwise S will request metadata from an online API, return the results as soon as possible and cache it
* The cache should uses a LRU caching policy

Metadata Output
============
```go
Lib,err := metadata.Init(1000000,"metadata.db")
if err == nil {
        data, err := Lib.GetMetadata("Breaking Bad","tv")
        if err == nil {
                fmt.Println(data)
        }
}
```

The above code will return a json string in the following format:-

```json
{
   "title":"Breaking Bad",
   "artwork":"http://thetvdb.com/banners/posters/81189-22.jpg",
   "year":"2008"
}
```
