package model

type CTNTrades struct {
	Success			bool
	instrument 		string
	Data	  		[]Data
	startDateTime	string
	endDateTime		string
}

type Data struct {
	Instrument		string
	Start			string
	End				string
	Low				float64
	High			float64
	Volume			float64
	Open			float64
	Close			float64
	GoinGeckoPrice	float64
}
