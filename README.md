# ფეხბურთის API

## Start up

აპლიკაციის გაშვება საკმაოდ მარტივია ყველაფერი არის docker-compose.yaml ფაილში გაწერილი, ჩვენი მხრიდამ მხოლოდ საჭირო
იქნება გავუშვათ ბრძანება

`docker-compose up --build`

ამით შიემქნება ორი კონტეინერი, თავად API და Postgre-ს ბაზა რომელშიც შევინახავთ მონაცემებს, ასევე ავტომატურად მოხდება ამ
ბაზის პოპულირება Dummy data-თი. (მეტი ინფორმაციისთვის ეწვიეთ 0002_seed.up.sql)

## Architecture

API-ში მთავარი ლოგიკა არის გადანაწილებული internal repo-ში, სტანდარტული naming convention-ებით.

## ნაკლი

მთავარი ნაკლი ლოკალიზაციის ვერ მოსწრება გახლავთ Requirment-ებიდან, i18n-ის ინტეგრაციას ვაპირებდი, მაგრამ ვეღარ მოესწრო.

## გადაწყვეტილებები

ნაწილობრივ მოგიყვებით რატო გადავწყვიტე,

### რეგისტრაცია იუზერის

განცალკევება იუზერის შექმნის და Assignment-ის თიმთან, რეალურად ეს უფრო რამდენიმე ნაბიჯიანი sign up უნდა იყოს რომელშიც
შეუძლია უზერს პირდპაირ გადაწყვიტოს აღარ სურს გაგრძელება გუნდის შემქნის, უნდა იმიჯნებოდეს ერთის შემთხვევაში Random სახელი
და ქვეყანა უნდა ქონდეს გუნდს ხოლო მეორე შემთხვევაში უფრო გრძელი sign up უნდა იყოს. ამიტომ ორი სხვადასხვა Creation func
არის Create და AssignTeamRandomly.

### ბაზის არქიტექტურა

ბაზა არის შემდგარი 5 Table-სგან

**USERS** - ადამიანები რომლებიც სარგებლობენ აპლიკაციით\
**TEAMS** - თამაშში მონაწილე გუნდები\
**PLAYERS** - ფეხბურთელები\
**TRANSFER_LISTINGS** - ტრანსფერებისთვის აღრიცხული მოთამაშეები\
**TRANSFERS** - უკვე დასრულებული ტრანსფერები რომელიც შეგვიძია ვნახოთ აუდიტისთვის

classDiagram
direction BT
class players {
bigint team_id
varchar(100) first_name
varchar(100) last_name
varchar(100) country
integer age
bigint market_value_cents
timestamp with time zone created_at
timestamp with time zone updated_at
player_position position
bigint id
}
class teams {
varchar(100) name
bigint budget_cents
timestamp with time zone created_at
varchar(100) country
bigint owner_id
bigint id
}
class transfer_listings {
bigint player_id
bigint seller_team_id
bigint sold_to_team_id
listing_status status
bigint asking_price_cents
timestamp with time zone listed_at
timestamp with time zone sold_at
bigint id
}
class transfers {
bigint player_id
bigint from_team_id
bigint to_team_id
bigint price_cents
timestamp with time zone created_at
bigint id
}
class users {
varchar(255) email
varchar(255) password_hash
timestamp with time zone created_at
bigint id
}

players -->  teams : team_id:id
teams -->  users : owner_id:id
transfer_listings -->  players : player_id:id
transfer_listings -->  teams : seller_team_id:id
transfer_listings -->  teams : sold_to_team_id:id
transfers -->  players : player_id:id
transfers -->  teams : from_team_id:id
transfers -->  teams : to_team_id:id

### PS

აქვე ნახავთ file-ებს რომელშიც არ წერია არაფერი, ეს რეალურად არის იდეები რომლების იმპლემენტაციაც მიმაჩნდა კარგი აზრი,
მაგრამ დროის დასრულებიდან გამომდინარე ვერ მოხდა იმპლემენტაცია, მაგრამ ვინმე თუ იკითხავს ეს უბრალოდ out of scope იყო და
შემდეგ სპრინტებში უნდა მოხდეს დაგეგმვა.

<img src="img.png" alt="drawing" width="500"/>