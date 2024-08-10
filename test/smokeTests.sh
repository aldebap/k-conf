#!  /usr/bin/ksh

./bin/kconf -help

./bin/kconf status

#   services
./bin/kconf add service --name=Produtos --url=http://192.168.68.107:8080/api/v1/produto

./bin/kconf query service --id=86e5be98-39fe-4035-b182-afc63553a027

./bin/kconf list service

#   routes
./bin/kconf add route --name=Produto --protocols=http --methods=GET,POST --paths=/gwa/v1/produtos --service-id=86e5be98-39fe-4035-b182-afc63553a027

./bin/kconf query route --id=3cb55349-1f3c-45da-9461-4457a7351ab6

./bin/kconf list route
