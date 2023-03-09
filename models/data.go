package models

import "time"

var (
	Access1 = Access{
		AccessID:     "ea42c4e7-151b-c327-e463-0ac7918cdc80",
		SessionID:    "879230a0-9236-53f0-1c17-dd4148371f49",
		UserName:     "Kyodai Taro",
		Organization: "Example Hospital",
		SAMLHash:     "B22F2528221F10A32B3A8D319E9263B458A769BAC83A3B54B9BC53D46B78D8A3EA77276A3AAD7921358B4A2C5E2B6DF40684A095E63383ABBD5ABB51FA99DCF8",
		Status:       "",
		AuthorizedAt: time.Now(),
		CreatedAt:    time.Now(),
	}

	Access2 = Access{
		AccessID:     "f8f94cb9-a0ee-e497-66a4-99dd02071c2d",
		SessionID:    "4ba30f6e-2aef-ba5a-d73f-b894a802cf66",
		UserName:     "Kyodai Hanako",
		Organization: "Example Hospital",
		SAMLHash:     "742E98945E11599BBD51C566BEDE26E041F11BDE5FA6DFCEEE615BD2B82FF7E71F9641E10EF3BF6299F69E333D32AB50F4FC93B8C3B0E9B563438BF9B7F8367B",
		Status:       "",
		AuthorizedAt: time.Now(),
		CreatedAt:    time.Now(),
	}
)

var (
	IdP1 = IdP{
		ID:        "694d9a03-c7ec-8c7e-33dc-bb7b9f2aa38f",
		Name:      "Example Hospital1",
		Url:       "https://example.com/idp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}

	IdP2 = IdP{
		ID:        "88595338-b732-84fc-b9be-7eef0838b01d",
		Name:      "Example Hospital2",
		Url:       "https://example.com/idp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}
)

var (
	PIdP1 = PIdP{
		ID:        "b400ed6f-bb03-c1cc-bbb9-24357a2dabe7",
		Name:      "Credit Card Company1",
		Url:       "https://example.org/pidp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}

	PIdP2 = PIdP{
		ID:        "413c1c33-2542-ef93-6625-ff8ff74052eb",
		Name:      "Credit Card Company2",
		Url:       "https://example.org/pidp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}
)

var (
	PDP1 = PDP{
		ID:        "251ad440-aa66-9840-a26f-41c346f49d73",
		Name:      "PHR Service1",
		Url:       "https://example.net/pdp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}

	PDP2 = PDP{
		ID:        "38ae56f0-b54a-3d95-f5c3-b44cb2148fa1",
		Name:      "PHR Service2",
		Url:       "https://example.net/pdp",
		ExpireAt:  time.Now(),
		CreatedAt: time.Now(),
	}
)
