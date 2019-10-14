
curl -s https://api.github.com/users/medetas | grep '"id"' | cut -d: -f2 | cut -d, -f1 | cut -d" " -f2
