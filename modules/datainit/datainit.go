package datainit

//Init test structures
type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Header      string `json:"header"`
	Description string `json:"description"`
	Address     string `json:"address"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ExternalURL string `json:"external_url"`
}

type IndexBanner struct {
	ID               string `json:"id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventDates       string `json:"event_dates"`
	IsHidden         bool   `json:"is_visible"`
}

type Contacts struct {
	Phone1  string `json:"phone1"`
	Phone2  string `json:"phone2"`
	Address string `json:"address"`
	WebURL  string `json:"web_url"`
	Email   string `json:"email"`
	MapCode string `json:"map_code"`
}

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Init test variables
var EventsData []Event
var IndexBannersData []IndexBanner
var ContactData Contacts
var UsersData []User

func InitTestData() {
	//Populate test variables
	EventsData = append(EventsData, Event{ID: "1", Name: "NTMEX 2003", Header: "NANO TECHNOLOGY EXPO", Description: "Нанотехнологии", Address: "Moscow", StartDate: "2003-12-20", EndDate: "2003-12-25", ExternalURL: "http://ntmex.ru"})
	EventsData = append(EventsData, Event{ID: "2", Name: "FESTIVAL NAUKI 2003", Header: "FESTIVAL NAUKI", Description: "Фестиваль науки", Address: "Moscow", StartDate: "2003-11-20", EndDate: "2003-11-25", ExternalURL: "http://festivalnauki.ru"})
	EventsData = append(EventsData, Event{ID: "3", Name: "МАКС 2003", Header: "МАКС", Description: "Международный авиасалон", Address: "Moscow", StartDate: "2003-10-20", EndDate: "2003-10-25", ExternalURL: "http://maks.ru"})
	EventsData = append(EventsData, Event{ID: "4", Name: "HIGHLOAD 2003", Header: "HIGHLOAD", Description: "Высоконагруженные ИТ-системы", Address: "Moscow", StartDate: "2003-09-20", EndDate: "2003-09-25", ExternalURL: "http://highload.ru"})

	IndexBannersData = append(IndexBannersData, IndexBanner{"1", "Highload 2019", "Highload++", "5-6 Ноября", false})
	IndexBannersData = append(IndexBannersData, IndexBanner{"2", "Jocker 2019", "Jocker", "5-6 Сентябрь", true})
	IndexBannersData = append(IndexBannersData, IndexBanner{"3", "РИТ 2019", "РИТ++", "5-6 Октябрь", false})
	IndexBannersData = append(IndexBannersData, IndexBanner{"4", "МАКС 2019", "МАКС", "5-6 Декабря", true})

	ContactData.Phone1 = "109857438946"
	ContactData.Phone2 = "109857438946"
	ContactData.Address = "Москва, Фиг-знает где 48/90"
	ContactData.Email = "mail@example.com"
	ContactData.WebURL = "https://example.com"
	ContactData.MapCode = "https://example.com/map?=12kflankjshiu34rcoqiy"

	UsersData = append(UsersData, User{"1", "test", "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"})
}
