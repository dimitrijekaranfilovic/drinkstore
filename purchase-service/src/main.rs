//TODO: 
//1.istorija kupovine, 
//2.najprodavaniji artikli, 
//3.mogucnost komentarisanja i ocjenjivanja 
//4.sama kupovina

#[macro_use] extern crate rocket;
mod tools;

use rocket::response::{content};
//use rocket::serde::json::Json;
// use rocket::{routes, Rocket, Build};
// use rocket::response::content::{self};
// use std::{thread};


//#[get("/history-for-user/<user_id>")]


//ostali servisi
#[get("/user-comment-and-grade?<user_id>&<drink_id>")] //treba drink_id i user_id
fn user_can_comment_and_grade(user_id: i32, drink_id: i32) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::get_user_purchase_count_for_drink(user_id, drink_id).unwrap())
    }).join().unwrap()
}



#[launch]
fn rocket() -> _ {
    std::thread::spawn(|| {
        tools::create_database();
    }).join().expect("Thread panicked.");
    rocket::build().mount("/api/purchases", routes![user_can_comment_and_grade])
}

