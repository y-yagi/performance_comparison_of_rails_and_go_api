class User < Sequel::Model
end

User.plugin :json_serializer
