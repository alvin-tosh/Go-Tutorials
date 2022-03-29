# Go-Tutorials
A series of Go basics tutorials

Open a command line in the directory of your package and run the ```go``` command with the ```build``` subcommand, followed by the name of the source file.
For example;
```$ go build Hello.go```

To execute the binary, type ```./``` followed by the name of the binary file. 
```
$ ./Hello.go
# output
    Hello World
 ```
OR ---   
use the ```go``` command with the ```run``` subcommand, followed by the name of the source file. This will combine the two steps outlined above and produce the same result, however, NO executable will be saved in the working directory. This method is mostly used for one-off snippets and experimental code that is unlikely to be needed in the future.
```
$ go run Hello.go

# output
Hello World

    
