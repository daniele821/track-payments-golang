# keep track of payments

## requirements

- python3 (`sqlite3` module is in standard library!!!)

## general ideas

- use python for the backend, and html + css for frontend and js for interactions
    - python: create a server which allows POST parameters, and uses them to make sqlite queries, and then answer via json
    - javascript: via `fetch()` make requests to python backend, and use it to get database data
    - NOTE: make python handle 2 servers: 1 for db interactions with js, and one to actually host html server (or compact the two into a single server)

- find way to decrypt sqlite file when starting python server, and decrypting it when closing the server
