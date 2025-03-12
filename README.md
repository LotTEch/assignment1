# Assignment 1 - CountryInfo REST API

This is a project for a Go-based REST service that meets the following requirements:

- **/countryinfo/v1/info/{:two_letter_country_code}{?limit=10}**  
  Returns general country information (name, continents, population, languages, borders, flag, capital city, and a list of cities).
  - `limit` is optional and restricts the number of cities.

- **/countryinfo/v1/population/{:two_letter_country_code}{?limit=startYear-endYear}**  
  Returns historical population data for the country, as well as an average value.  
  - `startYear` and `endYear` are optional. If not specified, all available years are returned.

- **/countryinfo/v1/status/**  
  Shows the status of dependent external services (CountriesNow and RestCountries) and the uptime of our service.

## Setup

1. Clone/copy the project.
2. (Optional) Add a `.env` file with `PORT=8080` if you want to change the port.
3. Run `go mod tidy` to download any missing dependencies.
4. Start the service with `go run main.go`.

## Usage

- **GET** `/countryinfo/v1/info/`  
  Displays a help text.
- **GET** `/countryinfo/v1/info/no`  
  Shows general info about Norway (ISO2 code: NO).
  - Optional: `?limit=5` to retrieve only 5 cities.

- **GET** `/countryinfo/v1/population/no`  
  Shows all available historical population data for Norway.  
  - Optional: `?limit=2010-2015` to only get data for the period 2010–2015.

- **GET** `/countryinfo/v1/status/`  
  Displays a JSON containing the status of CountriesNow, RestCountries, the service version, and uptime.

## Examples of PowerShell Commands

Below are some examples of how you can use PowerShell to send requests to the API and format the response:

## powershell
### Fetch general info about Norway with a limit of 10 cities
Invoke-WebRequest -Uri "http://localhost:8080/countryinfo/v1/info/no?limit=10" | ConvertFrom-Json | Format-List
