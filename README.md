# SHORTY

Tool to shorten long URL to short one.



## TECH STACK

- Golang
- React
- Mongo
- Echo (Golang Web Framework)


### Feature

- [x] Login Auth via JWT
- [x] URL Shortener
- [x] Custom URL for logged in user
- [ ] URL Click analytics


### Before run production or development build

> rename config/env.sample to config/.env
> update environment variable like Mongo URI

## How to run

>Docker Required

```sh
docker build -t shorty .
```

```sh
docker run -p 1234:1234 -it --rm --name myapp shorty
```

> Open localhost:1234 in browser


## For Development

- Golang 1.13+
- Node 10+
- Yarn 1.x

> Install golang dependencies

Run following commands from project root directory
```sh
go get -d -v ./...
go install -v ./...
```


To run Golang Server
```
go run .
```

> golang server run at localhost:1234

For React Web Build

> Install web dependencies

Run following commands from web directory
```sh
npm install # or yarn
```

for development react site
```sh
npm start #or yarn start
```

> react dev server run at localhost:3002

> API_BASE for dev is set to localhost:1234 default configuration is enough to run a working app. You can change web config from `web/config/constant.dev.js` 




## Links
* [https://wadehrarshpreet.com](https://wadehrarshpreet.com)
* [LinkedIn](https://www.linkedin.com/in/wadehrarshpreet/)
* [Twitter](https://twitter.com/wadehrarshpreet/)
