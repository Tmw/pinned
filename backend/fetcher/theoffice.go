package fetcher

// the OfficeFetcher is used to provide Pinned with Mock data for demo purposes.
// it is featuring The Office quotes from https://github.com/anderskristo/the-office-quotes
// and avatars used from wikipedia. (soon.)



import (
	"encoding/json"

	"github.com/tmw/pinned/backend/model"
)

const (
	quotesJson = `
	[
   {
      "id":"fee39bce-f367-42db-bb7c-cf5c41bc6729",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Jim is gone. He's gone. I miss him so much. Ooooh I cry myself to sleep, Jim! False. I do not miss him."
   },
   {
      "id":"5ef7e9be-d28d-4f97-bf1e-98f16b1829b5",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Would I ever leave this company? Look, I’m all about loyalty. In fact, I feel like part of what I’m being paid for here is my loyalty. But if there were somewhere else that valued loyalty more highly… I’m going wherever they value loyalty the most."
   },
   {
      "id":"63f2192d-cbf4-430a-94f6-b16d62c322b4",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"I am fast. To give you a reference point I am somewhere between a snake and a mongoose… And a panther."
   },
   {
      "id":"dae17ad8-eaf1-4065-90e5-1917bccb1056",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Security in this office park is a joke. Last year I came to work with my spud-gun in a duffel bag. I sat at my desk all day with a rifle that shoots potatoes at 60 pounds per square inch. Can you imagine if I was deranged?"
   },
   {
      "id":"c3577211-d0a1-482c-a352-f06273d27db2",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"I grew up on a farm. I have seen animals having sex in every position imaginable. Goat on chicken. Chicken on goat. Couple of chickens doing a goat, couple of pigs watching."
   },
   {
      "id":"c252b37c-939c-4b69-b03d-9605ef4ede97",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Through concentration, I can raise and lower my cholesterol at will."
   },
   {
      "id":"dc070eef-f736-4ee1-9395-cd63ecd9c631",
      "author_id":"e3cb89b7-2f06-4247-84ae-43265cd85f9d",
      "text":"Right now, this is just a job. If I advance any higher in this company, then this would be my career. And, well, if this were my career? I'd have to throw myself in front of a train."
   },
   {
      "id":"97fb01ee-a39d-4e19-8074-d7e794246b70",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Am I going to tell them? No, I am not going to tell them. I don't see the point of that. As a doctor, you would not tell a patient if they had cancer."
   },
   {
      "id":"e298b7d0-13af-48e7-ac50-fd70b90c66e4",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Wikipedia Is The Best Thing Ever. Anyone In The World Can Write Anything They Want About Any Subject, So You Know You Are Getting The Best Possible Information."
   },
   {
      "id":"a253e194-e212-4d3d-a228-3c8fa957695f",
      "author_id":"93c2daaa-084e-4228-8322-1abdb1d891cd",
      "text":"Saddle Shoes With Denim? I Will Literally Call Protective Services."
   },
   {
      "id":"8f4a2b85-b3be-47e3-94b2-c975c4c0a410",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Hate To See You Leave But Love To Watch You Go. ‘Cause Of Your Butt."
   },
   {
      "id":"a784afbb-6dcb-403e-839a-db04a9bf9ac8",
      "author_id":"1c8fd81f-757b-4ebf-af78-40ffc0268146",
      "text":"I Wonder What People Like About Me? Probably My Jugs."
   },
   {
      "id":"2ffcf519-59d1-404a-b868-1a99d93846ad",
      "author_id":"e3cb89b7-2f06-4247-84ae-43265cd85f9d",
      "text":"Right Now This Is Just A Job. If I Advance Any Higher In This Company, Then This Would Be My Career. And Well, If This Were My Career I’d Have To Throw Myself In Front Of A Train."
   },
   {
      "id":"08952ddc-5c81-4dd8-adde-3a62c4ae39e4",
      "author_id":"e3cb89b7-2f06-4247-84ae-43265cd85f9d",
      "text":"Bears. Beets. Battlestar Galactica."
   },
   {
      "id":"ea5a0737-5982-4501-8d54-5b90c6db74d1",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Through Concentration, I Can Raise And Lower My Cholesterol At Will."
   },
   {
      "id":"23c49c7e-55c3-420b-b479-c4f3b03c3e6a",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"I Love Inside Jokes. Love To Be A Part Of One Someday."
   },
   {
      "id":"79413d0f-383c-4dc1-96c8-720f53b17b27",
      "author_id":"7c1cdb63-a73a-4385-b23e-3362d7f613aa",
      "text":"I Wish There Was A Way To Know You’re In The Good Old Days, Before You’ve Actually Left Them."
   },
   {
      "id":"751b0f12-b108-4ee7-a0f9-5f7357510138",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Sometimes I’ll Start A Sentence And I Don’t Even Know Where It’s Going. I Just Hope I Find It Along The Way."
   },
   {
      "id":"9a8b6afc-f9b7-4565-b3cb-ddf82084f7c5",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"How Would I Describe Myself? Three Words. Hard-working. Alpha Male. Jackhammer. Merciless. Insatiable."
   },
   {
      "id":"f5913118-77a8-4b37-ab48-3bf26f1bc2a9",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"It’s Never Too Early For Ice-Cream, Jim."
   },
   {
      "id":"5280433d-92e6-412f-bc50-6e76a8c9cef0",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Why Are All These People Here? There Are Too Many People On This Earth. We Need A New Plague."
   },
   {
      "id":"4769a4c0-9dcc-4d2f-8ced-3a2d76fac056",
      "author_id":"7c1cdb63-a73a-4385-b23e-3362d7f613aa",
      "text":"Andy Bernard Does Not Lose Contests. He Wins Them… Or He Quits Them Because They Are Unfair."
   },
   {
      "id":"1058eab1-0ec7-4f0e-9c96-a0515a699692",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"You Couldn’t Handle My Undivided Attention."
   },
   {
      "id":"d1eb15fa-18ca-4f61-8b45-8c323c533320",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"That’s What She Said!"
   },
   {
      "id":"cb35d496-1359-41bd-8ff3-c90d3b96aeec",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"Welcome to the Hotel Hell. Check-in time in now, check-out time is never."
   },
   {
      "id":"8bd275b1-e0a3-4a06-9c4a-47d24f91d119",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"I'm the owner.. the co-owner. With Satan!"
   },
   {
      "id":"2889eead-7f8d-4b49-8efc-536f3f3bee96",
      "author_id":"32a73786-28e8-4756-a2e2-ec6e3b9b8342",
      "text":"I'm not gaining anything from this seminar. I'm a professional woman. The head of accounting. I'm in the healthiest relationship of my life. I just think it's insulting that Jan thinks we need this. And, apparently, judging from her outfit, Jan aspires to be a whore."
   },
   {
      "id":"0b491f4c-caa1-4d40-986d-58a69f4197f0",
      "author_id":"32a73786-28e8-4756-a2e2-ec6e3b9b8342",
      "text":"Poop is raining from the ceilings. Poop!"
   },
   {
      "id":"ef1c11ae-32ab-4cc1-a4ea-337d25e49555",
      "author_id":"99ec8981-c512-4047-b645-1e571aab8ce3",
      "text":"Stanley yelled at me today. That was one of the most frightening experiences of my life."
   },
   {
      "id":"1fd39e48-c73e-4beb-988a-3e867cd5ccf6",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Yes, it is true. I, Michael Scott, am signing up with an online dating service. Thousands of people have done it, and I am going to do it. I need a username. And... I have a great one [types]. Little kid lover. That way, people will know exactly where my priorities are at."
   },
   {
      "id":"dad98e0f-7831-4aa6-a5ab-68a66b2aa036",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"It's not that children make me uncomfortable, it's just that, why be a dad when you can be a fun uncle? I've never heard of anyone rebelling against their fun uncle."
   },
   {
      "id":"c3ea4dea-894c-4d49-81a1-78b1122c4311",
      "author_id":"a080e9be-3fd5-4d74-852d-bf048c00fff2",
      "text":"The Japanese camp guards of World War II always chose one man to kill whenever a batch of new prisoners arrived. I always wondered how they chose the man who was to die. I think I would have been good at choosing the person."
   },
   {
      "id":"761b6b70-d041-47ad-9bf0-ef6e237dedee",
      "author_id":"d2052497-d23c-459c-a290-c66ff931bbb6",
      "text":"Close your eyes. Picture a convict. What's he wearing? Nothing special, baseball cap on backwards, baggy pants... he says something ordinary like... 'yo, thats shizzle.' Okay. Now slowly open your eyes again. Who are you picturing? A black man? Wrong. That was a white woman. Surprised? Well, shame on you."
   }
]
`;
usersJson = `
[
  {
    "id": "a080e9be-3fd5-4d74-852d-bf048c00fff2",
    "name": "Dwight Schrute",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "e3cb89b7-2f06-4247-84ae-43265cd85f9d",
    "name": "Jim Halpert",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "d2052497-d23c-459c-a290-c66ff931bbb6",
    "name": "Michael Scott",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "93c2daaa-084e-4228-8322-1abdb1d891cd",
    "name": "Oscar Martinez",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "1c8fd81f-757b-4ebf-af78-40ffc0268146",
    "name": "Phyllis Vance",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "7c1cdb63-a73a-4385-b23e-3362d7f613aa",
    "name": "Andy Bernard",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "32a73786-28e8-4756-a2e2-ec6e3b9b8342",
    "name": "Angela Martin",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  },
  {
    "id": "99ec8981-c512-4047-b645-1e571aab8ce3",
    "name": "Ryan Howard",
    "is_bot": false,
    "avatar": "http://placekitten.com/128/128"
  }
]
`
)


// OfficeFetcher returns mocked data
type OfficeFetcher struct {}

// FetchUsers fetches users from Slack
func (of *OfficeFetcher) FetchUsers() ([]*model.User, error) {

	var users []*model.User
	if err := json.Unmarshal([]byte(usersJson), &users); err != nil {
		return nil, err;
	}

	return users, nil
}

// FetchPins fetches pins from Slack
func (of *OfficeFetcher) FetchPins() ([]*model.Pin, error) {

	pins := []*model.Pin{}
	if err := json.Unmarshal([]byte(quotesJson), &pins); err != nil {
		return nil, err;
	}

	return pins, nil
}

// New sets up and returns a new SlackFetcher
func NewStaticFetcher() *OfficeFetcher {
	return &OfficeFetcher{}
}
