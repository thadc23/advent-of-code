#! /bin/ruby

def solve
  input = get_input
  seen = Set.new
  plots = []
  for i in 0..input.length-1
    for j in 0..input[i].length-1
      next if seen.include?([i,j])
      plots << walk(input, i, j, seen)
    end
  end
  [plots.map { |fences, plots| fences*plots }.sum, plots.map { |_, plots, posts| plots*posts }.sum]
end

def walk(input, i, j, seen)
  fences = 0
  plots = 1
  plant = input[i][j]
  posts = 0
  seen.add([i,j])
  if !same_plant_up(input, i, j)
    fences += 1
  elsif !seen.include?([i-1, j])
    f, c, p, _ = walk(input, i-1, j, seen)
    fences += f
    plots += c
    posts += p
  end
  if !same_plant_down(input, i, j)
    fences += 1
  elsif !seen.include?([i+1, j])
    f, c, p, _ = walk(input, i+1, j, seen)
    fences += f
    plots += c
    posts += p
  end
  if !same_plant_left(input, i, j)
    fences += 1
  elsif !seen.include?([i, j-1])
    f, c, p, _ = walk(input, i, j-1, seen)
    fences += f
    plots += c
    posts += p
  end
  if !same_plant_right(input, i, j)
    fences += 1
  elsif !seen.include?([i, j+1])
    f, c, p, _ = walk(input, i, j+1, seen)
    fences += f
    plots += c
    posts += p
  end
  posts += find_posts(input, i, j)
  [fences, plots, posts, plant]
end

def same_plant_up(input, i, j)
  i > 0 && input[i-1][j] == input[i][j]
end

def same_plant_down(input, i, j)
  i < input.length-1 && input[i+1][j] == input[i][j]
end

def same_plant_left(input, i, j)  
  j > 0 && input[i][j-1] == input[i][j]
end

def same_plant_right(input, i, j)
  j < input[i].length-1 && input[i][j+1] == input[i][j]
end

def same_plant_up_left(input, i, j)
  i > 0 && j > 0 && input[i-1][j-1] == input[i][j]
end

def same_plant_up_right(input, i, j)
  i > 0 && j < input[i].length-1 && input[i-1][j+1] == input[i][j]
end

def same_plant_down_left(input, i, j)
  i < input.length-1 && j > 0 && input[i+1][j-1] == input[i][j]
end

def same_plant_down_right(input, i, j)
  i < input.length-1 && j < input[i].length-1 && input[i+1][j+1] == input[i][j]
end

def find_posts(input, i, j)
  posts = 0

  if !same_plant_left(input, i, j) && !same_plant_up(input, i, j)
    posts += 1
  end
  if !same_plant_left(input, i, j) && !same_plant_down(input, i, j)
    posts += 1
  end
  if !same_plant_right(input, i, j) && !same_plant_up(input, i, j)
    posts += 1
  end
  if !same_plant_right(input, i, j) && !same_plant_down(input, i, j)
    posts += 1
  end

  posts += 1 if same_plant_left(input, i, j) && same_plant_up(input, i, j) && !same_plant_up_left(input, i, j)
  posts += 1 if same_plant_left(input, i, j) && same_plant_down(input, i, j) && !same_plant_down_left(input, i, j)
  posts += 1 if same_plant_right(input, i, j) && same_plant_up(input, i, j) && !same_plant_up_right(input, i, j)
  posts += 1 if same_plant_right(input, i, j) && same_plant_down(input, i, j) && !same_plant_down_right(input, i, j)
  
  posts
end

def get_input
  input = []
  File.readlines("input.txt", chomp: true).each do |line|
    input << line.split("")
  end
  input
end

part1, part2 = solve
puts part1 # 1381056
puts part2 # 834828