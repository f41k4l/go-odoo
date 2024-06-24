package odoo

// HelpdeskTeam represents helpdesk.team model.
type HelpdeskTeam struct {
	AccessInstructionMessage     *String     `xmlrpc:"access_instruction_message,omitempty"`
	Active                       *Bool       `xmlrpc:"active,omitempty"`
	AliasBouncedContent          *String     `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact                 *Selection  `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults                *String     `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain                  *String     `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId                *Many2One   `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail                   *String     `xmlrpc:"alias_email,omitempty"`
	AliasEmailFrom               *String     `xmlrpc:"alias_email_from,omitempty"`
	AliasForceThreadId           *Int        `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasFullName                *String     `xmlrpc:"alias_full_name,omitempty"`
	AliasId                      *Many2One   `xmlrpc:"alias_id,omitempty"`
	AliasIncomingLocal           *Bool       `xmlrpc:"alias_incoming_local,omitempty"`
	AliasModelId                 *Many2One   `xmlrpc:"alias_model_id,omitempty"`
	AliasName                    *String     `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId           *Many2One   `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId          *Int        `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasStatus                  *Selection  `xmlrpc:"alias_status,omitempty"`
	AllowPortalTicketClosing     *Bool       `xmlrpc:"allow_portal_ticket_closing,omitempty"`
	AssignMethod                 *Selection  `xmlrpc:"assign_method,omitempty"`
	AutoAssignment               *Bool       `xmlrpc:"auto_assignment,omitempty"`
	AutoCloseDay                 *Int        `xmlrpc:"auto_close_day,omitempty"`
	AutoCloseTicket              *Bool       `xmlrpc:"auto_close_ticket,omitempty"`
	CanPublish                   *Bool       `xmlrpc:"can_publish,omitempty"`
	Color                        *Int        `xmlrpc:"color,omitempty"`
	CompanyId                    *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One   `xmlrpc:"create_uid,omitempty"`
	Description                  *String     `xmlrpc:"description,omitempty"`
	DisplayName                  *String     `xmlrpc:"display_name,omitempty"`
	FeatureFormUrl               *String     `xmlrpc:"feature_form_url,omitempty"`
	FromStageIds                 *Relation   `xmlrpc:"from_stage_ids,omitempty"`
	HasExternalMailServer        *Bool       `xmlrpc:"has_external_mail_server,omitempty"`
	HasMessage                   *Bool       `xmlrpc:"has_message,omitempty"`
	Id                           *Int        `xmlrpc:"id,omitempty"`
	IsPublished                  *Bool       `xmlrpc:"is_published,omitempty"`
	MemberIds                    *Relation   `xmlrpc:"member_ids,omitempty"`
	MessageAttachmentCount       *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Name                         *String     `xmlrpc:"name,omitempty"`
	OpenTicketCount              *Int        `xmlrpc:"open_ticket_count,omitempty"`
	PortalShowRating             *Bool       `xmlrpc:"portal_show_rating,omitempty"`
	PrivacyVisibility            *Selection  `xmlrpc:"privacy_visibility,omitempty"`
	PrivacyVisibilityWarning     *String     `xmlrpc:"privacy_visibility_warning,omitempty"`
	RatingAvg                    *Float      `xmlrpc:"rating_avg,omitempty"`
	RatingAvgPercentage          *Float      `xmlrpc:"rating_avg_percentage,omitempty"`
	RatingCount                  *Int        `xmlrpc:"rating_count,omitempty"`
	RatingIds                    *Relation   `xmlrpc:"rating_ids,omitempty"`
	RatingPercentageSatisfaction *Int        `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	ResourceCalendarId           *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	Sequence                     *Int        `xmlrpc:"sequence,omitempty"`
	ShowKnowledgeBase            *Bool       `xmlrpc:"show_knowledge_base,omitempty"`
	SlaFailed                    *Int        `xmlrpc:"sla_failed,omitempty"`
	SlaPolicyCount               *Int        `xmlrpc:"sla_policy_count,omitempty"`
	StageIds                     *Relation   `xmlrpc:"stage_ids,omitempty"`
	SuccessRate                  *Float      `xmlrpc:"success_rate,omitempty"`
	TicketClosed                 *Int        `xmlrpc:"ticket_closed,omitempty"`
	TicketIds                    *Relation   `xmlrpc:"ticket_ids,omitempty"`
	TicketProperties             interface{} `xmlrpc:"ticket_properties,omitempty"`
	ToStageId                    *Many2One   `xmlrpc:"to_stage_id,omitempty"`
	UnassignedTickets            *Int        `xmlrpc:"unassigned_tickets,omitempty"`
	UrgentTicket                 *Int        `xmlrpc:"urgent_ticket,omitempty"`
	UseAlias                     *Bool       `xmlrpc:"use_alias,omitempty"`
	UseCoupons                   *Bool       `xmlrpc:"use_coupons,omitempty"`
	UseCreditNotes               *Bool       `xmlrpc:"use_credit_notes,omitempty"`
	UseFsm                       *Bool       `xmlrpc:"use_fsm,omitempty"`
	UseHelpdeskSaleTimesheet     *Bool       `xmlrpc:"use_helpdesk_sale_timesheet,omitempty"`
	UseHelpdeskTimesheet         *Bool       `xmlrpc:"use_helpdesk_timesheet,omitempty"`
	UseProductRepairs            *Bool       `xmlrpc:"use_product_repairs,omitempty"`
	UseProductReturns            *Bool       `xmlrpc:"use_product_returns,omitempty"`
	UseRating                    *Bool       `xmlrpc:"use_rating,omitempty"`
	UseSla                       *Bool       `xmlrpc:"use_sla,omitempty"`
	UseTwitter                   *Bool       `xmlrpc:"use_twitter,omitempty"`
	UseWebsiteHelpdeskForm       *Bool       `xmlrpc:"use_website_helpdesk_form,omitempty"`
	UseWebsiteHelpdeskForum      *Bool       `xmlrpc:"use_website_helpdesk_forum,omitempty"`
	UseWebsiteHelpdeskKnowledge  *Bool       `xmlrpc:"use_website_helpdesk_knowledge,omitempty"`
	UseWebsiteHelpdeskLivechat   *Bool       `xmlrpc:"use_website_helpdesk_livechat,omitempty"`
	UseWebsiteHelpdeskSlides     *Bool       `xmlrpc:"use_website_helpdesk_slides,omitempty"`
	WebsiteFormViewId            *Many2One   `xmlrpc:"website_form_view_id,omitempty"`
	WebsiteId                    *Many2One   `xmlrpc:"website_id,omitempty"`
	WebsiteMenuId                *Many2One   `xmlrpc:"website_menu_id,omitempty"`
	WebsiteMessageIds            *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WebsitePublished             *Bool       `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                   *String     `xmlrpc:"website_url,omitempty"`
	WriteDate                    *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// HelpdeskTeams represents array of helpdesk.team model.
type HelpdeskTeams []HelpdeskTeam

// HelpdeskTeamModel is the odoo model name.
const HelpdeskTeamModel = "helpdesk.team"

// Many2One convert HelpdeskTeam to *Many2One.
func (ht *HelpdeskTeam) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeam(ht *HelpdeskTeam) (int64, error) {
	ids, err := c.CreateHelpdeskTeams([]*HelpdeskTeam{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeams(hts []*HelpdeskTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTeamModel, vv, nil)
}

// UpdateHelpdeskTeam updates an existing helpdesk.team record.
func (c *Client) UpdateHelpdeskTeam(ht *HelpdeskTeam) error {
	return c.UpdateHelpdeskTeams([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTeams updates existing helpdesk.team records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTeams(ids []int64, ht *HelpdeskTeam) error {
	return c.Update(HelpdeskTeamModel, ids, ht, nil)
}

// DeleteHelpdeskTeam deletes an existing helpdesk.team record.
func (c *Client) DeleteHelpdeskTeam(id int64) error {
	return c.DeleteHelpdeskTeams([]int64{id})
}

// DeleteHelpdeskTeams deletes existing helpdesk.team records.
func (c *Client) DeleteHelpdeskTeams(ids []int64) error {
	return c.Delete(HelpdeskTeamModel, ids)
}

// GetHelpdeskTeam gets helpdesk.team existing record.
func (c *Client) GetHelpdeskTeam(id int64) (*HelpdeskTeam, error) {
	hts, err := c.GetHelpdeskTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// GetHelpdeskTeams gets helpdesk.team existing records.
func (c *Client) GetHelpdeskTeams(ids []int64) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.Read(HelpdeskTeamModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeam finds helpdesk.team record by querying it with criteria.
func (c *Client) FindHelpdeskTeam(criteria *Criteria) (*HelpdeskTeam, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	return &((*hts)[0]), nil
}

// FindHelpdeskTeams finds helpdesk.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeams(criteria *Criteria, options *Options) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskTeamModel, criteria, options)
}

// FindHelpdeskTeamId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
