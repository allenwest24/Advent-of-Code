use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    // Variables needed.
    let mut total = 0;
    let mut total_tracker = Vec::new();

    loop {    
        let filename = "src/D1C2.txt";
        // Open the file in read-only mode (ignoring errors).
        let file = File::open(filename).unwrap();
        let reader = BufReader::new(file);

        // Read the file line by line using the lines() iterator from std::io::BufRead.
        for (index, line) in reader.lines().enumerate() {
            let line = line.unwrap(); // Ignore errors.
            total += line.parse::<i32>().unwrap();
            if total_tracker.contains(&total) {
                println!("Hit the following total twice before all other: {}", total);
                std::process::exit(1);
            }
            total_tracker.push(total);
            println!("{}", total);
        }
    }
}
