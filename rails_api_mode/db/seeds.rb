100.times do |i|
  User.create!(name: "dummy_#{i}", email: "dummy_#{i}@example.com")
end
