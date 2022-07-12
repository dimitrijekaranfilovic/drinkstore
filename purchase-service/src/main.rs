#[macro_use] extern crate rocket;
mod tools;
mod traits;

use rocket::request::Request;
use rocket::response::{self, Response, content, status};
use rocket::response::Responder;


use rocket::http::Status;
// use rocket::fairing::{Fairing, Info, Kind};
// use rocket::http::Header;


//ERROR CATCHERS
#[catch(404)]
fn catch_not_found() -> content::RawJson<String> {
    content::RawJson(serde_json::to_string(&tools::StatusMessage{message: "Not found.".to_string(), status_code: 404}).unwrap())
}

#[catch(500)]
fn catch_internal_server_error() -> content::RawJson<String> {
    content::RawJson(serde_json::to_string(&tools::StatusMessage{message: "Internal server error.".to_string(), status_code: 500}).unwrap())
}



struct OptionsResponder;

impl<'r> Responder<'r, 'static> for OptionsResponder {
    fn respond_to(self, _: &'r Request<'_>) -> response::Result<'static> {
        Ok(Response::build()
        .status(Status::NoContent)
        .raw_header("Access-Control-Allow-Origin", "http://localhost:9090")
        .raw_header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")
        .raw_header("Access-Control-Allow-Headers", "*")
        .raw_header("Access-Control-Allow-Credentials", "true")
        .finalize())
    }
}

//OPTIONS METHOD HANDLERS
#[options("/history-for-user/<_>")]
fn options_user_history() -> OptionsResponder{
    let response = OptionsResponder;
    response
}

#[options("/user-comment-and-grade/<_..>")]
fn options_user_can_comment_and_grade() -> OptionsResponder{
    let response = OptionsResponder;
    response
}


#[options("/")]
fn options_create_purchase() -> OptionsResponder {
    let response = OptionsResponder;
    response
}

#[options("/most-sold/<_..>")]
fn options_get_most_sold()-> OptionsResponder {
    let response = OptionsResponder;
    response
}




//FUNCTIONAL METHODS

//user
#[get("/history-for-user/<user_id>")]
fn user_history(user_id: i32) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::get_user_purchase_history(user_id).unwrap())
    }).join().unwrap()
}


//comment i drink servisi
#[get("/user-comment-and-grade?<user_id>&<drink_id>")]
fn user_can_comment_and_grade(user_id: i32, drink_id: i32) -> content::RawJson<String> {
    std::thread::spawn(move || {
        content::RawJson(tools::get_user_purchase_count_for_drink(user_id, drink_id).unwrap())
    }).join().unwrap()
}


//user
#[post("/", format="json", data="<purchase_creation_dto>")]
fn create_purchase(purchase_creation_dto: rocket::serde::json::Json<tools::PurchaseCreationDTO>) -> status::Custom<content::RawJson<String>> {
    std::thread::spawn(move || {
        status::Custom(Status::Created, content::RawJson(tools::create_purchase(purchase_creation_dto).unwrap()))

    }).join().unwrap()
}

//svi
#[get("/most-sold?<period>")]
fn get_most_sold(period: String) -> content::RawJson<String> {
 std::thread::spawn(move || {
        content::RawJson(tools::get_most_sold_drinks(period).unwrap())
    }).join().unwrap()
}




#[launch]
fn rocket() -> _ {
    // let cors = CORS{};
    std::thread::spawn(|| {
        tools::create_database().unwrap();
    }).join().expect("Thread panicked.");
    rocket::build()
    .mount("/api/purchases", routes![user_can_comment_and_grade, create_purchase, user_history, get_most_sold])
    .mount("/api/purchases", routes![options_user_can_comment_and_grade, options_create_purchase, options_user_history, options_get_most_sold])
    .register("/api/purchases", catchers![catch_not_found, catch_internal_server_error])
    // .attach(cors)
}

