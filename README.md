# RRRECSulator

The [RRECSulator](https://rrecsulator.com) helps USPS rural carriers calculate their pay check in RRECS.

## Source

This repository contains the source code for the back end server of the site.

Many eyes find many bugs.  Please report here, or email to [rrecsulator.com@gmail.com](mailto:rrecsulator.com@gmail.com)

## History

The calculator was developed from discussion in the [Rural Mail Talk Forum](https://www.ruralmailtalk.com/forums/rrecs-detailed-calculations-by-c.31/)

## Implementation Details

The server is written in go.  Each folder roughly corresonds to a different component of a carriers evaluation (e.g `flat` is for dealing with flats, like newspapers, magazines, etc.)  There are a few specialized folders:

- `standards` --> numeric constants are defined (like verifying letter addresses).
- `dataSets` --> glues together all of the other folders.
- `srv` --> code for the back end server.

## Running A Server

You can run your own RRECSulator calculator server!  

1.  Must have [golang](https://go.dev/doc/tutorial/getting-started) installed
2.  Change to the serve directory
3.  Build
4.  Run the server

The commands will look something like this, depending on your os:
```
cd srv
go build
./srv
```

## Using the RRECSulator.com server

You can also just make api request to the server runnign on RRECSulator.com  
This requires some familiarity with JSON.  Then build a request and point it at https://rrecsulator.com/api/daily
And the response will have all the details for your daily evaluation.

For instance, someone could make a plugin that integrates very cleanly with Excel and call out to the rrecsulator api for the daily eval calculations...


# License
GPL v3
