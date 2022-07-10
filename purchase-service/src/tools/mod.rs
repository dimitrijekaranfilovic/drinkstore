

use postgres::{Client, NoTls, Error};
use serde::{Serialize, Deserialize};
use chrono::prelude::{DateTime, Utc};


//Model
#[derive(Serialize, Deserialize)]
pub struct PurchaseCreationDTO {
  purchase_items: Vec<PurchaseItemCreationDTO>,
  user_id: i32,
}

#[derive(Serialize, Deserialize)]
pub struct PurchaseItemCreationDTO {
  name: String,
  drink_id: i32,
  num_items: i32,
  unit_price: i32,
}

#[derive(Serialize)]
pub struct StatusMessage {
  message: String,
}

#[derive(Serialize, Deserialize)]
pub struct UserPurchaseCount {
  num_purchases: i64,
}




//Functions
pub fn create_database() -> Result<(), Error> {
    //TODO: dodaj test podatke

    let mut db_client = Client::connect("postgresql://postgres:root@localhost:5432/ntp-purchase-service", NoTls)?; 
    
    db_client.batch_execute("DROP TABLE IF EXISTS purchase_items")?;
    db_client.batch_execute("DROP TABLE IF EXISTS purchases")?;
    
    db_client.batch_execute("
        CREATE TABLE IF NOT EXISTS purchases (
            id SERIAL PRIMARY KEY,
            user_id INTEGER
        )
        

    ")?;
    db_client.batch_execute("
        CREATE TABLE IF NOT EXISTS purchase_items (
            id SERIAL PRIMARY KEY,
            name VARCHAR,
            drink_id INTEGER,
            num_items INTEGER,
            unit_price INTEGER,
            user_id INTEGER,
            purchase_datetime VARCHAR,
            purchase_id INTEGER REFERENCES purchases(id)
        )
    ")?;


    db_client.close()?;

    Ok(())
}




pub fn get_user_purchase_count_for_drink(user_id: i32, drink_id: i32) -> Result<String, Error> {
    let mut db_client = Client::connect("postgresql://postgres:root@localhost:5432/ntp-purchase-service", NoTls)?; 
    let mut num_purchases: i64 = 0;
    for _ in db_client.query("SELECT * FROM purchase_items WHERE user_id = $1 AND drink_id = $2", &[&user_id, &drink_id])? {
      num_purchases += 1;
    }

    let value = UserPurchaseCount{
      num_purchases: num_purchases
    };
    Ok(serde_json::to_string(&value).unwrap())
}

pub fn create_purchase(purchase: rocket::serde::json::Json<PurchaseCreationDTO>) -> Result<String, Error> {
    let mut db_client = Client::connect("postgresql://postgres:root@localhost:5432/ntp-purchase-service", NoTls)?; 
    let user_id = purchase.user_id;
    let mut generated_id: i32 = 0;
    db_client.execute("INSERT INTO purchases (user_id) VALUES ($1)", &[&user_id])?;
    for row in db_client.query("SELECT id from purchases ORDER BY id DESC LIMIT 1", &[])? {
      generated_id = row.get(0);
      break;
    }
    let time_now = std::time::SystemTime::now();
    let formatter: DateTime<Utc> = time_now.clone().into();
    let iso_format_date = formatter.format("%+").to_string();
    for item in &purchase.purchase_items {
      db_client.execute("INSERT INTO purchase_items (name, drink_id, num_items, unit_price, purchase_id, user_id, purchase_datetime) VALUES ($1, $2, $3, $4, $5, $6, $7)", &[&item.name, &item.drink_id, &item.num_items, &item.unit_price, &generated_id, &user_id, &iso_format_date])?;
    }


  Ok(serde_json::to_string(&StatusMessage{message: "Purchase created successfully!".to_string()}).unwrap())
}