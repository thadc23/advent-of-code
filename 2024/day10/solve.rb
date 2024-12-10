#! /bin/ruby

class TrailFinder
  attr_reader :trails

  def initialize
    @trails = []
  end

  def solve
    input = get_input

    input.each_with_index do |line, i|
      line.each_with_index do |char, j|
        if char == 0
          find_trails([i,j], input, i, j)  
        end
      end
    end
  end

  def get_input
    input = []
    File.readlines("input.txt", chomp: true).each do |line|
      input << line.split("").map(&:to_i) 
    end
    input
  end

  def find_trails(start, input, i, j)
    @trails << [start, [i,j]] if input[i][j] == 9
    if i-1 >= 0 && input[i-1][j] == input[i][j] + 1
      find_trails(start, input, i-1, j)
    end
    if i+1 < input.length && input[i+1][j] == input[i][j] + 1
      find_trails(start, input, i+1, j)
    end
    if j-1 >= 0 && input[i][j-1] == input[i][j] + 1
      find_trails(start, input, i, j-1)
    end
    if j+1 < input[i].length && input[i][j+1] == input[i][j] + 1
      find_trails(start, input, i, j+1)
    end
  end
end


trail_finder = TrailFinder.new
trail_finder.solve
puts trail_finder.trails.uniq.length
puts trail_finder.trails.length