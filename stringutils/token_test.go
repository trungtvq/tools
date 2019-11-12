package stringutils_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trungtvq/go-utils/stringutils"
	"strings"
	"testing"
)

func TestGetLike(t *testing.T) {
	s1 := strings.Split(`app_trans_id
	status
	from_z_id
	from_zp_id
	from_avatar
	from_d_name
	from_message
	to_z_id
	to_zp_id
	to_avatar
	to_d_name
	to_message
	amount
	transfer_time
	open_time
	from_zp_trans
	to_zp_trans
	error_code
	error_message
	type
	bank_code
	last4_card_no
	first6_card_no
	last_account_no
	first_account_no
	qr_id
	xxx_unrecognized
	xxx_sizecache`, "\n\t")
	s2 := strings.Split(`dev_trans_statuses
	app_trans_id
	dev_trans_statuses
	zp_trans_id_submit_trans
	dev_trans_statuses
	zp_trans_id_update_order
	dev_trans_statuses
	sender_zalo_id
	dev_trans_statuses
	receiver_zalo_id
	dev_trans_statuses
	sender_zalo_pay_id
	dev_trans_statuses
	receiver_zalo_pay_id
	dev_trans_statuses
	amount
	dev_trans_statuses
	status
	dev_trans_statuses
	description
	dev_trans_statuses
	note
	dev_trans_statuses
	action_h5_fe
	dev_trans_statuses
	created_time
	dev_trans_statuses
	updated_time
	dev_trans_statuses
	sender_name
	dev_trans_statuses
	receiver_name
	dev_trans_statuses
	sender_avatar
	dev_trans_statuses
	receiver_avatar
	dev_trans_statuses
	zpi_type
	dev_trans_statuses
	zpi_id
	dev_trans_statuses
	is_no_confirm
	dev_trans_statuses
	error_code`, "\n\t")
	fmt.Printf("%s\n", stringutils.GetLike(s1, s2))
	assert.NotEqual(t,s1,s2)
}
func BenchmarkGetLike(b *testing.B) {
	s1 := strings.Split(`app_trans_id
	from_zp_trans`, "\n\t")
	s2 := strings.Split(`dev_trans_statuses
	app_trans_id`, "\n\t")
	for i:=0;i<b.N;i++{
		stringutils.GetLike(s1, s2)
	}
}
//1243