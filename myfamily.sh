curl -s -O https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
cat all.json | jq ".[] | select(.id==$HERO_ID) | .connections | .relatives " | tr -d '"'
