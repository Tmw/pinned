package fetcher

// the OfficeFetcher is used to provide Pinned with Mock data for demo purposes.
// it is featuring The Office quotes from https://github.com/anderskristo/the-office-quotes
// and avatars used from wikipedia. (soon.)

import (
	"github.com/satori/go.uuid"
	"github.com/tmw/pinned/backend/model"
)

var (
	pins = []*model.Pin{
		&model.Pin{
			ID:       uuid.FromStringOrNil("fee39bce-f367-42db-bb7c-cf5c41bc6729"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Jim is gone. He's gone. I miss him so much. Ooooh I cry myself to sleep, Jim! False. I do not miss him.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("5ef7e9be-d28d-4f97-bf1e-98f16b1829b5"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Would I ever leave this company? Look, I’m all about loyalty. In fact, I feel like part of what I’m being paid for here is my loyalty. But if there were somewhere else that valued loyalty more highly… I’m going wherever they value loyalty the most.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("63f2192d-cbf4-430a-94f6-b16d62c322b4"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "I am fast. To give you a reference point I am somewhere between a snake and a mongoose… And a panther.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("dae17ad8-eaf1-4065-90e5-1917bccb1056"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Security in this office park is a joke. Last year I came to work with my spud-gun in a duffel bag. I sat at my desk all day with a rifle that shoots potatoes at 60 pounds per square inch. Can you imagine if I was deranged?",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("c3577211-d0a1-482c-a352-f06273d27db2"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "I grew up on a farm. I have seen animals having sex in every position imaginable. Goat on chicken. Chicken on goat. Couple of chickens doing a goat, couple of pigs watching.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("c252b37c-939c-4b69-b03d-9605ef4ede97"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Through concentration, I can raise and lower my cholesterol at will.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("dc070eef-f736-4ee1-9395-cd63ecd9c631"),
			AuthorID: "e3cb89b7-2f06-4247-84ae-43265cd85f9d",
			Text:     "Right now, this is just a job. If I advance any higher in this company, then this would be my career. And, well, if this were my career? I'd have to throw myself in front of a train.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("97fb01ee-a39d-4e19-8074-d7e794246b70"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Am I going to tell them? No, I am not going to tell them. I don't see the point of that. As a doctor, you would not tell a patient if they had cancer.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("e298b7d0-13af-48e7-ac50-fd70b90c66e4"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Wikipedia Is The Best Thing Ever. Anyone In The World Can Write Anything They Want About Any Subject, So You Know You Are Getting The Best Possible Information.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("a253e194-e212-4d3d-a228-3c8fa957695f"),
			AuthorID: "93c2daaa-084e-4228-8322-1abdb1d891cd",
			Text:     "Saddle Shoes With Denim? I Will Literally Call Protective Services.",
		},
		&model.Pin{
			ID:       uuid.FromStringOrNil("8f4a2b85-b3be-47e3-94b2-c975c4c0a410"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Hate To See You Leave But Love To Watch You Go. ‘Cause Of Your Butt.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("a784afbb-6dcb-403e-839a-db04a9bf9ac8"),
			AuthorID: "1c8fd81f-757b-4ebf-af78-40ffc0268146",
			Text:     "I Wonder What People Like About Me? Probably My Jugs.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("2ffcf519-59d1-404a-b868-1a99d93846ad"),
			AuthorID: "e3cb89b7-2f06-4247-84ae-43265cd85f9d",
			Text:     "Right Now This Is Just A Job. If I Advance Any Higher In This Company, Then This Would Be My Career. And Well, If This Were My Career I’d Have To Throw Myself In Front Of A Train.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("08952ddc-5c81-4dd8-adde-3a62c4ae39e4"),
			AuthorID: "e3cb89b7-2f06-4247-84ae-43265cd85f9d",
			Text:     "Bears. Beets. Battlestar Galactica.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("ea5a0737-5982-4501-8d54-5b90c6db74d1"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Through Concentration, I Can Raise And Lower My Cholesterol At Will.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("23c49c7e-55c3-420b-b479-c4f3b03c3e6a"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "I Love Inside Jokes. Love To Be A Part Of One Someday.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("79413d0f-383c-4dc1-96c8-720f53b17b27"),
			AuthorID: "7c1cdb63-a73a-4385-b23e-3362d7f613aa",
			Text:     "I Wish There Was A Way To Know You’re In The Good Old Days, Before You’ve Actually Left Them.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("751b0f12-b108-4ee7-a0f9-5f7357510138"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Sometimes I’ll Start A Sentence And I Don’t Even Know Where It’s Going. I Just Hope I Find It Along The Way.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("9a8b6afc-f9b7-4565-b3cb-ddf82084f7c5"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "How Would I Describe Myself? Three Words. Hard-working. Alpha Male. Jackhammer. Merciless. Insatiable.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("f5913118-77a8-4b37-ab48-3bf26f1bc2a9"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "It’s Never Too Early For Ice-Cream, Jim.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("5280433d-92e6-412f-bc50-6e76a8c9cef0"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Why Are All These People Here? There Are Too Many People On This Earth. We Need A New Plague.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("4769a4c0-9dcc-4d2f-8ced-3a2d76fac056"),
			AuthorID: "7c1cdb63-a73a-4385-b23e-3362d7f613aa",
			Text:     "Andy Bernard Does Not Lose Contests. He Wins Them… Or He Quits Them Because They Are Unfair.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("1058eab1-0ec7-4f0e-9c96-a0515a699692"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "You Couldn’t Handle My Undivided Attention.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("d1eb15fa-18ca-4f61-8b45-8c323c533320"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "That’s What She Said!",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("cb35d496-1359-41bd-8ff3-c90d3b96aeec"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "Welcome to the Hotel Hell. Check-in time in now, check-out time is never.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("8bd275b1-e0a3-4a06-9c4a-47d24f91d119"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "I'm the owner.. the co-owner. With Satan!",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("2889eead-7f8d-4b49-8efc-536f3f3bee96"),
			AuthorID: "32a73786-28e8-4756-a2e2-ec6e3b9b8342",
			Text:     "I'm not gaining anything from this seminar. I'm a professional woman. The head of accounting. I'm in the healthiest relationship of my life. I just think it's insulting that Jan thinks we need this. And, apparently, judging from her outfit, Jan aspires to be a whore.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("0b491f4c-caa1-4d40-986d-58a69f4197f0"),
			AuthorID: "32a73786-28e8-4756-a2e2-ec6e3b9b8342",
			Text:     "Poop is raining from the ceilings. Poop!",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("ef1c11ae-32ab-4cc1-a4ea-337d25e49555"),
			AuthorID: "99ec8981-c512-4047-b645-1e571aab8ce3",
			Text:     "Stanley yelled at me today. That was one of the most frightening experiences of my life.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("1fd39e48-c73e-4beb-988a-3e867cd5ccf6"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Yes, it is true. I, Michael Scott, am signing up with an online dating service. Thousands of people have done it, and I am going to do it. I need a username. And... I have a great one [types]. Little kid lover. That way, people will know exactly where my priorities are at.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("dad98e0f-7831-4aa6-a5ab-68a66b2aa036"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "It's not that children make me uncomfortable, it's just that, why be a dad when you can be a fun uncle? I've never heard of anyone rebelling against their fun uncle.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("c3ea4dea-894c-4d49-81a1-78b1122c4311"),
			AuthorID: "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Text:     "The Japanese camp guards of World War II always chose one man to kill whenever a batch of new prisoners arrived. I always wondered how they chose the man who was to die. I think I would have been good at choosing the person.",
		},

		&model.Pin{
			ID:       uuid.FromStringOrNil("761b6b70-d041-47ad-9bf0-ef6e237dedee"),
			AuthorID: "d2052497-d23c-459c-a290-c66ff931bbb6",
			Text:     "Close your eyes. Picture a convict. What's he wearing? Nothing special, baseball cap on backwards, baggy pants... he says something ordinary like... 'yo, thats shizzle.' Okay. Now slowly open your eyes again. Who are you picturing? A black man? Wrong. That was a white woman. Surprised? Well, shame on you.",
		},
	}

	users = []*model.User{
		&model.User{
			ID:     "a080e9be-3fd5-4d74-852d-bf048c00fff2",
			Name:   "Dwight Schrute",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/c/cd/Dwight_Schrute.jpg",
		},

		&model.User{
			ID:     "e3cb89b7-2f06-4247-84ae-43265cd85f9d",
			Name:   "Jim Halpert",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/7/7e/Jim-halpert.jpg",
		},

		&model.User{
			ID:     "d2052497-d23c-459c-a290-c66ff931bbb6",
			Name:   "Michael Scott",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/d/dc/MichaelScott.png",
		},

		&model.User{
			ID:     "93c2daaa-084e-4228-8322-1abdb1d891cd",
			Name:   "Oscar Martinez",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/5/54/Oscar_Martinez_of_The_Office.jpg",
		},

		&model.User{
			ID:     "1c8fd81f-757b-4ebf-af78-40ffc0268146",
			Name:   "Phyllis Vance",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/f/ff/Phyllis_Lapin-Vance.jpg",
		},

		&model.User{
			ID:     "7c1cdb63-a73a-4385-b23e-3362d7f613aa",
			Name:   "Andy Bernard",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/8/84/Andy_Bernard_photoshot.jpg",
		},

		&model.User{
			ID:     "32a73786-28e8-4756-a2e2-ec6e3b9b8342",
			Name:   "Angela Martin",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/0/0b/Angela_Martin.jpg",
		},

		&model.User{
			ID:     "99ec8981-c512-4047-b645-1e571aab8ce3",
			Name:   "Ryan Howard",
			IsBot:  false,
			Avatar: "https://upload.wikimedia.org/wikipedia/en/9/91/Ryan_Howard_%28The_Office%29.jpg",
		},
	}
)

// OfficeFetcher returns mocked data
type OfficeFetcher struct{}

// FetchUsers returns the static user list defined above
func (of *OfficeFetcher) FetchUsers() ([]*model.User, error) {
	return users, nil
}

// FetchPins returns the static quotes as pins
func (of *OfficeFetcher) FetchPins() ([]*model.Pin, error) {
	return pins, nil
}

// NewOfficeFetcher sets up and returns new TheOffice fetcher
func NewOfficeFetcher() *OfficeFetcher {
	return &OfficeFetcher{}
}
