#! /bin/bash
curl https://api.github.com/users | jq ' .[] | select( .login == "projetfulle" )'
