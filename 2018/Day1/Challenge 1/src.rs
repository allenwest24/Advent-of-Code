use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "src/D1C1.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    // Variables needed.
    let mut total = 0;

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        total += line.parse::<i32>().unwrap();
        println!("{}", total);
    }
}
