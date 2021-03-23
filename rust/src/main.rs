mod offer;
mod question;

use question::q341;

fn main(){
    let nestedList= vec![q341::NestedInteger::Int(11)];
    let mut obj = q341::NestedIterator::new(nestedList);
    let ret_1: i32 = obj.next();
    let ret_2: bool = obj.has_next();
}