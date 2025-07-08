#!/bin/bash

# Create the target directory if it doesn't exist
mkdir -p ./tmp/config

echo "[copy-config] Copying config.yaml to ./tmp/config/"

# Copy the config.yaml file to the tmp/config directory
cp ./config/config.yaml ./tmp/config/

echo "[copy-config] Copy config.yaml to ./tmp/config/ successfully."