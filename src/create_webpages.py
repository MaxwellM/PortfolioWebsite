#!/usr/bin/python3
import os, sys, time, platform, shutil

pug_dir = "./pug/"
less_dir = "./less/"
public_dir = "../public/"
public_html_dir = "../public/html/"
css_dir = public_dir + "css/"
dir_to_be_moved = ["js", "json", "images"]

def determine_os():
    """determines the OS of the system"""
    thisOs = platform.system()
    if thisOs == "Windows":
        return True
    else:
        return False

def move_directories_to_public():
    """copies directories from the src dir to the public dir"""
    for directory in dir_to_be_moved:
        if determine_os():
            try:
                if os.path.exists("../public/"+directory):
                    print("Exists, removing: "+directory)
                    shutil.rmtree("../public/"+directory)
                shutil.copytree(directory, "../public/"+directory)
            except shutil.Error as e:
                print("Error 1: %s" %e)
            except OSError as e:
                print("Error 2: %s" %e)
        else:
            os.system("cp -r {} ../public/".format(directory))

def convert_pug_to_html():
    """processes the html files from the pug files"""
    os.system("pug -P {} --out {}".format(pug_dir, public_html_dir))
    #pug_files = os.listdir(pug_dir)
    #for pug_file in pug_files:
    #    if ".pug" in pug_file:
    #        os.system("pug -P -o {} {}{} ".format(public_dir, pug_dir, pug_file))

def convert_less_to_css():
    """Converting less files to css"""
    less_files = os.listdir(less_dir)
    for less_file in less_files:
        if ".less" in less_file:
            if "_" not in less_file:
                # removes the less from filename and adds css to the end
                css_file = less_file[0:len(less_file)-4] + "css"
                os.system("lessc {}{} {}{} ".format(less_dir, less_file, css_dir, css_file))

if __name__ == "__main__":
    determine_os()
    if len(sys.argv) == 1:
        print("moving {} to the public directory".format(dir_to_be_moved))
        move_directories_to_public()
        print("Converting pug files to HTML...")
        convert_pug_to_html()
        print("Converting less files to CSS...")
        convert_less_to_css()
    elif len(sys.argv) == 2:
        if sys.argv[1] == "-w":
            print("Watching...")
            while True:
                move_directories_to_public()
                convert_pug_to_html()
                convert_less_to_css()
                time.sleep(5)
    else:
        print("add -w to watch")
