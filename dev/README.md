# Instructions to start development setup

## Network up
```
./copyCC.sh & ./minifab up -i 2.4 -o warehousemen.onesheds.com -e true -s couchdb -n core -p '"Init"'
```

## Test Query
```
./minifab query -n core -o warehousemen.onesheds.com -p '"GetIdentity"'
```

## Copy chaincode
```
./copyCC.sh
```

## Update chaincode
```
./minifab ccup -v 1.1 -n core -p "Init"
```


## Open warehousemen state database
```
http://localhost:7008/_utils/#
```
