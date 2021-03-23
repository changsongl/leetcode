#[derive(Debug, PartialEq, Eq)]
pub enum NestedInteger {
  Int(i32),
  List(Vec<NestedInteger>)
}

pub struct NestedIterator {
    nest_int: Vec<NestedInteger>,
    sub_iterator: Vec<SubIterator>,
    start: bool,
}

struct SubIterator {
    v: &NestedInteger,
    i: i32,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NestedIterator {

    pub fn new(nestedList: Vec<NestedInteger>) -> Self {
        NestedIterator{
            nest_int: nestedList,
            sub_iterator: Vec::new(),
            start: false,
        }
    }

    pub fn next(&mut self) -> i32 {
        1
    }

    pub fn has_next(&mut self) -> bool {
        if self.start && self.sub_iterator.len() == 0{
            false
        }

        if !self.start{
            self.start = true;
            if self.nest_int.len() == 0{
                false
            }

            let int_vec = &self.nest_int[0];
            self.sub_iterator.push(
                SubIterator{
                    v: int_vec,
                    i: 0,
                }
            );
        }

        while self.sub_iterator.len() != 0 {
            let int_vec = self.sub_iterator.pop();
            match int_vec {
                None => false,
                Some(sub_iterator) => (
                    match &sub_iterator.v[i] {
                        NestedInteger::Int => {
                            self.sub_iterator.push(sub_iterator);
                            true
                        }
                        NestedInteger::List(nv) => {
                            if (nv.len() as i32) < sub_iterator.i{
                                self.sub_iterator.push(sub_iterator);
                                self.sub_iterator.push(
                                    SubIterator {
                                        v: &sub_iterator.v[i],
                                        i: i32,
                                    }
                                )
                            }
                        }
                    }
                ),
            }
        }

        true
    }

}