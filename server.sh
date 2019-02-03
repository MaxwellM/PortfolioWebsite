#!/bin/bash

set -e
set -x

export PGUSER=postgres
export PGPASSWORD=Maxwellmorin1
export PGPORT=5432
export PGDATABASE=portfoliowebsite_db

cd web
./buildwebsite.py
cd ..

#needs absolutes or else it hates life
/root/bin/go get
/root/bin/go run runTestServer.go

#go install
#/root/bin/PortfolioWebsite
