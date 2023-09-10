package ens

type Domain struct {
	Subdomains []Domain `json:"subdomains"`
}

const GET_DOMAINS_WITH_DOMAIN_NAME_AND_GET_SUBDOMAINS_WITH_OWNER_ADDRESSES = `
query GetDomainsWithDomainNameAndGetSubdomainsWithOwnerAddresses($domain_name: String, $owner_in: [String]) {
	domains(where: {name: $domain_name}, first: 1000) {
	  subdomains(where: { owner_in: $owner_in}) {
		owner {
		  id
		}
	  }
	}
}
`
