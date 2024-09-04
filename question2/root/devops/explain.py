fp = "./file.pin"
# using high level API and works perfectly
with open(fp, 'r') as f:
    for line in f:
        print(line.rstrip())

# override the high level API to low level API 
from os import *
# This will failed and still confused us we are using 
# high level API. 
with open(fp, 'r') as f:
    for line in f:
        print(line.rstrip())
