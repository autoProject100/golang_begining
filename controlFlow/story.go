package main

import "fmt"

type step struct {
	stepString   string
	stepQuestion *questions
}

type questions struct {
	stepQuestion    string
	one, two, three *step
}

func main() {
	//strings to variables
	gameOverString := "GAME OVER"
	zeroStepString := "Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. \n" +
		"Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.\n" +
		"У печері темно, тому Стівен іде стежкою, яка веде від печери в ліс.\n" +
		"У лісі Стівен натикається на мертве тіло дивної тварини.\n"
	zeroStepStringQuestion := "1 - з'їсти, щоб втамувати голод, скориставшись ножем та сірниками\n" +
		"2 - нічого не робити\n"
	firstStepString1 := "Стівен непритомніє від отруєння. А все могло бути зовсім інакше...\n"
	firstStepString2 := "Він обирає нічого з цим не робити й іти далі.\n" +
		"Через деякий час Стівен приходить до безлюдного табору.\n"
	firstStepStringQuestion := "1 - пройти повз і йти далі у ліс\n" +
		"2 - зайти у найближчий намет\n3 - підійти до дерева, на якому вирізане послання, щоб прочитати його\n"
	secondStepString1 := "Вирішує не зупинятись і на адреналіні йде далі, але ліс здається нескінченним. \n" +
		"Від холоду він непритомніє. А все могло бути зовсім інакше...\n"
	secondStepString2 := "У найближчому наметі він знаходить сейф з кодовим замком з двох чисел. \n" +
		"Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. \n" +
		"Стівен непритомніє. А все могло бути зовсім інакше...\n"
	secondStepString3 := "Підійшовши до дерева увімкнувши ліхтарик Стівен чинатє надпис:\n" +
		"'Компас і карта виходу з лісу у дуплі дерева'\n"
	secondStepStringQuestion := "1 - запалити сірник і кинути в дупло, щоб побачити у темряві що там всередині\n" +
		"2 - засунути руку у дупло, щоб впомацки знайти карту і компас\n"
	thirdStepString1 := "Карта з компасом згорають. Стівен вирішує шукати вихід з лісу сам, але ліс здається нескінченним.\n" +
		"Від холоду він непритомніє. А все могло бути зовсім інакше...\n"
	thirdStepString2 := "Використовуючи компас і карту Стівен знаходить вихід із лісу. Тут він зустрічає старця," +
		"який розкаже як його звуть і чому він опинився в печері. Але це вже зовсім інша історія...\n"

	//set all objects to create the story
	third2 := step{
		stepString:   thirdStepString2,
		stepQuestion: nil,
	}
	third1 := step{
		stepString:   thirdStepString1,
		stepQuestion: nil,
	}

	second3Question := questions{
		stepQuestion: secondStepStringQuestion,
		one:          &third1,
		two:          &third2,
		three:        nil,
	}

	second3 := step{
		stepString:   secondStepString3,
		stepQuestion: &second3Question,
	}
	second2 := step{
		stepString:   secondStepString2,
		stepQuestion: nil,
	}
	second1 := step{
		stepString:   secondStepString1,
		stepQuestion: nil,
	}

	first2Question := questions{
		stepQuestion: firstStepStringQuestion,
		one:          &second1,
		two:          &second2,
		three:        &second3,
	}

	first2 := step{
		stepString:   firstStepString2,
		stepQuestion: &first2Question,
	}
	first1 := step{
		stepString:   firstStepString1,
		stepQuestion: nil,
	}

	zeroQuestion := questions{
		stepQuestion: zeroStepStringQuestion,
		one:          &first1,
		two:          &first2,
		three:        nil,
	}

	zero := step{
		stepString:   zeroStepString,
		stepQuestion: &zeroQuestion,
	}

	//starting story
	var currentQuestion = &zeroQuestion
	var currentStep = &zero
	var i int
	for currentStep.stepQuestion != nil {
		fmt.Printf(currentStep.stepString)
		fmt.Printf(currentQuestion.stepQuestion)
	again:
		fmt.Scan(&i)
		switch i {
		case 1:
			currentStep = currentQuestion.one
		case 2:
			currentStep = currentQuestion.two
		case 3:
			currentStep = currentQuestion.three
			if currentStep == nil {
				fmt.Printf("No such way! Try input again!\n")
				goto again
			}
		default:
			fmt.Printf("No such way! Try input again!\n")
			goto again
		}
		//break from loop and do not find nil field from nil object
		if currentStep != nil {
			currentQuestion = currentStep.stepQuestion
		} else {
			break
		}
	}
	//do not find nil field from nil object
	if currentStep != nil {
		fmt.Printf(currentStep.stepString)
	}
	fmt.Println(gameOverString)
}
