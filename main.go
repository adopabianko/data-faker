package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
)

type User struct {
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	Jwt                string  `faker:"jwt"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	FirstName          string  `faker:"first_name"`
	LastName           string  `faker:"last_name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`
}

func main() {
	connStr := "postgres://postgres:password@localhost:9432/faker?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	for i := 1; i < 50000; i++ {
		err := faker.FakeData(&user)
		if err != nil {
			log.Fatal(err)
		}

		// insert to table users
		_, err = db.Exec(fmt.Sprintf(
			`INSERT INTO users(
				lat,long,cc_number,cc_type,email,domain_name,ipv4,ipv6,password,jwt,phone_number,mac_address,url,username,toll_free_number,e_164_phone_number,first_name,last_name,unix_time,date,time,month_name,year,day_of_week,day_of_month,timestamp,century,timezone,time_period,word,sentence,paragraph,currency,amount,amount_with_currency,uuid_hyphenated,uuid_digit
			)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37)`),
			user.Latitude,
			user.Longitude,
			user.CreditCardNumber,
			user.CreditCardType,
			user.Email,
			user.DomainName,
			user.IPV4,
			user.IPV6,
			user.Password,
			user.Jwt,
			user.PhoneNumber,
			user.MacAddress,
			user.URL,
			user.UserName,
			user.TollFreeNumber,
			user.E164PhoneNumber,
			user.FirstName,
			user.LastName,
			user.UnixTime,
			user.Date,
			user.Time,
			user.MonthName,
			user.Year,
			user.DayOfWeek,
			user.DayOfMonth,
			user.Timestamp,
			user.Century,
			user.TimeZone,
			user.TimePeriod,
			user.Word,
			user.Sentence,
			user.Paragraph,
			user.Currency,
			user.Amount,
			user.AmountWithCurrency,
			user.UUIDHypenated,
			user.UUID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d : success \n", i)
		i++
	}

}
