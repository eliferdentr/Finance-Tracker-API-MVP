package db

import (
	"log"
	"strconv"
	"time"

	"github.com/eliferdentr/finance-tracker-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to postgres using the info coming from the config struct, using gorm and return the gorm DB instance
func NewPostgres(cfg *config.Config) (*gorm.DB, error) {

	//örnek dsn : postgres://username:password@localhost:5432/dbname?sslmode=disable
	dsn := "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + strconv.Itoa(cfg.DBPort) +
		" sslmode=" + cfg.DBSSLMode +
		" TimeZone=Europe/Istanbul"

		//&gorm.Config{} kısmı GORM’un çalışma ayarlarını tuttuğun konfigürasyon nesnesi.
		//GORM’un log seviyesi, naming stratejisi, prepared statement cache’i, transaction ayarları gibi davranışlarını kontrol eder.
		//gorm.Open()’e 2. parametre olarak verilir ve opsiyoneldir. Boş {} verdiğinde GORM “default ayarlarla çalış” demiş olursun.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to Postgres:", err)
		return nil, err
	}
	log.Println("Connected to Postgres successfully")

	
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get sql.DB from GORM:", err)
		return nil, err
	}

	//pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	//ping
	err = sqlDB.Ping()
	if err != nil {
		log.Println("Postgres ping failed:", err)
		return nil, err
	}
	log.Println("Postgres ping successful")
	return db, nil
}
