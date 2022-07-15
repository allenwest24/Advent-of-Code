use std::fs::File;
use std::io::{BufRead, BufReader};

// Helper method that counts the number of differences in words by location and char.
fn count_diffs(a: &str, b: &str) -> i32 {
    let mut out = 0;
    let vec_a : Vec::<char> = a.chars().collect();
    let vec_b : Vec::<char> = b.chars().collect();
    for ii in 0..vec_a.len() {
        if vec_a[ii] != vec_b[ii] {out += 1;}
    }
    println!("{}", out);
    out
}

fn main() { 
    let filename = "src/D2.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    // Tracker.
    let mut seen = Vec::new();

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        if seen.len() == 0 {
            seen.push(line);
        } else {
            for ii in &seen {
                if count_diffs(&*line, &*ii) == 1 {
                    println!("{} and {}", line, ii);
                    std::process::exit(1);
                }
            }
            seen.push(line);
        }
    }
}
