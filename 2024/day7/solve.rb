#! /bin/ruby

def solve1
  results, numbers = get_input
  count = 0
  results.each_with_index do |result, i|
    count += result if find_ways(numbers[i][1..-1], result, numbers[i][0])
  end
  count
end

def solve2
  results, numbers = get_input
  count = 0
  results.each_with_index do |result, i|
    count += result if find_ways2(numbers[i][1..-1], result, numbers[i][0])
  end
  count
end

def get_input
  results = []
  numbers = []
  File.readlines('input.txt', chomp: true).map(&:strip).each do |line|
    results << line.split(":")[0].to_i
    numbers << line.split(":")[1].strip.split(" ").map(&:to_i)
  end
  [results, numbers]
end

def find_ways(numbers, target, current)
  return false if current > target
  return true if current == target && numbers.empty?
  return false if numbers.empty?
  return find_ways(numbers[1..-1], target, current + numbers[0]) || 
    find_ways(numbers[1..-1], target, current * numbers[0])
end

def find_ways2(numbers, target, current)
  return false if current > target
  return true if current == target && numbers.empty?
  return false if numbers.empty?
  return find_ways2(numbers[1..-1], target, current + numbers[0]) || 
    find_ways2(numbers[1..-1], target, current * numbers[0]) ||
    find_ways2(numbers[1..-1], target, (current.to_s + numbers[0].to_s).to_i)
end

puts solve1
puts solve2