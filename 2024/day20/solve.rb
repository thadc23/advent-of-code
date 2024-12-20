require 'set'

MIN_CHEAT_DISTANCE = 2
DIRECTIONS = [[-1, 0], [1, 0], [0, -1], [0, 1]]

def solve1
  track = get_input
  valid_cheats(track, 2)
end

def solve2
  track = get_input
  valid_cheats(track, 20)
end

def valid_cheats(track, max_distance)
  non_cheat = find_path(track)
  non_cheat.to_a.map do |pos, steps|
    valid_cheats_count(track, pos, max_distance, steps, non_cheat)
  end.sum
end

def find_start(track)
  track.each_with_index do |row, y|
    row.each_with_index do |cell, x|
      return [y, x] if cell == 'S'
    end
  end
end

def find_path(track)
  start = find_start(track)
  distances = { start => 0 }
  queue = [start]
  
  until queue.empty?
    pos = queue.shift
    DIRECTIONS.each do |dy, dx|
      next_pos = [pos[0] + dy, pos[1] + dx]
      next if track[next_pos[0]][next_pos[1]] == '#' || distances.key?(next_pos)
      
      distances[next_pos] = distances[pos] + 1
      queue.push(next_pos)
    end
  end
  
  distances
end

def find_path2(track)
  start = find_start(track)
  y, x = start
  path = []
  steps = 0
  while track[y][x] != 'E'
    path.push([y, x, steps])
    DIRECTIONS.each do |dy, dx|
      next_pos = [y + dy, x + dx]
      path.push(next_pos) if is_on_map?(track, next_pos)
    end
  end

  path
end

def valid_cheats_count(track, pos, max_distance, steps, non_cheat)
  count = 0
  cheats = Set.new
  for distance in MIN_CHEAT_DISTANCE..max_distance
    for dy in -distance..distance
      dx = distance - dy.abs
      next_pos = [pos[0] - dy, pos[1] - dx]
      if is_on_map?(track, next_pos) && saves_enough_time?(steps, distance, non_cheat, next_pos)
        cheats.add(next_pos)
      end
      next_pos = [pos[0] + dy, pos[1] + dx]
      if is_on_map?(track, next_pos) && saves_enough_time?(steps, distance, non_cheat, next_pos)
        cheats.add(next_pos)
      end
    end
  end
  cheats.count
end

def saves_enough_time?(steps, distance, non_cheat, next_pos)
  steps + 100 + distance <= (non_cheat[next_pos] || 0)
end

def is_on_map?(track, pos)
  pos[0] >= 0 && pos[1] >= 0 && pos[0] <= track.length - 1 && pos[1] <= track[0].length - 1
end

def get_input
  input = []
  File.readlines("input.txt", chomp: true).each do |line|
    input << line.split('')
  end
  input
end

puts solve1
puts solve2
