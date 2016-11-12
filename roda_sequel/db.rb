require 'sequel'

database_url = ENV.delete('DATABASE_URL') || "postgres:///performance_comparison"
DB = Sequel.connect(database_url)
