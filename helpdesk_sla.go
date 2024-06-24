package odoo

// HelpdeskSla represents helpdesk.sla model.
type HelpdeskSla struct {
	Active          *Bool      `xmlrpc:"active,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description     *String    `xmlrpc:"description,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	ExcludeStageIds *Relation  `xmlrpc:"exclude_stage_ids,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	PartnerIds      *Relation  `xmlrpc:"partner_ids,omitempty"`
	Priority        *Selection `xmlrpc:"priority,omitempty"`
	StageId         *Many2One  `xmlrpc:"stage_id,omitempty"`
	TagIds          *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId          *Many2One  `xmlrpc:"team_id,omitempty"`
	TicketCount     *Int       `xmlrpc:"ticket_count,omitempty"`
	TicketTypeIds   *Relation  `xmlrpc:"ticket_type_ids,omitempty"`
	Time            *Float     `xmlrpc:"time,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HelpdeskSlas represents array of helpdesk.sla model.
type HelpdeskSlas []HelpdeskSla

// HelpdeskSlaModel is the odoo model name.
const HelpdeskSlaModel = "helpdesk.sla"

// Many2One convert HelpdeskSla to *Many2One.
func (hs *HelpdeskSla) Many2One() *Many2One {
	return NewMany2One(hs.Id.Get(), "")
}

// CreateHelpdeskSla creates a new helpdesk.sla model and returns its id.
func (c *Client) CreateHelpdeskSla(hs *HelpdeskSla) (int64, error) {
	ids, err := c.CreateHelpdeskSlas([]*HelpdeskSla{hs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskSla creates a new helpdesk.sla model and returns its id.
func (c *Client) CreateHelpdeskSlas(hss []*HelpdeskSla) ([]int64, error) {
	var vv []interface{}
	for _, v := range hss {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskSlaModel, vv, nil)
}

// UpdateHelpdeskSla updates an existing helpdesk.sla record.
func (c *Client) UpdateHelpdeskSla(hs *HelpdeskSla) error {
	return c.UpdateHelpdeskSlas([]int64{hs.Id.Get()}, hs)
}

// UpdateHelpdeskSlas updates existing helpdesk.sla records.
// All records (represented by ids) will be updated by hs values.
func (c *Client) UpdateHelpdeskSlas(ids []int64, hs *HelpdeskSla) error {
	return c.Update(HelpdeskSlaModel, ids, hs, nil)
}

// DeleteHelpdeskSla deletes an existing helpdesk.sla record.
func (c *Client) DeleteHelpdeskSla(id int64) error {
	return c.DeleteHelpdeskSlas([]int64{id})
}

// DeleteHelpdeskSlas deletes existing helpdesk.sla records.
func (c *Client) DeleteHelpdeskSlas(ids []int64) error {
	return c.Delete(HelpdeskSlaModel, ids)
}

// GetHelpdeskSla gets helpdesk.sla existing record.
func (c *Client) GetHelpdeskSla(id int64) (*HelpdeskSla, error) {
	hss, err := c.GetHelpdeskSlas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hss)[0]), nil
}

// GetHelpdeskSlas gets helpdesk.sla existing records.
func (c *Client) GetHelpdeskSlas(ids []int64) (*HelpdeskSlas, error) {
	hss := &HelpdeskSlas{}
	if err := c.Read(HelpdeskSlaModel, ids, nil, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHelpdeskSla finds helpdesk.sla record by querying it with criteria.
func (c *Client) FindHelpdeskSla(criteria *Criteria) (*HelpdeskSla, error) {
	hss := &HelpdeskSlas{}
	if err := c.SearchRead(HelpdeskSlaModel, criteria, NewOptions().Limit(1), hss); err != nil {
		return nil, err
	}
	return &((*hss)[0]), nil
}

// FindHelpdeskSlas finds helpdesk.sla records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlas(criteria *Criteria, options *Options) (*HelpdeskSlas, error) {
	hss := &HelpdeskSlas{}
	if err := c.SearchRead(HelpdeskSlaModel, criteria, options, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHelpdeskSlaIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskSlaModel, criteria, options)
}

// FindHelpdeskSlaId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskSlaId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskSlaModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
