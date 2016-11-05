siege -r 100 -c 50 -b -q  http://localhost:3000/users > rails_api_result.log 2>&1
siege -r 100 -c 50 -b -q  http://localhost:4000/users > go_api_result.log 2>&1
