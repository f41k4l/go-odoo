package odoo

// HelpdeskTicketType represents helpdesk.ticket.type model.
type HelpdeskTicketType struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HelpdeskTicketTypes represents array of helpdesk.ticket.type model.
type HelpdeskTicketTypes []HelpdeskTicketType

// HelpdeskTicketTypeModel is the odoo model name.
const HelpdeskTicketTypeModel = "helpdesk.ticket.type"

// Many2One convert HelpdeskTicketType to *Many2One.
func (htt *HelpdeskTicketType) Many2One() *Many2One {
	return NewMany2One(htt.Id.Get(), "")
}

// CreateHelpdeskTicketType creates a new helpdesk.ticket.type model and returns its id.
func (c *Client) CreateHelpdeskTicketType(htt *HelpdeskTicketType) (int64, error) {
	ids, err := c.CreateHelpdeskTicketTypes([]*HelpdeskTicketType{htt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketType creates a new helpdesk.ticket.type model and returns its id.
func (c *Client) CreateHelpdeskTicketTypes(htts []*HelpdeskTicketType) ([]int64, error) {
	var vv []interface{}
	for _, v := range htts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketTypeModel, vv, nil)
}

// UpdateHelpdeskTicketType updates an existing helpdesk.ticket.type record.
func (c *Client) UpdateHelpdeskTicketType(htt *HelpdeskTicketType) error {
	return c.UpdateHelpdeskTicketTypes([]int64{htt.Id.Get()}, htt)
}

// UpdateHelpdeskTicketTypes updates existing helpdesk.ticket.type records.
// All records (represented by ids) will be updated by htt values.
func (c *Client) UpdateHelpdeskTicketTypes(ids []int64, htt *HelpdeskTicketType) error {
	return c.Update(HelpdeskTicketTypeModel, ids, htt, nil)
}

// DeleteHelpdeskTicketType deletes an existing helpdesk.ticket.type record.
func (c *Client) DeleteHelpdeskTicketType(id int64) error {
	return c.DeleteHelpdeskTicketTypes([]int64{id})
}

// DeleteHelpdeskTicketTypes deletes existing helpdesk.ticket.type records.
func (c *Client) DeleteHelpdeskTicketTypes(ids []int64) error {
	return c.Delete(HelpdeskTicketTypeModel, ids)
}

// GetHelpdeskTicketType gets helpdesk.ticket.type existing record.
func (c *Client) GetHelpdeskTicketType(id int64) (*HelpdeskTicketType, error) {
	htts, err := c.GetHelpdeskTicketTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*htts)[0]), nil
}

// GetHelpdeskTicketTypes gets helpdesk.ticket.type existing records.
func (c *Client) GetHelpdeskTicketTypes(ids []int64) (*HelpdeskTicketTypes, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.Read(HelpdeskTicketTypeModel, ids, nil, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHelpdeskTicketType finds helpdesk.ticket.type record by querying it with criteria.
func (c *Client) FindHelpdeskTicketType(criteria *Criteria) (*HelpdeskTicketType, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.SearchRead(HelpdeskTicketTypeModel, criteria, NewOptions().Limit(1), htts); err != nil {
		return nil, err
	}
	return &((*htts)[0]), nil
}

// FindHelpdeskTicketTypes finds helpdesk.ticket.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketTypes(criteria *Criteria, options *Options) (*HelpdeskTicketTypes, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.SearchRead(HelpdeskTicketTypeModel, criteria, options, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHelpdeskTicketTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HelpdeskTicketTypeModel, criteria, options)
}

// FindHelpdeskTicketTypeId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
