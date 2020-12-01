use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path,
};


// problem https://adventofcode.com/2018/day/1
pub fn fold_seq(freqs: &[&str]) -> i32 {

    let sum = (1..3).fold(0, |sum, x| sum + x);
    sum
}

#[test]
fn test_fold_seq() {
    let mut expected = 3;
    let freqs: &[&str] = &["+1","+1","+1"];
    let mut actual = fold_seq(freqs);
    assert_eq!(actual, expected);
}

fn lines_from_file<P>(filename: P) -> Vec<String>
where
    P: AsRef<Path>,
{
    let file = File::open(filename).expect("failed to read file");
    let buf = BufReader::new(file);
    buf.lines()
       .map(|l| l.expect("could not parse line"))
       .collect()
}

fn parse_str_to_int(str: &str) -> i32 {
    if str.is_empty() {
        0
    } else {
        let num: i32 = str.parse().unwrap();
        num
    }

}

fn main() {
    let lines = lines_from_file("src/q_inputs/q1.in");
    for ln in lines {
       println!("{:?}", parse_str_to_int(&ln));
    }
}
