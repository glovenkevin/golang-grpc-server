# Example Golang gRPC Server with Protobuf

This is my learning project for golang gRPC server with protobuf.

## Set Up Project

1. Init project module with this command

    ```bash
    make init 
    ```

2. Create protobuf file with this command
    ```bash
    make proto-gen
    ```

3. Run project with this command
    ```bash
    make run
    ```


## Set Up MongoDB Cluster With Docker

I presume you already have a docker installed in your PC.

1. Init docker needs
    ```bash
    make docker-prepare
    ```

2. Generate key for mongo cluster
    ```bash
    make mongo-gen-key
    ```

3. Start docker-compose
    ```bash
    docker-compose up -d
    ```

4. Initate the cluster
    ```bash
    make mongo-cluster-init
    ```

5. Add these dns on your hosts file (E.g. linux on /etc/hosts)
    ```
    127.0.0.1	mongo-c1
    127.0.0.1	mongo-c2
    ```

6. Use These connection string to connect from mongodb compas
    ```
    mongodb://root:password@localhost:27017,localhost:27018,localhost:27019/?readPreference=secondaryPreferred&ssl=false
    ```