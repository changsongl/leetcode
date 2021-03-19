mod offer;


fn main(){
    // let v = offer::q17::print_numbers(1);
    // for x in v {
    //     println!("n: {}", x)
    // }
    let v = vec![1,2,3,4];
    let s = offer::q21::Solution::exchange(v);
    println!("{:?}", s);
}