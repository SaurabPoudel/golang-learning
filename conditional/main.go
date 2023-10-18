package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Voting  Eligibility Check")
	fmt.Println("Are you Nepali? [Y/N]")
	reader := bufio.NewReader(os.Stdin)
	isNepali, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Aayein ! Kuirey ko biu")
		return
	}

	fmt.Println("Are you underage ? (below 18) [Y/N]")
	isUnderAge, err2 := reader.ReadString('\n')

	if err2 != nil {
		fmt.Println("Tumshey na ho payega beti/beti")
	}

	if strings.TrimSuffix(isNepali, "\n") == "N" {
		fmt.Println("Ghar janu hos kuireko biu ji, Pls ko to home foreigner's seed, Kripaya ghar pe jaalo pardesi's biu , Chhyang chhung ghintang ghintang")
	} else if strings.TrimSuffix(isUnderAge, "\n") == "Y" {
		fmt.Println("Bhai tero tero umer napugeko ki umer le talai napugeko")
	} else {
		fmt.Println("Hami hajur lai swagat garchham, We welcome you , Ham aapko swagat kartahey , Swagat gariyechha re timlai ")
	}
}
