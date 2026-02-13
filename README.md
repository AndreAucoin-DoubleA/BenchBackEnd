# BenchBackEnd

# 1. GoLang

Task 1:

- [x] Create a basic Go application that listens on port 7000 and has a status endpoint
- [x] Create a process to consume the wikipedia recent changes stream https://stream.wikimedia.org/v2/stream/recentchange and log these to stdout.
- [x] Replace the logs, with an in-memory /stats endpoint that a user can hit to get the latest stats on what we’ve processed
- [x] Create the following stats
    Number of messages consumed
    Number of distinct users
    Number of bots & Number of non-bots
    Count by distinct server URLs
- [x] Create tests for your application (if you didn’t already)
- [x] Run tests with the race detector on (-race)


Task 2: 

- [] Create a  Dockerfile for your application
- [] Build & Run your dockerized application
- [] Build a scratch container image of your application
- [] Use a file to set all the configurable items like ports, URLs and anything else that can be       dynamic, load these configs via the file

