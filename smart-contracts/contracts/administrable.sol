pragma solidity ^0.5.2;

contract administrable{
	address projManager;
	Admin[] admins;
	uint numadmins = 0;
	mapping(address => bool) hashmins;

	struct Admin {
		address addr;
		string name;
	}

	modifier restricted (){
		require(hashmins[msg.sender], "Access restricted!");
		_;
	}

	event newlyRecruited(address admin, string name);
	event honorablyDischarged(address admin, string name);

	// modify wwf employee list, only allowed by project manager
	function addEmployee(address account, string memory name) public returns (uint){
		require(msg.sender == projManager);
		emit newlyRecruited(account, name);
		hashmins[account] = true;
		uint id = admins.push(Admin(account, name));
		numadmins++;
		return id;
	}

	// modify wwf employee list, only allowed by project manager
	function delEmployee(uint id) public {
		require(msg.sender == projManager);
		require(id < numadmins, "Invalid admin id!");
		emit honorablyDischarged(admins[id].addr, admins[id].name);
		hashmins[admins[id].addr] = false;
		admins[id] = admins[numadmins-1];
		delete admins[numadmins-1];
		numadmins--;
	}

	// change project manager, careful (only reversibel by new address)
	function switchManager(address account) public {
		require(msg.sender == projManager);
		projManager = account;
	}

	function getNumEmp() public view returns(uint){
		return numadmins;
	}

	// only use this to iterate through array, employee ids are not static
	function getEmployee(uint id) public view returns (address, string memory) {
		return (admins[id].addr, admins[id].name);
	}
}