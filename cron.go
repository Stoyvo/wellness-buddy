package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"math/rand"
	"time"
)

func startJobs() {
	// Call hydrateReminder
	hydrateReminderJob := cron.New()
	_, err := hydrateReminderJob.AddFunc("@every 10s", hydrateReminder)

	if err != nil {
		return
	}

	hydrateReminderJob.Start()

	// Call breathingExercise every other weekday at 11AM
	breathingExerciseJob := cron.New()
	_, err = breathingExerciseJob.AddFunc("0 11 * * 1-5/2", breathingExercise)

	if err != nil {
		return
	}

	breathingExerciseJob.Start()

	// Call chairYoga twice a week at 1PM
	chairYogaJob := cron.New()
	_, err = chairYogaJob.AddFunc("0 13 * * 2,4", chairYoga)

	if err != nil {
		return
	}

	chairYogaJob.Start()

	// Call stretch at 10AM and 3PM each weekday
	stretchJob := cron.New()
	_, err = stretchJob.AddFunc("0 10,15 * * 1-5", stretch)

	if err != nil {
		return
	}

	stretchJob.Start()

	// Call dailyChallenge at 2PM
	dailyChallengeJob := cron.New()
	_, err = dailyChallengeJob.AddFunc("0 14 * * 1-5", dailyChallenge)

	if err != nil {
		return
	}

	dailyChallengeJob.Start()

	// Call take a walk once a week, on a wednesday at 3:30PM
	takeWalkJob := cron.New()
	_, err = takeWalkJob.AddFunc("30 15 * * 3", takeWalk)

	if err != nil {
		return
	}

	takeWalkJob.Start()

	// Call message a loved one action once a week, on a tuesday at 4:30PM
	messageJokeJob := cron.New()
	_, err = messageJokeJob.AddFunc("30 16 * * 2", messageJoke)

	if err != nil {
		return
	}

	messageJokeJob.Start()

	// Call healthy snack action once a day, at 12:30
	healthySnackJob := cron.New()
	_, err = healthySnackJob.AddFunc("30 12 * * 1-5", healthySnack)

	if err != nil {
		return
	}

	healthySnackJob.Start()
}

func hydrateReminder() {
	//display os notification to drink, clicking on the notification should:
	//open the app on the corresponding panel with button for 'done' or 'skip'
	//add the point if pressed 'done'
	//if 'skip' nothing should happen, cron should operate as usual
	fmt.Println("Drink a glass of water")
}

func dailyChallenge() {
	//display os notification to complete the challenge, clicking on the notification should:
	//open the app on the corresponding panel with button for 'done' or 'skip'
	//add the point if pressed 'done'
	//if 'skip' nothing should happen, cron should operate as usual
	var day = int(time.Now().Weekday())
	switch challenge := day; challenge {
	case 1:
		fmt.Println("Take a picture of a tree. Feel free to share on #opt_outside")
	case 2:
		fmt.Println("Burnout exercise (1 minute of squats or jumping jacks). Feel free to share on #b_active")
	case 3:
		fmt.Println("Jabs exercise. Feel free to share on #b_active")
	case 4:
		fmt.Println("Take a picture of a bird. Feel free to share on #photography")
	case 5:
		fmt.Println("Take a picture of a fire hydrant. Feel free to share on #photography")
	default:
		fmt.Println("No day argument was passed")
	}
}

func breathingExercise() {
	//display os notification to complete the breathing exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	fmt.Println("Stop and breathe!")
}

func chairYoga() {
	//display os notification to complete the chair yoga exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	fmt.Println("Stop and do some chair yoga!")
}

func stretch() {
	//display os notification to complete the stretch exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	fmt.Println("Stop and do some streching!")
}

func takeWalk() {
	//display os notification to take a walk, clicking on the notification should:
	//open the app on the corresponding panel with a button for 'start'
	//clicking the 'start' button will display/ enable a 'stop' button
	//and start a timer (with visuals if possible!)
	//add the point if pressed 'stop' and display 'last walked for X' where X is "2 hours"
	fmt.Println("Stop and go for a walk!")
}

func messageJoke() {
	//display os notification to message a joke to a loved one, clicking on the notification should:
	//open the app on the corresponding panel with a joke pulled from https://icanhazdadjoke.com/api
	//and a button for 'done'
	//clicking the 'done' button will add the point
	fmt.Println("Stop and message a joke to a loved one!")
}

func healthySnack() {
	snacks := make([]string, 0)
	snacks = append(snacks,
		"a Baby Carrot",
		"a Celery stick",
		"an Apple",
		"some Buckwheat",
		"a Cucumber")
	//display os notification to eat a healthy snack, clicking on the notification should:
	//open the app on the corresponding panel with a healthy snack suggestion from the random list up top
	//and a button for 'done'
	//clicking the 'done' button will add the point
	message := fmt.Sprint("Stop and eat something healthy, like ", snacks[rand.Intn(len(snacks))])
	fmt.Println(message)
}