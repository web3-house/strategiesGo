package abi

const GetEthBalance = `
	[
		{
			"constant": true,
			"inputs": [
				{
					"name": "addr",
					"type": "address"
				}
			],
			"name": "getEthBalance",
			"outputs": [
				{
					"name": "balance",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		}
	]
`
