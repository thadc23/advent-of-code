#! /bin/ruby

def solve1()
  locations_left, locations_right = get_input()

  locations_left.sort!
  locations_right.sort!

  sum = 0
  for i in 0..locations_left.length-1
    sum += (locations_left[i] - locations_right[i]).abs
  end
  sum
end

def solve2()
  locations_left, locations_right = get_input()
  
  locations_right_hash = Hash.new(0)
  locations_right.each do |num|
    locations_right_hash[num] += 1
  end
  sum = 0

  locations_left.each do |num|
    if locations_right_hash[num] > 0
      sum += num * locations_right_hash[num]
    end
  end
  sum
end

def get_input()
  locations_left = []
  locations_right = []
  File.readlines("input.txt", chomp: true).each do |line|
    locations_left << line.split(" ")[0].to_i
    locations_right << line.split(" ")[1].to_i
  end
  return locations_left, locations_right
end

puts solve1
puts solve2