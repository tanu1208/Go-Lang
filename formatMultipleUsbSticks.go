package main

import "os"
import "os/exec"
import "fmt"
import "bytes"
import "strconv"

func ejectDisk(stickCount int) {
	startDisk := 2

	for i := startDisk; i <= stickCount+1; i++ {
		//format := fmt.Sprintf("diskutil eraseDisk FAT32 %s /dev/disk%d", diskName, i)
		diskNr := fmt.Sprintf("/dev/disk%d", i)
		cmd := exec.Command("diskutil", "unmountdisk", diskNr)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    		return
		}
		fmt.Println(cmd)
	}

	fmt.Println("Done")
}

func formatDisk(stickCount int) {
	startDisk := 2
	diskName := "NONAME"

	for i := startDisk; i <= stickCount+1; i++ {
		//format := fmt.Sprintf("diskutil eraseDisk FAT32 %s /dev/disk%d", diskName, i)
		diskNr := fmt.Sprintf("/dev/disk%d", i)
		cmd := exec.Command("diskutil", "eraseDisk", "FAT32", diskName, diskNr)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    		return
		}
		fmt.Println(cmd)
	}

	fmt.Println("Done")
}

func main() {

	if len(os.Args) < 2{
		fmt.Println("* To few arguments")
		fmt.Println("- use 'h' for help \n")
	} else if len(os.Args) < 4 {

		if os.Args[1] == "h" {
			fmt.Println("=== Args: ===")
			fmt.Println("f: format n: number of usb sticks (default name: NONAME)")
			fmt.Println("e: eject n: number of usb sticks (ejects n amount of sticks) \n")
		} else if os.Args[1] == "f" {
			if len(os.Args) == 3 {
				usbStickCount := os.Args[2]
				stickCount, err := strconv.Atoi(usbStickCount)
				if err != nil {
				    // handle error
				    fmt.Println(err)
				    os.Exit(2)
				}

				formatDisk(stickCount)
			} else {
				fmt.Println("* To few arguments")
				fmt.Println("- use 'h' for help \n")
			}
		} else if os.Args[1] == "e" {
			if len(os.Args) == 3 {
				usbStickCount := os.Args[2]
				stickCount, err := strconv.Atoi(usbStickCount)
				if err != nil {
				    // handle error
				    fmt.Println(err)
				    os.Exit(2)
				}
				
				ejectDisk(stickCount)
			} else {
				fmt.Println("* To few arguments")
				fmt.Println("- use 'h' for help \n")
			}
		} else {
			fmt.Println("* Invalid argument(s)")
			fmt.Println("- use 'h' for help")
		}
	} else {
		fmt.Println("* Invalid argument(s)")
		fmt.Println("- use 'h' for help")
	}

}
