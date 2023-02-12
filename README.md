# DevOps
Itu miniTwit

## Notes on sunday, Convo between Amanda, Oliver and Bjarke.
After using most of the weekend understanding Golang and the Gin framework, we tried to convert the general functionality of ITU-minitwit.
We ended up not getting that far.

We currently only got the sqlite3-connection working and starting the public timeline. However the timeline is not rendered correctly due to lack of understanding the context handling and templating framework in Go. 

We furhtermore focused on getting a containered development setup, that we all could use - so the current setup uses a simple docker image. 

In general, due to our experience level with developing web-frameworks and Go, we where not able to execute on what was tasked and it has gotten us stuck at the moment. 

The general idea was to quickly mimic the Flask app and do some simple benchmarks in order to compare and justify the Go/Gin switch. Other considerations we had about choosing Go and the gin-framework; We wanted a compiled language and statically typed language, in order to handle errors better and possible obtain more speed. When reading about new frameworks, all of us had an interest in Go aswell, so as mentioned in the assignment we grabbed the opportunity to learn a new language.
