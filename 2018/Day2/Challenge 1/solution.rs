use std::fs::File;
use std::io::{BufRead, BufReader};
use character_frequency::*;
use std::collections::HashMap;

fn main() { 
    let filename = "src/D2.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    let mut twos = 0;
    let mut threes = 0;

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        let frequency_map = character_frequencies(&line);
        if frequency_map.values().any(|&val| val == 2) {twos += 1;}
        if frequency_map.values().any(|&val| val == 3) {threes += 1;}
    }
    println!("{}", twos * threes);
}
