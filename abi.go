package main

const BalanceOfABI = `
	[
		{
			"constant": true,
			"inputs": [
				{
					"name": "account",
					"type": "address"
				}
			],
			"name": "balanceOf",
			"outputs": [
				{
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		}
	]
`
