#!/usr/bin/env python3

# ./man.py
# Project management script
#
# (c) Justus Languell, 2021

import os, sys

def execute_command(command):
    if isinstance(command, str):
        os.system(command)
    elif isinstance(command, list):
        for s_command in command:
            execute_command(s_command)

def execute_commands(commands):
    if 'list' in sys.argv:
        print("list: available commands...")
        for command in commands:
            print("  * " + command)
    for command in sys.argv[1:]:
        if command in commands.keys():
            execute_command(commands[command])
        else:
            print("Unknown command: " + command)
            
if __name__ == '__main__':
    execute_commands({
        'build-server': 'cd server; go build ', 
        'run-server': 'cd server; ./app',
        'deploy-server': 'cd server; yes | gcloud app deploy; gcloud app browse',
        'browse-server': 'cd server; gcloud app browse',
        'build-client': 'rm -r server/public; cd client; yarn build; cp -r build ../server/public',
    })