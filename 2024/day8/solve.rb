#! /bin/ruby

class ResonanceFinder
  attr_reader :input, :rows, :cols

  def initialize(file_path)
    @input, @rows, @cols = get_input(file_path)
  end

  def solve1
    locations = Set.new
    input.each do |key, value|
      value.each_with_index do |location, i|
        l = find_locations(input, location, value[i+1..-1])
        l.each { |loc| locations.add(loc) }
      end
    end
    locations.length
  end
  
  def solve2
    locations = Set.new
    input.each do |key, value|
      value.each_with_index do |location, i|
        l = find_all_locations(input, location, value[i+1..-1])
        l.each { |loc| locations.add(loc) }
      end
    end
    locations.length
  end

  private

  def get_input(file_path)
    input = {}
    rows, cols = 0, 0
    File.readlines(file_path, chomp: true).each_with_index do |line, i|
      rows += 1
      cols = line.length
      line.split("").each_with_index do |col, j|
        input[col] = [] if input[col].nil?
        input[col] << [i, j]
      end
    end
    input.delete(".")
    [input, rows, cols]
  end

  def find_locations(input, location, remaining)
    locations = []
    remaining.each do |loc|
      d_i, d_j = loc[0] - location[0], loc[1] - location[1]
      locations << [location[0] - d_i, location[1] - d_j] if in_grid?(location[0] - d_i, location[1] - d_j)
      locations << [loc[0] + d_i, loc[1] + d_j] if in_grid?(loc[0] + d_i, loc[1] + d_j)
    end
    locations
  end
  
  def find_all_locations(input, location, remaining)
    locations = []
    remaining.each do |loc|
      d_i, d_j = loc[0] - location[0], loc[1] - location[1]
      locations << location
      locations << loc
      loops = 1
      while in_grid?(location[0] - (d_i*loops), location[1] - (d_j*loops))
        locations << [location[0] - (d_i*loops), location[1] - (d_j*loops)]
        loops += 1
      end
      loops = 1
      while in_grid?(loc[0] + (d_i*loops), loc[1] + (d_j*loops))
        locations << [loc[0] + (d_i*loops), loc[1] + (d_j*loops)]
        loops += 1
      end
    end
    locations
  end
  
  def in_grid?(i, j)
    i >= 0 && i < rows && j >= 0 && j < cols
  end

end

solver = ResonanceFinder.new("input.txt")
puts solver.solve1
puts solver.solve2