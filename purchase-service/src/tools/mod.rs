

use postgres::{Client, NoTls, Error};
use serde::{Serialize, Deserialize};



//Model
#[derive(Serialize, Deserialize)]
pub struct PurchaseCreationDTO {
  user_id: i64,
  purchase_items: Vec<PurchaseItemCreationDTO>,

}

#[derive(Serialize, Deserialize)]
pub struct PurchaseItemCreationDTO {
  drink_id: i64,
  num_items: u32,
  unit_price: u32,
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
            user_id INTEGER,
            purchase_datetime VARCHAR 
        )
        

    ")?;
    db_client.batch_execute("
        CREATE TABLE IF NOT EXISTS purchase_items (
            id SERIAL PRIMARY KEY,
            drink_id INTEGER,
            num_items INTEGER,
            unit_price INTEGER,
            user_id INTEGER,
            purchase_id INTEGER REFERENCES purchases(id)
        )
    ")?;


    db_client.close()?;

    Ok(())
}




pub fn get_user_purchase_count_for_drink(user_id: i32, drink_id: i32) -> Result<(String), Error> {
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