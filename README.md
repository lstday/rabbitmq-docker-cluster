# HAProxy and Clustering RabbitMQ docker containers

## Using `docker-compose`

This section details how to start the whole cluster using `docker-compose`.

1. Create a network shared by all containers
```bash
docker network create cluster-network
```

2. Start cluster:
```bash
docker-compose up -d
```

3. View logs for all containers
```bash
docker-compose logs -f
```

## Using `docker`

This section details how to start the whole cluster using `docker` command, starting containers one by one.

1. Create a network shared by all containers
```bash
docker network create cluster-network
```

2. Start master node:
```bash
docker run -d \
    --name="rabbit1" \
    --hostname="rabbit1"\
    -e RABBITMQ_ERLANG_COOKIE="12345" \
    -e RABBITMQ_NODENAME="rabbit1" \
    --volume=(pwd)/rabbitmq.config:/etc/rabbitmq/rabbitmq.config \
    --volume=(pwd)/definitions.json:/etc/rabbitmq/definitions.json \
    --volume=(pwd)/enabled_plugins:/etc/rabbitmq/enabled_plugins \
    --network=cluster-network \
    rabbitmq:3-management
```
3. Start node #1:
```bash
docker run -d \
    --name="rabbit2" \
    --hostname="rabbit2"\
    -e RABBITMQ_ERLANG_COOKIE="12345" \
    -e RABBITMQ_NODENAME="rabbit2" \
    --volume=(pwd)/rabbitmq.config:/etc/rabbitmq/rabbitmq.config \
    --volume=(pwd)/definitions.json:/etc/rabbitmq/definitions.json \
    --volume=(pwd)/enabled_plugins:/etc/rabbitmq/enabled_plugins \
    --network=cluster-network \
    rabbitmq:3-management
```

4. Start node #2:
```bash
docker run -d \
    --name="rabbit3" \
    --hostname="rabbit3"\
    -e RABBITMQ_ERLANG_COOKIE="12345" \
    -e RABBITMQ_NODENAME="rabbit3" \
    --volume=(pwd)/rabbitmq.config:/etc/rabbitmq/rabbitmq.config \
    --volume=(pwd)/definitions.json:/etc/rabbitmq/definitions.json \
    --volume=(pwd)/enabled_plugins:/etc/rabbitmq/enabled_plugins \
    --network=cluster-network \
    rabbitmq:3-management
```

5. Start haproxy:
```bash
docker run -d \
    --name="haproxy" \
    --hostname="haproxy"\
    --volume=(pwd)/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg \
    --publish="1936:1936" \
    --publish="1883:1883" \
    --publish="1883:1883" \
    --publish="4369:4369" \
    --publish="5672:5672" \
    --publish="15672:15672" \
    --network=cluster-network \
    haproxy:latest
```

## Checking the services

This section details how to check your cluster working

1. View container logs individually
```bash
docker logs -f $container_name
```

2. View conpose logs
```bash
docker-compose logs -f
```

3. Run sender / reciever to check
```
Send.py and recieve.py scripts was stolen from Internet. Use it as in example: "python send.py" and "python recieve.py" 
```

4. Web-interface
```
Use web interface to check status. Enter your host addres with port 15672, and choose user, which was created in config files.
```

## Generate password hash for rabbitmq config files

This section details how to generate has by password from console

1. Compile or run files in rabbpass. Run it with ```rabbpass --help```

