package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dagou8/go-programming-tour-book-tour/internal/timer"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Format time",
	Long:  "Format time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get now time",
	Long:  "Get now time and the timestamp",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("The output result is %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate time use duration",
	Long:  "Calculate time use duration",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")

			if space == 0 {
				layout = "2006-01-02"
			}

			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			log.Printf("Debug01: %s, %s", layout, calculateTime)

			currentTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				log.Printf("Debug02: %v", err)
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}
		log.Printf("debug: %v", currentTime)
		calc, err := timer.GetCalculateTime(currentTime, duration)
		if err != nil {
			log.Fatalf("calculateTimeCmd.err: %v", err)
		}

		log.Printf("The output result is %s, %d", calc.Format(layout), calc.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "Please enter the time need to calculate")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "Please enter duration time string")
}
