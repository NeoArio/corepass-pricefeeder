// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import "./Ownable.sol";

contract PriceFeed is Ownable {
    struct Price {
        uint256 timestamp;
        uint256 price;
    }

    uint256 public constant decimals = 18;
    string public constant name = "CoreToken";
    uint256 aggregationInSeconds = 30 * 60;
    mapping(address => bool) authorizedNodes;

    Price[] prices;

    function authorizeNode(address _nodeAddress) external onlyOwner {
        authorizedNodes[_nodeAddress] = true;
    }

    function deauthorizeNode(address _nodeAddress) external onlyOwner {
        authorizedNodes[_nodeAddress] = false;
    }

    function authorized(address _nodeAddress) public view returns (bool) {
        return authorizedNodes[_nodeAddress];
    }

    modifier onlyAuthorized() {
        require(
            authorized(_msgSender()) == true,
            "PriceFeed: caller is not authorized"
        );
        _;
    }

    function addPrice(uint256 _price) external onlyAuthorized {
        _addLatestPrice(_price);
    }

    function _addLatestPrice(uint256 _price) private {
        prices.push(Price(block.timestamp, _price));
    }

    function getLatestPrice() public view returns (uint256, uint256) {
        if (prices.length == 0) return (0, 0);

        Price memory price = prices[prices.length - 1];
        return (price.timestamp, price.price);
    }

    function getAggregatedPrice()
        public
        view
        returns (
            uint256,
            uint256,
            uint32
        )
    {
        if (prices.length == 0) {
            return (0, 0, 0);
        }

        uint256 timestampSegment = _aggregationSegment(block.timestamp);
        uint256 price = 0;
        uint256 index = prices.length;
        uint32 count = 0;
        while (
            index > 0 &&
            _aggregationSegment(prices[--index].timestamp) >= timestampSegment
        ) {
            price = (count * price + prices[index].price) / (count + 1);
            count++;
        }

        return (timestampSegment, price, count);
    }

    function _aggregationSegment(uint256 _timestamp)
        private
        view
        returns (uint256)
    {
        return _timestamp - (_timestamp % aggregationInSeconds);
    }
}
