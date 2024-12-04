#! /bin/ruby

def solve1
  input = get_input
  count = 0
  for i in 0..input.length-1
    for j in 0..input[i].length-1
      if input[i][j] == "X"
        count += find_xmas(input, i, j)
      end
    end
  end
  count
end

def solve2
  input = get_input
  count = 0
  for i in 0..input.length-1
    for j in 0..input[i].length-1
      if input[i][j] == "A"
        count += find_mas(input, i, j)
      end
    end
  end
  count
end

def get_input
  input = []
  File.open("input.txt", "r") do |f|
    f.each_line do |line|
      input << line.split("")
    end
  end
  input
end

def find_mas(input, i, j)
  count = 0
  count += 1 if check_X(input, i, j)
  count
end

def find_xmas(input, i, j)
  count = 0
  count += 1 if check_horizontal(input, i, j)
  count += 1 if check_vertical(input, i, j)
  count += 1 if check_diagonal(input, i, j)
  count += 1 if check_horizontal_backwards(input, i, j)
  count += 1 if check_vertical_backwards(input, i, j)
  count += 1 if check_diagonal_backwards(input, i, j)
  count += 1 if check_diagonal_up_right(input, i, j)
  count += 1 if check_diagonal_down_left(input, i, j)
  count
end

def check_horizontal(input, i, j)
  return false if j > input[0].length-4
  return input[i][j+1] == "M" && input[i][j+2] == "A" && input[i][j+3] == "S"
end

def check_vertical(input, i, j)
  return false if i > input.length-4
  return input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S"
end

def check_diagonal(input, i, j)
  return false if i > input.length-4 || j > input[0].length-4
  return input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S"
end

def check_horizontal_backwards(input, i, j)
  return false if j < 3
  return input[i][j-1] == "M" && input[i][j-2] == "A" && input[i][j-3] == "S"
end

def check_vertical_backwards(input, i, j)
  return false if i < 3
  return input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S"
end

def check_diagonal_backwards(input, i, j)
  return false if i < 3 || j < 3
  return input[i-1][j-1] == "M" && input[i-2][j-2] == "A" && input[i-3][j-3] == "S"
end

def check_diagonal_up_right(input, i, j)
  return false if i < 3 || j > input[0].length-4
  return input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S"
end

def check_diagonal_down_left(input, i, j)
  return false if i > input.length-4 || j < 3
  return input[i+1][j-1] == "M" && input[i+2][j-2] == "A" && input[i+3][j-3] == "S"
end

def check_X(input, i, j) 
  return false if i == 0 || i == input.length-1 || j == 0 || j == input[0].length-1

  return ((input[i-1][j-1] == "M" && input[i+1][j+1] == "S") || 
    (input[i+1][j+1] == "M" && input[i-1][j-1] == "S")) &&
    ((input[i-1][j+1] == "M" && input[i+1][j-1] == "S") || 
    (input[i+1][j-1] == "M" && input[i-1][j+1] == "S"))
end


puts solve1
puts solve2