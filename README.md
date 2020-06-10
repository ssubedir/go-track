
# About Go Track
API wrapper for package tracking services. 

### Supported Services
Currently supported services
|Service|Support| 
|---|---|
|Canada Post|✔|
|FedEx|✔|
|Purolator|✔|
|United Parcel Service|✔|
|DHL Express|✔|

### Download

Use git clone to get your local copy 
```
git clone https://github.com/ssubedir/go-track
```

### Build

Build API wrapper
```
cd go-track
go build -v main.go 
```

### Serve

Start API server on port 9000
```
./go-track
```

### End points

```
GET
  :9000/track/canadapost/<Tracking Number>
  :9000/track/dhl/<Tracking Number>

POST
  :9000/track/fedex/<Tracking Number>
  :9000/track/purolator/shipment/<Tracking Number>
  :9000/track/purolator/tracking/<Tracking Number>
  :9000/track/ups/<Tracking Number>
```

## Built With

* [GO](https://golang.org/) - Programming language


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/ssubedir/go-track/blob/master/LICENSE) file for details
