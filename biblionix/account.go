package biblionix

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

type AccountResponse struct {
	XMLName                      xml.Name `xml:"root"`
	Text                         string   `xml:",chardata"`
	PurchaseCount                string   `xml:"purchase_count,attr"`
	PurchaseCountThisMonth       string   `xml:"purchase_count_this_month,attr"`
	LibraryID                    string   `xml:"library_id,attr"`
	ID                           string   `xml:"id,attr"`
	ExpirationPrintable          string   `xml:"expiration_printable,attr"`
	LastName                     string   `xml:"last_name,attr"`
	IndividualCards              string   `xml:"individual_cards,attr"`
	AddressChangable             string   `xml:"address_changable,attr"`
	SelfRenewal                  string   `xml:"self_renewal,attr"`
	Card                         string   `xml:"card,attr"`
	CardPrintable                string   `xml:"card_printable,attr"`
	CardImage                    string   `xml:"card_image,attr"`
	ReserveLimit                 string   `xml:"reserve_limit,attr"`
	Year                         string   `xml:"year,attr"`
	Saved                        string   `xml:"saved,attr"`
	Type                         string   `xml:"type,attr"`
	Type2                        string   `xml:"type2,attr"`
	PasswordChanged              string   `xml:"password_changed,attr"`
	PasswordChangedPrintable     string   `xml:"password_changed_printable,attr"`
	Versacard                    string   `xml:"versacard,attr"`
	VersacardHomeLibrary         string   `xml:"versacard_home_library,attr"`
	DoHistory                    string   `xml:"do_history,attr"`
	CirchistoryDissociateDays    string   `xml:"circhistory_dissociate_days,attr"`
	MemberSince                  string   `xml:"member_since,attr"`
	ForcePasswordReset           string   `xml:"force_password_reset,attr"`
	IllAllowed                   string   `xml:"ill_allowed,attr"`
	SessionReferral              string   `xml:"session_referral,attr"`
	PreventSms                   string   `xml:"prevent_sms,attr"`
	IsViaStaffLogin              string   `xml:"is_via_staff_login,attr"`
	NeedEventTermsSigned         string   `xml:"need_event_terms_signed,attr"`
	NeedRoomTermsSigned          string   `xml:"need_room_terms_signed,attr"`
	MeetsRoomBookingRequirements string   `xml:"meets_room_booking_requirements,attr"`
	HistoryIneligible            string   `xml:"history_ineligible,attr"`
	LinkedCard                   []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"id,attr"`
		CardNumber string `xml:"card_number,attr"`
	} `xml:"linked_card"`
	Address struct {
		Text                string `xml:",chardata"`
		ID                  string `xml:"id,attr"`
		Mailing             string `xml:"mailing,attr"`
		AddressLine1        string `xml:"address_line_1,attr"`
		AddressLine2        string `xml:"address_line_2,attr"`
		City                string `xml:"city,attr"`
		State               string `xml:"state,attr"`
		Zip                 string `xml:"zip,attr"`
		AddressConcatenated string `xml:"address_concatenated,attr"`
		ZipPrintable        string `xml:"zip_printable,attr"`
	} `xml:"address"`
	Database []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr"`
		Name             string `xml:"name,attr"`
		Description      string `xml:"description,attr"`
		URL              string `xml:"url,attr"`
		LogoWidth        string `xml:"logo_width,attr"`
		LogoHeight       string `xml:"logo_height,attr"`
		IeOnly           string `xml:"ie_only,attr"`
		IeOnlyText       string `xml:"ie_only_text,attr"`
		PasswordText     string `xml:"password_text,attr"`
		UsernameText     string `xml:"username_text,attr"`
		Method           string `xml:"method,attr"`
		NeedsDataOnclick string `xml:"needs_data_onclick,attr"`
		Gratis           string `xml:"gratis,attr"`
		DisabledByIp     string `xml:"disabled_by_ip,attr"`
		Param            []struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Value     string `xml:"value,attr"`
			SearchBox string `xml:"search_box,attr"`
		} `xml:"param"`
	} `xml:"database"`
	Notifications struct {
		Text                  string `xml:",chardata"`
		OverdueID             string `xml:"overdue_id,attr"`
		ReserveID             string `xml:"reserve_id,attr"`
		OverdueType           string `xml:"overdue_type,attr"`
		ReserveType           string `xml:"reserve_type,attr"`
		OverdueWarningID      string `xml:"overdue_warning_id,attr"`
		OverdueWarningType    string `xml:"overdue_warning_type,attr"`
		CheckoutreceiptID     string `xml:"checkoutreceipt_id,attr"`
		Branch                string `xml:"branch,attr"`
		AlternatePickupmethod string `xml:"alternate_pickupmethod,attr"`
		Language              string `xml:"language,attr"`
	} `xml:"notifications"`
	Name struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Printable string `xml:"printable,attr"`
		Name      string `xml:"name,attr"`
		Type      string `xml:"type,attr"`
	} `xml:"name"`
	Phone struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"id,attr"`
		Printable  string `xml:"printable,attr"`
		FullNumber string `xml:"full_number,attr"`
		Area       string `xml:"area,attr"`
		Exchange   string `xml:"exchange,attr"`
		Number     string `xml:"number,attr"`
		Type       string `xml:"type,attr"`
		Primary    string `xml:"primary,attr"`
		Note       string `xml:"note,attr"`
		IsParent   string `xml:"is_parent,attr"`
	} `xml:"phone"`
	Email struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Address   string `xml:"address,attr"`
		Note      string `xml:"note,attr"`
		Printable string `xml:"printable,attr"`
		IsParent  string `xml:"is_parent,attr"`
	} `xml:"email"`
	Item []struct {
		Text            string `xml:",chardata"`
		ID              string `xml:"id,attr"`
		TitleSort       string `xml:"title_sort,attr"`
		Title           string `xml:"title,attr"`
		Medium          string `xml:"medium,attr"`
		Author          string `xml:"author,attr"`
		AuthorSort      string `xml:"author_sort,attr"`
		OutRaw          string `xml:"out_raw,attr"`
		DueRaw          string `xml:"due_raw,attr"`
		Out             string `xml:"out,attr"`
		Due             string `xml:"due,attr"`
		Renewable       string `xml:"renewable,attr"`
		IndicateReserve string `xml:"indicate_reserve,attr"`
		Overdue         string `xml:"overdue,attr"`
		Biblio          string `xml:"biblio,attr"`
	} `xml:"item"`
	Reserve []struct {
		Text                     string `xml:",chardata"`
		ID                       string `xml:"id,attr"`
		VersacatReserve          string `xml:"versacat_reserve,attr"`
		TitleRaw                 string `xml:"title_raw,attr"`
		Article                  string `xml:"article,attr"`
		Title                    string `xml:"title,attr"`
		Subtitle                 string `xml:"subtitle,attr"`
		Author                   string `xml:"author,attr"`
		Medium                   string `xml:"medium,attr"`
		Position                 string `xml:"position,attr"`
		PositionUnknowableReason string `xml:"position_unknowable_reason,attr"`
		Copies                   string `xml:"copies,attr"`
		Available                string `xml:"available,attr"`
		Due                      string `xml:"due,attr"`
		Biblio                   string `xml:"biblio,attr"`
		ContactID                string `xml:"contact_id,attr"`
		PlacedRaw                string `xml:"placed_raw,attr"`
		IsIllRequest             string `xml:"is_ill_request,attr"`
		Location                 string `xml:"location,attr"`
		ReserveCullTime          string `xml:"reserve_cull_time,attr"`
	} `xml:"reserve"`
}

func (c *Client) Account(session string) (*AccountResponse, error) {
	resp, err := c.httpClient().PostForm(c.URL("/catalog/ajax_backend/account.xml.pl"), url.Values{
		"session": {session},
	})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get account failed: %s", resp.Status)
	}

	defer func() { _ = resp.Body.Close() }()

	var account AccountResponse
	err = xml.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}

	// clean up obfuscated strings
	for i, item := range account.Item {
		account.Item[i].Title = deobfuscate(item.Title)
		account.Item[i].Author = deobfuscate(item.Author)
	}

	return &account, nil
}
