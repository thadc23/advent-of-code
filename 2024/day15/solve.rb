#! /bin/ruby


def solve1
  grid, moves = get_input
  x, y = 0, 0
  grid.each_with_index do |row, i|
    grid[i].each_with_index do |cell, j|
      if cell == "@"
        x, y = j, i
        break
      end
    end
  end

  moves.each do |move|
    case move
    when "^"
      y -= 1 if move_up(grid, x, y)
    when "v"
      y += 1 if move_down(grid, x, y)
    when "<"
      x -= 1 if move_left(grid, x, y)
    when ">"
      x += 1 if move_right(grid, x, y)
    end
  end
  print_grid(grid)
  total = 0
  grid.each_with_index do |row, i|
    grid[i].each_with_index do |cell, j|
      total += (100 * i) + j if cell == "O"
    end
  end
  total
end

def solve2
  grid, moves = get_input
  new_grid = Array.new(grid.length) { Array.new(grid[0].length * 2, ".") }
  grid.each_with_index do |row, i|
    grid[i].each_with_index do |cell, j|
      if cell == "#"
        new_grid[i][j*2] = "#"
        new_grid[i][(j*2)+1] = "#"
      end
      if cell == "O"
        new_grid[i][j*2] = "["
        new_grid[i][(j*2)+1] = "]"
      end
      if cell == "@"
        new_grid[i][j*2] = "@"
      end
    end
  end

  x, y = 0, 0
  grid.each_with_index do |row, i|
    grid[i].each_with_index do |cell, j|
      if cell == "@"
        x, y = j, i
        break
      end
    end
  end

  moves.each do |move|
    case move
    when "^"
      y -= 1 if move_up(new_grid, x, y)
    when "v"
      y += 1 if move_down(new_grid, x, y)
    when "<"
      x -= 1 if move_left(new_grid, x, y)
    when ">"
      x += 1 if move_right(new_grid, x, y)
    end
  end
  print_grid(new_grid)
  total = 0
  new_grid.each_with_index do |row, i|
    new_grid[i].each_with_index do |cell, j|
      total += (100 * i) + j if cell == "["
    end
  end
  total
end

def move_up(grid, x, y)
  i = 1
  while y-i >0 && grid[y - i][x] != "." && grid[y-i][x] != "#" do
    i+=1
  end
  if y-i >0 && grid[y - i][x] == "."
    while i > 0 do
      temp = grid[y - i][x]
      grid[y - i][x] = grid[y - i + 1][x]
      grid[y - i + 1][x] = temp
      i -= 1
    end
    return true
  end
  return false
end

def move_down(grid, x, y)
  i = 1
  while y+i < grid.length && grid[y + i][x] != "." && grid[y+i][x] != "#" do
    i+=1
  end
  if y+i < grid.length && grid[y + i][x] == "."
    while i > 0 do
      temp = grid[y + i][x]
      grid[y + i][x] = grid[y + i - 1][x]
      grid[y + i - 1][x] = temp
      i -= 1
    end
    return true
  end
  return false
end

def move_left(grid, x, y)
  i = 1
  while x-i > 0 && grid[y][x - i] != "." && grid[y][x-i] != "#" do
    i+=1
  end
  if x-i > 0 && grid[y][x - i] == "."
    while i > 0 do
      temp = grid[y][x - i]
      grid[y][x - i] = grid[y][x - i + 1]
      grid[y][x - i + 1] = temp
      i -= 1
    end
    return true
  end
  return false
end

def move_right(grid, x, y)
  i = 1
  while x+i < grid[y].length && grid[y][x + i] != "." && grid[y][x+i] != "#" do
    i+=1
  end
  if x+i < grid[y].length && grid[y][x + i] == "."
    while i > 0 do
      temp = grid[y][x + i]
      grid[y][x + i] = grid[y][x + i - 1]
      grid[y][x + i - 1] = temp
      i -= 1
    end
    return true
  end
  return false
end

def get_input
  grid, moves = [], []
  File.readlines("sample_large.txt", chomp: true).each do |line|
    if line.include? "#"
      grid.push(line.split(""))
    elsif !line.empty?
      line.split("").each do |char|
        moves.push(char)
      end
    end
  end
  [grid, moves]
end

def print_grid(grid)
  grid.each do |row|
    puts row.join
  end
  puts
end

puts solve1 
puts solve2