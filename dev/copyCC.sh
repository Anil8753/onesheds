#!/bin/sh

echo "-- Copy started"

rm -rf vars/chaincode/core/go
mkdir -p vars/chaincode/core/
cp -R -v ../chaincode/core/go ./vars/chaincode/core/go

echo "-- Copy ended"
echo ""