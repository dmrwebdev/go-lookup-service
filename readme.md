# Go Lookup Service

Just playing around with Go for the first time. Decided to recreate an exercise I had done in the past fetching IP
address from an api, caching them for future requests and allow viewing the history of searched IP's.

## Usage
Start the HTTP server with ```go run main.go``` or by running the binary using ```./golookupservice```.

The default port assigned is ```3333``` however you may change that by modifying the argument passed to ```.NewServer``` in ```main.go```, and rebuilding the binary if using that.

| Choice             | Endpoint                                                      |
| ------------------ | ------------------------------------------------------------- |
| IP City and State: | http://localhost:3333/lookup/{ipAddress}                      |
| Cached IP's        | http://localhost:8015?city={city}&country={country}           |

#### Checking an IP address:

To find the city and country of an IP, start the application and simply visit http://localhost:3333/lookup/{ipAddressOfYourChoice}. The server will first check the cache for the value and if not present query GeoJS for the IP's information.

If you would like to view all the previously searched and cached IP address, simply vist http://localhost:3333. If you would like to query by a specific city, country, or both you may include them as query parameters.

- If the port variable has been modified, swap out the port above for the one you have chosen

Todo:
  - Testing
  - Further refactors
