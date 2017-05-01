# Mock server


This app will mock http api server, listen default on :7890, it accept any post/get

## Run

To quick run a mock server

```bash
docker run --restart=always -d --name ms -p 7890:7890 netroby/mock-server 
```

## Road map

1. provided toml configure syntax
2. allow  key=/path/to/api.json style configure
3. if request match "/api/call" configure section, it will try to load /path/to/api.json file


