#!/usr/bin/env

import os
from dotenv import load_dotenv
import json
load_dotenv()

PROXY_NAME_LIST = json.loads(os.getenv('PROXY_NAME_LIST'))
IP_CHANGE_SECONDS = int(os.getenv('IP_CHANGE_SECONDS'))
STARTING_PORT = int(os.getenv('STARTING_PORT'))

WARNING = "# Generated by generate-proxy-configurations script.\n\n"

# Generate docker-compose.yml.
#
with open("docker-compose.yml","w") as f:
    f.write(WARNING)
    f.write("version: '3'\n\nservices:\n")

    for index, name in enumerate(PROXY_NAME_LIST):
        f.write(f"  tor-{name}:\n")
        f.write(f"    container_name: tor-{name}\n")
        f.write(f"    image: 'pickapp/tor-proxy:latest'\n")
        f.write(f"    ports:\n")
        f.write(f"      - '{STARTING_PORT + index}:8888'\n")
        f.write(f"    environment:\n")
        f.write(f"      - 'IP_CHANGE_INTERVAL={IP_CHANGE_SECONDS}'\n")   
        # write retart 
        f.write("    restart: always\n")


# Generate proxy-list.txt
#
with open("proxy-list.txt", "w") as f:
    f.write(WARNING)
    
    for index, name in enumerate(PROXY_NAME_LIST):
        f.write(f'http://127.0.0.1:{STARTING_PORT+index}\n')