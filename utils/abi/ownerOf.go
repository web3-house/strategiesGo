package abi

const OwnerOf = `
	[
		{
			"constant": true,
			"inputs": [
				{
					"name": "tokenId",
					"type": "uint256"
				}
			],
			"name": "ownerOf",
			"outputs": [
				{
					"name": "owner",
					"type": "address"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		}
	]
`
