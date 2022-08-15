



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


#[derive(Serialize, Deserialize)]
pub struct PurchaseDTO {
  id: i32,
  user_id: i32,
  purchase_items: Vec<PurchaseItemDTO>
}


#[derive(Serialize, Deserialize)]
pub struct PurchaseItemDTO {
  id: i32,
  name: String,
  drink_id: i32,
  num_items: i32,
  unit_price: i32,
  purchase_datetime: String
}

#[derive(Serialize)]
pub struct StatusMessage {
  pub message: String,
  pub status_code: i32
}

#[derive(Serialize, Deserialize)]
pub struct UserPurchaseCount {
  num_purchases: i64,
}



#[derive(Serialize, Deserialize)]
pub struct DrinkSoldCount {
  drink_id: i32,
  sold_items: i64,
}








//Functions
pub fn create_database() -> Result<(), Error> {

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


  Ok(serde_json::to_string(&StatusMessage{message: "Purchase created successfully!".to_string(), status_code: 201}).unwrap())
}

pub fn get_user_purchase_history(user_id: i32) -> Result<String, Error> {
    let mut db_client = Client::connect("postgresql://postgres:root@localhost:5432/ntp-purchase-service", NoTls)?; 
    let mut purchase_history: Vec<PurchaseDTO> = Vec::new();
    for purchase_row in db_client.query("SELECT id from purchases WHERE user_id = $1", &[&user_id])? {
        let purchase_id: i32 = purchase_row.get(0);
        let mut purchase: PurchaseDTO = PurchaseDTO { id:purchase_id, user_id: user_id, purchase_items: Vec::new() };
        for item_row in db_client.query("SELECT * FROM purchase_items WHERE user_id = $1 AND purchase_id = $2", &[&user_id, &purchase_id])? {
          let item_id: i32 = item_row.get(0);
          let item_name: String = item_row.get(1);
          let item_drink_id: i32 = item_row.get(2);
          let item_num_items: i32 = item_row.get(3);
          let item_unit_price: i32 = item_row.get(4);
          let item_purchase_date_time: String = item_row.get(6);

          let purchase_item: PurchaseItemDTO = PurchaseItemDTO { id:item_id, name: item_name, drink_id: item_drink_id, num_items: item_num_items, unit_price: item_unit_price, purchase_datetime: item_purchase_date_time };
          purchase.purchase_items.push(purchase_item);
        }
        purchase_history.push(purchase);
    }

    Ok(serde_json::to_string(&purchase_history).unwrap())


}


pub fn get_most_sold_drinks(period: String) -> Result<String, Error> {

  let mut to_sub: u64 = 0;
  if period == "day" {
    to_sub = 86400;
  } 
  else if period == "week" {
    to_sub = 604800;
  }
  else if period == "month" {
    to_sub = 2629743;
  }
  else if period == "year" {
    to_sub = 31556926;
  } 
  


  let mut time_to_check = std::time::SystemTime::now();
  match time_to_check.checked_sub(std::time::Duration::new(to_sub, 0)) {
    Some(value) => time_to_check = value,
    None => println!("ERROR")
  }

  let formatter: DateTime<Utc> = time_to_check.clone().into();
  let iso_format_date = formatter.format("%+").to_string(); 

  println!("Checking date: {}", iso_format_date);
  
  
  
  let mut db_client = Client::connect("postgresql://postgres:root@localhost:5432/ntp-purchase-service", NoTls)?; 
  let mut most_sold: Vec<DrinkSoldCount> = Vec::new();
  for row in db_client.query("SELECT drink_id, SUM(num_items) as total_sold from purchase_items WHERE purchase_datetime >= $1 GROUP BY drink_id ORDER BY total_sold DESC", &[&iso_format_date])? {
    let drink_id: i32 = row.get(0);
    let total_sold: i64 = row.get(1);

    let drink_sold: DrinkSoldCount = DrinkSoldCount { drink_id: drink_id, sold_items: total_sold };
    most_sold.push(drink_sold);
  }
  Ok(serde_json::to_string(&most_sold).unwrap())

}