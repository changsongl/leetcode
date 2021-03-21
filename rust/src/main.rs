mod offer;
mod question;


fn main(){
    let mut v = Vec::new();
    v.push(vec![1,1,1]);
    v.push(vec![1,0,1]);
    v.push(vec![1,1,1]);
    let ans = question::q73::Solution::set_zeroes(&mut v);
    println!("{:?}", v);
}