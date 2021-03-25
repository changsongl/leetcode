mod offer;
mod question;

use offer::q22;

#[derive(PartialEq, Eq, Clone, Debug)]
struct A {
    b: String
}

fn main()
{
    let mut a = A{b: String::from("ja")};
    let mut aa = Some(a);
    let mut aaa = aa.as_mut().unwrap();
    aaa.b = String::from("xxxx");
    println!("{:?}, {:?}", aa, aaa);
}