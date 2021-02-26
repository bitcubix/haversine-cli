package cmd

import (
	"fmt"
	"github.com/bitcubix/haversine-cli/haversine"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var out string
var short bool

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc [lat1] [lng1] [lat2] [lng2]",
	Short: "calculate distance between two coordinates",
	Run: func(cmd *cobra.Command, args []string) {
		coords := make([]float64, 4)
		var count int
		for c, arg := range args {
			count = c
			if count > 3 {
				fmt.Println("to many arguments")
				os.Exit(1)
			}
			coords[count], _ = strconv.ParseFloat(arg, 64)
		}

		if count < 3 {
			fmt.Println("missing arguments")
			os.Exit(1)
		}

		distance := haversine.CalculateDistance(coords[0], coords[1], coords[2], coords[3])

		switch out {
		case "m":
			break
		case "km":
			distance = distance / 1000
		case "cm":
			distance = distance * 100
		default:
			log.Fatal("invalid output format")
		}

		if !short {
			fmt.Printf("Distance between %f, %f and %f, %f\n", coords[0], coords[1], coords[2], coords[3])
			fmt.Printf("%.3f%s\n", distance, out)
		} else {
			fmt.Printf("%.3f\n", distance)
		}
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
	calcCmd.Flags().StringVarP(&out, "out", "o", "m", "output format (km, m, cm) (default 'm')")
	calcCmd.Flags().BoolVarP(&short, "short", "s", false, "prints only the value")
}
