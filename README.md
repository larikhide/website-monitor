# website-monitor
Website Monitor is a program to check website availability and gather response time statistics. It scans sites every minute, provides response times for specific sites, and identifies sites with the shortest and longest response times. Administrators receive usage stats.

## How to use 

```
go run cmd/websiteMonitor/main.go
```

**Instructions for Sending Requests to Your Localhost:8000 Service**

To interact with your service running on localhost:8000, follow the instructions below. This guide outlines the six available endpoints and their respective functionalities.

**For Regular Users:**
1. `/ping?url=<domain name without protocol>`: Use this endpoint to retrieve the response time of a specific website. Replace `<domain name without protocol>` with the desired website's domain name (e.g., example.com). The response will provide the response time of the requested website.

2. `/minping`: This endpoint returns the website with the minimum response time. Send a GET request to this endpoint to obtain the website that exhibits the fastest response time.

3. `/maxping`: Similarly, this endpoint returns the website with the maximum response time. Make a GET request to this endpoint to obtain the website with the slowest response time.

**For Administrators:**  
4. `/ping/stats?url=<domain name without protocol>`: Utilize this endpoint to retrieve the number of requests made to ping a specific website. Replace `<domain name without protocol>` with the domain name of the desired website. The response will contain the total number of ping requests made for that website.

5. `/minping/stats`: By sending a GET request to this endpoint, you can acquire the number of requests made to the website with the fastest response time. The response will provide the total count of requests for the website with the minimum response time.

6. `/maxping/stats`: This endpoint returns the number of requests made to the website with the slowest response time. Send a GET request to this endpoint to obtain the total count of requests for the website with the maximum response time.

Ensure that your service is running on localhost:8000 before sending any requests. Use an appropriate tool (e.g., cURL, Postman, browser) to send HTTP requests and receive responses from the specified endpoints.

Please note that the domain name should be provided without the protocol (e.g., http:// or https://).

If you have any further questions or require assistance, please don't hesitate to reach out.
