# internship-web-app

## Installation

Application requires Docker and Docker-Compose to be installed.

For MongoDB initialization please download MongoDB sample database content from:
https://github.com/neelabalan/mongodb-sample-dataset

Then store the folders with .json files into /mongo-sample folder of the root directory of the project nearside init-mongo.sh file.

To start up the project run the following command from within the project directory:

```sh
docker-compose up -d
```

Your docker volume for storing database state is located at ~/mongodata folder on your local machine, for configuration revise volumes in docker-compose.yml file.