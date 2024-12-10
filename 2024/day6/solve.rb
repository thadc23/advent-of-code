#! /bin/ruby

def solve1
  input = get_input
  walk(input)
end

def solve2
  input = get_input
  path, _ = walk(input)
  count = 0
  path.each do |i, j, _|
    next if input[i][j] == "^"
    input[i][j] = "#"
    _, is_loop = walk(input)
    count += 1 if is_loop
    input[i][j] = "."
  end
  count
end

def walk(input)
  i,j = find_start(input)
  direction = "up"
  visited = Set.new
  while i > 0 && i < input.length-1 && j > 0 && j < input[i].length-1 && !visited.include?([i, j, direction])
    visited.add([i, j, direction])

    case direction
    when "up"
      if input[i-1][j] == "#"
        direction = "right"
      else
        i -= 1
      end
    when "down"
      if input[i+1][j] == "#"
        direction = "left"
      else
        i += 1
      end
    when "left"
      if input[i][j-1] == "#"
        direction = "up"
      else
        j -= 1
      end
    when "right"
      if input[i][j+1] == "#"
        direction = "down"
      else
        j += 1
      end
    end
  end

  return visited.map{ |i, j, _| [i,j]}.to_set, visited.include?([i, j, direction])
end

def get_input
  input = []
  File.readlines("input.txt", chomp: true).each do |line|
    input << line.split("")
  end
  input
end

def find_start(input)
  for i in 0..input.length-1
    for j in 0..input[i].length-1
      if input[i][j] == "^"
        return i,j
      end
    end
  end
  [0,0]
end

puts "Part 1: #{solve1[0].to_set.length+1}" # 4758
puts "Part 2: #{solve2}" # 1670