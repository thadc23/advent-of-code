#! /usr/bin/env ruby

class Computer
  attr_accessor :a, :b, :c, :program, :output

  def initialize(a, b, c, program)
    @a = a
    @b = b
    @c = c
    @program = program
    @output = []
  end

  def to_s
    "a: #{a}, b: #{b}, c: #{c}, program: #{program.join(",")}, output: #{output.join(",")}"
  end

  def reset(a)
    @output = []
    @a = a
    @b = 0
    @c = 0
  end
end

def solve(computer)
  i = 0
  while i+1 < computer.program.length
    case computer.program[i]
    when 0
      computer.a /= 2**combo_operand(computer, computer.program[i+1])
    when 1
      computer.b ^= computer.program[i+1]
    when 2
      computer.b = combo_operand(computer, computer.program[i+1]) % 8
    when 3
      i = computer.program[i+1]-2 if computer.a != 0
    when 4
      computer.b ^= computer.c
    when 5
      computer.output << combo_operand(computer, computer.program[i+1]) % 8
    when 6
      computer.b = computer.a / 2**combo_operand(computer, computer.program[i+1])
    when 7
      computer.c = computer.a / 2**combo_operand(computer, computer.program[i+1])
    end
    i += 2
  end
end

def solve2(computer)
  i = computer.program.length
  a = 0
  while i >= 0
    for j in 0..63
      computer.reset(a+j)
      solve(computer)
      if computer.output == computer.program[i..-1]
        a += j
        break
      end
    end
    a *= 8 if computer.output.length != computer.program.length
    i -= 1
  end
  computer.a = a
end

def combo_operand(computer,value)
  case value
  when 0, 1, 2, 3
    value
  when 4
    computer.a
  when 5
    computer.b
  when 6
    computer.c
  end
end

def get_input
  a,b,c, program = nil, nil, nil, []
  File.readlines("test.txt", chomp: true).each do |line|
    if line.start_with?("Register A:")
      a = line.split(":").last.strip.to_i
    elsif line.start_with?("Register B:")
      b = line.split(":").last.strip.to_i
    elsif line.start_with?("Register C:")
      c = line.split(":").last.strip.to_i
    elsif line.start_with?("Program")
      program = line.split(":").last.strip.split(",").map(&:to_i)
    end
  end
  Computer.new(a, b, c, program)
end

computer = get_input
solve(computer)
puts computer.output.join(",") # 6,4,6,0,4,5,7,2,7

computer = get_input
solve2(computer)
puts computer.a # 164541160582845