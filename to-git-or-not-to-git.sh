#! /bin/bash
id=$(curl https://api.github.com/users/projetfulle | jq '.id')
echo $id

