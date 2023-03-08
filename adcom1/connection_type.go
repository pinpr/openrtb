package adcom1

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ConnectionType represents options for the type of device connectivity.
type ConnectionType int8

// Options for the type of device connectivity.
const (
	ConnectionUnknown  ConnectionType = 0 // 0 Unknown
	ConnectionEthernet ConnectionType = 1 // 1	Ethernet; Wired Connection
	ConnectionWIFI     ConnectionType = 2 // 2	WIFI
	ConnectionCellular ConnectionType = 3 // 3	Cellular Network - Unknown Generation
	Connection2G       ConnectionType = 4 // 4	Cellular Network - 2G
	Connection3G       ConnectionType = 5 // 5	Cellular Network - 3G
	Connection4G       ConnectionType = 6 // 6	Cellular Network - 4G
	Connection5G       ConnectionType = 7 // 7	Cellular Network - 5G

	connUnknown  string = "Unknown"
	connEthernet string = "Ethernet"
	connWIFI     string = "WIFI"
	connCellular string = "Cellular"
	conn2G       string = "2G"
	conn3G       string = "3G"
	conn4G       string = "4G"
	conn5G       string = "5G"
)

// Ptr returns pointer to own value.
func (c ConnectionType) Ptr() *ConnectionType {
	return &c
}

// Val safely dereferences pointer, returning default value (ConnectionUnknown) for nil.
func (c *ConnectionType) Val() ConnectionType {
	if c == nil {
		return ConnectionUnknown
	}
	return *c
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (c *ConnectionType) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Wrong connection type value: ", value))
	}

	switch v {
	case connUnknown:
		*c = ConnectionUnknown
	case connEthernet:
		*c = ConnectionEthernet
	case connWIFI:
		*c = ConnectionWIFI
	case connCellular:
		*c = ConnectionCellular
	case conn2G:
		*c = Connection2G
	case conn3G:
		*c = Connection3G
	case conn4G:
		*c = Connection4G
	case conn5G:
		*c = Connection5G
	}

	return nil
}

// Value return json value, implement driver.Valuer interface
func (c *ConnectionType) Value() (driver.Value, error) {
	var v string

	switch *c {
	case ConnectionUnknown:
		v = connUnknown
	case ConnectionEthernet:
		v = connEthernet
	case ConnectionWIFI:
		v = connWIFI
	case ConnectionCellular:
		v = connCellular
	case Connection2G:
		v = conn2G
	case Connection3G:
		v = conn3G
	case Connection4G:
		v = conn4G
	case Connection5G:
		v = conn5G
	}

	return v, nil
}
