pragma solidity ^0.5.2;

contract FundManager {
	// Accept donations with a cerntain % allocated to community
	// development ('perc'), and the rest to conservation efforts 
	function donate(uint8 perc) public payable;

	// exchange coins received as payment for hunter services for ether
	// only the contract representing the hunter payment can invoke this
	function redeemEther(address payable recipient, uint amount) public;

	// allows wwf employees to withdraw funds to cover operation costs
	// of conservation efforts, has to provide a category to track fund allocation:
	// e.g. one of the impact goals, salary, etc..., also provide detailed purpose
	function payOpCosts(address payable recipient, uint amount, string memory category, string memory purpose) public;

	// allows wwf employees to disburs bonus community development fund
	// ****add interaction with voting, community, etc...****
	function payCommProj(address payable recipient, uint amount, string memory purpose) public;

	// change project manager, careful (only reversibel by new address)
	function switchManager(address account) public;

	// provide viewing access to funds
	function getConFund() public view returns(uint);

	// provide viewing access to funds
	function getDevFund() public view returns(uint);

	function setHuntContrAddr(address account) public;

	function setProjManager(address account) public;
}