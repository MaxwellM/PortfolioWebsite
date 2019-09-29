#!/bin/bash

set -e
set -x

export GOPATH=/root
export PGUSER=postgres
export PGPASSWORD=Maxwellmorin1
export PGPORT=5432
export PGDATABASE=portfoliowebsite_db
export PATH=$PATH:/usr/local/go/bin
export GOOGLE_APPLICATION_CREDENTIALS="googleKey.json"

whoami

cd web
./buildwebsite.py
cd ..

#needs absolutes or else it hates life
cd /root/src/PortfolioWebsite
go get
#/root/src/PortfolioWebsite/ 
go run main.go

#go install
#/root/bin/PortfolioWebsite
