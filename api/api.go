package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	apiURL       = "https://api.open.fec.gov/v1/"
	schedulePath = "schedules/schedule_a"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ScheduleAResponse struct {
	APIVersion string     `json:"api_version"`
	Results    []Result   `json:"results"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	PerPage     int `json:"per_page"`
	LastIndexes struct {
		LastIndex                   string `json:"last_index"`
		LastContributionReceiptDate string `json:"last_contribution_receipt_date"`
	} `json:"last_indexes"`
	Pages int `json:"pages"`
	Count int `json:"count"`
}

type Result struct {
	ContributorName                    string    `json:"contributor_name"`
	ReceiptTypeFull                    string    `json:"receipt_type_full"`
	CandidateOffice                    string    `json:"candidate_office"`
	MemoedSubtotal                     bool      `json:"memoed_subtotal"`
	Contributor                        string    `json:"contributor"`
	ContributorStreet1                 string    `json:"contributor_street_1"`
	CommitteeID                        string    `json:"committee_id"`
	DonorCommitteeName                 string    `json:"donor_committee_name"`
	UnusedContbrID                     string    `json:"unused_contbr_id"`
	RecipientCommitteeDesignation      string    `json:"recipient_committee_designation"`
	MemoCodeFull                       string    `json:"memo_code_full"`
	ContributorState                   string    `json:"contributor_state"`
	ContributionReceiptDate            string    `json:"contribution_receipt_date"`
	ConduitCommitteeState              string    `json:"conduit_committee_state"`
	ScheduleType                       string    `json:"schedule_type"`
	ContributorMiddleName              string    `json:"contributor_middle_name"`
	AmendmentIndicator                 string    `json:"amendment_indicator"`
	MemoText                           string    `json:"memo_text"`
	EntityType                         string    `json:"entity_type"`
	AmendmentIndicatorDesc             string    `json:"amendment_indicator_desc"`
	LinkID                             int64     `json:"link_id"`
	LineNumberLabel                    string    `json:"line_number_label"`
	ConduitCommitteeStreet1            string    `json:"conduit_committee_street1"`
	ConduitCommitteeID                 string    `json:"conduit_committee_id"`
	ContributorStreet2                 string    `json:"contributor_street_2"`
	LineNumber                         string    `json:"line_number"`
	ContributorEmployer                string    `json:"contributor_employer"`
	ElectionType                       string    `json:"election_type"`
	FecElectionTypeDesc                string    `json:"fec_election_type_desc"`
	ContributorSuffix                  string    `json:"contributor_suffix"`
	ContributorOccupation              string    `json:"contributor_occupation"`
	BackReferenceScheduleName          string    `json:"back_reference_schedule_name"`
	ContributorPrefix                  string    `json:"contributor_prefix"`
	CandidateOfficeState               string    `json:"candidate_office_state"`
	FileNumber                         int       `json:"file_number"`
	CandidateSuffix                    string    `json:"candidate_suffix"`
	CommitteeName                      string    `json:"committee_name"`
	ContributionReceiptAmount          float32   `json:"contribution_receipt_amount"`
	ReceiptTypeDesc                    string    `json:"receipt_type_desc"`
	CandidateMiddleName                string    `json:"candidate_middle_name"`
	CandidateFirstName                 string    `json:"candidate_first_name"`
	CandidateLastName                  string    `json:"candidate_last_name"`
	ElectionTypeFull                   string    `json:"election_type_full"`
	FecElectionYear                    string    `json:"fec_election_year"`
	CandidateOfficeDistrict            string    `json:"candidate_office_district"`
	CandidateID                        string    `json:"candidate_id"`
	ConduitCommitteeName               string    `json:"conduit_committee_name"`
	IsIndividual                       bool      `json:"is_individual"`
	SubID                              string    `json:"sub_id"`
	ContributorAggregateYtd            float32   `json:"contributor_aggregate_ytd"`
	ReportType                         string    `json:"report_type"`
	ConduitCommitteeStreet2            string    `json:"conduit_committee_street2"`
	OriginalSubID                      string    `json:"original_sub_id"`
	EntityTypeDesc                     string    `json:"entity_type_desc"`
	ReportYear                         int       `json:"report_year"`
	ScheduleTypeFull                   string    `json:"schedule_type_full"`
	CandidateName                      string    `json:"candidate_name"`
	LoadDate                           time.Time `json:"load_date"`
	ContributorZip                     string    `json:"contributor_zip"`
	NationalCommitteeNonfederalAccount string    `json:"national_committee_nonfederal_account"`
	IncreasedLimit                     string    `json:"increased_limit"`
	BackReferenceTransactionID         string    `json:"back_reference_transaction_id"`
	CandidateOfficeFull                string    `json:"candidate_office_full"`
	ContributorLastName                string    `json:"contributor_last_name"`
	ContributorID                      string    `json:"contributor_id"`
	TwoYearTransactionPeriod           int       `json:"two_year_transaction_period"`
	ImageNumber                        string    `json:"image_number"`
	RecipientCommitteeOrgType          string    `json:"recipient_committee_org_type"`
	ContributorFirstName               string    `json:"contributor_first_name"`
	ConduitCommitteeCity               string    `json:"conduit_committee_city"`
	PdfURL                             string    `json:"pdf_url"`
	FilingForm                         string    `json:"filing_form"`
	CandidatePrefix                    string    `json:"candidate_prefix"`
	ContributorCity                    string    `json:"contributor_city"`
	Committee                          struct {
		Party                   string   `json:"party"`
		PartyFull               string   `json:"party_full"`
		Street1                 string   `json:"street_1"`
		State                   string   `json:"state"`
		OrganizationType        string   `json:"organization_type"`
		TreasurerName           string   `json:"treasurer_name"`
		CommitteeID             string   `json:"committee_id"`
		Designation             string   `json:"designation"`
		Cycles                  []int    `json:"cycles"`
		Zip                     string   `json:"zip"`
		OrganizationTypeFull    string   `json:"organization_type_full"`
		FilingFrequency         string   `json:"filing_frequency"`
		StateFull               string   `json:"state_full"`
		AffiliatedCommitteeName string   `json:"affiliated_committee_name"`
		CandidateIds            []string `json:"candidate_ids"`
		Name                    string   `json:"name"`
		DesignationFull         string   `json:"designation_full"`
		CommitteeType           string   `json:"committee_type"`
		City                    string   `json:"city"`
		Street2                 string   `json:"street_2"`
		CommitteeTypeFull       string   `json:"committee_type_full"`
		Cycle                   int      `json:"cycle"`
	} `json:"committee"`
	CandidateOfficeStateFull string `json:"candidate_office_state_full"`
	ConduitCommitteeZip      string `json:"conduit_committee_zip"`
	TransactionID            string `json:"transaction_id"`
	RecipientCommitteeType   string `json:"recipient_committee_type"`
	MemoCode                 string `json:"memo_code"`
	ReceiptType              string `json:"receipt_type"`
}

// ErrEOR is the api end-of-records
var ErrEOR = errors.New("end of records")
var perPage = "50"

const committeeID = "C00580100"

// GetContributions makes an fec API request for a single page
func GetContributions(c Doer, zip, year, lastIndex, lastContributionReceiptDate, apiKey string) (*ScheduleAResponse, error) {
	url := fmt.Sprintf("%s%s?per_page=%s&api_key=%s&committee_id=%s&contributor_zip=%s&two_year_transaction_period=%s",
		apiURL,
		schedulePath,
		perPage,
		apiKey,
		committeeID,
		zip,
		year,
	)
	if lastIndex != "" && lastContributionReceiptDate != "" {
		url += fmt.Sprintf("&last_index=%s&last_contribution_receipt_date=%s", lastIndex, lastContributionReceiptDate)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var s ScheduleAResponse
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	if len(s.Results) == 0 {
		return nil, ErrEOR
	}
	return &s, nil
}

// GetContributionsPaged makes an fec API request for all pages
func GetContributionsPaged(c Doer, zip, year, apiKey string) ([]Result, error) {
	var lastIndex string
	var lastContributionReceiptDate string
	var results []Result
	for {
		resp, err := GetContributions(c, zip, year, lastIndex, lastContributionReceiptDate, apiKey)
		if err != nil {
			if err == ErrEOR {
				break
			} else {
				return nil, err
			}
		}
		lastIndex = resp.Pagination.LastIndexes.LastIndex
		lastContributionReceiptDate = resp.Pagination.LastIndexes.LastContributionReceiptDate
		results = append(results, resp.Results...)
	}
	fmt.Println(results)
	return results, nil
}
