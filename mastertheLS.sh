#!/bin/bash
ls -p -A -t --ignore={.,..} | tr '\n' ',' | sed 's/,$//'