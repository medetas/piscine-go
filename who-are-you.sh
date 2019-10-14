curl -O https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
cat all.json | jq '.[] | select(.id==70).name' 


