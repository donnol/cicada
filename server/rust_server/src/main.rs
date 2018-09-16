fn main() {
    // 1
    // let str = "hello";

    // 2
    {
        let str = String::from("hello");
        println!("{}", str) // 块内最后一个分号可以忽略，作为返回值
    } // rust run drop() to free str.

    // 3
    {
        let s1 = String::from("hello");
        let s2 = s1; // move
        // println!("{}", s1) // use of moved value: `s1`
        // println!("{}, {}", s1, s2); // use of moved value: `s1`
        println!("{}", s2)
    } // free one memory twice.

    // 4
    {
        // heap
        let str = String::from("hello");
        takes_ownership(str);
        // println!("{}", str); // use of moved value: `str`

        // stack
        let i = 1;
        makes_copy(i);
        println!("{}", i)
    }

    // 5
    {
        let s1 = gives_ownership();
        println!("{}", s1);

        let s2 = String::from("hello");
        let s3 = takes_and_gives_back(s2);
        println!("{}", s3);
    }

    // 6
    {
        let str = String::from("hello");
        let l = calculate_length(&str); // 传入引用(允许你使用值但不获取其所有权)
        println!("{}, {}", l, str); // str依然可用
    }

    // 7
    {
        let str = String::from("hello, world.");
        let str = first_word(&str[..str.len()-1]);
        println!("{}", str);

        let a = [1, 2, 3, 4, 5];
        let a = array_slice(&a[..a.len()-1]);
        println!("{:?}", a)
    }
}

fn takes_ownership(str: String) {
    println!("{}", str)
}

fn makes_copy(i: i32) {
    println!("{}", i)
}

fn gives_ownership() -> String {
    let str = String::from("hello");
    str
}

fn takes_and_gives_back(str: String) -> String {
    str
}

fn calculate_length(str: &String) -> usize { // 将引用作为函数参数称为借用（borrowing）
    // str.push_str(", world."); // 修改借用的变量时会报错：cannot borrow immutable borrowed content `*str` as mutable
    // 要想修改，必须在参数声明里添加mut：&mut String
    // 可变引用有一个很大的限制：在特定作用域中的特定数据有且只有一个可变引用
    // 在拥有不可变引用的同时不能再拥有可变引用
    // -> 读写: 可以有多个读同时存在，但不允许读写共存
    str.len()
}

fn first_word(s: &str) -> &str {
    s
}

fn array_slice(a: &[i32]) -> &[i32] {
    a
}