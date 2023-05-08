package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// new model
type Account_Infos struct {
	Account_id     int    `json:"account_id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Is_pass_change int    `json:"is_pass_change"`
}

type Login_Request_Data struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Remember_me bool   `json:"rememberme"`
}

type Login_Response_Data struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Remember string `json:"remember"`
}

type Change_Pass struct {
	User_id          int    `json:"userid"`
	Username         string `json:"username"`
	Old_password     string `json:"oldpassword"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirmpassword"`
}

// new model
type Dashboard_Info struct {
	Dashboard_id    int    `json:"dashboardid"`
	Bank_id         int    `json:"bankid"`
	Dashboard_title string `json:"dashboardtitle"`
	Dashboard_link  string `json:"dashboardlink"`
}

type Access_Dashboard_Apps struct {
	Access_dashboard_apps_id int `json:"accessdashboardappsid"`
	Dashboard_id             int `json:"dashboardid"`
	App_id                   int `json:"appid"`
}

type Access_Accounts_Banks struct {
	Access_accounts_banks_id int   `json:"accessaccountsbanksid"`
	Account_id               int   `json:"accountid"`
	Bank_id_array            []int `json:"bankidarray"`
}

// dashboard apps and tags view
type View_Access_Dashboard_Info struct {
	Dashboard_id    int         `json:"dashboardid"`
	Bank_id         int         `json:"bankin"`
	Dashboard_title string      `json:"dashboardtitle"`
	Dashboard_link  string      `json:"dashboardlink"`
	Tag_id_array    StringArray `json:"tagidarray"`
	Tag_list        StringArray `json:"taglist"`
}

// dashboard apps and tags view
type View_Access_Account_Info struct {
	Account_id int         `json:"accountd"`
	Username   string      `json:"username"`
	Bank_list  StringArray `json:"banklist"`
	Tag_list   StringArray `json:"taglist"`
}

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan StringArray value: %v", value)
	}
	s = strings.Trim(s, "{}")
	*a = strings.Split(s, ",")
	return nil
}

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	values := make([]string, len(a))
	for i, v := range a {
		values[i] = fmt.Sprintf("%v", v)
	}
	return "{" + strings.Join(values, ",") + "}", nil
}

type Email_Sender_Details struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Email_Reciever_Details struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Sendinblue_Request_Body struct {
	Sender       Email_Sender_Details     `json:"sender"`
	To           []Email_Reciever_Details `json:"to"`
	Subject      string                   `json:"subject"`
	Html_content string                   `json:"htmlcontent"`
}

type ContactUs_Request_Body struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
