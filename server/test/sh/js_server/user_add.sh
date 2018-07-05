#!/bin/bash

curl http://localhost:3000/users/add -d images[0][name]=first\&images[0][url]=d.png
