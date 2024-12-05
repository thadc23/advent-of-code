

def solve1
  rules, updates = get_input
  valid_updates, _ = separate_updates(rules, updates)
  sum(valid_updates)
end

def solve2
  rules, updates = get_input
  _, invalid_updates = separate_updates(rules, updates)
  sum(invalid_updates)
end

def get_input
  rules = {}
  updates = []

  File.readlines("input.txt").each do |line|
    line = line.strip
    if line.include?('|')
      key, value = line.split('|')
      if rules[key.to_i].nil?
        rules[key.to_i] = []
      end
      rules[key.to_i].push(value.to_i)
    elsif line.include?(',')
      updates << line.split(',').map(&:to_i)
    end
  end
  return rules, updates
end

def separate_updates(rules, updates)
  valid_updates = []
  invalid_updates = []
  for i in 0..updates.length-1
    pages_seen = []
    valid = true
    for j in 0..updates[i].length-1
      pages_seen << updates[i][j]
      if rules[updates[i][j]] != nil && pages_seen & rules[updates[i][j]] != []
        valid = false
        fix_update(updates[i], pages_seen & rules[updates[i][j]], j)
      end
    end
    valid_updates << updates[i] if valid
    invalid_updates << updates[i] if !valid
  end
  return valid_updates, invalid_updates
end

def fix_update(update, intersection, current_idx)
  new_spot = find_lowest_index(update, intersection)
  update.insert(new_spot, update.delete_at(current_idx))
end

def find_lowest_index(update, intersection)
  lowest = 0
  for i in 0..update.length-1
    if intersection.include?(update[i])
      lowest = i
      break
    end
  end
  lowest
end

def sum(updates)
  sum = 0
  updates.each do |update|
    sum += update[update.length/2]
  end
  sum
end

puts solve1
puts solve2