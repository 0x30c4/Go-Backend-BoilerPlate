
<h1 align="center">Go Backend BoilerPlate</h1>
<p align="center">
  <img src="https://img.shields.io/github/languages/top/0x30c4/GhostBin?style=flat-square" alt="Test">
</p>
<p align="justify">
  Simple Golang Backend BoilerPlate code. With the least amount of third party modules. I use this BoilerPlate backend for almost all of my golang backends.
</p>
<p align="center">
  <a href="mailto:sanaf@0x30c4.dev"> sanaf@0x30c4.dev </a>
</p>

## Table of Contents

- [How To Use](#how-to-use)
  - [Basic Usage](#basic-usage)
- [Deployment](#deployment)
  - [Built With](#built-with)
  - [Prerequisites](#prerequisites)
  - [Docker Compose](#docker-compose)
  - [Contributing](#contributing)
  - [License](#license)
- [Test](#test)
  - [Test Coverage](#current-test-coverage)
- [TODO](#todo)
- [Contribution](#contribution)
- [Donate](#donate)

## How To Use

### Basic Usage


find . -type f -exec  sed -i "s/github.com\/0x30c4\/Go-Backend-BoilerPlate/NAME_OF_YOUR_PROJECT/g" {} +

```bash
$ git clone https://github.com/0x30c4/Go-Backend-BoilerPlate.git
```

## Deployment
Want to run a server like this? clone it! Remember centralization is bad.

### Built With.

* [Docker](https://www.docker.com) - Platform and Software Deployment
* [Go](https://go.dev) - Backend Frame-work.
* [Redis](https://redis.io/) - DataStore DataStore

### Prerequisites.

Make sure you have [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git), [make](https://tldp.org/HOWTO/Software-Building-HOWTO-3.html) and [Docker](https://www.docker.com/products/docker-desktop) installed.

### Docker Compose

GhostBin can be easily deployed using Docker Compose. Follow these steps to deploy GhostBin:

1. **Clone Repository**: Clone the Go-Backend-BoilerPlate repository to your box.

2. **Configuration**: Duplicate the `env-example` file and rename it as `.env.dev` for local development or `.env.prod` for the production environment. Customize the contents of these files according to your requirements.

3. **Build**: Build the Docker images for your Project using the provided Makefile command:

    ```bash
    make build
    ```

4. **Development Environment**:

    ```bash
    make up-dev
    ```

5. **Production Environment**:

    ```bash
    make up-prod
    ```

6. **Access Logs**:

    To access logs, you can use:

    ```bash
    make logs
    ```

    To tail logs in real-time:

    ```bash
    make logs-tail
    ```

7. **Additional Commands**:

    - `make down-dev` / `make down-prod`: Shutdown the development/production environment.
    - `make restart-dev` / `make restart-prod`: Restart the development/production environment.
    - `make exec-dev` / `make exec-prod`: Access the shell of the development/production container.

## Test
Unit Test for the redis connection and logger is provided.

## Contributing

Contributions to Go-Backend-BoilerPlate are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on our [GitHub repository](https://github.com/0x30c4/Go-Backend-BoilerPlate).

## License

Go-Backend-BoilerPlate is licensed under the [BSD 3-Clause License](https://github.com/0x30c4/Go-Backend-BoilerPlate/blob/main/LICENSE).


## Contribution
Pull requests are welcome.
For major changes, please open an issue first to discuss what you would like to change.

## Donate
You can support this project via Liberapay.
<br>
<a target="_blank" href="https://liberapay.com/sanaf/donate"><img src="https://img.shields.io/liberapay/gives/1"></a>

Monero wallet address: 83BDAy6tN99PVud2sUnjyoMzsUDdXJCoMjjwJ59cVwPF91RccxLWCVsfD9imMqxUaMhMG1brzuVBeAM4KREUSf9U9efbKx1
