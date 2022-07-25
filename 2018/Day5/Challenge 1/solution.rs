use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() { 
    let filename = "src/D5.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        let mut as_chars : Vec<char> = line.chars().collect();
        println!("{:?}", as_chars);
        loop {
            let mut temp : Vec<char> = Vec::new();
            let mut to_remove = 0;
            let mut prev = ' ';
            for ii in 0..as_chars.len() {
                if as_chars[ii] != prev && (as_chars[ii] == prev.to_uppercase().next().unwrap() || as_chars[ii] == prev.to_lowercase().next().unwrap()) {
                    temp.pop();
                    prev = ' ';
                } else {
                    temp.push(as_chars[ii]);
                    prev = as_chars[ii];
                }
            }
            as_chars = temp;
            println!("{}", as_chars.len());
        }
    }
}
