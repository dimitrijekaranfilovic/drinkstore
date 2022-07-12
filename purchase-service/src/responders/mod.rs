use rocket::response::{self, Response, content};
use rocket::http::{Status, Header};
use rocket::request::Request;


pub struct OptionsResponder;

#[derive(Responder)]
#[response(status = 200, content_type="json")]
pub struct OptionsResponderWithContentStatus200 {
    pub content: content::RawJson<String>,
    pub access_control_allow_origin: Header<'static>,
    pub access_control_allow_methods: Header<'static>,
    pub access_control_allow_headers: Header<'static>,
    pub access_control_allow_credentials: Header<'static>
}


#[derive(Responder)]
#[response(status = 201, content_type="json")]
pub struct OptionsResponderWithContentStatus201 {
    pub content: content::RawJson<String>,
    pub access_control_allow_origin: Header<'static>,
    pub access_control_allow_methods: Header<'static>,
    pub access_control_allow_headers: Header<'static>,
    pub access_control_allow_credentials: Header<'static>
}

impl<'r> response::Responder<'r, 'static> for OptionsResponder {
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

pub fn add_content_to_options_responder_status200(content: String) -> OptionsResponderWithContentStatus200 {
     OptionsResponderWithContentStatus200{
            content: content::RawJson(content),
            access_control_allow_credentials: Header::new("Access-Control-Allow-Origin", "http://localhost:9090"),
            access_control_allow_headers: Header::new("Access-Control-Allow-Headers", "*"),
            access_control_allow_methods: Header::new("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD"),
            access_control_allow_origin: Header::new("Access-Control-Allow-Origin", "http://localhost:9090")
        }
  
}

pub fn add_content_to_options_responder_status201(content: String) -> OptionsResponderWithContentStatus201 {
     OptionsResponderWithContentStatus201{
            content: content::RawJson(content),
            access_control_allow_credentials: Header::new("Access-Control-Allow-Origin", "http://localhost:9090"),
            access_control_allow_headers: Header::new("Access-Control-Allow-Headers", "*"),
            access_control_allow_methods: Header::new("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD"),
            access_control_allow_origin: Header::new("Access-Control-Allow-Origin", "http://localhost:9090")
        }
  
}