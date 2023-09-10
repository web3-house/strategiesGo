package abi

func GetABI(name string) string {
	switch name {
	case BALANCE_OF:
		return BalanceOf
	case GET_ETH_BALANCE:
		return GetEthBalance
	case TOKEN_ID:
		return TokenId
	}
	return ""
}

func GetFuncName(name string) string {
	switch name {
	case BALANCE_OF, TOKEN_ID:
		return "balanceOf"
	case GET_ETH_BALANCE:
		return "getEthBalance"
	}
	return ""

}
