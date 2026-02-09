package kap

import "encoding/json"

// TokenResponse is returned by the GenerateToken endpoint.
type TokenResponse struct {
	Token string `json:"token"`
}

// Disclosure represents a single item in the disclosure list.
type Disclosure struct {
	DisclosureIndex       string   `json:"disclosureIndex"`
	DisclosureType        string   `json:"disclosureType"`
	DisclosureClass       string   `json:"disclosureClass"`
	SubReportIDs          []string `json:"subReportIds"`
	Title                 string   `json:"title"`
	CompanyID             string   `json:"companyId"`
	FundID                string   `json:"fundId,omitempty"`
	FundCode              string   `json:"fundCode,omitempty"`
	AcceptedDataFileTypes []string `json:"acceptedDataFileTypes"`
}

// DisclosureListParams holds optional filters for the Disclosures endpoint.
type DisclosureListParams struct {
	DisclosureClass string
	DisclosureType  string
	CompanyID       string
}

// LocalizedText holds Turkish and English translations.
type LocalizedText struct {
	TR *string `json:"tr"`
	EN *string `json:"en"`
}

// AttachmentURL represents a disclosure attachment link.
type AttachmentURL struct {
	URL      string `json:"url"`
	FileName string `json:"fileName"`
}

// PresentationItem holds structured presentation data. The Content field
// varies by disclosure type, so it is kept as raw JSON.
type PresentationItem struct {
	ID      string          `json:"id"`
	Content json.RawMessage `json:"content"`
}

// FlatDataItem holds flat data content. The Content field varies by
// disclosure type.
type FlatDataItem struct {
	ID      string          `json:"id"`
	Content json.RawMessage `json:"content"`
}

// HTMLMessage holds base64-encoded HTML content in Turkish and English.
type HTMLMessage struct {
	ID string  `json:"id"`
	TR *string `json:"tr"`
	EN *string `json:"en"`
}

// RelatedStock represents a related company or fund in a disclosure.
type RelatedStock struct {
	Code string `json:"code"`
}

// DisclosureDetail holds full details for a single disclosure.
type DisclosureDetail struct {
	DisclosureIndex        string             `json:"disclosureIndex"`
	SenderID               string             `json:"senderId"`
	SenderTitle            string             `json:"senderTitle"`
	SenderExchCodes        []string           `json:"senderExchCodes"`
	BehalfSenderID         string             `json:"behalfSenderId,omitempty"`
	BehalfSenderTitle      string             `json:"behalfSenderTitle,omitempty"`
	BehalfSenderExchCodes  []string           `json:"behalfSenderExchCodes,omitempty"`
	BehalfFundCode         string             `json:"behalfFundCode,omitempty"`
	BehalfFundTitle        string             `json:"behalfFundTitle,omitempty"`
	DisclosureReason       string             `json:"disclosureReason"`
	DisclosureDelayStatus  string             `json:"disclosureDelayStatus,omitempty"`
	RelatedDisclosureIndex string             `json:"relatedDisclosureIndex,omitempty"`
	DisclosureType         string             `json:"disclosureType"`
	DisclosureClass        string             `json:"disclosureClass"`
	Subject                LocalizedText      `json:"subject"`
	Consolidation          string             `json:"consolidation,omitempty"`
	Year                   string             `json:"year,omitempty"`
	Period                 *LocalizedText     `json:"period,omitempty"`
	RelatedStocks          []RelatedStock     `json:"relatedStocks"`
	Summary                LocalizedText      `json:"summary"`
	Time                   string             `json:"time"`
	Link                   string             `json:"link"`
	AttachmentURLs         []AttachmentURL    `json:"attachmentUrls"`
	EventType              string             `json:"eventType,omitempty"`
	EventID                int                `json:"eventId,omitempty"`
	Presentation           []PresentationItem `json:"presentation"`
	FlatData               []FlatDataItem     `json:"flatData"`
	HTMLMessages           []HTMLMessage      `json:"htmlMessages"`
}

// LastDisclosureIndexResponse is returned by the LastDisclosureIndex endpoint.
type LastDisclosureIndexResponse struct {
	LastDisclosureIndex string `json:"lastDisclosureIndex"`
}

// CAEventStatus holds the status of a corporate action event.
type CAEventStatus struct {
	RefID        string `json:"refId"`
	Status       string `json:"status"`
	StatusReason string `json:"statusReason,omitempty"`
	CompleteDate string `json:"completeDate,omitempty"`
}

// Member represents a KAP member company.
type Member struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	StockCode  string `json:"stockCode"`
	MemberType string `json:"memberType"`
	KFIFUrl    string `json:"kfifUrl,omitempty"`
}

// Security holds security information for a listed company.
type Security struct {
	ISIN              string  `json:"isin"`
	ISINDesc          string  `json:"isinDesc"`
	BorsaKodu         string  `json:"borsaKodu"`
	TakasKodu         string  `json:"takasKodu"`
	TertipGroup       string  `json:"tertipGroup"`
	Capital           float64 `json:"capital"`
	CurrentCapital    float64 `json:"currentCapital"`
	GroupCode         string  `json:"groupCode"`
	GroupCodeDesc     string  `json:"groupCodeDesc"`
	BorsadaIslemeAcik bool    `json:"borsadaIslemeAcik"`
}

// CompanyInfo holds summary information about a company in the member
// securities response.
type CompanyInfo struct {
	ID                     string  `json:"id"`
	MemberType             string  `json:"memberType"`
	SermayeSistemi         string  `json:"sermayeSistemi,omitempty"`
	KayitliSermayeTavani   float64 `json:"kayitliSermayeTavani,omitempty"`
	KstSonGecerlilikTarihi string  `json:"kstSonGecerlilikTarihi,omitempty"`
	SirketUnvan            string  `json:"sirketUnvan,omitempty"`
	MksMbrID               string  `json:"mksMbrId,omitempty"`
}

// MemberSecurities pairs a company with its securities.
type MemberSecurities struct {
	Member     CompanyInfo `json:"member"`
	Securities []Security  `json:"securities"`
}

// DetailField holds a single detail field for company or fund detail
// responses. The Value field varies by field key.
type DetailField struct {
	NameTR          string          `json:"nameTr"`
	NameEN          string          `json:"nameEn"`
	Key             string          `json:"key"`
	PublishDateTime *string         `json:"publishDateTime"`
	Value           json.RawMessage `json:"value"`
	CodeKey         string          `json:"codeKey,omitempty"`
}

// Fund represents a fund in the fund list.
type Fund struct {
	FundID           int    `json:"fundId"`
	FundName         string `json:"fundName"`
	FundCode         string `json:"fundCode"`
	FundType         string `json:"fundType"`
	FundClass        string `json:"fundClass"`
	FundExpiry       string `json:"fundExpiry"`
	FundState        string `json:"fundState"`
	Title            string `json:"title"`
	UmbMemberTypes   string `json:"umbMemberTypes"`
	FundMemberTypes  string `json:"fundMemberTypes"`
	KAPUrl           string `json:"kapUrl"`
	NonInactiveCount int    `json:"nonInactiveCount"`
	FundCompanyID    string `json:"fundCompanyId"`
	FundCompanyTitle string `json:"fundCompanyTitle"`
}

// FundListParams holds optional filters for the Funds endpoint.
type FundListParams struct {
	FundState []string
	FundClass []string
	FundType  []string
}
