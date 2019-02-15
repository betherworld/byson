pragma solidity ^0.5.2;

import "./administrable.sol";
import "./FundManagerInt.sol";
import "./VotingInt.sol";


contract Hunters is administrable{
	address fmContrAddress;	// address of contract holding all funds and responsible for distribution
	address htContrAddress;	// address of contract responsible for community project voting
	FundManager fm;
	HunterToken ht;

	Activity[] public activities;
	uint activityCount = 0;
	uint[] proofHashes;
	uint proofCount = 0;

	// this represents an action hunters can perform for ether & HTT
	// 3 states are defined:
	//  initially the contract is active (true), i.e. available for anyone to sign up
	//  once a hunter has signed up, it becomes inactive, blocking other sign ups
	//  then the hunter submits a proof to put a claim on the reward, now claimed (true)
	//  lastly if the claim is approved, the contract becomes completed (true), otherwise reset state
	struct Activity {
		uint ETHReward;
		uint RTReward;
		bool active;
		bool claimed;
		bool completed;
		string description;
		address payable user;
	}

	event newActivity (uint activNum);
	event actBooked (uint activNum, address hunt);
	event proofSent (uint activNum, uint hash);
	event proofValidated (uint ETHrew, uint RTrew, uint activNum, address hunt);
	event proofRejected(uint ETHrew, uint RTrew, uint activNum, address hunt);

	constructor (address fundManagerAddr, address votingAddr) public {
		fmContrAddress = fundManagerAddr;
		htContrAddress = votingAddr;
		fm = FundManager(fmContrAddress);
		ht = HunterToken(htContrAddress);
	}

	// wwf employees can post new tasks for hunters to complete, specifying the reward
	// in voting tokens and ether
	function createActivity(uint _rew, uint _ranRew, string memory  _description) public restricted{
		activities.push(activ(_rew,_ranRew,true,false,false,_description,address(0)));
		activityCount++;
		emit newActivity(activityCount - 1);
	}

	// hunters can block a task for themselves so it's not completed twice
	function bookActivity(uint _index) public {
		require(activities[_index].active);
		activities[_index].user = msg.sender;
		activities[_index].active = false;
		emit actBooked(_index,activities[_index].user);
	}

	// here hunters claim to have completed the task and submit proof of it
	function sendProof(uint _index, uint _hash) public {
		require(activities[_index].user == msg.sender);
		proofHashes.push(_hash);
		proofCount++;
		activities[_index].claimed = true;
		emit proofSent(_index,_hash);
	}

	// wwf employess either accept or decline the proof (oracle)
	function validateProof(uint _index, bool approved) public restricted {
		require(!activities[_index].active && activities[_index].claimed);
		if(approved){
			activities[_index].completed = true;
			emit proofValidated(activities[_index].ETHReward,activities[_index].RTReward,_index,activities[_index].user);
			ht.mint(activities[_index].user, activities[_index].RTReward);
			fm.redeemEther(activities[_index].user, activities[_index].ETHReward);
		}else{
			activities[_index].claimed = false;
			activities[_index].active = true;
			activities[_index].user = address(0);
			emit proofRejected(activities[_index].ETHReward, activities[_index].RTReward, _index, activities[_index].user);
		}
	}

	// use this to loop through activities or get a specific one, task ids are static here
	function displayActivity(uint id) public view returns(uint,uint,bool,bool,bool,string memory,address) {
		require(id < activityCount);
		return(activities[id].ETHReward,activities[id].RTReward,activities[id].active,activities[id].claimed,activities[id].completed,activities[id].description,activities[id].user);
	}

	function getActivityCount() public view returns(uint){
		return activityCount;
	}
}