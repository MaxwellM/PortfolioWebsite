#!/usr/bin/python
import sys, os

less_files = os.listdir("./less")
for file_name in less_files:
    css_file = file_name.split(".")[0] + ".css"
    print(css_file)
    os.system("lessc ./less/{} > ./css/{}".format(file_name, css_file))

pug_files = os.listdir("./pug")
pug_files.remove("resources")
for file_name in pug_files:
    os.system("pug ./pug/{} -o ./html/".format(file_name))