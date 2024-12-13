#! /bin/ruby

class Machine
  attr_accessor :a, :b, :prize

  def initialize(a, b, prize)
    @a = a
    @b = b
    @prize = prize
  end
end

def solve1
  input = get_input
  cost = 0
  input.each do |machine|
    cost += get_cost_cramer(machine)
  end
  cost
end

def solve2
  input = get_input
  cost = 0
  input.map { |m| m.prize = [m.prize[0] + 10000000000000, m.prize[1] + 10000000000000] }
  input.each do |machine|
    cost += get_cost_cramer(machine)
  end
  cost
end

def get_input
  input = []
  ax = 0
  ay = 0
  bx = 0
  by = 0
  prizex = 0
  prizey = 0
  File.readlines("input.txt", chomp: true).each do |line|
    if line.include?("Button A")
      ax = line.split(",")[0].split("+")[1].to_i
      ay = line.split(",")[1].split("+")[1].to_i
    elsif line.include?("Button B")
      bx = line.split(",")[0].split("+")[1].to_i
      by = line.split(",")[1].split("+")[1].to_i
    elsif line.include?("Prize")
      prizex = line.split(",")[0].split("=")[1].to_i
      prizey = line.split(",")[1].split("=")[1].to_i
      input << Machine.new([ax, ay], [bx, by], [prizex, prizey])
    end
  end
  input
end

def get_cost_cramer(machine)
  # a1*x + b1*y = c1
  # a2*x + b2*y = c2
  # https://en.wikipedia.org/wiki/Cramer%27s_rule#Explicit_formulas_for_small_systems
  
  a1 = machine.a[0]
  b1 = machine.b[0]
  c1 = machine.prize[0]

  a2 = machine.a[1]
  b2 = machine.b[1]
  c2 = machine.prize[1]

  denominator = a1*b2 - b1*a2
  return 0 if denominator == 0

  x_numerator = c1*b2 - b1*c2
  y_numerator = a1*c2 - c1*a2
  return 0 if x_numerator % denominator != 0 || y_numerator % denominator != 0

  x = x_numerator / denominator
  y = y_numerator / denominator

  return 3*x + y
end

puts solve1
puts solve2