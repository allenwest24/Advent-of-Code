use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() { 
    let filename = "src/D3.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    // Build the nasty matrix of values.
    let mut cross = vec!(0; 1000 as usize);
    let mut grid : Vec<Vec<i32>> = vec!(cross; 1000 as usize);

    // Read the file line by line using the lines() iterator from std::io::BufRead.
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        let mut split = line.split(" ").collect::<Vec<&str>>();

        let coords = split[2];
        let coords_split = coords.split(",").collect::<Vec<&str>>();
        let x : i32 = coords_split[0].parse().unwrap();
        let y : i32 = *&coords_split[1][..&coords_split[1].len() - 1].parse().unwrap();

        let size = split[3];
        let size_split = size.split("x").collect::<Vec<&str>>();
        let width : i32= size_split[0].parse().unwrap();
        let height : i32 = size_split[1].parse().unwrap();

        for ii in x..x+width {
            for jj in y..y+height {
                grid[jj as usize][ii as usize] += 1;
            }
        }
    }

    // Read again, this time looking for a claim that overlaps the lonely zone.
    let filename = "src/D3.txt";
    // Open the file in read-only mode (ignoring errors).
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);
    for (index, line) in reader.lines().enumerate() {
        let line = line.unwrap(); // Ignore errors.
        let mut split = line.split(" ").collect::<Vec<&str>>();

        let coords = split[2];
        let coords_split = coords.split(",").collect::<Vec<&str>>();
        let x : i32 = coords_split[0].parse().unwrap();
        let y : i32 = *&coords_split[1][..&coords_split[1].len() - 1].parse().unwrap();

        let size = split[3];
        let size_split = size.split("x").collect::<Vec<&str>>();
        let width : i32= size_split[0].parse().unwrap();
        let height : i32 = size_split[1].parse().unwrap();

        let mut good = true;
        for ii in x..x+width {
            for jj in y..y+height {
                if grid[jj as usize][ii as usize] != 1 {
                    good = false;
                }
            }
        }
        if good {println!("{}", split[0]);}
    }
}
