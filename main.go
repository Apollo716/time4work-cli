package main

import (
	"fmt"
	"log"
	"time"

	conf "github.com/Apollo716/time4work-cli/config"
	time4work "github.com/Apollo716/time4work-cli/time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	initConf()
	db = initDB()

	var rootCmd = &cobra.Command{Use: "myapp"}
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(endCmd)
	rootCmd.Execute()
}

// declare Cmd
/*-------------------------------------------------------*/

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command",
	Long:  `Start the application`,
	Run: func(cmd *cobra.Command, args []string) {
		time4work.InsertTimeRecord(db, "start", time.Now())
	},
}

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End command",
	Long:  `End the application`,
	Run: func(cmd *cobra.Command, args []string) {
		time4work.InsertTimeRecord(db, "end", time.Now())
	},
}

/*-------------------------------------------------------*/

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", conf.DBUser(), conf.DBPass(), conf.DBName(), conf.DBPort())

	// データベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("データベース接続に失敗しました: " + err.Error())
	}
	return db
}

func initConf() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
