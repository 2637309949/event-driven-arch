// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IUniswapV2Router {
    function swapExactETHForTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable returns (uint[] memory amounts);

    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
}

contract SimpleSwap {
    event SwapSucceeded(address indexed user, uint ethAmount, uint tokenAmount);
    event SwapFailed(address indexed user, uint ethAmount, string reason);

    // Sepolia 测试网 Uniswap V2 Router 地址
    address public constant UNISWAP_ROUTER = 0xeE567Fe1712Faf6149d80dA1E6934E354124CfE3;
    // Sepolia WETH 地址
    address public constant WETH = 0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14;
    // Sepolia USDC 地址
    address public constant USDC = 0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238;

    function swapETHForUSDC(uint amountOutMin, uint deadline) external payable {
        require(msg.value > 0, "need ETH to swap");
        require(amountOutMin > 0, "amountOutMin must be greater than 0");
        require(deadline >= block.timestamp, "invalid deadline");

        address[] memory path = new address[](2);
        path[0] = WETH;
        path[1] = USDC;

        // 检查预期输出，防止无效路径
        uint[] memory amountsOut = IUniswapV2Router(UNISWAP_ROUTER).getAmountsOut(msg.value, path);
        require(amountsOut[amountsOut.length - 1] >= amountOutMin, "Insufficient output amount");

        IUniswapV2Router router = IUniswapV2Router(UNISWAP_ROUTER);

        try router.swapExactETHForTokens{value: msg.value}(
            amountOutMin,
            path,
            msg.sender,
            deadline
        ) returns (uint[] memory amounts) {
            emit SwapSucceeded(msg.sender, msg.value, amounts[amounts.length - 1]);
        } catch Error(string memory reason) {
            emit SwapFailed(msg.sender, msg.value, reason);
            revert(reason);
        } catch {
            emit SwapFailed(msg.sender, msg.value, "Unknown error");
            revert("Swap failed with unknown error");
        }
    }

    receive() external payable {}
}