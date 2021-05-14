package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"github.com/robfig/cron/v3"
	"github.com/stoyvo/wellness-buddy/panel/breathingexercises"
	"github.com/stoyvo/wellness-buddy/panel/chairyoga"
	"github.com/stoyvo/wellness-buddy/panel/connect"
	"github.com/stoyvo/wellness-buddy/panel/dailychallenge"
	"github.com/stoyvo/wellness-buddy/panel/exercise"
	"github.com/stoyvo/wellness-buddy/panel/hydrate"
	"github.com/stoyvo/wellness-buddy/panel/snacks"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func startJobs() {
	// Call dailyChallenge at 2PM
	dailyChallengeJob := cron.New()
	_, err := dailyChallengeJob.AddFunc("00 14 * * 1-5", dailyChallengeAction)

	if err != nil {
		return
	}

	dailyChallengeJob.Start()

	// Call breathingExercise every other weekday at 11AM
	breathingExerciseJob := cron.New()
	_, err = breathingExerciseJob.AddFunc("0 11 * * 1-5/2", breathingExerciseAction)

	if err != nil {
		return
	}

	breathingExerciseJob.Start()

	// Call take a walk once a week, on a wednesday at 3:30PM
	takeWalkJob := cron.New()
	_, err = takeWalkJob.AddFunc("30 15 * * 3", takeWalkAction)

	if err != nil {
		return
	}

	takeWalkJob.Start()

	// Call stretch at 10AM and 3PM each weekday
	stretchJob := cron.New()
	_, err = stretchJob.AddFunc("0 10,15 * * 1-5", stretchAction)

	if err != nil {
		return
	}

	stretchJob.Start()

	// Call hydrateReminder
	hydrateReminderJob := cron.New()
	_, err = hydrateReminderJob.AddFunc("@every 30m", hydrateReminderAction)

	if err != nil {
		return
	}

	hydrateReminderJob.Start()

	// Call chairYoga twice a week at 1PM
	chairYogaJob := cron.New()
	_, err = chairYogaJob.AddFunc("0 13 * * 2,4", chairYogaAction)

	if err != nil {
		return
	}

	chairYogaJob.Start()

	// Call healthy snack action once a day, at 12:30
	healthySnackJob := cron.New()
	_, err = healthySnackJob.AddFunc("30 12 * * 1-5", healthySnackAction)

	if err != nil {
		return
	}

	healthySnackJob.Start()

	// Call message a loved one action once a week, on a tuesday at 4:30PM
	messageJokeJob := cron.New()
	_, err = messageJokeJob.AddFunc("30 16 * * 2", messageJokeAction)

	if err != nil {
		return
	}

	messageJokeJob.Start()

	ondemandJokeAction()
	// Call ondemand joke renewal
	ondemandJokeJob := cron.New()
	_, err = ondemandJokeJob.AddFunc("@every 5m", ondemandJokeAction)

	if err != nil {
		return
	}

	ondemandJokeJob.Start()
}

func dailyChallengeAction() {
	//display os notification to complete the challenge, clicking on the notification should:
	//open the app on the corresponding panel with button for 'done' or 'skip'
	//add the point if pressed 'done'
	//if 'skip' nothing should happen, cron should operate as usual
	var challengeText string
	var day = int(time.Now().Weekday())

	switch challenge := day; challenge {
	case 1:
		challengeText = "Take a picture of a tree. Feel free to share it on #opt_outside"
		note := gosxnotifier.NewNotification(challengeText)
		//Optionally, set a title
		note.Title = "Wellness Buddy"
		//Optionally, set a sound from a predefined set.
		note.Sound = gosxnotifier.Default
		note.Link  = "com.bounteous.wellness-buddy"
		//Then, push the notification
		err := note.Push()

		//If necessary, check error
		if err != nil {
			fmt.Println(err.Error())
		}
	case 2:
		challengeText = "Have a burnout exercise (1 minute of squats or jumping jacks). Feel free to share your experience on #b_active"
		note := gosxnotifier.NewNotification(challengeText)
		//Optionally, set a title
		note.Title = "Wellness Buddy"
		//Optionally, set a sound from a predefined set.
		note.Sound = gosxnotifier.Default
		note.Link  = "com.bounteous.wellness-buddy"
		//Then, push the notification
		err := note.Push()

		//If necessary, check error
		if err != nil {
			fmt.Println(err.Error())
		}
	case 3:
		challengeText = "Have a jabs exercise. Feel free to share your experience #b_active"
		note := gosxnotifier.NewNotification(challengeText)
		//Optionally, set a title
		note.Title = "Wellness Buddy"
		//Optionally, set a sound from a predefined set.
		note.Sound = gosxnotifier.Default
		note.Link  = "com.bounteous.wellness-buddy"
		//Then, push the notification
		err := note.Push()

		//If necessary, check error
		if err != nil {
			fmt.Println(err.Error())
		}
	case 4:
		challengeText = "Take a picture of a bird. Feel free to share it on #photography"
		note := gosxnotifier.NewNotification(challengeText)
		//Optionally, set a title
		note.Title = "Wellness Buddy"
		//Optionally, set a sound from a predefined set.
		note.Sound = gosxnotifier.Default
		note.Link  = "com.bounteous.wellness-buddy"
		//Then, push the notification
		err := note.Push()

		//If necessary, check error
		if err != nil {
			fmt.Println(err.Error())
		}
	case 5:
		challengeText = "Take a picture of a fire hydrant. Feel free to share it on #photography"
		note := gosxnotifier.NewNotification(challengeText)
		//Optionally, set a title
		note.Title = "Wellness Buddy"
		//Optionally, set a sound from a predefined set.
		note.Sound = gosxnotifier.Default
		note.Link  = "com.bounteous.wellness-buddy"
		//Then, push the notification
		err := note.Push()

		//If necessary, check error
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("No day argument was passed!")
	}

	//load the daily challenge panel in the app
	dailychallenge.ChallengeAction = challengeText
	navList.Select(1)
}

func breathingExerciseAction() {
	//display os notification to complete the breathing exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	note := gosxnotifier.NewNotification("Check out this breathing exercise - your body sure wants to!")
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load breathing exercises panel in the app
	breathingexercises.Active = true
	navList.Select(2)
}

func takeWalkAction() {
	//display os notification to take a walk, clicking on the notification should:
	//open the app on the corresponding panel with a button for 'start'
	//clicking the 'start' button will display/ enable a 'stop' button
	//and start a timer (with visuals if possible!)
	//add the point if pressed 'stop' and display 'last walked for X' where X is "2 hours"
	note := gosxnotifier.NewNotification("Go ahead and take a walk - you've earned it!")
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load exercise panel
	exercise.TakeWalk = true
	navList.Select(3)
}


func stretchAction() {
	//display os notification to complete the stretch exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	note := gosxnotifier.NewNotification("Stop and do some stretches - you know you want to!")
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load exercise panel
	exercise.Active = true
	navList.Select(3)
}

func hydrateReminderAction() {
	//display os notification to drink, clicking on the notification should:
	//open the app on the corresponding panel with button for 'done' or 'skip'
	//add the point if pressed 'done'
	//if 'skip' nothing should happen, cron should operate as usual
	note := gosxnotifier.NewNotification("Go ahead and have a glass of water - your body will thank you!")
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load the hydration panel in the app
	hydrate.Active = true
	navList.Select(4)
}

func chairYogaAction() {
	//display os notification to complete the chair yoga exercise, clicking on the notification should:
	//open the app on the corresponding panel with the youtube link and a button for 'done'
	//add the point if pressed 'done' and disable that button
	//then, run cron for 2 hours which will re-enable the button
	note := gosxnotifier.NewNotification("Check out this chair yoga exercise!")
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load chair yoga panel
	chairyoga.Active = true
	navList.Select(5)
}

func randomizeSnack() string {
	snacksOpts := make([]string, 0)
	snacksOpts = append(snacksOpts,
		"some strawberries",
		"some Greek yogurt",
		"a banana",
		"a granola bar",
		"some kale chips",
		"some trail mix",
		"some dark chocolate",
		"a baby carrot",
		"a celery stick",
		"an apple",
		"some buckwheat",
		"a cucumber")

	return snacksOpts[rand.Intn(len(snacksOpts))]
}

func healthySnackAction() {
	//display os notification to eat a healthy snack, clicking on the notification should:
	//open the app on the corresponding panel with a healthy snack suggestion from the random list up top
	//and a button for 'done'
	//clicking the 'done' button will add the point
	selectedSnack := randomizeSnack()
	message := fmt.Sprint("Stop and eat something healthy, like ", selectedSnack, "- you deserve it!")
	note := gosxnotifier.NewNotification(message)
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load snack panel
	snacks.Active = true
	snacks.Snack = selectedSnack
	navList.Select(6)
}

func fetchDadJoke() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	req.Header.Set("User-Agent", "Wellness Buddy (https://github.com/Stoyvo/wellness-buddy)")//required header
	req.Header.Set("Accept", "text/plain")
	res, _ := client.Do(req)
	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	return string(responseData)
}

func messageJokeAction() {
	//display os notification to message a joke to a loved one, clicking on the notification should:
	//open the app on the corresponding panel with a joke pulled from https://icanhazdadjoke.com/api
	//and a button for 'done'
	//clicking the 'done' button will add the point
	joke := fetchDadJoke()
	message := fmt.Sprint("Reach out to a loved one, perhaps - with a joke? \r\n \"", joke, "\"")
	note := gosxnotifier.NewNotification(message)
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err := note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load connect panel
	connect.Active = true
	connect.Joke = joke
	navList.Select(7)
}

func ondemandJokeAction() {
	if !connect.Active { //if currently not completing other cron action
		joke := fetchDadJoke()
		connect.Joke = joke
	}
}