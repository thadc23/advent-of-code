#! /bin/ruby

require "set"

def solve
  input = get_input
  grid = Array.new(71) { Array.new(71, ".") }
  for i in 0..1023
    grid[input[i][1]][input[i][0]] = "#"
  end
  walk(grid, 0, 0)
end

def solve2
  input = get_input
  grid = Array.new(71) { Array.new(71, ".") }
  for i in 0..1023
    grid[input[i][1]][input[i][0]] = "#"
  end
  for i in 1024..input.length-1
    grid[input[i][1]][input[i][0]] = "#"
    steps = walk(grid, 0, 0)
    return [input[i][0], input[i][1]].join(",") if steps == -1
  end
end

def get_input
  input = []
  File.readlines("input.txt", chomp: true).each do |line|
    input << line.split(",").map(&:to_i)
  end
  input
end

def walk(grid, start_x, start_y)
  goal_x, goal_y = 70, 70
  queue = [[start_x, start_y, 0]]  # [x, y, distance]
  visited = Set.new
  visited.add([start_x, start_y])
  
  while !queue.empty?
    x, y, dist = queue.shift
    return dist if x == goal_x && y == goal_y
    
    [[0,1], [1,0], [0,-1], [-1,0]].each do |dx, dy|
      new_x, new_y = x + dx, y + dy
      next if new_x < 0 || new_x > 70 || new_y < 0 || new_y > 70
      next if grid[new_y][new_x] == "#"
      next if visited.include?([new_x, new_y])
      
      visited.add([new_x, new_y])
      queue << [new_x, new_y, dist + 1]
    end
  end
  
  -1  # No path found
end

puts solve
puts solve2
