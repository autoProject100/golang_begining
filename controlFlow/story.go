package main

import "fmt"

type stepString struct {
	nextChosenNumber  int
	currentStepString string
	caseNumber               string
	nextStep                  *stepString
}

func main() {
	gameOverString := "GAME OVER"
	zeroStepString := "Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. \n" +
		"Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.\n" +
		"У печері темно, тому Стівен іде стежкою, яка веде від печери в ліс.\n" +
		"У лісі Стівен натикається на мертве тіло дивної тварини.\n"
	zeroStepStringQuestion := "1 - з'їсти, щоб втамувати голод, скориставшись ножем та сірниками\n" +
		"2 - нічого не робити\n"
	firstStepString := "Стівен непритомніє від отруєння. А все могло бути зовсім інакше...\n"
	firstStepString1 := "Він обирає нічого з цим не робити й іти далі.\n" +
		"Через деякий час Стівен приходить до безлюдного табору.\n"
	firstStepStringQuestion := "1 - пройти повз і йти далі у ліс\n" +
		"2 - зайти у найближчий намет\n3 - підійти до дерева, на якому вирізане послання, щоб прочитати його"
	secondStepString := "Вирішує не зупинятись і на адреналіні йде далі, але ліс здається нескінченним. \n" +
		"Від холоду він непритомніє. А все могло бути зовсім інакше...\n"
	secondStepString1 := "У найближчому наметі він знаходить сейф з кодовим замком з двох чисел. \n" +
		"Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. \n" +
		"Стівен непритомніє. А все могло бути зовсім інакше...\n"
	secondStepString2 := "Підійшовши до дерева увімкнувши ліхтарик Стівен чинатє надпис:\n" +
		"'Компас і карта виходу з лісу у дуплі дерева'\n"
	secondStepStringQuestion := "1 - запалити сірник і кинути в дупло, щоб побачити у темряві що там всередині\n" +
		"2 - засунути руку у дупло, щоб впомацки знайти карту і компас\n"
	thirdStepString := "Карта з компасом згорають. Стівен вирішує шукати вихід з лісу сам, але ліс здається нескінченним.\n " +
		"Від холоду він непритомніє. А все могло бути зовсім інакше...\n"
	thirdStepString1 := "Використовуючи компас і карту Стівен знаходить вихід із лісу. Тут він зустрічає старця," +
		"який розкаже як його звуть і чому він опинився в печері. Але це вже зовсім інша історія..."

	thirdStep:=stepString{nextChosenNumber: 1,
		currentStepString:thirdStepString,
		caseNumber:1,
	}
	thirdStep1:=stepString{nextChosenNumber: 2,
		currentStepString:thirdStepString1,
		caseNumber:2,
	}
	secondStep:=stepString{nextChosenNumber: 1,
		currentStepString:secondStepString,
		caseNumber:2,
	}
	secondStep1:=stepString{nextChosenNumber: 2,
		currentStepString:secondStepString1,
		caseNumber:2,
	}
	secondStep2:=stepString{nextChosenNumber: 3,
		currentStepString:secondStepString2,
		caseNumber:2,
		nextStep:&thirdStep
	}
	secondStep3:=stepString{nextChosenNumber: 3,
		currentStepString:secondStepString2,
		caseNumber:2,
		nextStep:&thirdStep1
	}
	firstStep:=stepString{nextChosenNumber: 1,
		currentStepString:firstStepString,
		caseNumber:1,
		nextStep:&thirdStep1
	}
	firstStep1:=stepString{nextChosenNumber: 2,
		currentStepString:firstStepString1,
		caseNumber:2,
		nextStep:&thirdStep1
	}
	zeroStep:=stepString{nextChosenNumber: 1,
		currentStepString:firstStepString,
		caseNumber:1,
	}
	zeroStep1:=stepString{nextChosenNumber: 2,
		currentStepString:firstStepString1,
		caseNumber:firstStepStringQuestion,
		nextStep:&firstStep1
	}
	zeroStory:=stepString{nextChosenNumber: 0,
		currentStepString:zeroStepString,
		caseNumber:1,
	}

	fmt.Printf(zeroStory.currentStepString)
	var i int
	var currentStep *stepString = &zeroStep
	fmt.Print(zeroStepStringQuestion)
	fmt.Scan(&i)
	switch i {
	case 1:
		currentStep := &zeroStep
	case 2:
		currentStep := &zeroStep1
	default:
		fmt.Printf(gameOverString)
	}
	for currentStep != nil{
		fmt.Printf(currentStep.currentStepString)
		fmt.Printf(currentStep.caseNumber)
		fmt.Scan(&i)
		switch i {
		case 1:
			currentStep := &currentStep.nextStep
		case 2:
			currentStep := &zeroStep1
		default:
			fmt.Printf(gameOverString)
		}
		currentStep = currentStep.nextStep
	}

	//fmt.Printf(zeroStepString)
	//fmt.Print(zeroStepStringQuestion)
	//fmt.Scan(&i)
	//if i == 1 {
	//	fmt.Printf(firstStepString)
	//} else if i == 2 {
	//	fmt.Printf(firstStepString1)
	//}
	//fmt.Print(firstStepStringQuestion)
	//fmt.Scan(&i)
	//switch i {
	//case 1:
	//	fmt.Printf(secondStepString)
	//case 2:
	//	fmt.Printf(secondStepString1)
	//case 3:
	//	fmt.Printf(secondStepString2)
	//}
	//fmt.Print(secondStepStringQuestion)
	//fmt.Scan(&i)
	//switch i {
	//case 1:
	//	fmt.Printf(thirdStepString)
	//case 2:
	//	fmt.Printf(thirdStepString1)
	//}
	//fmt.Println(gameOverString)
}
