#!/bin/bash

tar -czf input.tar.gz example_inputs

curl -F input=@input.tar.gz http://localhost:8888
