# earl

earl is a tiny app for shortening URLs. It's built for fun and to keep my Go skills fresh.
It exposes a couple of REST endpoints that handle the functionality. The shortened links are stored in Redis.  

It also serves a very small web page on the root url (as seen on [earl.sh](https://earl.sh)).

## Getting Started

With the project being mostly for fun, the most important information is how to run it locally and play with it.

### Run locally

#### Prerequisites
- [Docker](https://www.docker.com)

#### Start the development server

```
docker-compose up
```
By default the app should start on [localhost:8000](localhost:8000).  
The server is rebuilding on change by using [reflex](https://github.com/cespare/reflex), so feel free to break stuff and watch the the app panic.

## Usage

Aside from visiting [earl.sh](earl.sh) there are two other ways to play with it:

### curl

The easiest way to test it out is by using the REST API directly.
```
$ curl -X POST -H "Content-Type: application/json" --data '{"url": "https://earl.sh"}' https://earl.sh/api/shorten
```

### Command-line

A quick improvisation on how one can use it instantly on the command-line.

```
$ go install github.com/shiftingphotons/earl@latest
$ earl shorten --url=https://earl.sh
```

If you install and build the executable from this repository, it points by default to `https://earl.sh`, but if you have deployed the app yourself, and want to shorten links from the command line - you have to point the executable by setting a `EARL_HOST` environment variable that points to your domain, or passing it as a `--host` flag when running the `shorten` command.

```
$ earl shorten --url=https://earl.sh --host=http://example.com
```

## Make it yours
The app is lightweight and easy to deploy it yourself. It makes a couple of assumptions which can be changed and tailored to your needs.

### Redis
The links are kept in Redis, so it's essential to set a `REDIS_HOST` envrionment variable that points to your storage.

### Life cycle
By default the links are kept for 72 hours because I don't want to provide guarantees on [earl.sh](https://earl.sh) because it's a showcase site, not built for actual users.

This can be changed by setting up a `EARL_LIFECYCLE` environment variables and give it a number of hours for which the links should be accessible.
Another possibility for configuring this is by passing a `--lifecycle` flag to the executable, and giving it an amount of hours.
