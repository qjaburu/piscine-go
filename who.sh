#!/bin/bash
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json |jq -r '.[]| select (id==70).name'

