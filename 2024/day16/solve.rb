#! /bin/ruby

class Direction
  NORTH = 0
  EAST = 1
  SOUTH = 2
  WEST = 3
end
def solve1
  map = get_input
  x,y = 1, map.length-2
  scores = walk(map, x, y, Direction::EAST, 0, [])
  min = scores.min
  scores = walk2(map, x, y, Direction::EAST, 0, {}, [], min)
  puts min
  seats = Set.new

  scores[min].each do |path|
    path.each do |x,y|
      map[y][x] = "O"
      seats.add([x,y])
    end
  end
  puts seats.length
end

def get_input
  map = []
  File.readlines("input.txt", chomp: true).each do |line|
    map << line.split("")
  end
  map
end

@seen = {}

def walk(map, x, y, dir, score, scores)
  if map[y][x] == "E"
    scores << score
  end
  if map[y][x] != "#" && map[y][x] != "E"
    @seen[[x,y]] = score if @seen[[x,y]].nil? || @seen[[x,y]] > score
    new_x, new_y = find_new_xy(x, y, dir)
    walk(map, new_x, new_y, dir, score+1, scores) if !@seen[[new_x,new_y]] || @seen[[new_x,new_y]] > score+1
    new_x, new_y = find_new_xy(x, y, (dir+1) % 4)
    walk(map, new_x, new_y, (dir+1) % 4, score+1001, scores) if !@seen[[new_x,new_y]] || @seen[[new_x,new_y]] > score+1001
    new_x, new_y = find_new_xy(x, y, (dir+3) % 4)
    walk(map, new_x, new_y, (dir+3) % 4, score+1001, scores) if !@seen[[new_x,new_y]] || @seen[[new_x,new_y]] > score+1001
  end

  scores
end

def walk2(map, x, y, dir, score, scores, path, min)
  path << [x,y]
  if map[y][x] == "E"
    scores[score] = [] if scores[score].nil?
    scores[score] << path
    return scores
  end
  if score > min
    return scores
  end
  @seen[[x,y, dir]] = score if @seen[[x,y, dir]].nil? || @seen[[x,y, dir]] > score
  new_x, new_y = find_new_xy(x, y, dir)
  walk2(map, new_x, new_y, dir, score+1, scores, path.clone, min) if (!@seen[[new_x,new_y, dir]] || @seen[[new_x,new_y, dir]] >= score+1) && map[new_y][new_x] != "#"
  next_dir = (dir+1) % 4
  new_x, new_y = find_new_xy(x, y, next_dir)
  walk2(map, new_x, new_y, next_dir, score+1001, scores, path.clone, min) if (!@seen[[new_x,new_y, next_dir]] || @seen[[new_x,new_y, next_dir]] >= score+1001) && map[new_y][new_x] != "#"
  next_dir = (dir+3) % 4
  new_x, new_y = find_new_xy(x, y, next_dir)
  walk2(map, new_x, new_y, next_dir, score+1001, scores, path.clone, min) if (!@seen[[new_x,new_y, next_dir]] || @seen[[new_x,new_y, next_dir]] >= score+1001) && map[new_y][new_x] != "#"

  scores
end

def find_new_xy(x, y, dir)
  case dir
  when Direction::NORTH
    return x, y-1
  when Direction::EAST
    return x+1, y
  when Direction::SOUTH
    return x, y+1
  when Direction::WEST
    return x-1, y
  end
end

solve1