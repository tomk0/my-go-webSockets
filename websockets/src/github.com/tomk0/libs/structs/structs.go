package structs

type MenuItem struct {
	ID       string  `json:"Id"`
	Name     string  `json:"Name"`
	Disc     string  `json:"Disc"`
	Price    float64 `json:"Price"`
	Amount   int     `json:"Amount"`
	Category string  `json:"Category"`
}

type MenuOut struct {

	ItemsAry []MenuItem `json:"Items"`
}

type DataOut struct {

	Used bool `json:"Used"`
	Menu MenuOut `json:"Menu"`

}

type CmdOut struct {

	Cmd string `json:"Cmd"`
	Data DataOut `json:"Data"`

}

type OrderIn struct {

	ITM_ID string `json:"ITM_ID"`
	FILL_ID string `json:"FILL_ID"`
	Amount int `json:"Amount"`
	Price float64 `json:"Price"`
	Notes string `json:"Notes"`

}

type DataIn struct {

	Order []OrderIn `json:"Order"`
}

type CmdIn struct {
	Cmd  string `json:"Cmd"`
	Data DataIn   `json:"Data"`
}