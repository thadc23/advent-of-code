#! /bin/ruby

@seen = {}
def solve(blinks)
  input = get_input
  total = 0
  input.each do |num|
    total += split_stone(num, blinks)
  end
  total
end

def split_stone(num, blinks)
    if blinks == 0
      return 1
    end
    
    if @seen.has_key?([num, blinks])
      return @seen[[num, blinks]]
    end

    total = 0
    if num == 0
      total = split_stone(1, blinks-1)
    elsif num.to_s.length % 2 == 0
      left = num.to_s[0..num.to_s.length/2-1]
      right = num.to_s[num.to_s.length/2..num.to_s.length-1]
      total = split_stone(left.to_i, blinks-1) + split_stone(right.to_i, blinks-1)
    else
      total = split_stone(num*2024, blinks-1)
    end
    @seen[[num,blinks]] = total
    total
end

def get_input
  input = []
  File.readlines("input.txt", chomp: true).each do |line|
    input = line.split(" ").map(&:to_i)
  end
  input
end

puts solve(25) # 183435
@seen.clear
puts solve(75) # 218279375708592