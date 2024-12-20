#! /bin/ruby

def solve
  grid, moves = get_input
  output = []
  [grid, grid.gsub(/[#\.O@]/, "#" => "##", "." => "..", "O" => "[]", "@" => "@.")].each do |current_grid|
    grid_hash = {}
    current_grid.split("\n").each_with_index do |row, j|
      row.chars.each_with_index do |c, i|
        grid_hash[Complex(i, j)] = c
      end
    end

    pos = grid_hash.find { |p, c| c == '@' }&.first

    moves.tr("\n", '').chars.each do |m|
      dir = { '<' => -1, '>' => 1, '^' => -Complex(0, 1), 'v' => Complex(0, 1) }[m]
      copy = grid_hash.dup

      if move(grid_hash, pos, dir)
        pos += dir
      else
        grid_hash = copy
      end
    end

    ans = grid_hash.sum do |pos, c| 
      c.match?(/[O\[]/) ? pos : 0
    end
    
    output.push((ans.real + ans.imag * 100).to_i)
  end
  output
end

def get_input
  grid, moves = File.read('input.txt').split("\n\n")
  [grid, moves]
end

def move(grid, p, d)
  p += d
  if [
    grid[p] != '[' || (move(grid, p + 1, d) && move(grid, p, d)),
    grid[p] != ']' || (move(grid, p - 1, d) && move(grid, p, d)),
    grid[p] != 'O' || move(grid, p, d),
    grid[p] != '#'
  ].all?
    grid[p], grid[p - d] = grid[p - d], grid[p]
    true
  end
end

answer = solve
puts answer[0]
puts answer[1]