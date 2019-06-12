#! /bin/bash
id=$(curl https://api.github.com/users/Mousouba | jq '.id')
echo $id

