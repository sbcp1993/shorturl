# Introduction 
Shorturl is an api developed in Go to short long urls. 

# Getting Started
steps to run 
- Clone this repository
- cd shorturl
- docker-compose up -d (to up shorturl and it's dependent services)

# Build and Test
When the service is up you can do POST request to generate a short url

example:
```
url : http://localhost:8080/generate

```
request body:
```json
{
	"url": "https://docs.microsoft.com/en-us/azure/devops/report/powerbi/data-connector-connect?view=azure-devops"
}

```
This request return a response with the short url like the following 
```json
{
    "url":"http://localhost:8080/e2bc8299"
}
```
you can now use this short url instead of your long url.
