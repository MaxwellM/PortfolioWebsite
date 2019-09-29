#!/usr/bin/python3
import os

pug_dir = "./pug/"
less_dir = "./less/"
public_dir = "../public/"
css_dir = public_dir + "css/"

# copies directories from the src dir to the public dir with no clobber
dir_to_be_moved = ["js", "json", "images", "unityGames", "node_modules"]
for directory in dir_to_be_moved:
    print("moving {} to public/{}".format(directory, directory))
    os.system("cp -n -r {} ../public/".format(directory))

# processes the html files from the pug files
print("Converting pug files to HTML...")
pug_files = os.listdir(pug_dir)
print("Creating html files from {}".format(pug_files))
for pug_file in pug_files:
    if ".pug" in pug_file:
        os.system("pug -P -o {} {}{} ".format(public_dir, pug_dir, pug_file))

# Converting less files to css
print("Converting less files to CSS...")
less_files = os.listdir(less_dir)
print("Creating css files from {}".format(less_files))
for less_file in less_files:
    if ".less" in less_file:
        if "_" not in less_file:
            # removes the less from filename and adds css to the end
            css_file = less_file[0:len(less_file)-4] + "css"
            os.system("lessc {}{} {}{} ".format(less_dir, less_file, css_dir, css_file))
