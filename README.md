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
curl -X POST -H "Content-Type: application/json" --data '{"url": "https://earl.sh"}' https://earl.sh/api/shorten
```

### Command-line

This one is work in progress...

```
go install github.com/shiftingphotons/earl@latest
TODO
```
