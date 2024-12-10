#! /bin/ruby


def solve1
  input = get_input
  final = []
  end_ptr = input.length-1
  for i in 0..input.length-1
    final << input[i] unless input[i] == "."
    if input[i] == "."
      while input[end_ptr] == "."
        input[end_ptr] = "_"
        end_ptr -= 1
      end
      break if end_ptr <= i
      final << input[end_ptr]
      input[end_ptr] = "_"
      end_ptr -= 1
    end
  end

  checksum = 0
  final.each_with_index do |f, i|
    checksum += (i * f.to_i)
  end
  checksum
end

def solve2
  disk = get_disk
  disk.files.reverse.each do |file|
    disk.free_space.each do |space|
      break if space.start_idx > file.start_idx
      if space.size >= file.size
        space.size -= file.size
        file.start_idx = space.start_idx
        space.start_idx += file.size
        disk.free_space.delete(space) if space.size == 0
        break
      end
    end
  end
  sum = 0
  disk.files.each do |file|
    for i in file.start_idx..file.start_idx+file.size-1
      sum += i * file.id
    end
  end
  sum
end

def get_input
  input=[]
  file_id = 0
  File.readlines("input.txt", chomp: true).each do |line|
    line.split("").map(&:to_i).each_with_index do |size, i|
      add_file(input, file_id, size) if i % 2 == 0
      add_free_space(input, size) if i % 2 == 1
      file_id += 1 if i % 2 == 0
    end
  end
  input
end

def get_disk
  disk = Disk.new
  file_id = 0
  curr_idx = 0
  File.readlines("input.txt", chomp: true).each do |line|
    line.split("").map(&:to_i).each_with_index do |size, i|
      disk.files << MyFile.new(size, file_id, curr_idx) if i % 2 == 0
      disk.free_space << Block.new(size, curr_idx) if i % 2 == 1
      file_id += 1 if i % 2 == 0
      curr_idx += size
    end
  end
  disk
end

def add_file(input, file_id, size)
  for _ in 0..size-1
    input << file_id.to_s
  end
end

def add_free_space(input, size)
  for _ in 0..size-1
    input << "."
  end
end

class Block
  attr_accessor :size, :start_idx
  def initialize(size, start_idx)
    @size = size
    @start_idx = start_idx
  end
end

class MyFile < Block
  attr_accessor :size, :id, :start_idx
  def initialize(size, id, start_idx)
    super(size, start_idx)
    @id = id
  end
end

class Disk 
  attr_accessor :files, :free_space

  def initialize
    @files = []
    @free_space = []
  end

end

puts solve1 # 6446899523367
puts solve2 # 6478232739671
