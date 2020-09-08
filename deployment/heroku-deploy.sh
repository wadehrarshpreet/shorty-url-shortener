#!/bin/sh
sudo heroku container:push web -a washorty
sudo heroku container:release web -a washorty