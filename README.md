# Taxi Fare Backend

## Running with Docker
In order to run the taxi fare backend, you need to build a docker image:

### `docker build -t taxi-fare .`

Then you need to run it in a docker container: 

### `docker run -p 8080:8080 taxi-fare`

Once the docker is container is running, you can run the Taxi Fare UI.

## Running without Docker
In order to run the server locally, you need to edit the `data_path` 
config in the `tf.toml` configuration file.

Then you can build:

### `go build .`

And then run the server app:

### `./TaxiFare`

Finally in order to run tests, you can run the following command:

### `go test -v ./endpoints`
