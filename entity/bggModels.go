package entity

import "encoding/xml"

type BGGCollection struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Totalitems string   `xml:"totalitems,attr"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Pubdate    string   `xml:"pubdate,attr"`
	Item       []struct {
		Text       string `xml:",chardata"`
		Objecttype string `xml:"objecttype,attr"`
		Objectid   string `xml:"objectid,attr"`
		Subtype    string `xml:"subtype,attr"`
		Collid     string `xml:"collid,attr"`
		Name       struct {
			Text      string `xml:",chardata"`
			Sortindex string `xml:"sortindex,attr"`
		} `xml:"name"`
		Yearpublished string `xml:"yearpublished"`
		Image         string `xml:"image"`
		Thumbnail     string `xml:"thumbnail"`
		Stats         struct {
			Text        string `xml:",chardata"`
			Minplayers  string `xml:"minplayers,attr"`
			Maxplayers  string `xml:"maxplayers,attr"`
			Minplaytime string `xml:"minplaytime,attr"`
			Maxplaytime string `xml:"maxplaytime,attr"`
			Playingtime string `xml:"playingtime,attr"`
			Numowned    string `xml:"numowned,attr"`
			Rating      struct {
				Text       string `xml:",chardata"`
				Value      string `xml:"value,attr"`
				Usersrated struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"usersrated"`
				Average struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"average"`
				Bayesaverage struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"bayesaverage"`
				Stddev struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"stddev"`
				Median struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"median"`
				Ranks struct {
					Text string `xml:",chardata"`
					Rank []struct {
						Text         string `xml:",chardata"`
						Type         string `xml:"type,attr"`
						ID           string `xml:"id,attr"`
						Name         string `xml:"name,attr"`
						Friendlyname string `xml:"friendlyname,attr"`
						Value        string `xml:"value,attr"`
						Bayesaverage string `xml:"bayesaverage,attr"`
					} `xml:"rank"`
				} `xml:"ranks"`
			} `xml:"rating"`
		} `xml:"stats"`
		Status struct {
			Text             string `xml:",chardata"`
			Own              string `xml:"own,attr"`
			Prevowned        string `xml:"prevowned,attr"`
			Fortrade         string `xml:"fortrade,attr"`
			Want             string `xml:"want,attr"`
			Wanttoplay       string `xml:"wanttoplay,attr"`
			Wanttobuy        string `xml:"wanttobuy,attr"`
			Wishlist         string `xml:"wishlist,attr"`
			Preordered       string `xml:"preordered,attr"`
			Lastmodified     string `xml:"lastmodified,attr"`
			Wishlistpriority string `xml:"wishlistpriority,attr"`
		} `xml:"status"`
		Numplays string `xml:"numplays"`
	} `xml:"item"`
}

type BGGThing struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description   string `xml:"description"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
		Minplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplayers"`
		Maxplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplayers"`
		Poll []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Title      string `xml:"title,attr"`
			Totalvotes string `xml:"totalvotes,attr"`
			Results    []struct {
				Text       string `xml:",chardata"`
				Numplayers string `xml:"numplayers,attr"`
				Result     []struct {
					Text     string `xml:",chardata"`
					Value    string `xml:"value,attr"`
					Numvotes string `xml:"numvotes,attr"`
					Level    string `xml:"level,attr"`
				} `xml:"result"`
			} `xml:"results"`
		} `xml:"poll"`
		Playingtime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"playingtime"`
		Minplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplaytime"`
		Maxplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplaytime"`
		Minage struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minage"`
		Link []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			ID      string `xml:"id,attr"`
			Value   string `xml:"value,attr"`
			Inbound string `xml:"inbound,attr"`
		} `xml:"link"`
		Statistics struct {
			Text    string `xml:",chardata"`
			Page    string `xml:"page,attr"`
			Ratings struct {
				Text       string `xml:",chardata"`
				Usersrated struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"usersrated"`
				Average struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"average"`
				Bayesaverage struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"bayesaverage"`
				Ranks struct {
					Text string `xml:",chardata"`
					Rank []struct {
						Text         string `xml:",chardata"`
						Type         string `xml:"type,attr"`
						ID           string `xml:"id,attr"`
						Name         string `xml:"name,attr"`
						Friendlyname string `xml:"friendlyname,attr"`
						Value        string `xml:"value,attr"`
						Bayesaverage string `xml:"bayesaverage,attr"`
					} `xml:"rank"`
				} `xml:"ranks"`
				Stddev struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"stddev"`
				Median struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"median"`
				Owned struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"owned"`
				Trading struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"trading"`
				Wanting struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"wanting"`
				Wishing struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"wishing"`
				Numcomments struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"numcomments"`
				Numweights struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"numweights"`
				Averageweight struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"averageweight"`
			} `xml:"ratings"`
		} `xml:"statistics"`
	} `xml:"item"`
}

type BGGErrors struct {
	XMLName xml.Name `xml:"errors"`
	Text    string   `xml:",chardata"`
	Error   struct {
		Text    string `xml:",chardata"`
		Message string `xml:"message"`
	} `xml:"error"`
}
