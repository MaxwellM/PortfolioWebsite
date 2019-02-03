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

go install
PortfolioWebsite
