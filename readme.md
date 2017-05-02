# Mock server


This app will mock http api server, listen default on :7890, it accept any post/get

## Run


To quick run a mock server

```bash
docker run --restart=always -d --name ms -v $(pwd)/www:/www -p 7890:7890 netroby/mock-server 
```

make sure ./www exists, and mock-server will try to find exists file to handler api request.

such as ./www/api/version.json will handle as http://example.com/api/version.json

The default MIME type for response, will be application/json


