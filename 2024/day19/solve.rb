#! /usr/bin/env ruby

@cache = {}
def solve1
  towels, patterns = get_input
  count = 0
  patterns.each do |pattern|
    count += 1 if is_possible?(pattern, towels)
  end
  count
end

def solve2
  towels, patterns = get_input
  count = 0
  patterns.each do |pattern|
    count += all_possible(pattern, towels)
  end
  count
end

def is_possible?(pattern, towels)
  return true if pattern.empty?

  towels.each do |towel|
    return true if pattern.start_with?(towel) && 
      is_possible?(pattern[towel.length..], towels)
  end
  false
end

def all_possible(pattern, towels)
  return 1 if pattern.empty?
  return @cache[pattern] if @cache.key?(pattern)
  count = 0
  towels.each do |towel|
    if pattern.start_with?(towel)
      count += all_possible(pattern[towel.length..], towels)
    end
    @cache[pattern] = count
  end
  count
end

def get_input
  towels = []
  patterns = []
  File.readlines("input.txt", chomp: true).each do |line|
    if line.include?(",")
      towels = line.split(",").map(&:strip)
    elsif !line.empty?
      patterns << line
    end
  end
  [towels, patterns]
end

puts solve1
puts solve2
