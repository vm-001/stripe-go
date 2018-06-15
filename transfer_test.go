package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestTransfer_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Transfer
		err := json.Unmarshal([]byte(`"tr_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "tr_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Transfer{ID: "tr_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "tr_123", v.ID)
	}
}

func TestTransferDestination_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v TransferDestination
		err := json.Unmarshal([]byte(`"acct_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "acct_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because TransferDestination
		// has a custom MarshalJSON implementation as well, and it'll turn into
		// a string if we marshaled a struct instance. This ensures that we're
		// working with a JSON objects.
		data := []byte(`{"id":"acct_123"}`)

		var v TransferDestination
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "acct_123", v.ID)

		// The child Account field should also be expanded. For legacy reasons
		// it's a different object.
		assert.Equal(t, "acct_123", v.Account.ID)
	}
}
