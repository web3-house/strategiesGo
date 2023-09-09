package abi

func GetABI(name string) string {
	switch name {
	case BALANCE_OF:
		return BalanceOf
	case GET_ETH_BALANCE:
		return GetEthBalance
	}
	return ""
}
