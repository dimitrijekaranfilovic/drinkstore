#[macro_use] extern crate rocket;
mod tools;
mod traits;
mod responders;

use rocket::response::content;

//ERROR CATCHERS
#[catch(404)]
fn catch_not_found() -> content::RawJson<String> {
    content::RawJson(serde_json::to_string(&tools::StatusMessage{message: "Not found.".to_string(), status_code: 404}).unwrap())
}

#[catch(500)]
fn catch_internal_server_error() -> content::RawJson<String> {
    content::RawJson(serde_json::to_string(&tools::StatusMessage{message: "Internal server error.".to_string(), status_code: 500}).unwrap())
}



//OPTIONS METHOD HANDLERS
#[options("/history-for-user/<_>")]
fn options_user_history() -> responders::OptionsResponder{
    let response = responders::OptionsResponder;
    response
}

#[options("/user-comment-and-grade/<_..>")]
fn options_user_can_comment_and_grade() -> responders::OptionsResponder{
    let response = responders::OptionsResponder;
    response
}


#[options("/")]
fn options_create_purchase() -> responders::OptionsResponder {
    let response = responders::OptionsResponder;
    response
}

#[options("/most-sold/<_..>")]
fn options_get_most_sold()-> responders::OptionsResponder {
    let response = responders::OptionsResponder;
    response
}






//FUNCTIONAL METHODS

//user
#[get("/history-for-user/<user_id>")]
fn user_history(user_id: i32) -> responders::OptionsResponderWithContentStatus200 {
    std::thread::spawn(move || {
        responders::add_content_to_options_responder_status200(tools::get_user_purchase_history(user_id).unwrap())
    }).join().unwrap()
}


//comment i drink servisi
#[get("/user-comment-and-grade?<user_id>&<drink_id>")]
fn user_can_comment_and_grade(user_id: i32, drink_id: i32) -> responders::OptionsResponderWithContentStatus200 {
    std::thread::spawn(move || {
        responders::add_content_to_options_responder_status200(tools::get_user_purchase_count_for_drink(user_id, drink_id).unwrap())
    }).join().unwrap()
}

//user
#[post("/", format="json", data="<purchase_creation_dto>")]
fn create_purchase(purchase_creation_dto: rocket::serde::json::Json<tools::PurchaseCreationDTO>) -> responders::OptionsResponderWithContentStatus201 {
    std::thread::spawn(move || {
        responders::add_content_to_options_responder_status201(tools::create_purchase(purchase_creation_dto).unwrap())

    }).join().unwrap()
}


//svi
#[get("/most-sold?<period>")]
fn get_most_sold(period: String) -> responders::OptionsResponderWithContentStatus200 {
 std::thread::spawn(move || {
    responders::add_content_to_options_responder_status200(tools::get_most_sold_drinks(period).unwrap())
    }).join().unwrap()
}




#[launch]
fn rocket() -> _ {
    std::thread::spawn(|| {
        tools::create_database().unwrap();
    }).join().expect("Thread panicked.");
    rocket::build()
    .mount("/api/purchases", routes![user_can_comment_and_grade, create_purchase, user_history, get_most_sold])
    .mount("/api/purchases", routes![options_user_can_comment_and_grade, options_create_purchase, options_user_history, options_get_most_sold])
    .register("/api/purchases", catchers![catch_not_found, catch_internal_server_error])
}

