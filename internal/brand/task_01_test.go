package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask01(t *testing.T) {
	now := time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC)
	s := NewService(NewRegistry(), func() time.Time { return now })
	l := activeLicense(now)
	l.Status = LicensePending
	_, err := s.ActivateLicense(context.Background(), l, compliantStore(now))
	require.NoError(t, err)
}
