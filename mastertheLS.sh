#!/bin/bash
ls -PAt --ignore={.,..} | tr '\n' ',' | sed 's/,$//'