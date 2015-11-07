package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Verbose bool
var action string = "up/down"

var HugoCmd = &cobra.Command{
	Use:   "pfs",
	Short: "A personal file server",
	Long: `A personal file server for sharing and receiving files from friends.
            Complete documentation is available at http://github.com/joncrlsn/pfs`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Allow friend to (only) upload files to your computer",
	Long: `Starts web server that only allows uploading of files to 
	the current directory.  No downloading is allowed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uploading", args)
		action = "up"
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Allow friend to download files from the current directory",
	Long: `Starts web server that only allows downloading of files from 
	the current directory.  No uploading is allowed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("downloading", args)
		action = "down"
	},
}

func init() {
	HugoCmd.AddCommand(uploadCmd)
	HugoCmd.AddCommand(downloadCmd)
	HugoCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	HugoCmd.Execute()
}

func main() {
	fmt.Println("verbose", Verbose)
	fmt.Println("action", action)
}
