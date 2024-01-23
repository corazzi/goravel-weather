# Goravel Weather

A very simple, very basic weather application.

## Get Started
Clone the repository
```shell
git clone git@github.com:corazzi/goravel-weather.git
```

Install dependencies
```shell
go mod tidy
```

Copy the `.env` file and generate your application key
```shell
cp .env.example .env

go run . artisan key:generate
```

## Configuration
You will need an [OpenWeatherMap API key](https://openweathermap.org/api)

Once you have that, pop it in your `.env` file
```dotenv
OPENWEATHERMAP_API_KEY=XXXXXXXXXXXX
```

## Build & Run

```shell
./build.sh
./goweather
```
You can also run without compiling:

```shell
go run .
```

## Bug Reports
Please file bug reports in `/dev/null`
