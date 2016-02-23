# GOOF Talk
![Gooftalk_Logo](https://raw.githubusercontent.com/GOOFS/go-goof-Talk/master/Images/Gooftalk_Logo.jpg)
 A simple chat application built with golang.

 For more details and the Design Document, go to our [Wiki](https://github.com/GOOFS/go-goof-Talk/wiki/Design-Document)

# Usage
Currently our code can get the server to running state and enables other clients to connect to it. Once connected, the username will be shown in server log and the online user list will be updated. Run the `server.go` with no parameters to run accept the messages in default port i.e. 3410.
To connect to the running server, use

`./client -user username -host 192.168.1.3:3410`

To connect to the server running in same machine, don't write -host.

# Demonstration

#####Client Window:
```sh
 $ ./client -user vishwas
  _____    ____     ____    ______         _______           _   _
 / ____|  / __ \   / __ \  |  ____|       |__   __|         | | | |
| |  __  | |  | | | |  | | | |__             | |      __ _  | | | | __
| | |_ | | |  | | | |  | | |  __|            | |     / _` | | | | |/ /
| |__| | | |__| | | |__| | | |               | |    | (_| | | | |   <
 \_____|  \____/   \____/  |_|               |_|     \__,_| |_| |_|\_\  v1.0
List of GOOFS online:
vishwas
manju
presely
listGoofs
Current online Goofs:
vishwas
manju
presely
oiwruowieroiwqru
Invalid function, try 'help' to list all available functions
help
Welcome to GOOFtalk help:
<<<<<<< HEAD
List of functions,
1. listGoofs
2. logout
3. shutdown
=======
List of funcitons,
1. listGoofs
2. logout
>>>>>>> origin/master
logout
2016/02/23 22:29:26 Logged out Succesfully
```

#####Server Window:

```sh
$ ./server
2016/02/23 22:42:09 Listening on port :3410...
2016/02/23 22:42:21 vishwas has joined the chat.
2016/02/23 22:42:28 manju has joined the chat.
2016/02/23 22:42:34 presely has joined the chat.
2016/02/23 22:42:45 Dumped list of Goofs to client output
2016/02/23 22:44:49 vishwas has left the chat
```
