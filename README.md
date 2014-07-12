Metadata Handler
==========
We are creating a metadata library to be used in a server that helps us retrieve metadata information and artwork about media files. Here is a high level picture of how the library works: 

![Metadata Library Architecture](docs/img/metadata_server.jpg)

Details
=======
* In our implementation of this library L, the origin server acts as a client to the library
* The main entry point to L will be a method/function call with a media name(string) whose metadata we return in a structured way
* Format of the metadata returned will be a json structure with details such as cast, director, and url of cover art images that are related to it. There may be more than one call for this information separately
* The library keeps a transparent cache as part of the implementation
* Whenever L gets a request from the client, it queries a local database (the cache) to find if the requested information is present in cache already
* If it is present, L reads the data from cache and returns it
* Otherwise S will request metadata from an online API, return the results as soon as possible and cache it
* The cache should use a caching policy, like LRU

TV Metadata
============
```go
data, err := metadata.GetMetadata("Modern Family - 2x17 - Two Monkeys and a Panda.avi","tv")
```

The above code will return a json string in the following format:-

```json
{
   "Media_type":"tv",
   "SeriesName":"Modern Family",
   "Banner_Url":"http://thetvdb.com/banners/",
   "Actors":"|Julie Bowen|Ty Burrell|Jesse Tyler Ferguson|Eric Stonestreet|Sofia Vergara|Ed O'Neill|Rico Rodriguez|Nolan Gould|Sarah Hyland|Ariel Winter|Aubrey Anderson-Emmons|",
   "Overview":"This mockumentary explores the many different types of a modern family through the stories of a gay couple, comprised of Mitchell and Cameron, and their daughter Lily, a straight couple, comprised of Phil and Claire, and their three kids, Haley, Alex, and Luke, and a multi-ethnic couple, which is comprised of Jay and Gloria, and her son Manny.",
   "Banner":"graphical/95011-g11.jpg",
   "FanArt":"fanart/original/95011-2.jpg",
   "Poster":"posters/95011-3.jpg",
   "Rating":"8.8",
   "FirstAired":"2009-09-23"
}
```

Movie Metadata
============

```go
data, err := metadata.GetMetadata("The.Prince.and.the.Pauper.avi","movie")
```

The above code will return a json string in the following format:-

```json
{
   "Id":58831,
   "Media_type":"movie",
   "Backdrop_path":"/zTy1rX6MTcWUH20koNz6qqCOsMa.jpg",
   "Poster_path":"/1qWQK5f53MiUWoNDEihvRMnVYwV.jpg",
   "Credits":{
      "Id":58831,
      "Cast":[
         {
            "Character":"",
            "Name":"Cole Sprouse",
            "Profile_path":"/ecR3jwrg0nSeXZQ4l7a4y73Coul.jpg"
         },
         {
            "Character":"",
            "Name":"Dylan Sprouse",
            "Profile_path":"/86nkUScYbpUzEDTQOFipiFvdfQi.jpg"
         },
         {
            "Character":"",
            "Name":"Kay Panabaker",
            "Profile_path":"/ucba2M5Mt1Roa1fJVNA7PRybjMm.jpg"
         },
         {
            "Character":"Miles",
            "Name":"Vincent Spano",
            "Profile_path":"/ktwG3Xfk4toIBAG0PcbJcjEBICx.jpg"
         }
      ],
      "Crew":[
         {
            "Department":"Directing",
            "Name":"James Quattrochi",
            "Job":"Director",
            "Profile_path":""
         }
      ]
   },
   "Config":{
      "Images":{
         "Base_url":"http://image.tmdb.org/t/p/",
         "Secure_base_url":"https://image.tmdb.org/t/p/",
         "Backdrop_sizes":[
            "w300",
            "w780",
            "w1280",
            "original"
         ],
         "Logo_sizes":[
            "w45",
            "w92",
            "w154",
            "w185",
            "w300",
            "w500",
            "original"
         ],
         "Poster_sizes":[
            "w92",
            "w154",
            "w185",
            "w342",
            "w500",
            "w780",
            "original"
         ],
         "Profile_sizes":[
            "w45",
            "w185",
            "h632",
            "original"
         ],
         "Still_sizes":[
            "w92",
            "w185",
            "w300",
            "original"
         ]
      }
   },
   "Imdb_id":"tt0874424",
   "Overview":"A modern day telling of the Mark Twain classic, The Prince and the Pauper.",
   "Title":"The Prince and the Pauper",
   "Release_date":"2007-11-11"
}
```


