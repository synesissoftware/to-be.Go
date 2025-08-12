package to_be_test

import (
	"github.com/stretchr/testify/require"
	. "github.com/synesissoftware/to-be.Go"

	"testing"
)

func Test_Truthy_StringIsFalsey(t *testing.T) {

	require.Equal(t, false, StringIsFalsey(""))
	require.Equal(t, false, StringIsFalsey("Copyright ©"))

	require.Equal(t, true, StringIsFalsey("0"))
	require.Equal(t, true, StringIsFalsey("false"))
	require.Equal(t, true, StringIsFalsey(" FALSE"))
	require.Equal(t, true, StringIsFalsey("False"))
	require.Equal(t, true, StringIsFalsey("FaLSe"))
	require.Equal(t, true, StringIsFalsey("no"))
	require.Equal(t, true, StringIsFalsey("No "))
	require.Equal(t, true, StringIsFalsey("NO"))
	require.Equal(t, true, StringIsFalsey(" Off "))
	require.Equal(t, true, StringIsFalsey("off"))
	require.Equal(t, true, StringIsFalsey("OFF"))

	require.Equal(t, false, StringIsFalsey("1"))
	require.Equal(t, false, StringIsFalsey("true"))
	require.Equal(t, false, StringIsFalsey("TRUE"))
	require.Equal(t, false, StringIsFalsey("True"))
	require.Equal(t, false, StringIsFalsey("tRuE"))
	require.Equal(t, false, StringIsFalsey("yes"))
	require.Equal(t, false, StringIsFalsey(" YES"))
	require.Equal(t, false, StringIsFalsey("Yes   "))
	require.Equal(t, false, StringIsFalsey("yEs"))
}

func Test_Truthy_StringIsTruey(t *testing.T) {

	require.Equal(t, false, StringIsTruey(""))
	require.Equal(t, false, StringIsTruey("Copyright ©"))

	require.Equal(t, false, StringIsTruey("0"))
	require.Equal(t, false, StringIsTruey("false"))
	require.Equal(t, false, StringIsTruey(" FALSE"))
	require.Equal(t, false, StringIsTruey("False"))
	require.Equal(t, false, StringIsTruey("FaLSe"))
	require.Equal(t, false, StringIsTruey("no"))
	require.Equal(t, false, StringIsTruey("No "))
	require.Equal(t, false, StringIsTruey("NO"))
	require.Equal(t, false, StringIsTruey(" Off "))
	require.Equal(t, false, StringIsTruey("off"))
	require.Equal(t, false, StringIsTruey("OFF"))

	require.Equal(t, true, StringIsTruey("1"))
	require.Equal(t, true, StringIsTruey("true"))
	require.Equal(t, true, StringIsTruey("TRUE"))
	require.Equal(t, true, StringIsTruey("True"))
	require.Equal(t, true, StringIsTruey("tRuE"))
	require.Equal(t, true, StringIsTruey("yes"))
	require.Equal(t, true, StringIsTruey(" YES"))
	require.Equal(t, true, StringIsTruey("Yes   "))
	require.Equal(t, true, StringIsTruey("yEs"))
}

func Test_Truthy_StringIsTruthy(t *testing.T) {

	require.Equal(t, false, StringIsTruthy(""))
	require.Equal(t, false, StringIsTruthy("Copyright ©"))

	require.Equal(t, true, StringIsTruthy("0"))
	require.Equal(t, true, StringIsTruthy("false"))
	require.Equal(t, true, StringIsTruthy(" FALSE"))
	require.Equal(t, true, StringIsTruthy("False"))
	require.Equal(t, true, StringIsTruthy("FaLSe"))
	require.Equal(t, true, StringIsTruthy("no"))
	require.Equal(t, true, StringIsTruthy("No "))
	require.Equal(t, true, StringIsTruthy("NO"))
	require.Equal(t, true, StringIsTruthy(" Off "))
	require.Equal(t, true, StringIsTruthy("off"))
	require.Equal(t, true, StringIsTruthy("OFF"))

	require.Equal(t, true, StringIsTruthy("1"))
	require.Equal(t, true, StringIsTruthy("true"))
	require.Equal(t, true, StringIsTruthy("TRUE"))
	require.Equal(t, true, StringIsTruthy("True"))
	require.Equal(t, true, StringIsTruthy("tRuE"))
	require.Equal(t, true, StringIsTruthy("yes"))
	require.Equal(t, true, StringIsTruthy(" YES"))
	require.Equal(t, true, StringIsTruthy("Yes   "))
	require.Equal(t, true, StringIsTruthy("yEs"))
}
