# Metric distributed system

This is a distributed application composed of several parts:
1. An nginx server that proxys request from the outside
2. A javascript server that serves static files for our frontend
3. An Elixir application implementing websockets for real time data streaming
4. A golang app that acomplishes two things:
    
    1. Generate data each second and insert it on a sqlite db
    2. Open an endpoint to consult the stored data

The lifecycle of this app is very simple, golang will generate and insert data into sqlite, next elixir will consult the go app via http each second to get the last metric and broadcast that info to the phoenix channel, finally in the frontend each new metric receive from the channel will update the html page in realtime. 

The nginx server is used here to have a common ingress to distribute our traffic, we'll have two locations for our ingress, one to get the static files and another one to connect via sockets to elixir.

## Run the app

This app uses Docker to build and orchestrate this services. With Docker installed on your computer build each image with its correspondent tag, the tag should be the same as the directory you are in: 

example:
```
> cd metric-broadcaster
> docker build -t metric-broadcaster .
```
alternatively if you dont want to navigate directories, pass the dockerfile path with the -f and use its directory as reference:
```
> docker build -t metric-broadcaster -f ./metric-broadcaster/Dockerfile ./metric-broadcaster
```

After building all docker images you can execute
```
docker-compose up
```

## Considerations

This is an MVP and as it, it's built arround some considerations:

1. Despite only exposing only one mocked server data the metrics-generator is built to concurrently generate N server data with the use of WaitGroups. The ammount of servers is hardcoded to 1 but in the future will be accesible via ENV varibales.

2. To update the numbers of servers to track there are a few steps to it. `metric-generator` should have a new endpoint that retreives the available `serverIds` so that the elixir app can spawn one worker for each one of those servers and make the consultation for the latest metrics to post to the channel. Lastly the frontend should be updated to iterate over the metrics instead of accessing the first one.

   
