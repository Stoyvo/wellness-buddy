package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func startJobs() {
	// Call dailyChallenge at 2PM
	dailyChallengeJob := cron.New()
	_, err := dailyChallengeJob.AddFunc("0 14 * * 1-5", dailyChallenge)

	if err != nil {
		return
	}

	dailyChallengeJob.Start()

	// Call breathingExercise every other weekday at 11AM
	breathingExerciseJob := cron.New()
	_, err = breathingExerciseJob.AddFunc("0 11 * * 1-5/2", breathingExercise)

	if err != nil {
		return
	}

	breathingExerciseJob.Start()

	// Call take a walk once a week, on a wednesday at 3:30PM
	takeWalkJob := cron.New()
	_, err = takeWalkJob.AddFunc("30 15 * * 3", takeWalk)

	if err != nil {
		return
	}

	takeWalkJob.Start()

	// Call hydrateReminder
	hydrateReminderJob := cron.New()
	_, err = hydrateReminderJob.AddFunc("@every 30m", hydrateReminder)

	if err != nil {
		return
	}

	hydrateReminderJob.Start()

	// Call chairYoga twice a week at 1PM
	chairYogaJob := cron.New()
	_, err = chairYogaJob.AddFunc("0 13 * * 2,4", chairYoga)

	if err != nil {
		return
	}

	chairYogaJob.Start()

	// Call healthy snack action once a day, at 12:30
	healthySnackJob := cron.New()
	_, err = healthySnackJob.AddFunc("30 12 * * 1-5", healthySnack)

	if err != nil {
		return
	}

	healthySnackJob.Start()

	// Call stretch at 10AM and 3PM each weekday
	stretchJob := cron.New()
	_, err = stretchJob.AddFunc("0 10,15 * * 1-5", stretch)

	if err != nil {
		return
	}

	stretchJob.Start()

	// Call message a loved one action once a week, on a tuesday at 4:30PM
	messageJokeJob := cron.New()
	_, err = messageJokeJob.AddFunc("11 20 * * 4", messageJoke)

	if err != nil {
		return
	}

	messageJokeJob.Start()
}

func dailyChallenge() {
	//display os notification to complete the challenge, clicking on the notification should:
	//open the app on the corresponding panel with button for 'done' or 'skip'
	//add the point if pressed 'done'
	//if 'skip' nothing should happen, cron should operate as usual
	var day = int(time.Now().Weekday())
	switch challenge := day; challenge {
	case 1:
		note := gosxnotifier.NewNotification("Take a picture of a tree. Feel free to share it on #opt_outside")
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
		note := gosxnotifier.NewNotification("Have a burnout exercise (1 minute of squats or jumping jacks). Feel free to share your experience on #b_active")
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
		note := gosxnotifier.NewNotification("Have a jabs exercise. Feel free to share your experience #b_active")
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
		note := gosxnotifier.NewNotification("Take a picture of a bird. Feel free to share it on #photography")
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
		note := gosxnotifier.NewNotification("Take a picture of a fire hydrant. Feel free to share it on #photography")
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
	navList.Select(1)
}

func breathingExercise() {
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

	//load breathing panel in the app
	navList.Select(2)
}

func takeWalk() {
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

	//load walk panel
	navList.Select(3)
}

func hydrateReminder() {
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
	navList.Select(4)
}

func chairYoga() {
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
	navList.Select(5)
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
	message := fmt.Sprint("Stop and eat something healthy, like ", snacks[rand.Intn(len(snacks))], "- you deserve it!")
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
	navList.Select(6)
}

func stretch() {
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

	//load summary panel
	navList.Select(0)
}

func messageJoke() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	req.Header.Set("User-Agent", "Wellness Buddy (https://github.com/Stoyvo/wellness-buddy)")//required header
	req.Header.Set("Accept", "text/plain")
	res, _ := client.Do(req)
	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	//display os notification to message a joke to a loved one, clicking on the notification should:
	//open the app on the corresponding panel with a joke pulled from https://icanhazdadjoke.com/api
	//and a button for 'done'
	//clicking the 'done' button will add the point
	joke := string(responseData)
	fmt.Println(joke)
	message := fmt.Sprint("Reach out to a loved one, perhaps - with a joke? \r\n \"", joke, "\"")
	note := gosxnotifier.NewNotification(message)
	//Optionally, set a title
	note.Title = "Wellness Buddy"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Default
	note.Link  = "com.bounteous.wellness-buddy"
	//Then, push the notification
	err = note.Push()

	//If necessary, check error
	if err != nil {
		fmt.Println(err.Error())
	}

	//load summary panel
	navList.Select(0)
}