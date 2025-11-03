#!/bin/bash -ex
cd "$(dirname "$0")"
go install tool
mage build coverage
cat coverage.out