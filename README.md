# power-factor
### How to deploy
please run --> docker compose up --build
### Tests
run go test ./app
## Sample requests

#### 1 month period
localhost:3003/ptlist?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z

#### 1 year period
localhost:3003/ptlist?period=1y&tz=Europe/Athens&t1=20110214T204603Z&t2=20211115T123456Z

#### 1 day period

localhost:3003/ptlist?period=1d&tz=Europe/Athens&t1=20211014T204603Z&t2=20211115T123456Z

#### 1 hour period

localhost:3003/ptlist?period=1h&tz=Europe/Athens&t1=20211114T204603Z&t2=20211115T123456Z
