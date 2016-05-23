spac
====

Simple resource packager for golang projects.

Usage
-----

`spac` expects two arguments:

    Usage: spac <directory> <filename>

        <directory> : directory containing the resources to package
        <filename> : name of the golong source file packaging all the resoucres

Spac functions by traversing the directory pointed out by the first argument, 
adding the content of each file in a golang source file pointed by the second 
parameter in a map named `SpacContent`. Each file is stored as byte array 
described by a properly escaped string contained in a map.

Workflow
--------

A good way to employ `spac` in your project's workflow is to use it in 
conjunction with the `go generate` command. For instance, add the following 
line at the beginning of your `main.go`:

    //go:generate spac static spacContent.go

To package the content of your resource directory you will have to run 
`go generate` in your project. 

This will package the content of the directory named "static" into a source 
file named "spacContent.go" containting the resource map named "SpacContent".  

You can then use all the resources packaged as byte arrays from within your
go code.

License
-------

This project is under the [WTFPL](http://www.wtfpl.net/), see 
[LICENSE](https://github.com/Marneus68/spac/blob/master/LICENSE) for more 
details.
