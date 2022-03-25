package localStorage

import (
	"anyData/internal/app/anyData/model"
)

/*
Example of the model:

{Id: 1, Address: "0xd2dF44930A7D43716c48c438B920563F9Bdb88db", Date: "2022-02-10", Amount: 100000000000000000},
{Id: 2, Address: "0x009c939223bC965D1b44F31d2eAd0DbAEB46Be53", Date: "2022-02-10", Amount: 120000000000000000},
{Id: 3, Address: "0x3a3F0892f33C1360884bC133af21805cB45726A2", Date: "2022-02-11", Amount: 110000000000000000},
{Id: 4, Address: "0xC9C79BAeB42Eb8B7D27D40771DD406e620F60334", Date: "2022-02-15", Amount: 130000000000000000},

*/

var DataStorage []model.AnyData
