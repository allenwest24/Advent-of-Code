use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::HashMap;

fn main() { 
    let filename = "src/D4.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    // Data sturcture to hold attributions for calculations later.
    let mut guards : HashMap<&str, i64> = HashMap::new();
    let mut minutes : HashMap<&str, Vec<i64>> = HashMap::new();
    let mut ordered_times : Vec<i64> = Vec::new();
    let mut ordered_actions : Vec<String> = Vec::new();

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (_index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        let year : i64 = line[1..5].parse().unwrap();
        let month : i64 = line[6..8].parse().unwrap();
        let day : i64 = line[9..11].parse().unwrap();
        let hour : i64 = line[12..14].parse().unwrap();
        let minute : i64 = line[15..17].parse().unwrap();
        let action = line[19..line.len()].to_string();
        let full_time : i64 = (year * 100000000) + (month * 1000000) + (day * 24 * 60) + (hour * 60) + minute;
        if ordered_times.len() == 0 {
            ordered_times.push(full_time);
            ordered_actions.push(action);
        } else {
            for ii in 0..ordered_times.len() {
                if full_time < ordered_times[ii] {
                    // insert before in time
                    ordered_times.insert(ii, full_time);
                    // insert before in actions
                    ordered_actions.insert(ii, action);
                    break;
                } else if ii == ordered_times.len() - 1 {
                    ordered_times.push(full_time);
                    ordered_actions.push(action);
                    break;
                }
            }
        }
    }
    
    // Parse ordered data into maps.
    let mut current_guard : &str = "";
    let mut sleep = 0;
    for aa in 0..ordered_actions.len() {
        let mut current_minutes : Vec<i64> = Vec::new();
        if ordered_actions[aa].contains("Guard") {
            current_guard = ordered_actions[aa].split(" ").collect::<Vec<&str>>()[1];
        } else if ordered_actions[aa].contains("falls asleep") {
            sleep = ordered_times[aa];
        } else if ordered_actions[aa].contains("wakes up") {
            if sleep == 0 {
                // If events are out of order.
                println!("{} {}", ordered_actions[aa-1], ordered_actions[aa]);
            }
            let mut to_add = 0;
            if guards.contains_key(&current_guard) {
                to_add = *guards.get(&current_guard).unwrap();
                let temp = &*minutes.get(&current_guard).unwrap();
                for m in temp {
                    current_minutes.push(*m);
                }
            }
            guards.insert(current_guard, to_add + ordered_times[aa] - sleep);
            for ii in sleep..ordered_times[aa] {
                current_minutes.push(ii % 60);
            }
            minutes.insert(current_guard, current_minutes);
            sleep = 0;
        }
    }

    // Find sleepiest guard.
    let mut most_sleep : i64 = 0;
    let mut worst_employee = "";

    for (key, val) in guards {
        if val > most_sleep {
            most_sleep = val;
            worst_employee = key;
        }
    }

    // Calculate sleepiest minute.
    let worst = &*minutes.get(&worst_employee).unwrap();
    let mut highest_freq = 0;
    let mut best_minute = 0;
    for entry in 0..worst.len() {
        let mut counter = 0;
        for others in entry..worst.len() {
            if worst[entry] == worst[others] {
                counter += 1;
            }
        }
        if counter > highest_freq {
            highest_freq = counter;
            best_minute = worst[entry];
        }
    }

    // Calculate the final code.
    let bad_employee_code : i64 = worst_employee[1..].parse().unwrap();
    println!("{}", best_minute * bad_employee_code);
}
