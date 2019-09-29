#!/usr/bin/python3
import sys, os

user = "postgres"
db_name = "portfoliowebsite_db"
sql_dir = "./sql/"

def print_help_menu():
    print("-c will create the database and tables")
    print("-cd will create the database")
    print("-ct will create the tables")
    print("-dd will drop the database")
    print("-dt will drop the tables")
    print("-a will allow you to add to a table")

if len(sys.argv) < 2:
    print_help_menu()
else:
    # Create table args
    if sys.argv[1] == "-c":
        os.system("psql -U {} -f {}/createDatabase.sql".format(user, sql_dir))
        os.system("psql -U {} -d {} -f {}/createTables.sql".format(user, db_name, sql_dir))
    elif sys.argv[1] == "-cd":
        os.system("psql -U {} -f {}/createDatabase.sql".format(user, sql_dir))
    elif sys.argv[1] == "-ct":
        os.system("psql -U {} -d {} -f {}/createTables.sql".format(user, db_name, sql_dir))
    # Drop table args
    elif sys.argv[1] == "-dd":
        #os.system("psql -U {} -d {} -f {}/deleteTables.sql".format(user, db_name, sql_dir))
        os.system("psql -U {} -f {}/deleteDatabase.sql".format(user, sql_dir))
    elif sys.argv[1] == "-dt":
        os.system("psql -U {} -d {} -f {}/deleteTables.sql".format(user, db_name, sql_dir))

