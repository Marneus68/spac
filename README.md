spac
====

Spac is a simple resource packager for golang projects.

Usage
-----

`spac` expects two arguments:

    Usage: spac <directory> <filename>

        <directory> : directory containing the resources to package
        <filename> : name of the golong source file packaging all the resoucres

Spac functions by traversing the directory pointed out by the first argument, 
adding the content of each file in a golang source file pointed by the second 
parameter in a map named `SpacContent`.


