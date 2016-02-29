# GOOF Talk
![Gooftalk_Logo](https://raw.githubusercontent.com/GOOFS/go-goof-Talk/master/img/Gooftalk_Logo.jpg)
 A simple chat application built with golang.

 For more details and the Design Document, go to our [Wiki](https://github.com/GOOFS/go-goof-Talk/wiki/Design-Document)

# Usage
Currently our code can get the server to running state and enables other clients to connect to it. Once connected, the username will be shown in server log and the online user list will be updated. Run the `server.go` with no parameters to run accept the messages in default port i.e. 3410.
To connect to the running server, use

`./client -user <username> -host 192.168.1.3:3410`

To connect to the server running in same machine, don't write -host.


# Documentation
Each of the exported functions present in custom packages `goofserver` and `goofclient` contains comments about their fuctionalities. To view them in command line, use :<br/>
` $ godoc goofclient <function name> `
<br/>OR
<br/>` $ godoc goofserver <function name> `
<br/>
Please note that to make godoc read from our package, the project path should be assigned as GOPATH.<br/>
Alternatively, you can view the html documents present in `doc` folder which are generated using `godoc -html` command
