#! /bin/ruby
#

MAX_X = 101
MAX_Y = 103
class Robot
  attr_accessor :x, :y, :vx, :vy

  def initialize(x, y, vx, vy)
    @x = x
    @y = y
    @vx = vx
    @vy = vy
  end

  def to_s
    "Robot at (#{x}, #{y}) with velocity (#{vx}, #{vy})"
  end
end

def solve1
  robots = get_input
  for _ in 0..99 do
    robots.each do |robot|
      robot.x = (robot.x + robot.vx) % MAX_X
      robot.y = (robot.y + robot.vy) % MAX_Y
    end
  end

  get_safety_score(robots)
end

def get_safety_score(robots)
  quadrants = {1=> 0, 2=> 0, 3=> 0, 4=> 0}
  robots.each do |robot|
    q, included = find_quadrant(robot)
    quadrants[q] += 1 if included
  end
  total = 1
  quadrants.values.map { |v| total *= v }
  total
end

def solve2
  robots = get_input
  grid = Array.new(MAX_Y) { Array.new(MAX_X, ".") }
  robots.each do |robot|
    grid[robot.y][robot.x] = "#"
  end
  min = 100000000000000
  min_sec = MAX_X * MAX_Y
  for i in 0..MAX_X * MAX_Y do
    ss = get_safety_score(robots)
    if ss < min
      min = ss if get_safety_score(robots) < min
      min_sec = i
      # print_grid(grid)
    end
    robots.each do |robot|
      grid[robot.y][robot.x] = "."
      robot.x = (robot.x + robot.vx) % MAX_X
      robot.y = (robot.y + robot.vy) % MAX_Y
      grid[robot.y][robot.x] = "#"
    end
  end
  min_sec
end

def print_grid(grid)
  grid.each do |row|
    puts row.join
  end
  puts
end

def find_quadrant(robot)
  if robot.x == MAX_X / 2 || robot.y == MAX_Y / 2
    return 0, false
  end

  if robot.x < MAX_X / 2
    if robot.y < MAX_Y / 2
      return 1, true
    else
      return 3, true
    end
  else
    if robot.y < MAX_Y / 2
      return 2, true
    else
      return 4, true
    end
  end
end

def get_input
  robots = []
  File.readlines("input.txt", chomp: true).each do |line|
    start = line.split(" ")[0].tr("p=", "")
    velocity = line.split(" ")[1].tr("v=", "")
    x = start.split(",")[0].to_i
    y = start.split(",")[1].to_i
    vx = velocity.split(",")[0].to_i
    vy = velocity.split(",")[1].to_i
    robot = Robot.new(x, y, vx, vy)
    robots.push(robot)
  end
  robots
end

puts solve1
puts solve2