#!/usr/bin/python3

import os, sys

if len(sys.argv) > 1:
    if sys.argv[1] == "-r":
        print("hit")
        os.system("go run go/runServer.go")
    if sys.argv[1] == "-b":
        os.system("go build -o ../runServer go/runServer.go")
else:
    os.system("go build -o ../runServer go/runServer.go")
