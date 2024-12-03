#! /bin/ruby

def solve1()
  reports = get_input()
  count = 0

  reports.each do |report|
    valid, _ = check_report(report)
    count += 1 if valid
  end
  count
end

def solve2()
  reports = get_input()
  count = 0

  reports.each do |report|
    valid, _ = check_report(report)
    count += 1 if valid
    for i in 0..report.length-1
      report_copy = report.dup
      report_copy.delete_at(i)
      valid, _ = check_report(report_copy)
      count += 1 if valid
      break if valid
    end if !valid
  end
  count
end

def get_input()
  # read input.txt
  # Each line is a report
  # each report has levels that are numbers and are separated by whitespace
  # Read the file and create a 2d array of reports
  
  reports = []
  File.readlines("input.txt", chomp: true).each do |line|
    reports << line.split(" ").map(&:to_i)
  end
  reports
end

def check_report(report)
  ascending = report[1] > report[0]
  descending = report[1] < report[0]  
  for i in 0..report.length-2
    if (report[i] - report[i+1]).abs > 3 || (ascending && report[i] > report[i+1]) || (descending && report[i] < report[i+1]) || report[i] == report[i+1]
      return false, i
    end
  end
  return true, -1
end

puts solve1
puts solve2