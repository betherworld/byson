pragma solidity ^0.5.2;

import "./administrable.sol";

contract FundManager is administrable {
	uint devFund = 0;			// community development fund
	uint conFund = 0;			// conservation efforts fund
	address huntContrAddress;	// address of contract responsible for hunter/gatherer interactions
	address votingContrAddress; // address of contract responsible for community project voting

	event donationReceived(address donor, uint amount, uint8 perc);
	event costsPayed(address amdin, uint amount, string category, string purpose);
	event bonusDisbursed(address admin, uint amount, string purpose);

	// Accept donations with a cerntain % allocated to community
	// development ('perc'), and the rest to conservation efforts 
	function donate(uint8 perc) public payable {
		require(msg.value % 100 == 0, "Donations needs to be multiple of 100!");
		require(perc >= 0 && perc <= 100, "Invalid percentage!");
		devFund += (msg.value*perc)/100;
		conFund += (msg.value*(100-perc))/100;
		emit donationReceived(msg.sender, msg.value, perc);
	}

	// exchange coins received as payment for hunter services for ether
	// only the contract representing the hunter payment can invoke this
	function redeemEther(address payable recipient, uint amount) public {
		require(msg.sender == huntContrAddress, "Access restricted!");
		require(amount <= conFund, "Out of funds!");
		conFund -= amount;
		recipient.transfer(amount);
	}

	// allows wwf employees to withdraw funds to cover operation costs
	// of conservation efforts, has to provide a category to track fund allocation:
	// e.g. one of the impact goals, salary, etc..., also provide detailed purpose
	function payOpCosts(address payable recipient, uint amount, string memory category, string memory purpose) public restricted{
		require(amount <= conFund, "Out of funds!");
		conFund -= amount;
		emit costsPayed(msg.sender, amount, category, purpose);
		recipient.transfer(amount);
	}

	// allows wwf employees to disburs bonus community development fund
	// ****add interaction with voting, community, etc...****
	function payCommProj(address payable recipient, uint amount, string memory purpose) public {
		require(msg.sender == votingContrAddress);
		require(amount <= conFund, "Out of funds!");
		emit bonusDisbursed(msg.sender, amount, purpose);
		devFund -= amount;
		recipient.transfer(amount);
	}

	// provide viewing access to funds
	function getConFund() public view returns(uint){
		return conFund;
	}

	// provide viewing access to funds
	function getDevFund() public view returns(uint){
		return devFund;
	}

	// one-time executable functions to set priviliged contract addresses
	function setHuntContrAddr(address account) public {
	    require(huntContrAddress == address(0), "Can only be set once!");
	    huntContrAddress = account;
	}

	// one-time executable functions to set priviliged contract addresses
	function setVotingContrAddr(address account) public {
	    require(votingContrAddress == address(0), "Can only be set once!");
	    votingContrAddress = account;
	}

	// one-time executable functions to set priviliged project manager
	function setProjManager(address account) public {
	    require(projManager == address(0), "Can only be set once!");
	    projManager = account;
	}
}