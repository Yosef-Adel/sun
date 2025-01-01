# Weather Forecast CLI

This is a simple command-line interface (CLI) application written in Go that fetches and displays weather forecast information for a specified location.

## Features

- Fetches current weather and hourly forecast for a given location
- Displays temperature, conditions, and chance of rain
- Highlights high chances of rain in red
- Supports custom location input

## Prerequisites

- Go 1.x
- Weather API key from [WeatherAPI.com](https://www.weatherapi.com/)

## Installation

1. Clone this repository
2. Run `go mod tidy` to install dependencies

## Configuration

Create a `.env` file in the project root with your Weather API key:

```

WEATHER_API=your_api_key_here

```

## Usage

Run the application with:

1. Build the binary
2. ./sun [location]
   If no location is provided, it defaults to "Cairo".

## Example Output

```

Egypt, Cairo: 25C, Sunny
14:00 - 26C, 0%, Sunny
15:00 - 26C, 0%, Sunny
16:00 - 25C, 0%, Sunny
17:00 - 24C, 0%, Clear
18:00 - 22C, 0%, Clear
...
```

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - For colored console output
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - For loading environment variables

## License

This project is open source and available under the [MIT License](LICENSE).
