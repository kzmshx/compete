use std::io::{BufRead, stdin, stdout, Write};
use std::str::FromStr;

fn solve<T: BufRead, U: Write>(input: &mut T, output: &mut U) {
    let mut buf = String::new();

    input.read_line(&mut buf).unwrap();
    let n = i32::from_str(buf.trim()).unwrap();

    for i in (0..=n).rev() {
        output.write_fmt(format_args!("{}\n", i)).unwrap();
    }
}

fn main() {
    solve(&mut stdin().lock(), &mut stdout())
}

#[cfg(test)]
mod tests {
    use std::io::{BufReader, BufWriter};

    use super::*;

    #[test]
    fn solve_1() {
        let input = String::from(r"5
");
        let expected = String::from(r"5
4
3
2
1
0
");

        let mut in_buf = BufReader::new(input.as_bytes());
        let mut out_buf = BufWriter::new(vec![]);
        solve(&mut in_buf, &mut out_buf);

        assert_eq!(expected, String::from_utf8(out_buf.into_inner().unwrap().clone()).unwrap());
    }
}
