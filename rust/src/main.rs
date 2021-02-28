mod offer;
use offer::q17::print_numbers;

fn main(){
    let v = print_numbers(1);
    for x in v {
        println!("n: {}", x)
    }
}