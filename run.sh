# excecute before run
# ./bin/rails s -e production
# ./go_api
# RACK_ENV=production be puma -C puma.rb

siege -r 100 -c 50 -b -q  http://localhost:3000/users > rails_api_result.log 2>&1
siege -r 100 -c 50 -b -q  http://localhost:4000/users > go_api_result.log 2>&1
siege -r 100 -c 50 -b -q  http://localhost:9292/users > roda_sequel_result.log 2>&1
