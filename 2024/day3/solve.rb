#! /bin/ruby

def solve1
  contents = get_input
  # I need a regex to match the pattern mul(x,y) where x and y are integers from 1 to 3 digits
  # I need to extract the x and y values and multiply them and add them to a sum
  # I need to return the sum
  calculate(contents)
end

def solve2
  contents = get_input
  total = 0
  contents.split(/don't\(\)(.|\n)*?do\(\)/).each do |section|
    # puts "Section: #{section}\n\n"
    total += calculate(section)
  end
  total
end

def calculate(contents)
  sum = 0
  contents.scan(/mul\((\d{1,3}),(\d{1,3})\)/).each do |match|
    sum += match[0].to_i * match[1].to_i
  end
  sum
end

def get_input
  #Read the contents of input.txt into a single String
  File.read("input.txt")
end

puts solve1
puts solve2