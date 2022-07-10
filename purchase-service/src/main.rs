//TODO: 
//1.istorija kupovine, 
//2.najprodavaniji artikli, 
//3.mogucnost komentarisanja i ocjenjivanja => DODATO
//4.sama kupovina => DODATO

#[macro_use] extern crate rocket;
mod tools;

use rocket::response::{content};
//use rocket::serde::json::Json;
// use rocket::{routes, Rocket, Build};
// use rocket::response::content::{self};
// use std::{thread};


#[get("/history-for-user/<user_id>")]
fn user_history(user_id: i32) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::get_user_purchase_history(user_id).unwrap())
    }).join().unwrap()
}


//ostali servisi
#[get("/user-comment-and-grade?<user_id>&<drink_id>")]
fn user_can_comment_and_grade(user_id: i32, drink_id: i32) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::get_user_purchase_count_for_drink(user_id, drink_id).unwrap())
    }).join().unwrap()
}

#[post("/", format="json", data="<purchase_creation_dto>")]
fn create_purchase(purchase_creation_dto: rocket::serde::json::Json<tools::PurchaseCreationDTO>) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::create_purchase(purchase_creation_dto).unwrap())
    }).join().unwrap()
}


#[get("/most-sold?<period>")]
fn get_most_sold(period: String) -> content::RawJson<String> {
 std::thread::spawn(move || {
        content::RawJson(tools::get_most_sold_drinks(period).unwrap())
    }).join().unwrap()
}




#[launch]
fn rocket() -> _ {
    std::thread::spawn(|| {
        tools::create_database().unwrap();
    }).join().expect("Thread panicked.");
    rocket::build().mount("/api/purchases", routes![user_can_comment_and_grade, create_purchase, user_history, get_most_sold])
}

