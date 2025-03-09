# Assignment 1 – Country Information and Population API

**[English](#english) | [Norsk](#norsk)**

---

<a id="english"></a>

## English

### Overview
This is a REST web application built with Go. It retrieves country information and historical population data from external APIs, and provides a status endpoint for monitoring the availability of these external services.

**Key Features:**
1. **Country Info Endpoint**  
   - Returns general country information including name, continent, population, languages, borders, flag, capital, and an optional limited list of cities.
2. **Population Data Endpoint**  
   - Returns historical population data for a given country and calculates the average population over a specified year range.
3. **Status Endpoint**  
   - Reports the health status (HTTP status codes) of external APIs, along with service version and uptime.

### Requirements
- [Go](https://golang.org/) (1.18 or later recommended)
- Environment variables set in a `.env` file (see [Setup](#setup) below)
- Internet connection to reach external APIs

### External APIs
1. **REST Countries API**  
   - Hosted at: `http://129.241.150.113:8080/v3.1/`
2. **CountriesNow API**  
   - Hosted at: `http://129.241.150.113:3500/api/v0.1/`




